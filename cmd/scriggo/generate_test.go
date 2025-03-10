//go:build go1.23

package main

import (
	"bytes"
	"go/types"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"testing"

	pkgs "golang.org/x/tools/go/packages"
)

func Test_renderPackages(t *testing.T) {
	// NOTE: these tests ignores whitespaces, imports and comments.
	cases := map[string]struct {
		sf       *scriggofile
		goos     string
		expected string
	}{
		"Importing fmt with an alternative path": {
			sf: &scriggofile{
				pkgName:  "test",
				variable: "packages",
				imports: []*importCommand{
					{path: "fmt", asPath: "custom/fmt/path"},
				},
			},
			expected: `package test

			import (
				"fmt"
			)

			import "github.com/open2b/scriggo/native"
			import "reflect"

			func init() {
				packages = make(native.Packages, 1)
				var decs native.Declarations
				// "custom/fmt/path"
				decs = make(native.Declarations, 29)
				decs["Append"] = fmt.Append
				decs["Appendf"] = fmt.Appendf
				decs["Appendln"] = fmt.Appendln
				decs["Errorf"] = fmt.Errorf
				decs["FormatString"] = fmt.FormatString
				decs["Formatter"] = reflect.TypeOf((*fmt.Formatter)(nil)).Elem()
				decs["Fprint"] = fmt.Fprint
				decs["Fprintf"] = fmt.Fprintf
				decs["Fprintln"] = fmt.Fprintln
				decs["Fscan"] = fmt.Fscan
				decs["Fscanf"] = fmt.Fscanf
				decs["Fscanln"] = fmt.Fscanln
				decs["GoStringer"] = reflect.TypeOf((*fmt.GoStringer)(nil)).Elem()
				decs["Print"] = fmt.Print
				decs["Printf"] = fmt.Printf
				decs["Println"] = fmt.Println
				decs["Scan"] = fmt.Scan
				decs["ScanState"] = reflect.TypeOf((*fmt.ScanState)(nil)).Elem()
				decs["Scanf"] = fmt.Scanf
				decs["Scanln"] = fmt.Scanln
				decs["Scanner"] = reflect.TypeOf((*fmt.Scanner)(nil)).Elem()
				decs["Sprint"] = fmt.Sprint
				decs["Sprintf"] = fmt.Sprintf
				decs["Sprintln"] = fmt.Sprintln
				decs["Sscan"] = fmt.Sscan
				decs["Sscanf"] = fmt.Sscanf
				decs["Sscanln"] = fmt.Sscanln
				decs["State"] = reflect.TypeOf((*fmt.State)(nil)).Elem()
				decs["Stringer"] = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
				packages["custom/fmt/path"] = native.Package{
					Name: "fmt",
					Declarations: decs,
				}
			}`,
		},
		"Importing archive/tar simple": {
			sf: &scriggofile{
				pkgName:  "test",
				variable: "packages",
				imports:  []*importCommand{{path: "archive/tar"}},
			},
			expected: `package test

			import (
				tar "archive/tar"
			)

			import "github.com/open2b/scriggo/native"
			import "reflect"

			func init() {
				packages = make(native.Packages, 1)
				var decs native.Declarations
				// "archive/tar"
				decs = make(native.Declarations, 31)
				decs["ErrFieldTooLong"] = &tar.ErrFieldTooLong
				decs["ErrHeader"] = &tar.ErrHeader
				decs["ErrInsecurePath"] = &tar.ErrInsecurePath
				decs["ErrWriteAfterClose"] = &tar.ErrWriteAfterClose
				decs["ErrWriteTooLong"] = &tar.ErrWriteTooLong
				decs["FileInfoHeader"] = tar.FileInfoHeader
				decs["FileInfoNames"] = reflect.TypeOf((*tar.FileInfoNames)(nil)).Elem()
				decs["Format"] = reflect.TypeOf((*tar.Format)(nil)).Elem()
				decs["FormatGNU"] = tar.FormatGNU
				decs["FormatPAX"] = tar.FormatPAX
				decs["FormatUSTAR"] = tar.FormatUSTAR
				decs["FormatUnknown"] = tar.FormatUnknown
				decs["Header"] = reflect.TypeOf((*tar.Header)(nil)).Elem()
				decs["NewReader"] = tar.NewReader
				decs["NewWriter"] = tar.NewWriter
				decs["Reader"] = reflect.TypeOf((*tar.Reader)(nil)).Elem()
				decs["TypeBlock"] = native.UntypedNumericConst("52")
				decs["TypeChar"] = native.UntypedNumericConst("51")
				decs["TypeCont"] = native.UntypedNumericConst("55")
				decs["TypeDir"] = native.UntypedNumericConst("53")
				decs["TypeFifo"] = native.UntypedNumericConst("54")
				decs["TypeGNULongLink"] = native.UntypedNumericConst("75")
				decs["TypeGNULongName"] = native.UntypedNumericConst("76")
				decs["TypeGNUSparse"] = native.UntypedNumericConst("83")
				decs["TypeLink"] = native.UntypedNumericConst("49")
				decs["TypeReg"] = native.UntypedNumericConst("48")
				decs["TypeRegA"] = native.UntypedNumericConst("0")
				decs["TypeSymlink"] = native.UntypedNumericConst("50")
				decs["TypeXGlobalHeader"] = native.UntypedNumericConst("103")
				decs["TypeXHeader"] = native.UntypedNumericConst("120")
				decs["Writer"] = reflect.TypeOf((*tar.Writer)(nil)).Elem()
				packages["archive/tar"] = native.Package{
					Name: "tar",
					Declarations: decs,
				}
			}`,
		},
		"Importing fmt simple": {
			sf: &scriggofile{
				pkgName:  "test",
				variable: "packages",
				imports:  []*importCommand{{path: "fmt"}},
			},
			expected: `package test

			import (
				"fmt"
			)

			import "github.com/open2b/scriggo/native"
			import "reflect"

			func init() {
				packages = make(native.Packages, 1)
				var decs native.Declarations
				// "fmt"
				decs = make(native.Declarations, 29)
				decs["Append"] = fmt.Append
				decs["Appendf"] = fmt.Appendf
				decs["Appendln"] = fmt.Appendln
				decs["Errorf"] = fmt.Errorf
				decs["FormatString"] = fmt.FormatString
				decs["Formatter"] = reflect.TypeOf((*fmt.Formatter)(nil)).Elem()
				decs["Fprint"] = fmt.Fprint
				decs["Fprintf"] = fmt.Fprintf
				decs["Fprintln"] = fmt.Fprintln
				decs["Fscan"] = fmt.Fscan
				decs["Fscanf"] = fmt.Fscanf
				decs["Fscanln"] = fmt.Fscanln
				decs["GoStringer"] = reflect.TypeOf((*fmt.GoStringer)(nil)).Elem()
				decs["Print"] = fmt.Print
				decs["Printf"] = fmt.Printf
				decs["Println"] = fmt.Println
				decs["Scan"] = fmt.Scan
				decs["ScanState"] = reflect.TypeOf((*fmt.ScanState)(nil)).Elem()
				decs["Scanf"] = fmt.Scanf
				decs["Scanln"] = fmt.Scanln
				decs["Scanner"] = reflect.TypeOf((*fmt.Scanner)(nil)).Elem()
				decs["Sprint"] = fmt.Sprint
				decs["Sprintf"] = fmt.Sprintf
				decs["Sprintln"] = fmt.Sprintln
				decs["Sscan"] = fmt.Sscan
				decs["Sscanf"] = fmt.Sscanf
				decs["Sscanln"] = fmt.Sscanln
				decs["State"] = reflect.TypeOf((*fmt.State)(nil)).Elem()
				decs["Stringer"] = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
				packages["fmt"] = native.Package{
					Name: "fmt",
					Declarations: decs,
				}
			}`,
		},
		"Importing only Println from fmt": {
			sf: &scriggofile{
				pkgName:  "test",
				variable: "packages",
				imports: []*importCommand{
					{
						path:      "fmt",
						including: []string{"Println"},
					},
				},
			},
			expected: `package test

			import (
				"fmt"
			)

			import "github.com/open2b/scriggo/native"

			func init() {
				packages = make(native.Packages, 1)
				var decs native.Declarations
				// "fmt"
				decs = make(native.Declarations, 1)
				decs["Println"] = fmt.Println
				packages["fmt"] = native.Package{
					Name:      "fmt",
					Declarations: decs,
				}
			}`,
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if c.goos == "" {
				c.goos = os.Getenv("GOOS")
				if c.goos == "" {
					c.goos = runtime.GOOS
				}
			}
			b := bytes.Buffer{}
			err := renderPackages(&b, "", c.sf, c.goos, buildFlags{})
			if err != nil {
				t.Fatal(err, c.sf)
			}
			got := _cleanOutput(b.String())
			c.expected = _cleanOutput(c.expected)
			if got != c.expected {
				if testing.Verbose() {
					t.Fatalf("expecting:\n\n%s\n\ngot:\n\n%s", c.expected, got)
				}
				t.Fatalf("expecting %q, got %q", c.expected, got)
			}
		})
	}
}

