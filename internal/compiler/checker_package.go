// Copyright (c) 2019 Open2b Software Snc. All rights reserved.
// https://www.open2b.com

// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package compiler

import (
	"fmt"
	"reflect"

	"scriggo/ast"
)

// Sync with scriggo.Package.
type scriggoPackage interface {
	Name() string
	Lookup(declName string) interface{}
	DeclarationNames() []string
}

// toTypeCheckerScope generates a type checker scope given a predefined package.
// depth must be 0 unless toTypeCheckerScope is called recursively.
func toTypeCheckerScope(pp predefinedPackage, depth int, opts checkerOptions) typeCheckerScope {
	pkgName := pp.Name()
	declarations := pp.DeclarationNames()
	s := make(typeCheckerScope, len(declarations))
	for _, ident := range declarations {
		value := pp.Lookup(ident)
		// Import an auto-imported package. This is supported in package-less
		// programs and templates only.
		if p, ok := value.(scriggoPackage); ok {
			if opts.SyntaxType == ProgramSyntax && !opts.PackageLess {
				panic(fmt.Errorf("scriggo: auto-imported packages are supported only for package-less programs and templates"))
			}
			if depth > 0 {
				panic(fmt.Errorf("scriggo: cannot have an auto-imported package inside another auto-imported package"))
			}
			autoPkg := &packageInfo{}
			autoPkg.Declarations = map[string]*typeInfo{}
			for n, d := range toTypeCheckerScope(p, depth+1, opts) {
				autoPkg.Declarations[n] = d.t
			}
			autoPkg.Name = p.Name()
			s[ident] = scopeElement{t: &typeInfo{
				value:      autoPkg,
				Properties: propertyIsPackage | propertyHasValue,
			}}
			continue
		}
		// Import a type.
		if t, ok := value.(reflect.Type); ok {
			s[ident] = scopeElement{t: &typeInfo{
				Type:              t,
				Properties:        propertyIsType | propertyIsPredefined,
				PredefPackageName: pkgName,
			}}
			continue
		}
		// Import a variable.
		if reflect.TypeOf(value).Kind() == reflect.Ptr {
			v := reflect.ValueOf(value)
			s[ident] = scopeElement{t: &typeInfo{
				Type:              reflect.TypeOf(value).Elem(),
				value:             &v,
				Properties:        propertyAddressable | propertyIsPredefined | propertyHasValue,
				PredefPackageName: pkgName,
			}}
			continue
		}
		// Import a function.
		if typ := reflect.TypeOf(value); typ.Kind() == reflect.Func {
			s[ident] = scopeElement{t: &typeInfo{
				Type:              removeEnvArg(typ, false),
				value:             reflect.ValueOf(value),
				Properties:        propertyIsPredefined | propertyHasValue,
				PredefPackageName: pkgName,
			}}
			continue
		}
		// Import an untyped constant.
		if c, ok := value.(UntypedConstant); ok {
			constant, typ, err := parseConstant(string(c))
			if err != nil {
				panic(fmt.Errorf("scriggo: invalid untyped constant %q for %s.%s", c, pp.Name(), ident))
			}
			s[ident] = scopeElement{t: &typeInfo{
				Type:              typ,
				Properties:        propertyUntyped,
				Constant:          constant,
				PredefPackageName: pkgName,
			}}
			continue
		}
		// Import a typed constant.
		constant := convertToConstant(value)
		if constant == nil {
			panic(fmt.Errorf("scriggo: invalid constant value %v for %s.%s", value, pkgName, ident))
		}
		s[ident] = scopeElement{t: &typeInfo{
			Type:              reflect.TypeOf(value),
			Constant:          constant,
			PredefPackageName: pkgName,
		}}
	}
	return s
}

// pkgPathToIndex maps a package path to an unique identifier.
// TODO(Gianluca): it's not safe to use a global variable. More than this, we need a structure in common among all type checkers, which should hold:
//
// - pkgPathToIndex
// - all the collected pkgInfos
// - information about which tree has already been checked.
//
var pkgPathToIndex = map[string]int{}

type packageInfo struct {
	Name         string
	Declarations map[string]*typeInfo
	IndirectVars map[*ast.Identifier]bool
	TypeInfos    map[ast.Node]*typeInfo
}

func depsOf(name string, deps PackageDeclsDeps) []*ast.Identifier {
	for g, d := range deps {
		if g.Name == name {
			return d
		}
	}
	return nil
}

func checkDepsPath(path []*ast.Identifier, deps PackageDeclsDeps) []*ast.Identifier {
	last := path[len(path)-1]
	for _, dep := range depsOf(last.Name, deps) {
		for _, p := range path {
			if p.Name == dep.Name {
				return append(path, dep)
			}
		}
		loopPath := checkDepsPath(append(path, dep), deps)
		if loopPath != nil {
			return loopPath
		}
	}
	return nil
}

