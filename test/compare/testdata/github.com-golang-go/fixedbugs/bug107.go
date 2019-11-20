// skip : compiler/checker: cannot reference to an imported type from a package level type declaration https://github.com/open2b/scriggo/issues/466

// errorcheck

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main
import os "os"
type _ os.FileInfo
func f() (os int) {
	 // In the next line "os" should refer to the result variable, not
	 // to the package.
	 v := os.Open("", 0, 0);	// ERROR "undefined"
	 return 0
}
func main() { _ = os.Args }