func _cleanOutput(s string) string {
	re := regexp.MustCompile(`(?s)import \(.*?\)`)
	s = re.ReplaceAllString(s, "")
	lines := []string{}
	for _, l := range strings.Split(s, "\n") {
		l := strings.TrimSpace(l)
		if l != "" && !strings.HasPrefix(l, "//") && !strings.HasPrefix(l, "import ") {
			l := strings.Join(strings.Fields(l), " ")
			lines = append(lines, l)
		}
	}
	return strings.Join(lines, "\n")
}

func Test_parseGoPackage(t *testing.T) {
	cases := map[string]struct {
		name  string            // package name.
		decls map[string]string // package declarations.
	}{
		"fmt": {
			name: "fmt",
			decls: map[string]string{
				"Append":       "fmt.Append",
				"Appendf":      "fmt.Appendf",
				"Appendln":     "fmt.Appendln",
				"Errorf":       "fmt.Errorf",
				"FormatString": "fmt.FormatString",
				"Formatter":    "reflect.TypeOf((*fmt.Formatter)(nil)).Elem()",
				"Fprint":       "fmt.Fprint",
				"Fprintf":      "fmt.Fprintf",
				"Fprintln":     "fmt.Fprintln",
				"Fscan":        "fmt.Fscan",
				"Fscanf":       "fmt.Fscanf",
				"Fscanln":      "fmt.Fscanln",
				"GoStringer":   "reflect.TypeOf((*fmt.GoStringer)(nil)).Elem()",
				"Print":        "fmt.Print",
				"Printf":       "fmt.Printf",
				"Println":      "fmt.Println",
				"Scan":         "fmt.Scan",
				"ScanState":    "reflect.TypeOf((*fmt.ScanState)(nil)).Elem()",
				"Scanf":        "fmt.Scanf",
				"Scanln":       "fmt.Scanln",
				"Scanner":      "reflect.TypeOf((*fmt.Scanner)(nil)).Elem()",
				"Sprint":       "fmt.Sprint",
				"Sprintf":      "fmt.Sprintf",
				"Sprintln":     "fmt.Sprintln",
				"Sscan":        "fmt.Sscan",
				"Sscanf":       "fmt.Sscanf",
				"Sscanln":      "fmt.Sscanln",
				"State":        "reflect.TypeOf((*fmt.State)(nil)).Elem()",
				"Stringer":     "reflect.TypeOf((*fmt.Stringer)(nil)).Elem()",
			},
		},
		"archive/tar": {
			name: "tar",
			decls: map[string]string{
				"ErrFieldTooLong":    "&tar.ErrFieldTooLong",
				"ErrHeader":          "&tar.ErrHeader",
				"ErrInsecurePath":    "&tar.ErrInsecurePath",
				"ErrWriteAfterClose": "&tar.ErrWriteAfterClose",
				"ErrWriteTooLong":    "&tar.ErrWriteTooLong",
				"FileInfoHeader":     "tar.FileInfoHeader",
				"FileInfoNames":      "reflect.TypeOf((*tar.FileInfoNames)(nil)).Elem()",
				"Format":             "reflect.TypeOf((*tar.Format)(nil)).Elem()",
				"FormatGNU":          "tar.FormatGNU",
				"FormatPAX":          "tar.FormatPAX",
				"FormatUSTAR":        "tar.FormatUSTAR",
				"FormatUnknown":      "tar.FormatUnknown",
				"Header":             "reflect.TypeOf((*tar.Header)(nil)).Elem()",
				"NewReader":          "tar.NewReader",
				"NewWriter":          "tar.NewWriter",
				"Reader":             "reflect.TypeOf((*tar.Reader)(nil)).Elem()",
				"TypeBlock":          "native.UntypedNumericConst(\"52\")",
				"TypeChar":           "native.UntypedNumericConst(\"51\")",
				"TypeCont":           "native.UntypedNumericConst(\"55\")",
				"TypeDir":            "native.UntypedNumericConst(\"53\")",
				"TypeFifo":           "native.UntypedNumericConst(\"54\")",
				"TypeGNULongLink":    "native.UntypedNumericConst(\"75\")",
				"TypeGNULongName":    "native.UntypedNumericConst(\"76\")",
				"TypeGNUSparse":      "native.UntypedNumericConst(\"83\")",
				"TypeLink":           "native.UntypedNumericConst(\"49\")",
				"TypeReg":            "native.UntypedNumericConst(\"48\")",
				"TypeRegA":           "native.UntypedNumericConst(\"0\")",
				"TypeSymlink":        "native.UntypedNumericConst(\"50\")",
				"TypeXGlobalHeader":  "native.UntypedNumericConst(\"103\")",
				"TypeXHeader":        "native.UntypedNumericConst(\"120\")",
				"Writer":             "reflect.TypeOf((*tar.Writer)(nil)).Elem()",
			},
		},
	}
	goos := "linux" // paths in this test should be OS-independent.
	for path, expected := range cases {
		t.Run(path, func(t *testing.T) {
			gotName, gotDecls, _, _, err := loadGoPackage(path, "", goos, buildFlags{}, nil, nil, newPackageNameCache())
			if err != nil {
				t.Fatal(err)
			}
			if gotName != expected.name {
				t.Fatalf("path %q: expecting name %q, got %q", path, expected.name, gotName)
			}
			if len(gotDecls) != len(expected.decls) {
				t.Fatalf("path %q: expecting %#v, got %#v", path, expected.decls, gotDecls)
			}
			if !reflect.DeepEqual(gotDecls, expected.decls) {
				t.Fatalf("path %q: expecting %#v, got %#v", path, expected.decls, gotDecls)
			}
		})
	}
}

