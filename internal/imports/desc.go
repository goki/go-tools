// Copyright 2023 The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imports

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
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
	ast.Inspect(f, func(n ast.Node) bool {
		if st, ok := n.(*ast.StructType); ok {
			for _, field := range st.Fields.List {
				if field.Tag == nil {
					continue
				}
				rst := reflect.StructTag(field.Tag.Value)
				desc, ok := rst.Lookup("desc")
				if ok {
					fmt.Println(desc)
					field.Comment = &ast.CommentGroup{
						List: []*ast.Comment{
							{
								Slash: field.End(),
								Text:  desc,
							},
						},
					}
				}
			}
		}
		return true
	})
	return nil
}
