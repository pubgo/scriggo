// Copyright (c) 2019 Open2b Software Snc. All rights reserved.
// https://www.open2b.com

// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"scriggo"
	"scriggo/runtime"
	"scriggo/template"
)

const usage = "usage: %s [-S] [-mem 250K] [-time 50ms] filename\n"

var packages scriggo.Packages
var Main *scriggo.Package

type mainLoader []byte

func (b mainLoader) Load(path string) (interface{}, error) {
	if path == "main" {
		return bytes.NewReader(b), nil
	}
	return nil, nil
}

func renderPanics(p *runtime.Panic) string {
	var msg string
	for ; p != nil; p = p.Next() {
		msg = "\n" + msg
		if p.Recovered() {
			msg = " [recovered]" + msg
		}
		msg = p.String() + msg
		if p.Next() != nil {
			msg = "\tpanic: " + msg
		}
	}
	return msg
}

func run() {

	var asm = flag.Bool("S", false, "print assembly listing")
	var timeout = flag.String("time", "", "limit the execution time; zero is no limit")
	var mem = flag.String("mem", "", "limit the allocable memory; zero is no limit")

	flag.Parse()

	var loadOptions = &scriggo.LoadOptions{}
	var runOptions = &scriggo.RunOptions{}

	if *timeout != "" {
		d, err := time.ParseDuration(*timeout)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, usage, os.Args[0])
			flag.PrintDefaults()
			os.Exit(-1)
		}
		if d != 0 {
			var cancel context.CancelFunc
			runOptions.Context, cancel = context.WithTimeout(context.Background(), d)
			defer cancel()
		}
	}

	if *mem != "" {
		loadOptions.LimitMemorySize = true
		var unit = (*mem)[len(*mem)-1]
		if unit > 'Z' {
			unit -= 'z' - 'Z'
		}
		switch unit {
		case 'B', 'K', 'M', 'G':
			*mem = (*mem)[:len(*mem)-1]
		}
		var err error
		runOptions.MaxMemorySize, err = strconv.Atoi(*mem)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, usage, os.Args[0])
			flag.PrintDefaults()
			os.Exit(-1)
		}
		switch unit {
		case 'K':
			runOptions.MaxMemorySize *= 1024
		case 'M':
			runOptions.MaxMemorySize *= 1024 * 1024
		case 'G':
			runOptions.MaxMemorySize *= 1024 * 1024 * 1024
		}
	}

	var args = flag.Args()

	if len(args) != 1 {
		_, _ = fmt.Fprintf(os.Stderr, usage, os.Args[0])
		flag.PrintDefaults()
		os.Exit(-1)
	}

	file := args[0]
	ext := filepath.Ext(file)
	if ext != ".go" && ext != ".sgo" && ext != ".html" {
		fmt.Printf("%s: extension must be \".go\" for main packages, \".sgo\" for scripts and \".html\" for template pages\n", file)
		os.Exit(-1)
	}

	absFile, err := filepath.Abs(file)
	if err != nil {
		fmt.Printf("%s: %s\n", file, err)
		os.Exit(-1)
	}

	switch ext {
	case ".sgo":
		r, err := os.Open(absFile)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "scriggo: %s\n", err)
			os.Exit(2)
		}
		loadOptions.AllowShebangLine = true
		script, err := scriggo.LoadScript(r, packages, loadOptions)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "scriggo: %s\n", err)
			os.Exit(2)
		}
		_ = r.Close()
		if *asm {
			_, err := script.Disassemble(os.Stdout)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "scriggo: %s\n", err)
				os.Exit(2)
			}
		} else {
			err = script.Run(nil, runOptions)
			if err != nil {
				if p, ok := err.(*runtime.Panic); ok {
					panic(renderPanics(p))
				}
				_, _ = fmt.Fprintf(os.Stderr, "scriggo: %s\n", err)
				os.Exit(2)
			}
		}
		os.Exit(0)
	case ".go":
		main, err := ioutil.ReadFile(absFile)
		if err != nil {
			panic(err)
		}
		program, err := scriggo.Load(scriggo.Loaders(mainLoader(main), packages), loadOptions)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "scriggo: %s\n", err)
			os.Exit(2)
		}
		if *asm {
			_, err := program.Disassemble(os.Stdout, "main")
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "scriggo: %s\n", err)
				os.Exit(2)
			}
		} else {
			err = program.Run(runOptions)
			if err != nil {
				if p, ok := err.(*runtime.Panic); ok {
					panic(renderPanics(p))
				}
				if err == context.DeadlineExceeded {
					err = errors.New("process took too long")
				}
				_, _ = fmt.Fprintf(os.Stderr, "scriggo: %s\n", err)
				os.Exit(2)
			}
		}
		os.Exit(0)
	case ".html":
		r := template.DirReader(filepath.Dir(absFile))
		path := "/" + filepath.Base(absFile)
		builtins := template.Builtins()
		loadedMain, err := packages.Load("main")
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		var main scriggo.Package
		if loadedMain == nil {
			main = builtins
		} else {
			mp := &scriggo.MapPackage{PkgName: "main", Declarations: map[string]interface{}{}}
			for _, name := range builtins.DeclarationNames() {
				mp.Declarations[name] = builtins.Lookup(name)
			}
			for _, name := range loadedMain.(scriggo.Package).DeclarationNames() {
				mp.Declarations[name] = builtins.Lookup(name)
			}
			main = mp
		}
		t, err := template.Load(path, r, main, template.ContextHTML, &template.LoadOptions{LimitMemorySize: loadOptions.LimitMemorySize})
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		if *asm {
			_, err := t.Disassemble(os.Stdout)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "scriggo: %s\n", err)
				os.Exit(2)
			}
		} else {
			options := &template.RenderOptions{
				Context:       runOptions.Context,
				MaxMemorySize: runOptions.MaxMemorySize,
			}
			err = t.Render(os.Stdout, nil, options)
			if err != nil {
				if p, ok := err.(*runtime.Panic); ok {
					panic(renderPanics(p))
				}
				fmt.Println(err)
				os.Exit(-1)
			}
		}
		os.Exit(0)
	}
}
	