var testpkgDecl map[string]types.Object

// init populates 'testpkgDecl'.
func init() {
	conf := &pkgs.Config{
		Mode: 1023, // this value has been taken from 'generate.go'.
	}
	packages, err := pkgs.Load(conf, "github.com/open2b/scriggo/cmd/scriggo/testpkg")
	if err != nil {
		panic(err)
	}
	testpkgDecl = map[string]types.Object{}
	for _, def := range packages[0].TypesInfo.Defs {
		if def == nil || !def.Exported() {
			continue
		}
		if def.Parent() == nil || def.Parent().Parent() != types.Universe {
			continue
		}
		testpkgDecl[def.Name()] = def
	}
}

func Test_isGenericFunction(t *testing.T) {
	tests := map[string]bool{
		// Function declarations.
		"F1": false,
		"F2": true,

		// Other declarations, which should be false.
		"EmptyInterface": false,
		"FunctionType1":  false,
		"GeneralInter1":  false,
		"GeneralInter2":  false,
		"GeneralInter3":  false,
		"GeneralInter4":  false,
		"GeneralInter5":  false,
		"GeneralInter6":  false,
		"GenericList":    false,
		"Int1":           false,
		"Int2":           false,
		"Interface1":     false,
		"Interface2":     false,
		"Interface3":     false,
		"Receiver":       false,
		"T1":             false,
		"T2":             false,
		"V1":             false,
		"V2":             false,
	}
	if len(testpkgDecl) < 5 {
		t.Fatal("it seems like 'testpkgDecl' has not been initialized. This is a bug in the tests")
	}
	for decl, typ := range testpkgDecl {
		t.Run(decl, func(t *testing.T) {
			expected, ok := tests[decl]
			if !ok {
				t.Fatalf("declaration %q, which is declared in 'testpkg', should also be added to this test", decl)
			}
			got := isGenericFunction(typ)
			if got != expected {
				t.Fatalf("%s: expected isGenericFunction = %t, got %t", decl, expected, got)
			}
		})
		delete(tests, decl)
	}
	if len(tests) > 0 {
		t.Fatalf("there are declarations in 'tests' that do not correspond to any declaration within 'testpkg': %v", tests)
	}
}

