// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !go1.18
// +build !go1.18

package infertypeargs

import (
	"go/token"
	"go/types"

	"github.com/goki/go-tools/go/analysis"
	"github.com/goki/go-tools/go/ast/inspector"
)

// DiagnoseInferableTypeArgs returns an empty slice, as generics are not supported at
// this go version.
func DiagnoseInferableTypeArgs(fset *token.FileSet, inspect *inspector.Inspector, start, end token.Pos, pkg *types.Package, info *types.Info) []analysis.Diagnostic {
	return nil
}