// An initLoopError is a error representing an initialization loop.
type initLoopError struct {
	node ast.Node
	msg  string
}

func (loopErr initLoopError) Error() string {
	return loopErr.node.Pos().String() + ": " + loopErr.msg
}

func detectConstantsLoop(consts []*ast.Const, deps PackageDeclsDeps) error {
	for _, c := range consts {
		path := []*ast.Identifier{c.Lhs[0]}
		loopPath := checkDepsPath(path, deps)
		if loopPath != nil {
			msg := "constant definition loop\n"
			for i := 0; i < len(loopPath)-1; i++ {
				msg += "\t" + loopPath[i].Pos().String() + ": "
				msg += loopPath[i].String() + " uses " + loopPath[i+1].String()
				msg += "\n"
			}
			return initLoopError{node: c, msg: msg}
		}
	}
	return nil
}

func detectVarsLoop(vars []*ast.Var, deps PackageDeclsDeps) error {
	for _, v := range vars {
		for _, left := range v.Lhs {
			path := []*ast.Identifier{left}
			loopPath := checkDepsPath(path, deps)
			if loopPath != nil {
				msg := "typechecking loop involving " + v.String() + "\n"
				for _, p := range loopPath {
					msg += "\t" + p.Pos().String() + ": " + p.String() + "\n"
				}
				return initLoopError{node: v, msg: msg}
			}
		}
	}
	return nil
}

func detectTypeLoop(types []*ast.TypeDeclaration, deps PackageDeclsDeps) error {
	for _, t := range types {
		if !t.IsAliasDeclaration {
			// TODO: currently the initialization loop check for type
			// definitions is the same as for type declarations. Is this
			// correct?
		}
		path := []*ast.Identifier{t.Identifier}
		loopPath := checkDepsPath(path, deps)
		if loopPath != nil {
			msg := "invalid recursive type alias " + t.String() + "\n"
			for _, p := range loopPath {
				msg += "\t" + p.Pos().String() + ": " + p.String() + "\n"
			}
			return initLoopError{node: t, msg: msg}
		}
	}
	return nil
}

