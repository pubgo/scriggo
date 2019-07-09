// Copyright (c) 2019 Open2b Software Snc. All rights reserved.
// https://www.open2b.com

// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"fmt"
	"go/constant"
	"go/types"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	pkgs "golang.org/x/tools/go/packages"
)

const (
	dirPerm  = 0775 // default new directory permission.
	filePerm = 0644 // default new file permission.
)

// renderPackages renders a Scriggofile. It also returns a boolean indicating
// if the content contains packages. Ignores all main packages contained in
// the Scriggofile.
func renderPackages(sf *scriggofile, goos string) (string, bool, error) {

	type packageType struct {
		name string
		decl map[string]string
	}

	importReflect := false

	explicitImports := strings.Builder{}
	for _, imp := range sf.imports {
		uniqueName := uniquePackageName(imp.path)
		if uniqueName != imp.path {
			explicitImports.WriteString(uniqueName + ` "` + imp.path + `"` + "\n")
		} else {
			explicitImports.WriteString(`"` + imp.path + `"` + "\n")
		}
		if imp.path == "reflect" {
			importReflect = true
		}
	}

	packages := map[string]*packageType{}
	for _, imp := range sf.imports {
		pkgName, decls, err := loadGoPackage(imp.path, goos)
		if err != nil {
			panic(err) // TODO(Gianluca).
		}
		// No declarations at path: move on to next import path.
		if len(decls) == 0 {
			continue
		}
		if len(imp.including) > 0 {
			decls, err = filterIncluding(decls, imp.including)
			if err != nil {
				return "", false, err
			}
		} else if len(imp.excluding) > 0 {
			decls, err = filterExcluding(decls, imp.excluding)
			if err != nil {
				return "", false, err
			}
		}
		// Convert all declaration name to "unexported" if requested.
		// TODO(Gianluca): if uncapitalized name conflicts with a Go builtin
		//  return an error. Note that the builtin functions 'print' and
		//  'println' should be handled as special case:
		if imp.notCapitalized {
			tmp := map[string]string{}
			for name, decl := range decls {
				newName := uncapitalize(name)
				if newName == "main" || newName == "init" {
					return "", false, fmt.Errorf("%q is not a valid identifier: remove 'uncapitalized' or change declaration name in package %q", newName, imp.path)
				}
				if isGoKeyword(newName) {
					return "", false, fmt.Errorf("%q is not a valid identifier as it conflicts with Go keyword %q: remove 'uncapitalized' or change declaration name in package %q", newName, newName, imp.path)
				}
				if isPredeclaredIdentifier(newName) {
					if newName == "print" || newName == "println" {
						// TODO(Gianluca): If the signature is the same of Go
						//  'print' and 'println', shadowing is allowed. Else, an
						//  error must be returned.
					} else {
						return "", false, fmt.Errorf("%q is not a valid identifier as it conflicts with Go predeclared identifier %q: remove 'uncapitalized' or change declaration name in package %q", newName, newName, imp.path)
					}
				}
				tmp[newName] = decl
			}
			decls = tmp
		}

		// Determine which import path use: the default (the one specified in
		// source, used by Go) or a new one indicated in Scriggo comments.
		path := imp.path
		if imp.asPath != "" {
			// TODO(Gianluca): if asPath points to standard lib, return error.
			//  Else, package mixing is allowed.
			path = imp.asPath
		}

		switch imp.asPath {
		case "main": // Add read declarations to package main as builtins.
			if packages["main"] == nil {
				packages["main"] = &packageType{"main", map[string]string{}}
			}
			for name, decl := range decls {
				if _, ok := packages["main"].decl[name]; ok {
					return "", false, fmt.Errorf("declaration name collision in package main: %q", name)
				}
				packages["main"].decl[name] = decl
			}
		default: // Add read declarations to the specified package.
			if packages[path] == nil {
				packages[path] = &packageType{
					decl: map[string]string{},
				}
			}
			for name, decl := range decls {
				if _, ok := packages[path].decl[name]; ok {
					return "", false, fmt.Errorf("declaration name collision: %q in package %q", name, imp.path)
				}
				packages[path].decl[name] = decl
				packages[path].name = pkgName
				name = filepath.Base(path)
			}
		}

	}

	// If no packages have been declared, just return.
	if len(packages) == 0 {
		return "", false, nil
	}

	// Render package content.
	allPkgsContent := strings.Builder{}
	paths := make([]string, 0, len(packages))
	hasMain := false
	for path := range packages {
		if path == "main" {
			hasMain = true
			continue
		}
		paths = append(paths, path)
	}
	sort.Strings(paths)
	if hasMain {
		paths = append([]string{"main"}, paths...)
	}
	for _, path := range paths {
		pkg := packages[path]
		declarations := strings.Builder{}
		names := make([]string, 0, len(pkg.decl))
		for name := range pkg.decl {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			decl := pkg.decl[name]
			declarations.WriteString(strconv.Quote(name) + ": " + decl + ",\n")
		}
		out := `[path]: {
			Name: [pkgName],
			Declarations: map[string]interface{}{
				[declarations]
			},
		},
		`
		out = strings.Replace(out, "[path]", strconv.Quote(path), 1)
		out = strings.Replace(out, "[pkgName]", strconv.Quote(pkg.name), 1)
		out = strings.Replace(out, "[declarations]", declarations.String(), 1)
		allPkgsContent.WriteString(out)
	}

	// TODO(Gianluca): package main must be first.

	// Skeleton for a package group.
	const pkgsSkeleton = `package [pkgName]

		import (
			[explicitImports]
		)

		import . "scriggo"
		[reflectImport]

		func init() {
			[variable] = Packages{
				[pkgContent]
			}
		}`

	var reflectImport string
	if !importReflect {
		reflectImport = `import "reflect"`
	}

	pkgOutput := strings.NewReplacer(
		"[pkgName]", sf.pkgName,
		"[explicitImports]", explicitImports.String(),
		"[reflectImport]", reflectImport,
		"[variable]", sf.variable,
		"[pkgContent]", allPkgsContent.String(),
	).Replace(pkgsSkeleton)
	pkgOutput = genHeader(sf, goos) + pkgOutput

	return pkgOutput, true, nil
}

