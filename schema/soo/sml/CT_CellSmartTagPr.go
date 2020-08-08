// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// commercial license can be purchased via https://unidoc.io website.

package sml

import (
	"encoding/xml"
	"fmt"
)

type CT_CellSmartTagPr struct {
	// Key Name
	KeyAttr string
	// Value
	ValAttr string
}

func NewCT_CellSmartTagPr() *CT_CellSmartTagPr {
	ret := &CT_CellSmartTagPr{}
	return ret
}

func (m *CT_CellSmartTagPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "key"},
		Value: fmt.Sprintf("%v", m.KeyAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CellSmartTagPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "key" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.KeyAttr = parsed
			continue
		}
		if attr.Name.Local == "val" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ValAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_CellSmartTagPr: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_CellSmartTagPr and its children
func (m *CT_CellSmartTagPr) Validate() error {
	return m.ValidateWithPath("CT_CellSmartTagPr")
}

// ValidateWithPath validates the CT_CellSmartTagPr and its children, prefixing error messages with path
func (m *CT_CellSmartTagPr) ValidateWithPath(path string) error {
	return nil
}
