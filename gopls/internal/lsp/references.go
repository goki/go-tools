// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lsp

import (
	"context"

	"github.com/goki/go-tools/gopls/internal/lsp/protocol"
	"github.com/goki/go-tools/gopls/internal/lsp/source"
	"github.com/goki/go-tools/gopls/internal/lsp/template"
	"github.com/goki/go-tools/internal/event"
	"github.com/goki/go-tools/internal/event/tag"
)

func (s *Server) references(ctx context.Context, params *protocol.ReferenceParams) ([]protocol.Location, error) {
	ctx, done := event.Start(ctx, "lsp.Server.references", tag.URI.Of(params.TextDocument.URI))
	defer done()

	snapshot, fh, ok, release, err := s.beginFileRequest(ctx, params.TextDocument.URI, source.UnknownKind)
	defer release()
	if !ok {
		return nil, err
	}
	if snapshot.View().FileKind(fh) == source.Tmpl {
		return template.References(ctx, snapshot, fh, params)
	}
	return source.References(ctx, snapshot, fh, params.Position, params.Context.IncludeDeclaration)
}
