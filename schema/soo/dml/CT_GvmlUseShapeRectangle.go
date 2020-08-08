// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// commercial license can be purchased via https://unidoc.io website.

package dml

import (
	"encoding/xml"
	"fmt"
)

type CT_GvmlUseShapeRectangle struct {
}

func NewCT_GvmlUseShapeRectangle() *CT_GvmlUseShapeRectangle {
	ret := &CT_GvmlUseShapeRectangle{}
	return ret
}

func (m *CT_GvmlUseShapeRectangle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GvmlUseShapeRectangle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_GvmlUseShapeRectangle: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_GvmlUseShapeRectangle and its children
func (m *CT_GvmlUseShapeRectangle) Validate() error {
	return m.ValidateWithPath("CT_GvmlUseShapeRectangle")
}

// ValidateWithPath validates the CT_GvmlUseShapeRectangle and its children, prefixing error messages with path
func (m *CT_GvmlUseShapeRectangle) ValidateWithPath(path string) error {
	return nil
}
