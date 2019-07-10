// Copyright (c) 2019 Open2b Software Snc. All rights reserved.
// https://www.open2b.com

// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

const helpInstall = `
usage: sgc install [-v] [-work] [ package | file ]
Install compiles and installs an interpreter for Scriggo programs, scripts
and templates from a package or single file.

Executables are installed in the directory GOBIN as for the go install
command.

For more about the GOBIN directory, see 'go help install'.

If a package has a file named "Scriggofile" in its directory, an interpreter
is build and installed from the instructions in this file according to a
specific format. For example:

sgc install github.com/organization/example

will install an interpreter named "example" (or "example.exe") from the commands
in the file "github.com/organization/example/Scriggofile".

For more about the Scriggofile specific format, see 'sgc help Scriggofile'.

If a file is given, instead of a package, the file must contains the commands
to build the interpreter. The name of the executable is the same of the file
without the file extension. For example if the file is "example.Scriggofile"
the executable will be named "example" (or "example.exe").

The -v flag prints the imported packages as defined in the Scriggofile.

The -work flag prints the name of the temporary work directory of a package
used to build and install the interpreter. The directory will not be deleted.

See also: sgc embed.
`

const helpScriggofile = `
A Scriggo descriptor file consits of a valid Go package source code containing one
Scriggo file comment and one or more imports, which may in turn have a Scriggo import comment

An example Scriggo descriptor is:

	//scriggo: interpreters:"script"
	
	package x
	
	import (
		_ "fmt"
		_ "math" //scriggo: main uncapitalize
	)

This Scriggo descriptor describes a Scriggo interpreter provides package "fmt"
(available through an import statement) and package "math" as "builtin", with
all names "uncapitalized".

Each import statement should have a name _, which prevents tools like goimports from removing import.

Options available in the Scriggo file comment are:

	interpreters:targets  describe an interpreter for targets. Valid targets are "template, "script" and "program"
	interpreter           install all kinds of interpreters
	embedded              describe an embedded packages declaration
	output                select output file/directory
	goos:GOOSs            force GOOS to the specified value. More than one value can be provided

Options available as Scriggo import comments are:

	main                    import as package main. Only available in scripts an templates
	uncapitalize            declarations imported as main are "uncapitalized"
	path                    change Scrigo import path
	export:names            only export names
	noexport:names          export everything excluding names

Example import comments

Default. Makes "fmt" available in Scriggo as would be available in Go:

	import _ "fmt" //scriggo:

Import all declarations from "fmt" in package main, making them accessible
without a selector:

	import _ "fmt" //scriggo: main

Import all declarations from "fmt" in package main with uncapitalized names,
making them accessible without a selector:

	import _ "fmt" //scriggo: main uncapitalize

Import all declarations from "fmt" excluding "Print" and Println":

	import _ "fmt" //scriggo: noexport:"Print,Println"
`
