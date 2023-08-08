// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The nilness command applies the github.com/goki/go-tools/go/analysis/passes/nilness
// analysis to the specified packages of Go source code.
package main

import (
	"github.com/goki/go-tools/go/analysis/passes/nilness"
	"github.com/goki/go-tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(nilness.Analyzer) }