func sortDeclarations(pkg *ast.Package) error {
	var extends *ast.Extends
	types := []*ast.TypeDeclaration{}
	consts := []*ast.Const{}
	vars := []*ast.Var{}
	imports := []*ast.Import{}
	funcs := []*ast.Func{}

	// Fragments global declarations.
	for _, decl := range pkg.Declarations {
		switch decl := decl.(type) {
		case *ast.Extends:
			extends = decl
		case *ast.Import:
			imports = append(imports, decl)
		case *ast.Func:
			funcs = append(funcs, decl)
		case *ast.Macro:
			funcs = append(funcs, macroToFunc(decl))
		case *ast.Const:
			if len(decl.Rhs) == 0 {
				for i := range decl.Lhs {
					consts = append(consts, ast.NewConst(decl.Pos(), decl.Lhs[i:i+1], decl.Type, nil, decl.Index))
				}
			} else {
				for i := range decl.Lhs {
					consts = append(consts, ast.NewConst(decl.Pos(), decl.Lhs[i:i+1], decl.Type, decl.Rhs[i:i+1], decl.Index))
				}
			}
		case *ast.TypeDeclaration:
			types = append(types, decl)
		case *ast.Var:
			if len(decl.Lhs) == len(decl.Rhs) {
				for i := range decl.Lhs {
					vars = append(vars, ast.NewVar(decl.Pos(), decl.Lhs[i:i+1], decl.Type, decl.Rhs[i:i+1]))
				}
			} else {
				vars = append(vars, decl)
			}
		}
	}

	// Removes from deps all non-global dependencies.
	deps := AnalyzeTree(pkg)
	for decl, ds := range deps {
		newDs := []*ast.Identifier{}
		for _, d := range ds {
			// Is d a global identifier?
			for _, c := range consts {
				if c.Lhs[0].Name == d.Name {
					newDs = append(newDs, d)
				}
			}
			for _, t := range types {
				if t.Identifier.Name == d.Name {
					newDs = append(newDs, d)
				}
			}
			for _, f := range funcs {
				if f.Ident.Name == d.Name {
					newDs = append(newDs, d)
				}
			}
			for _, v := range vars {
				for _, left := range v.Lhs {
					if left.Name == d.Name {
						newDs = append(newDs, d)
					}
				}
			}
		}
		deps[decl] = newDs
	}

	err := detectConstantsLoop(consts, deps)
	if err != nil {
		return err
	}
	err = detectVarsLoop(vars, deps)
	if err != nil {
		return err
	}
	err = detectTypeLoop(types, deps)
	if err != nil {
		return err
	}

	// Sort types declarations.
	sortedTypes := []*ast.TypeDeclaration{}
typesLoop:
	for len(types) > 0 {
		// Searches for next type declaration with resolved deps.
		for i, t := range types {
			depsOk := true
			for _, dep := range deps[t.Identifier] {
				found := false
				for _, resolvedT := range sortedTypes {
					if dep.Name == resolvedT.Identifier.Name {
						// This dependency has been resolved: move
						// on checking for next one.
						found = true
						break
					}
				}
				if !found {
					// A dependency of t cannot be found.
					depsOk = false
					break
				}
			}
			if depsOk {
				// All dependencies of types are resolved.
				sortedTypes = append(sortedTypes, t)
				types = append(types[:i], types[i+1:]...)
				continue typesLoop
			}
		}
		// All remaining constants are added as "sorted"; a later stage of
		// the type checking will report the error (usually a not defined
		// symbol).
		sortedTypes = append(sortedTypes, types...)
		types = nil
	}

	// Sorts constants.
	sortedConsts := []*ast.Const{}
constsLoop:
	for len(consts) > 0 {
		// Searches for next constant with resolved deps.
		for i, c := range consts {
			depsOk := true
			for _, dep := range deps[c.Lhs[0]] {
				found := false
				for _, resolvedC := range sortedConsts {
					if dep.Name == resolvedC.Lhs[0].Name {
						// This dependency has been resolved: move
						// on checking for next one.
						found = true
						break
					}
				}
				for _, t := range sortedTypes {
					if dep.Name == t.Identifier.Name {
						// This dependency has been resolved: move
						// on checking for next one.
						found = true
						break
					}
				}
				if !found {
					// A dependency of c cannot be found.
					depsOk = false
					break
				}
			}
			if depsOk {
				// All dependencies of consts are resolved.
				sortedConsts = append(sortedConsts, c)
				consts = append(consts[:i], consts[i+1:]...)
				continue constsLoop
			}
		}
		if true {
			// All remaining constants are added as "sorted"; a later stage of
			// the type checking will report the error (usually a not defined
			// symbol).
			sortedConsts = append(sortedConsts, consts...)
			consts = nil
		}
	}

	// Sorts variables.
	sortedVars := []*ast.Var{}
varsLoop:
	for len(vars) > 0 {
		// Searches for next variable with resolved deps.
		for i, v := range vars {
			depsOk := true
			for _, dep := range deps[v.Lhs[0]] {
				found := false
			resolvedLoop:
				for _, resolvedV := range sortedVars {
					for _, left := range resolvedV.Lhs {
						if dep.Name == left.Name {
							// This dependency has been resolved: move
							// on checking for next one.
							found = true
							break resolvedLoop
						}
					}
				}
				for _, resolvedC := range sortedConsts {
					if dep.Name == resolvedC.Lhs[0].Name {
						// This dependency has been resolved: move
						// on checking for next one.
						found = true
						break
					}
				}
				for _, f := range funcs {
					if dep.Name == f.Ident.Name {
						// This dependency has been resolved: move
						// on checking for next one.
						found = true
						break
					}
				}
				for _, t := range sortedTypes {
					if dep.Name == t.Identifier.Name {
						// This dependency has been resolved: move
						// on checking for next one.
						found = true
						break
					}
				}
				if !found {
					// A dependency of v cannot be found.
					depsOk = false
					break
				}
			}
			if depsOk {
				// All dependencies of vars are resolved.
				sortedVars = append(sortedVars, v)
				vars = append(vars[:i], vars[i+1:]...)
				continue varsLoop
			}
		}
		// All remaining vars are added as "sorted"; a later stage of the
		// type checking will report the error (usually a not defined
		// symbol).
		sortedVars = append(sortedVars, vars...)
		vars = nil
	}

	sorted := []ast.Node{}
	if extends != nil {
		sorted = append(sorted, extends)
	}
	for _, imp := range imports {
		sorted = append(sorted, imp)
	}
	for _, t := range sortedTypes {
		sorted = append(sorted, t)
	}
	for _, c := range sortedConsts {
		sorted = append(sorted, c)
	}
	for _, v := range sortedVars {
		sorted = append(sorted, v)
	}
	for _, f := range funcs {
		sorted = append(sorted, f)
	}
	pkg.Declarations = sorted

	return nil

}

