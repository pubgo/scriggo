// Copyright (c) 2019 Open2b Software Snc. All rights reserved.
// https://www.open2b.com

// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"context"
	"io"
	"reflect"
	"sync"

	"github.com/open2b/scriggo/ast"
	"github.com/open2b/scriggo/native"
)

type PrintFunc func(interface{})

// Context represents a context in Show and Text instructions.
type Context byte

type Renderer interface {
	Show(env native.Env, v interface{}, ctx Context)
	Text(env native.Env, txt []byte, inURL, isSet bool)
	Out() io.Writer
	WithOut(out io.Writer) Renderer
	WithConversion(fromFormat, toFormat ast.Format) Renderer
	Close() error
}

// The env type implements the native.Env interface.
type env struct {
	ctx     context.Context // context.
	globals []reflect.Value // global variables.
	print   PrintFunc       // custom print builtin.
	typeof  TypeOfFunc      // typeof function.

	// Only exited, exits and filePath fields can be changed after the vm has
	// been started and access to these three fields must be done with this
	// mutex.
	mu       sync.Mutex
	exited   bool     // reports whether it is exited.
	exits    []func() // exit functions.
	filePath string   // path of the file where the main goroutine is in.
}

func (env *env) Context() context.Context {
	return env.ctx
}

func (env *env) Exit(code int) {
	panic(exitError(code))
}

func (env *env) Exited() bool {
	var exited bool
	env.mu.Lock()
	exited = env.exited
	env.mu.Unlock()
	return exited
}

func (env *env) ExitFunc(f func()) {
	env.mu.Lock()
	if env.exited {
		go f()
	} else {
		env.exits = append(env.exits, f)
	}
	env.mu.Unlock()
	return
}

func (env *env) Fatal(v interface{}) {
	panic(&FatalError{env: env, msg: v})
}

func (env *env) FilePath() string {
	env.mu.Lock()
	filePath := env.filePath
	env.mu.Unlock()
	return filePath
}

func (env *env) Print(args ...interface{}) {
	for _, arg := range args {
		env.doPrint(arg)
	}
}

func (env *env) Println(args ...interface{}) {
	for i, arg := range args {
		if i > 0 {
			env.doPrint(" ")
		}
		env.doPrint(arg)
	}
	env.doPrint("\n")
}

func (env *env) TypeOf(v reflect.Value) reflect.Type {
	return env.typeof(v)
}

func typeOfFunc(v reflect.Value) reflect.Type {
	return v.Type()
}

func (env *env) doPrint(arg interface{}) {
	if env.print != nil {
		env.print(arg)
		return
	}
	r := reflect.ValueOf(arg)
	switch r.Kind() {
	case reflect.Invalid, reflect.Array, reflect.Func, reflect.Struct:
		print(hex(reflect.ValueOf(&arg).Elem().InterfaceData()[1]))
	case reflect.Bool:
		print(r.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		print(r.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		print(r.Uint())
	case reflect.Float32, reflect.Float64:
		print(r.Float())
	case reflect.Complex64, reflect.Complex128:
		print(r.Complex())
	case reflect.Chan, reflect.Map, reflect.UnsafePointer:
		print(hex(r.Pointer()))
	case reflect.Interface, reflect.Ptr:
		print(arg)
	case reflect.Slice:
		print("[", r.Len(), "/", r.Cap(), "]", hex(r.Pointer()))
	case reflect.String:
		print(r.String())
	}
}

func (env *env) exit() {
	env.mu.Lock()
	if !env.exited {
		for _, f := range env.exits {
			go f()
		}
		env.exits = nil
		env.exited = true
	}
	env.mu.Unlock()
}