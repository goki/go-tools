// Copyright 2023 The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imports

import (
	"go/ast"
	"go/token"
	"reflect"
	"strings"
)

// This file modifies the fixImports function to also
// add a comment to all struct fields with a 'desc' tag.

func init() {
	fixImports = func(fset *token.FileSet, f *ast.File, filename string, env *ProcessEnv) error {
		err := fixImportsDefault(fset, f, filename, env)
		if err != nil {
			return err
		}
		return addDescComments(fset, f, filename, env)
	}
}

// addDescComments adds a comment to all struct fields with a 'desc' tag.
func addDescComments(fset *token.FileSet, f *ast.File, filename string, env *ProcessEnv) error {
	cm := ast.NewCommentMap(fset, f, f.Comments)
	ast.Inspect(f, func(n ast.Node) bool {
		if st, ok := n.(*ast.StructType); ok {
			for _, field := range st.Fields.List {
				if field.Tag == nil {
					continue
				}
				// need to get rid of backquotes around tag value
				tv := strings.TrimPrefix(field.Tag.Value, "`")
				tv = strings.TrimSuffix(tv, "`")
				rst := reflect.StructTag(tv)
				desc, ok := rst.Lookup("desc")
				if ok {
					field.Doc = &ast.CommentGroup{
						List: []*ast.Comment{
							{
								Slash: field.Pos() - 1,
								Text:  "// " + desc,
							},
						},
					}
					cm[field] = []*ast.CommentGroup{field.Doc}
				}
			}
		}
		return true
	})
	f.Comments = cm.Comments()
	return nil
}
