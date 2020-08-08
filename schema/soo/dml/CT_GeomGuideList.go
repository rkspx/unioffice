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

	"github.com/unidoc/unioffice"
)

type CT_GeomGuideList struct {
	Gd []*CT_GeomGuide
}

func NewCT_GeomGuideList() *CT_GeomGuideList {
	ret := &CT_GeomGuideList{}
	return ret
}

func (m *CT_GeomGuideList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Gd != nil {
		segd := xml.StartElement{Name: xml.Name{Local: "a:gd"}}
		for _, c := range m.Gd {
			e.EncodeElement(c, segd)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GeomGuideList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_GeomGuideList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "gd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "gd"}:
				tmp := NewCT_GeomGuide()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Gd = append(m.Gd, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_GeomGuideList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GeomGuideList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GeomGuideList and its children
func (m *CT_GeomGuideList) Validate() error {
	return m.ValidateWithPath("CT_GeomGuideList")
}

// ValidateWithPath validates the CT_GeomGuideList and its children, prefixing error messages with path
func (m *CT_GeomGuideList) ValidateWithPath(path string) error {
	for i, v := range m.Gd {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Gd[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
