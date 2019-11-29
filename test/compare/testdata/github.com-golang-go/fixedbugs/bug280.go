// skip : panic instead of returning error 'use of [...] array outside of array literal' https://github.com/open2b/scriggo/issues/467

// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// https://golang.org/issue/808

package main

type A [...]int	// ERROR "outside of array literal"

func main() { }