// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package composite_test

import (
	"testing"

	"github.com/goki/go-tools/go/analysis/analysistest"
	"github.com/goki/go-tools/go/analysis/passes/composite"
	"github.com/goki/go-tools/internal/typeparams"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	pkgs := []string{"a"}
	if typeparams.Enabled {
		pkgs = append(pkgs, "typeparams")
	}
	analysistest.RunWithSuggestedFixes(t, testdata, composite.Analyzer, pkgs...)
}