func Test_isGeneralInterface(t *testing.T) {
	tests := map[string]bool{
		// Interface declarations.
		"EmptyInterface": false,
		"GeneralInter1":  true,
		"GeneralInter2":  true,
		"GeneralInter3":  true,
		"GeneralInter4":  true,
		"GeneralInter5":  true,
		"GeneralInter6":  true,
		"Interface1":     false,
		"Interface2":     true,
		"Interface3":     false,

		// Other declarations, which should be false.
		"F1":            false,
		"F2":            false,
		"FunctionType1": false,
		"GenericList":   false,
		"Int1":          false,
		"Int2":          false,
		"Receiver":      false,
		"T1":            false,
		"T2":            false,
		"V1":            false,
		"V2":            false,
	}
	if len(testpkgDecl) < 5 {
		t.Fatal("it seems like 'testpkgDecl' has not been initialized. This is a bug in the tests")
	}
	for decl, typ := range testpkgDecl {
		t.Run(decl, func(t *testing.T) {
			expected, ok := tests[decl]
			if !ok {
				t.Fatalf("declaration %q, which is declared in 'testpkg', should also be added to this test", decl)
			}
			got := isGeneralInterface(typ)
			if got != expected {
				t.Fatalf("%s: expected isGeneralInterface = %t, got %t", decl, expected, got)
			}
		})
		delete(tests, decl)
	}
	if len(tests) > 0 {
		t.Fatalf("there are declarations in 'tests' that do not correspond to any declaration within 'testpkg': %v", tests)
	}
}

