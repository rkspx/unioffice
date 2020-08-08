// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// commercial license can be purchased via https://unidoc.io website.

package math

import (
	"encoding/xml"

	"github.com/unidoc/unioffice"
)

type CT_SPrePr struct {
	CtrlPr *CT_CtrlPr
}

func NewCT_SPrePr() *CT_SPrePr {
	ret := &CT_SPrePr{}
	return ret
}

func (m *CT_SPrePr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.CtrlPr != nil {
		sectrlPr := xml.StartElement{Name: xml.Name{Local: "m:ctrlPr"}}
		e.EncodeElement(m.CtrlPr, sectrlPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SPrePr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SPrePr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "ctrlPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "ctrlPr"}:
				m.CtrlPr = NewCT_CtrlPr()
				if err := d.DecodeElement(m.CtrlPr, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_SPrePr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SPrePr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SPrePr and its children
func (m *CT_SPrePr) Validate() error {
	return m.ValidateWithPath("CT_SPrePr")
}

// ValidateWithPath validates the CT_SPrePr and its children, prefixing error messages with path
func (m *CT_SPrePr) ValidateWithPath(path string) error {
	if m.CtrlPr != nil {
		if err := m.CtrlPr.ValidateWithPath(path + "/CtrlPr"); err != nil {
			return err
		}
	}
	return nil
}
