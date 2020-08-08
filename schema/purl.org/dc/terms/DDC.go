// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// commercial license can be purchased via https://unidoc.io website.

package terms

import (
	"encoding/xml"
	"fmt"
)

type DDC struct {
}

func NewDDC() *DDC {
	ret := &DDC{}
	return ret
}

func (m *DDC) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "DDC"
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *DDC) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing DDC: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the DDC and its children
func (m *DDC) Validate() error {
	return m.ValidateWithPath("DDC")
}

// ValidateWithPath validates the DDC and its children, prefixing error messages with path
func (m *DDC) ValidateWithPath(path string) error {
	return nil
}
