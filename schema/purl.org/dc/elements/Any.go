// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// commercial license can be purchased via https://unidoc.io website.

package elements

import (
	"encoding/xml"
	"fmt"
)

type Any struct {
	SimpleLiteral
}

func NewAny() *Any {
	ret := &Any{}
	ret.SimpleLiteral = *NewSimpleLiteral()
	return ret
}

func (m *Any) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.SimpleLiteral.MarshalXML(e, start)
}

func (m *Any) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.SimpleLiteral = *NewSimpleLiteral()
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing Any: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the Any and its children
func (m *Any) Validate() error {
	return m.ValidateWithPath("Any")
}

// ValidateWithPath validates the Any and its children, prefixing error messages with path
func (m *Any) ValidateWithPath(path string) error {
	if err := m.SimpleLiteral.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