// checkPackage type checks a package.
func checkPackage(pkg *ast.Package, path string, imports PackageLoader, pkgInfos map[string]*packageInfo, opts checkerOptions, globalScope typeCheckerScope) (err error) {

	defer func() {
		if r := recover(); r != nil {
			if rerr, ok := r.(*CheckingError); ok {
				err = rerr
			} else {
				panic(r)
			}
		}
	}()

	tc := newTypechecker(path, opts, globalScope)

	err = sortDeclarations(pkg)
	if err != nil {
		loopErr := err.(initLoopError)
		return tc.errorf(loopErr.node, loopErr.msg)
	}

	// First: import packages.
	for _, d := range pkg.Declarations {
		if d, ok := d.(*ast.Import); ok {
			err := tc.checkImport(d, imports, pkgInfos, true)
			if err != nil {
				return err
			}
		}
	}

	// Second: check all type declarations.
	for _, d := range pkg.Declarations {
		if td, ok := d.(*ast.TypeDeclaration); ok {
			name, ti := tc.checkTypeDeclaration(td)
			if ti != nil {
				tc.assignScope(name, ti, td.Identifier)
			}
		}
	}

	// Defines functions in file/package block before checking all
	// declarations.
	for _, d := range pkg.Declarations {
		if f, ok := d.(*ast.Func); ok {
			if isBlankIdentifier(f.Ident) {
				continue
			}
			if f.Ident.Name == "init" || f.Ident.Name == "main" {
				if len(f.Type.Parameters) > 0 || len(f.Type.Result) > 0 {
					return tc.errorf(f.Ident, "func %s must have no arguments and no return values", f.Ident.Name)
				}
			}
			// Function type must be checked for every function, including
			// 'init's functions.
			funcType := tc.checkType(f.Type).Type
			if f.Ident.Name == "init" {
				// Do not add the 'init' function to the file/package block.
				continue
			}
			if _, ok := tc.filePackageBlock[f.Ident.Name]; ok {
				return tc.errorf(f.Ident, "%s redeclared in this block", f.Ident.Name)
			}
			tc.filePackageBlock[f.Ident.Name] = scopeElement{t: &typeInfo{Type: funcType}}
		}
	}

	// Type check and defined functions, variables and constants.
	for _, d := range pkg.Declarations {
		switch d := d.(type) {
		case *ast.Func:
			tc.enterScope()
			tc.addToAncestors(d)
			// Adds parameters to the function body scope.
			tc.checkDuplicateParams(d.Type)
			tc.addMissingTypes(d.Type)
			isVariadic := d.Type.IsVariadic
			for i, param := range d.Type.Parameters {
				t := tc.checkType(param.Type)
				if param.Ident != nil && !isBlankIdentifier(param.Ident) {
					if isVariadic && i == len(d.Type.Parameters)-1 {
						tc.assignScope(param.Ident.Name, &typeInfo{Type: tc.types.SliceOf(t.Type), Properties: propertyAddressable}, nil)
					} else {
						tc.assignScope(param.Ident.Name, &typeInfo{Type: t.Type, Properties: propertyAddressable}, nil)
					}
				}
			}
			// Adds named return values to the function body scope.
			for _, ret := range d.Type.Result {
				t := tc.checkType(ret.Type)
				if ret.Ident != nil && !isBlankIdentifier(ret.Ident) {
					tc.assignScope(ret.Ident.Name, &typeInfo{Type: t.Type, Properties: propertyAddressable}, nil)
				}
			}
			d.Body.Nodes, err = tc.checkNodesError(d.Body.Nodes)
			if err != nil {
				return err
			}
			// «If the function's signature declares result parameters, the
			// function body's statement list must end in a terminating
			// statement.»
			if len(d.Type.Result) > 0 {
				if !tc.terminating {
					panic(tc.errorf(d, "missing return at end of function"))
				}
			}
			tc.ancestors = tc.ancestors[:len(tc.ancestors)-1]
			tc.exitScope()
		case *ast.Const:
			tc.checkConstantDeclaration(d)
		case *ast.Var:
			tc.checkVariableDeclaration(d)
		}
	}

	if !tc.opts.AllowNotUsed {
		for pkg := range tc.unusedImports {
			// https://github.com/open2b/scriggo/issues/309.
			return tc.errorf(new(ast.Position), "imported and not used: \"%s\"", pkg)
		}
	}

	if pkg.Name == "main" {
		main, ok := tc.filePackageBlock["main"]
		if !ok {
			return tc.errorf(new(ast.Position), "function main is undeclared in the main package")
		}
		if main.t.Type.Kind() != reflect.Func || main.t.Addressable() {
			return tc.errorf(new(ast.Position), "cannot declare main - must be func")
		}
	}

	pkgInfo := &packageInfo{
		Name:         pkg.Name,
		Declarations: make(map[string]*typeInfo, len(pkg.Declarations)),
		TypeInfos:    tc.typeInfos,
	}
	pkgInfo.Declarations = make(map[string]*typeInfo)
	for ident, ti := range tc.filePackageBlock {
		pkgInfo.Declarations[ident] = ti.t
	}
	pkgInfo.IndirectVars = tc.indirectVars

	pkgInfos[path] = pkgInfo

	return nil
}
