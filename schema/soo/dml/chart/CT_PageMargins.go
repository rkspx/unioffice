// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// commercial license can be purchased via https://unidoc.io website.

package chart

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_PageMargins struct {
	LAttr      float64
	RAttr      float64
	TAttr      float64
	BAttr      float64
	HeaderAttr float64
	FooterAttr float64
}

func NewCT_PageMargins() *CT_PageMargins {
	ret := &CT_PageMargins{}
	return ret
}

func (m *CT_PageMargins) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "l"},
		Value: fmt.Sprintf("%v", m.LAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r"},
		Value: fmt.Sprintf("%v", m.RAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "t"},
		Value: fmt.Sprintf("%v", m.TAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "b"},
		Value: fmt.Sprintf("%v", m.BAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "header"},
		Value: fmt.Sprintf("%v", m.HeaderAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "footer"},
		Value: fmt.Sprintf("%v", m.FooterAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PageMargins) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "l" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.LAttr = parsed
			continue
		}
		if attr.Name.Local == "r" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.RAttr = parsed
			continue
		}
		if attr.Name.Local == "t" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.TAttr = parsed
			continue
		}
		if attr.Name.Local == "b" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.BAttr = parsed
			continue
		}
		if attr.Name.Local == "header" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.HeaderAttr = parsed
			continue
		}
		if attr.Name.Local == "footer" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.FooterAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_PageMargins: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_PageMargins and its children
func (m *CT_PageMargins) Validate() error {
	return m.ValidateWithPath("CT_PageMargins")
}

// ValidateWithPath validates the CT_PageMargins and its children, prefixing error messages with path
func (m *CT_PageMargins) ValidateWithPath(path string) error {
	return nil
}
