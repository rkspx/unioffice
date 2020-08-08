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

	"github.com/unidoc/unioffice"
)

type CT_Record struct {
	// No Value
	M []*CT_Missing
	// Numeric Value
	N []*CT_Number
	// Boolean
	B []*CT_Boolean
	// Error Value
	E []*CT_Error
	// Character Value
	S []*CT_String
	// Date Time
	D []*CT_DateTime
	// Shared Items Index
	X []*CT_Index
}

func NewCT_Record() *CT_Record {
	ret := &CT_Record{}
	return ret
}

func (m *CT_Record) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.M != nil {
		sem := xml.StartElement{Name: xml.Name{Local: "ma:m"}}
		for _, c := range m.M {
			e.EncodeElement(c, sem)
		}
	}
	if m.N != nil {
		sen := xml.StartElement{Name: xml.Name{Local: "ma:n"}}
		for _, c := range m.N {
			e.EncodeElement(c, sen)
		}
	}
	if m.B != nil {
		seb := xml.StartElement{Name: xml.Name{Local: "ma:b"}}
		for _, c := range m.B {
			e.EncodeElement(c, seb)
		}
	}
	if m.E != nil {
		see := xml.StartElement{Name: xml.Name{Local: "ma:e"}}
		for _, c := range m.E {
			e.EncodeElement(c, see)
		}
	}
	if m.S != nil {
		ses := xml.StartElement{Name: xml.Name{Local: "ma:s"}}
		for _, c := range m.S {
			e.EncodeElement(c, ses)
		}
	}
	if m.D != nil {
		sed := xml.StartElement{Name: xml.Name{Local: "ma:d"}}
		for _, c := range m.D {
			e.EncodeElement(c, sed)
		}
	}
	if m.X != nil {
		sex := xml.StartElement{Name: xml.Name{Local: "ma:x"}}
		for _, c := range m.X {
			e.EncodeElement(c, sex)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Record) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Record:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "m"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "m"}:
				tmp := NewCT_Missing()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.M = append(m.M, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "n"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "n"}:
				tmp := NewCT_Number()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.N = append(m.N, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "b"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "b"}:
				tmp := NewCT_Boolean()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.B = append(m.B, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "e"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "e"}:
				tmp := NewCT_Error()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.E = append(m.E, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "s"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "s"}:
				tmp := NewCT_String()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.S = append(m.S, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "d"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "d"}:
				tmp := NewCT_DateTime()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.D = append(m.D, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "x"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "x"}:
				tmp := NewCT_Index()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.X = append(m.X, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_Record %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Record
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Record and its children
func (m *CT_Record) Validate() error {
	return m.ValidateWithPath("CT_Record")
}

// ValidateWithPath validates the CT_Record and its children, prefixing error messages with path
func (m *CT_Record) ValidateWithPath(path string) error {
	for i, v := range m.M {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/M[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.N {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/N[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.B {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/B[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.E {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/E[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.S {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/S[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.D {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/D[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.X {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/X[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
