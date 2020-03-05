// Copyright (c) 2019 Open2b Software Snc. All rights reserved.
// https://www.open2b.com

// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package templates

import (
	"context"
	"errors"
	"fmt"
	"io"
	"reflect"
	_sort "sort"

	"scriggo/compiler"
	"scriggo/compiler/ast"
	"scriggo/runtime"
)

var (
	// ErrInvalidPath is returned from the Load function and a FileReader when
	// the path argument is not valid.
	ErrInvalidPath = errors.New("scriggo: invalid path")

	// ErrNotExist is returned from the Load function when the path does not
	// exist.
	ErrNotExist = errors.New("scriggo: path does not exist")

	// ErrReadTooLarge is returned from the Load function when a limit is
	// exceeded reading a path.
	ErrReadTooLarge = errors.New("scriggo: read too large")
)

// HTMStringer is implemented by values that are not escaped in HTML context.
type HTMLStringer interface {
	HTML() string
}

// CSSStringer is implemented by values that are not escaped in CSS context.
type CSSStringer interface {
	CSS() string
}

// JavaScriptStringer is implemented by values that are not escaped in
// JavaScript context.
type JavaScriptStringer interface {
	JavaScript() string
}

// HTML implements the HTMLStringer interface.
type HTML string

func (html HTML) HTML() string {
	return string(html)
}

// CSS implements the CSSStringer interface.
type CSS string

func (css CSS) CSS() string {
	return string(css)
}

// JavaScript implements the JavaScriptStringer interface.
type JavaScript string

func (js JavaScript) JavaScript() string {
	return string(js)
}

// A Language represents a source language.
type Language int

const (
	LanguageText Language = iota
	LanguageHTML
	LanguageCSS
	LanguageJavaScript
)

func (language Language) String() string {
	return ast.Language(language).String()
}

type LoadOptions struct {
	DisallowGoStmt  bool
	LimitMemory     bool                  // limit the execution memory size.
	TreeTransformer func(*ast.Tree) error // if not nil transforms tree after parsing.
}

// Declarations.
type Declarations map[string]interface{}

type RenderOptions struct {
	Context       context.Context
	MemoryLimiter runtime.MemoryLimiter
	PrintFunc     runtime.PrintFunc
}

type Template struct {
	fn          *runtime.Function
	globals     []compiler.Global
	limitMemory bool
}

// CompilerError represents an error returned by the compiler.
type CompilerError interface {
	error
	Position() ast.Position
	Path() string
	Message() string
}

// Load loads a template given its file name. Load calls the method ReadFile of
// files to read the files of the template. builtins declares constants, types,
// variables and functions that are accessible from the code in the template.
func Load(name string, files FileReader, builtins Declarations, lang Language, options *LoadOptions) (*Template, error) {
	co := compiler.Options{
		Builtins:       compiler.Declarations(builtins),
		RelaxedBoolean: true,
	}
	if options != nil {
		co.LimitMemory = options.LimitMemory
		co.TreeTransformer = options.TreeTransformer
		co.DisallowGoStmt = options.DisallowGoStmt
	}
	code, err := compiler.CompileTemplate(name, files, ast.Language(lang), co)
	if err != nil {
		if err == compiler.ErrInvalidPath {
			return nil, ErrInvalidPath
		}
		if err == compiler.ErrNotExist {
			return nil, ErrNotExist
		}
		return nil, err
	}
	return &Template{fn: code.Main, globals: code.Globals, limitMemory: co.LimitMemory}, nil
}

var emptyVars = map[string]interface{}{}

// Render renders the template and write the output to out. vars contains the
// values of the template builtin variables.
func (t *Template) Render(out io.Writer, vars map[string]interface{}, options *RenderOptions) error {
	if options != nil && options.MemoryLimiter != nil && !t.limitMemory {
		panic("scriggo: template not loaded with LimitMemory option")
	}
	writeFunc := out.Write
	renderFunc := render
	uw := &urlEscaper{w: out}
	t.globals[0].Value = &out
	t.globals[1].Value = &writeFunc
	t.globals[2].Value = &renderFunc
	t.globals[3].Value = &uw
	if vars == nil {
		vars = emptyVars
	}
	vm := newVM(options)
	_, err := vm.Run(t.fn, initGlobals(t.globals, vars))
	return err
}

// MustRender is like Render but panics if the rendering fails.
func (t *Template) MustRender(out io.Writer, vars map[string]interface{}, options *RenderOptions) {
	err := t.Render(out, vars, options)
	if err != nil {
		panic(err)
	}
}

// Disassemble disassembles a template.
func (t *Template) Disassemble(w io.Writer) (int64, error) {
	return compiler.DisassembleFunction(w, t.fn, t.globals)
}

// Vars returns the names of the template builtin variables that are used in
// the template.
func (t *Template) Vars() []string {
	vars := make([]string, len(t.globals)-4)
	for i, global := range t.globals[4:] {
		vars[i] = global.Name
	}
	_sort.Strings(vars)
	return vars
}

// newVM returns a new vm with the given options.
func newVM(options *RenderOptions) *runtime.VM {
	vm := runtime.NewVM()
	if options != nil {
		if options.Context != nil {
			vm.SetContext(options.Context)
		}
		vm.SetMemoryLimiter(options.MemoryLimiter)
		if options.PrintFunc != nil {
			vm.SetPrint(options.PrintFunc)
		}
	}
	return vm
}

// initGlobals initializes the global variables and returns the values. It
// panics if init is not valid.
//
// This function is a copy of the function in the scriggo package.
func initGlobals(globals []compiler.Global, init map[string]interface{}) []interface{} {
	n := len(globals)
	if n == 0 {
		return nil
	}
	values := make([]interface{}, n)
	for i, global := range globals {
		if global.Pkg == "main" {
			if value, ok := init[global.Name]; ok {
				if global.Value != nil {
					panic(fmt.Sprintf("variable %q already initialized", global.Name))
				}
				if value == nil {
					panic(fmt.Sprintf("variable initializer %q cannot be nil", global.Name))
				}
				val := reflect.ValueOf(value)
				if typ := val.Type(); typ == global.Type {
					v := reflect.New(typ).Elem()
					v.Set(val)
					values[i] = v.Addr().Interface()
				} else {
					if typ.Kind() != reflect.Ptr || typ.Elem() != global.Type {
						panic(fmt.Sprintf("variable initializer %q must have type %s or %s, but have %s",
							global.Name, global.Type, reflect.PtrTo(global.Type), typ))
					}
					if val.IsNil() {
						panic(fmt.Sprintf("variable initializer %q cannot be a nil pointer", global.Name))
					}
					values[i] = value
				}
				continue
			}
		}
		if global.Value == nil {
			values[i] = reflect.New(global.Type).Interface()
		} else {
			values[i] = global.Value
		}
	}
	return values
}