// skip : panics the type checker https://github.com/open2b/scriggo/issues/528

// errorcheck

// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

type t struct {
	x int  // GCCGO_ERROR "duplicate field name .x."
	x int  // GC_ERROR "duplicate field x"
}

func f(t *t) int {
	return t.x  // GC_ERROR "ambiguous selector t.x"
}

func main() { }