// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package pml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_TLAnimVariantBooleanVal struct {
	// Value
	ValAttr bool
}

func NewCT_TLAnimVariantBooleanVal() *CT_TLAnimVariantBooleanVal {
	ret := &CT_TLAnimVariantBooleanVal{}
	return ret
}

func (m *CT_TLAnimVariantBooleanVal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
		Value: fmt.Sprintf("%d", b2i(m.ValAttr))})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLAnimVariantBooleanVal) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			parsed, err := strconv.ParseBool(attr.Value)
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
			return fmt.Errorf("parsing CT_TLAnimVariantBooleanVal: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TLAnimVariantBooleanVal and its children
func (m *CT_TLAnimVariantBooleanVal) Validate() error {
	return m.ValidateWithPath("CT_TLAnimVariantBooleanVal")
}

// ValidateWithPath validates the CT_TLAnimVariantBooleanVal and its children, prefixing error messages with path
func (m *CT_TLAnimVariantBooleanVal) ValidateWithPath(path string) error {
	return nil
}
