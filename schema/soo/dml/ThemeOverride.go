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

	"github.com/unidoc/unioffice"
)

type ThemeOverride struct {
	CT_BaseStylesOverride
}

func NewThemeOverride() *ThemeOverride {
	ret := &ThemeOverride{}
	ret.CT_BaseStylesOverride = *NewCT_BaseStylesOverride()
	return ret
}

func (m *ThemeOverride) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "a:themeOverride"
	return m.CT_BaseStylesOverride.MarshalXML(e, start)
}

func (m *ThemeOverride) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_BaseStylesOverride = *NewCT_BaseStylesOverride()
lThemeOverride:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "clrScheme"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "clrScheme"}:
				m.ClrScheme = NewCT_ColorScheme()
				if err := d.DecodeElement(m.ClrScheme, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "fontScheme"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "fontScheme"}:
				m.FontScheme = NewCT_FontScheme()
				if err := d.DecodeElement(m.FontScheme, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "fmtScheme"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "fmtScheme"}:
				m.FmtScheme = NewCT_StyleMatrix()
				if err := d.DecodeElement(m.FmtScheme, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on ThemeOverride %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lThemeOverride
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the ThemeOverride and its children
func (m *ThemeOverride) Validate() error {
	return m.ValidateWithPath("ThemeOverride")
}

// ValidateWithPath validates the ThemeOverride and its children, prefixing error messages with path
func (m *ThemeOverride) ValidateWithPath(path string) error {
	if err := m.CT_BaseStylesOverride.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
