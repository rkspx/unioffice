// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// commercial license can be purchased via https://unidoc.io website.

package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_DataBinding struct {
	// XML Namespace Prefix Mappings
	PrefixMappingsAttr *string
	// XPath
	XpathAttr string
	// Custom XML Data Storage ID
	StoreItemIDAttr string
}

func NewCT_DataBinding() *CT_DataBinding {
	ret := &CT_DataBinding{}
	return ret
}

func (m *CT_DataBinding) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.PrefixMappingsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:prefixMappings"},
			Value: fmt.Sprintf("%v", *m.PrefixMappingsAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:xpath"},
		Value: fmt.Sprintf("%v", m.XpathAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:storeItemID"},
		Value: fmt.Sprintf("%v", m.StoreItemIDAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DataBinding) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "prefixMappings" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PrefixMappingsAttr = &parsed
			continue
		}
		if attr.Name.Local == "xpath" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.XpathAttr = parsed
			continue
		}
		if attr.Name.Local == "storeItemID" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StoreItemIDAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_DataBinding: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_DataBinding and its children
func (m *CT_DataBinding) Validate() error {
	return m.ValidateWithPath("CT_DataBinding")
}

// ValidateWithPath validates the CT_DataBinding and its children, prefixing error messages with path
func (m *CT_DataBinding) ValidateWithPath(path string) error {
	return nil
}