func Test_isGenericType(t *testing.T) {
	tests := map[string]bool{
		// Type declarations.
		"FunctionType1": false,
		"GenericList":   true,
		"Int1":          false,
		"Int2":          false,
		"Receiver":      false,
		"T1":            false,
		"T2":            true,

		// Other declarations, which should be false.
		"EmptyInterface": false,
		"F1":             false,
		"F2":             false,
		"GeneralInter1":  false,
		"GeneralInter2":  false,
		"GeneralInter3":  false,
		"GeneralInter4":  false,
		"GeneralInter5":  false,
		"GeneralInter6":  false,
		"Interface1":     false,
		"Interface2":     false,
		"Interface3":     false,
		"V1":             false,
		"V2":             false,
	}
	if len(testpkgDecl) < 5 {
		t.Fatal("it seems like 'testpkgDecl' has not been initialized. This is a bug in the tests")
	}
	for decl, typ := range testpkgDecl {
		t.Run(decl, func(t *testing.T) {
			expected, ok := tests[decl]
			if !ok {
				t.Fatalf("declaration %q, which is declared in 'testpkg', should also be added to this test", decl)
			}
			got := isGenericType(typ)
			if got != expected {
				t.Fatalf("%s: expected isGenericType = %t, got %t", decl, expected, got)
			}
		})
		delete(tests, decl)
	}
	if len(tests) > 0 {
		t.Fatalf("there are declarations in 'tests' that do not correspond to any declaration within 'testpkg': %v", tests)
	}
}
