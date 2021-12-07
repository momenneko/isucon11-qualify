// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

// Gen generates sais2.go by duplicating functions in sais.go
// using different input types.
// See the comment at the top of sais.go for details.
package main

import (
	"bytes"
	"log"
	"os"
	"strings"
)

func main() {
	log.SetPrefix("gen: ")
	log.SetFlags(0)

	data, err := os.ReadFile("sais.go")
	if err != nil {
		log.Fatal(err)
	}

	x := bytes.Index(data, []byte("\n\n"))
	if x < 0 {
		log.Fatal("cannot find blank line after copyright comment")
	}

	var buf bytes.Buffer
	buf.Write(data[:x])
	buf.WriteString("\n\n// Code generated by go generate; DO NOT EDIT.\n\npackage suffixarray\n")

	for {
		x := bytes.Index(data, []byte("\nfunc "))
		if x < 0 {
			break
		}
		data = data[x:]
		p := bytes.IndexByte(data, '(')
		if p < 0 {
			p = len(data)
		}
		name := string(data[len("\nfunc "):p])

		x = bytes.Index(data, []byte("\n}\n"))
		if x < 0 {
			log.Fatalf("cannot find end of func %s", name)
		}
		fn := string(data[:x+len("\n}\n")])
		data = data[x+len("\n}"):]

		if strings.HasSuffix(name, "_32") {
			buf.WriteString(fix32.Replace(fn))
		}
		if strings.HasSuffix(name, "_8_32") {
			// x_8_32 -> x_8_64 done above
			fn = fix8_32.Replace(stripByteOnly(fn))
			buf.WriteString(fn)
			buf.WriteString(fix32.Replace(fn))
		}
	}

	if err := os.WriteFile("sais2.go", buf.Bytes(), 0666); err != nil {
		log.Fatal(err)
	}
}

var fix32 = strings.NewReplacer(
	"32", "64",
	"int32", "int64",
)

var fix8_32 = strings.NewReplacer(
	"_8_32", "_32",
	"byte", "int32",
)

func stripByteOnly(s string) string {
	lines := strings.SplitAfter(s, "\n")
	w := 0
	for _, line := range lines {
		if !strings.Contains(line, "256") && !strings.Contains(line, "byte-only") {
			lines[w] = line
			w++
		}
	}
	return strings.Join(lines[:w], "")
}
