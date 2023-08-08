// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unitchecker_test

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"

	"github.com/goki/go-tools/go/analysis/passes/asmdecl"
	"github.com/goki/go-tools/go/analysis/passes/assign"
	"github.com/goki/go-tools/go/analysis/passes/atomic"
	"github.com/goki/go-tools/go/analysis/passes/bools"
	"github.com/goki/go-tools/go/analysis/passes/buildtag"
	"github.com/goki/go-tools/go/analysis/passes/cgocall"
	"github.com/goki/go-tools/go/analysis/passes/composite"
	"github.com/goki/go-tools/go/analysis/passes/copylock"
	"github.com/goki/go-tools/go/analysis/passes/defers"
	"github.com/goki/go-tools/go/analysis/passes/directive"
	"github.com/goki/go-tools/go/analysis/passes/errorsas"
	"github.com/goki/go-tools/go/analysis/passes/framepointer"
	"github.com/goki/go-tools/go/analysis/passes/httpresponse"
	"github.com/goki/go-tools/go/analysis/passes/ifaceassert"
	"github.com/goki/go-tools/go/analysis/passes/loopclosure"
	"github.com/goki/go-tools/go/analysis/passes/lostcancel"
	"github.com/goki/go-tools/go/analysis/passes/nilfunc"
	"github.com/goki/go-tools/go/analysis/passes/printf"
	"github.com/goki/go-tools/go/analysis/passes/shift"
	"github.com/goki/go-tools/go/analysis/passes/sigchanyzer"
	"github.com/goki/go-tools/go/analysis/passes/stdmethods"
	"github.com/goki/go-tools/go/analysis/passes/stringintconv"
	"github.com/goki/go-tools/go/analysis/passes/structtag"
	"github.com/goki/go-tools/go/analysis/passes/testinggoroutine"
	"github.com/goki/go-tools/go/analysis/passes/tests"
	"github.com/goki/go-tools/go/analysis/passes/timeformat"
	"github.com/goki/go-tools/go/analysis/passes/unmarshal"
	"github.com/goki/go-tools/go/analysis/passes/unreachable"
	"github.com/goki/go-tools/go/analysis/passes/unusedresult"
	"github.com/goki/go-tools/go/analysis/unitchecker"
)

// vet is the entrypoint of this executable when ENTRYPOINT=vet.
// Keep consistent with the actual vet in GOROOT/src/cmd/vet/main.go.
func vet() {
	unitchecker.Main(
		asmdecl.Analyzer,
		assign.Analyzer,
		atomic.Analyzer,
		bools.Analyzer,
		buildtag.Analyzer,
		cgocall.Analyzer,
		composite.Analyzer,
		copylock.Analyzer,
		defers.Analyzer,
		directive.Analyzer,
		errorsas.Analyzer,
		framepointer.Analyzer,
		httpresponse.Analyzer,
		ifaceassert.Analyzer,
		loopclosure.Analyzer,
		lostcancel.Analyzer,
		nilfunc.Analyzer,
		printf.Analyzer,
		shift.Analyzer,
		sigchanyzer.Analyzer,
		stdmethods.Analyzer,
		stringintconv.Analyzer,
		structtag.Analyzer,
		testinggoroutine.Analyzer,
		tests.Analyzer,
		timeformat.Analyzer,
		unmarshal.Analyzer,
		unreachable.Analyzer,
		// unsafeptr.Analyzer, // currently reports findings in runtime
		unusedresult.Analyzer,
	)
}

// TestVetStdlib runs the same analyzers as the actual vet over the
// standard library, using go vet and unitchecker, to ensure that
// there are no findings.
func TestVetStdlib(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in -short mode")
	}
	if version := runtime.Version(); !strings.HasPrefix(version, "devel") {
		t.Skipf("This test is only wanted on development branches where code can be easily fixed. Skipping because runtime.Version=%q.", version)
	}

	cmd := exec.Command("go", "vet", "-vettool="+os.Args[0], "std")
	cmd.Env = append(os.Environ(), "ENTRYPOINT=vet")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Errorf("go vet std failed (%v):\n%s", err, out)
	}
}
