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
	"strconv"
)

type CT_SheetCalcPr struct {
	// Full Calculation On Load
	FullCalcOnLoadAttr *bool
}

func NewCT_SheetCalcPr() *CT_SheetCalcPr {
	ret := &CT_SheetCalcPr{}
	return ret
}

func (m *CT_SheetCalcPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.FullCalcOnLoadAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fullCalcOnLoad"},
			Value: fmt.Sprintf("%d", b2i(*m.FullCalcOnLoadAttr))})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SheetCalcPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "fullCalcOnLoad" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FullCalcOnLoadAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_SheetCalcPr: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_SheetCalcPr and its children
func (m *CT_SheetCalcPr) Validate() error {
	return m.ValidateWithPath("CT_SheetCalcPr")
}

// ValidateWithPath validates the CT_SheetCalcPr and its children, prefixing error messages with path
func (m *CT_SheetCalcPr) ValidateWithPath(path string) error {
	return nil
}
