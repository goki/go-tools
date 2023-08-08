// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package misc

import (
	"testing"

	"github.com/goki/go-tools/gopls/internal/bug"
	"github.com/goki/go-tools/gopls/internal/hooks"
	"github.com/goki/go-tools/gopls/internal/lsp/regtest"
)

func TestMain(m *testing.M) {
	bug.PanicOnBugs = true
	regtest.Main(m, hooks.Options)
}
