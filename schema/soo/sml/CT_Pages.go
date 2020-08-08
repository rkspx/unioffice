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

	"github.com/unidoc/unioffice"
)

type CT_Pages struct {
	// Page Item String Count
	CountAttr *uint32
	// Page Items
	Page []*CT_PCDSCPage
}

func NewCT_Pages() *CT_Pages {
	ret := &CT_Pages{}
	return ret
}

func (m *CT_Pages) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	sepage := xml.StartElement{Name: xml.Name{Local: "ma:page"}}
	for _, c := range m.Page {
		e.EncodeElement(c, sepage)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Pages) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
			continue
		}
	}
lCT_Pages:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "page"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "page"}:
				tmp := NewCT_PCDSCPage()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Page = append(m.Page, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_Pages %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Pages
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Pages and its children
func (m *CT_Pages) Validate() error {
	return m.ValidateWithPath("CT_Pages")
}

// ValidateWithPath validates the CT_Pages and its children, prefixing error messages with path
func (m *CT_Pages) ValidateWithPath(path string) error {
	for i, v := range m.Page {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Page[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
