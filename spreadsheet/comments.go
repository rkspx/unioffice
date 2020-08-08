// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io.

package spreadsheet

import (
	"github.com/unidoc/unioffice/color"
	"github.com/unidoc/unioffice/schema/soo/sml"
	"github.com/unidoc/unioffice/spreadsheet/reference"
	"github.com/unidoc/unioffice/vmldrawing"
)

// Comments is the container for comments for a single sheet.
type Comments struct {
	w *Workbook
	x *sml.Comments
}

// MakeComments constructs a new Comments wrapper.
func MakeComments(w *Workbook, x *sml.Comments) Comments {
	return Comments{w, x}
}

// X returns the inner wrapped XML type.
func (c Comments) X() *sml.Comments {
	return c.x
}

// Comments returns the list of comments for this sheet
func (c Comments) Comments() []Comment {
	ret := []Comment{}
	for _, cmt := range c.x.CommentList.Comment {
		ret = append(ret, Comment{c.w, cmt, c.x})
	}
	return ret
}

func (c Comments) getOrCreateAuthor(author string) uint32 {
	for i, knownAuthor := range c.x.Authors.Author {
		if knownAuthor == author {
			return uint32(i)
		}
	}

	// didn't find the author, so add a new one
	authIdx := uint32(len(c.x.Authors.Author))
	c.x.Authors.Author = append(c.x.Authors.Author, author)
	return authIdx
}

// AddComment adds a new comment and returns a RichText which will contain the
// styled comment text.
func (c Comments) AddComment(cellRef string, author string) RichText {

	cmt := sml.NewCT_Comment()
	c.x.CommentList.Comment = append(c.x.CommentList.Comment, cmt)
	cmt.RefAttr = cellRef
	cmt.AuthorIdAttr = c.getOrCreateAuthor(author)
	cmt.Text = sml.NewCT_Rst()
	return RichText{cmt.Text}
}

// AddCommentWithStyle adds a new comment styled in a default way
func (c Comments) AddCommentWithStyle(cellRef string, author string, comment string) error {
	rt := c.AddComment(cellRef, author)
	run := rt.AddRun()
	run.SetBold(true)
	run.SetSize(10)
	run.SetColor(color.Black)
	run.SetFont("Calibri")
	run.SetText(author + ":")

	run = rt.AddRun()
	run.SetSize(10)
	run.SetFont("Calibri")
	run.SetColor(color.Black)
	run.SetText("\r\n" + comment + "\r\n")

	cref, err := reference.ParseCellReference(cellRef)
	if err != nil {
		return err
	}
	c.w.vmlDrawings[0].Shape = append(c.w.vmlDrawings[0].Shape, vmldrawing.NewCommentShape(int64(cref.ColumnIdx), int64(cref.RowIdx-1)))
	return nil
}
