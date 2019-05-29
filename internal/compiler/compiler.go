// Copyright (c) 2019 Open2b Software Snc. All rights reserved.
// https://www.open2b.com

// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The compiler implements parsing, type-checking and emitting of sources.
//
// Parsing
//
// Parsing is done using
//
//	ParseTemplate(...)
//	ParseProgram(...)
//	ParseScript(...) (currently not implemented)
//
// Typechecking
//
// When parsing is done, tree can be type-checked by:
//
// 	Typecheck(...)
//
// Emitting
//
// To emit a type-checked tree, use:
//
//    EmitSingle(...)
//    EmitPackageMain(...)
//
package compiler

import (
	"scrigo/internal/compiler/ast"
)

type Options struct {

	// AllowImports makes import statements available.
	AllowImports bool

	// NotUsedError returns a checking error if a variable is declared and not
	// used or a package is imported and not used.
	NotUsedError bool

	// IsPackage indicate if it's a package. If true, all sources must start
	// with "package" and package-level declarations will be sorted according
	// to Go package inizialization specs.
	IsPackage bool

	// DisallowGoStmt disables the "go" statement.
	DisallowGoStmt bool
}

func Typecheck(opts *Options, tree *ast.Tree, main *PredefinedPackage, imports map[string]*PredefinedPackage, deps GlobalsDependencies, customBuiltins TypeCheckerScope) (_ map[string]*PackageInfo, err error) {
	if opts.IsPackage && main != nil {
		panic("cannot have package main with option IsPackage enabled")
	}
	if opts.IsPackage && customBuiltins != nil {
		panic("cannot have customBuiltins with option IsPackage enabled")
	}
	if imports != nil && !opts.IsPackage {
		panic("cannot have imports when checking a non-package")
	}
	if deps != nil && !opts.IsPackage {
		panic("cannot have deps when checking a non-package")
	}
	defer func() {
		if r := recover(); r != nil {
			if rerr, ok := r.(*CheckingError); ok {
				err = rerr
			} else {
				panic(r)
			}
		}
	}()
	tc := newTypechecker(tree.Path, true, opts.DisallowGoStmt)
	tc.Universe = universe
	if customBuiltins != nil {
		tc.Scopes = append(tc.Scopes, customBuiltins)
	}
	if main != nil {
		tc.Scopes = append(tc.Scopes, ToTypeCheckerScope(main))
	}
	if opts.IsPackage {
		pkgInfos := map[string]*PackageInfo{}
		err := checkPackage(tree, deps, imports, pkgInfos, opts.DisallowGoStmt)
		if err != nil {
			return nil, err
		}
		return pkgInfos, nil
	}
	tc.CheckNodesInNewScope(tree.Nodes)
	mainPkgInfo := &PackageInfo{}
	mainPkgInfo.IndirectVars = tc.IndirectVars
	mainPkgInfo.TypeInfo = tc.TypeInfo
	return map[string]*PackageInfo{"main": mainPkgInfo}, nil
}