// loadGoPackage loads the Go package with the given path and returns its name
// and its exported declarations.
func loadGoPackage(path, goos string) (string, map[string]string, error) {

	fmt.Printf("generating package %q (GOOS=%q)...", path, goos)

	declarations := map[string]string{}
	// TODO(marco): remove the global cache of package names.
	pkgBase := uniquePackageName(path)

	conf := &pkgs.Config{
		Mode: 1023,
	}
	if goos != "" {
		// TODO(Gianluca): enable support for different GOOSes.
		//  conf.Env = append(os.Environ(), "GOOS=", goos)
	}

	packages, err := pkgs.Load(conf, path)
	if err != nil {
		return "", nil, err
	}

	if pkgs.PrintErrors(packages) > 0 {
		return "", nil, errors.New("error")
	}

	if len(packages) > 1 {
		return "", nil, errors.New("package query returned more than one package")
	}

	if len(packages) != 1 {
		panic("bug")
	}

	name := packages[0].Name

	for _, v := range packages[0].TypesInfo.Defs {
		// Include only exported names. Do not take into account whether the
		// object is in a local (function) scope or not.
		if v == nil || !v.Exported() {
			continue
		}
		// Include only package-level names.
		if v.Parent() == nil || v.Parent().Parent() != types.Universe {
			continue
		}
		switch v := v.(type) {
		case *types.Const:
			switch v.Val().Kind() {
			case constant.String, constant.Bool:
				declarations[v.Name()] = fmt.Sprintf("ConstValue(%s.%s)", pkgBase, v.Name())
				continue
			case constant.Int:
				// Most cases fall here.
				if len(v.Val().ExactString()) < 7 {
					declarations[v.Name()] = fmt.Sprintf("ConstValue(%s.%s)", pkgBase, v.Name())
					continue
				}
			}
			typ := v.Type().String()
			if strings.HasPrefix(typ, "untyped ") {
				typ = "nil"
			} else {
				if strings.Contains(typ, ".") {
					typ = pkgBase + filepath.Ext(typ)
				}
				typ = fmt.Sprintf("reflect.TypeOf(new(%s)).Elem()", typ)
			}
			exact := v.Val().ExactString()
			quoted := strconv.Quote(exact)
			if v.Val().Kind() == constant.String && len(quoted) == len(exact)+4 {
				quoted = "`" + exact + "`"
			}
			declarations[v.Name()] = fmt.Sprintf("ConstLiteral(%v, %s)", typ, quoted)
		case *types.Func:
			if v.Type().(*types.Signature).Recv() == nil {
				declarations[v.Name()] = fmt.Sprintf("%s.%s", pkgBase, v.Name())
			}
		case *types.Var:
			if !v.Embedded() && !v.IsField() {
				declarations[v.Name()] = fmt.Sprintf("&%s.%s", pkgBase, v.Name())
			}
		case *types.TypeName:
			if parts := strings.Split(v.String(), " "); len(parts) >= 3 {
				if strings.HasPrefix(parts[2], "struct{") {
					declarations[v.Name()] = fmt.Sprintf("reflect.TypeOf(%s.%s{})", pkgBase, v.Name())
					continue
				}
			}
			declarations[v.Name()] = fmt.Sprintf("reflect.TypeOf(new(%s.%s)).Elem()", pkgBase, v.Name())
		}
	}

	fmt.Printf("done!\n")

	return name, declarations, nil
}