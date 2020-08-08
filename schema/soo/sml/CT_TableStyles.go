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

type CT_TableStyles struct {
	// Table Style Count
	CountAttr *uint32
	// Default Table Style
	DefaultTableStyleAttr *string
	// Default Pivot Style
	DefaultPivotStyleAttr *string
	// Table Style
	TableStyle []*CT_TableStyle
}

func NewCT_TableStyles() *CT_TableStyles {
	ret := &CT_TableStyles{}
	return ret
}

func (m *CT_TableStyles) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	if m.DefaultTableStyleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "defaultTableStyle"},
			Value: fmt.Sprintf("%v", *m.DefaultTableStyleAttr)})
	}
	if m.DefaultPivotStyleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "defaultPivotStyle"},
			Value: fmt.Sprintf("%v", *m.DefaultPivotStyleAttr)})
	}
	e.EncodeToken(start)
	if m.TableStyle != nil {
		setableStyle := xml.StartElement{Name: xml.Name{Local: "ma:tableStyle"}}
		for _, c := range m.TableStyle {
			e.EncodeElement(c, setableStyle)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TableStyles) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		if attr.Name.Local == "defaultTableStyle" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DefaultTableStyleAttr = &parsed
			continue
		}
		if attr.Name.Local == "defaultPivotStyle" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DefaultPivotStyleAttr = &parsed
			continue
		}
	}
lCT_TableStyles:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "tableStyle"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "tableStyle"}:
				tmp := NewCT_TableStyle()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.TableStyle = append(m.TableStyle, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_TableStyles %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TableStyles
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TableStyles and its children
func (m *CT_TableStyles) Validate() error {
	return m.ValidateWithPath("CT_TableStyles")
}

// ValidateWithPath validates the CT_TableStyles and its children, prefixing error messages with path
func (m *CT_TableStyles) ValidateWithPath(path string) error {
	for i, v := range m.TableStyle {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/TableStyle[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
