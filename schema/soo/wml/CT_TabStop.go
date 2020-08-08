// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package wml

import (
	"encoding/xml"
	"fmt"
)

type CT_TabStop struct {
	// Tab Stop Type
	ValAttr ST_TabJc
	// Tab Leader Character
	LeaderAttr ST_TabTlc
	// Tab Stop Position
	PosAttr ST_SignedTwipsMeasure
}

func NewCT_TabStop() *CT_TabStop {
	ret := &CT_TabStop{}
	ret.ValAttr = ST_TabJc(1)
	return ret
}

func (m *CT_TabStop) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.ValAttr.MarshalXMLAttr(xml.Name{Local: "w:val"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	if m.LeaderAttr != ST_TabTlcUnset {
		attr, err := m.LeaderAttr.MarshalXMLAttr(xml.Name{Local: "w:leader"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:pos"},
		Value: fmt.Sprintf("%v", m.PosAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TabStop) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ValAttr = ST_TabJc(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "val" {
			m.ValAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "leader" {
			m.LeaderAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "pos" {
			parsed, err := ParseUnionST_SignedTwipsMeasure(attr.Value)
			if err != nil {
				return err
			}
			m.PosAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TabStop: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TabStop and its children
func (m *CT_TabStop) Validate() error {
	return m.ValidateWithPath("CT_TabStop")
}

// ValidateWithPath validates the CT_TabStop and its children, prefixing error messages with path
func (m *CT_TabStop) ValidateWithPath(path string) error {
	if m.ValAttr == ST_TabJcUnset {
		return fmt.Errorf("%s/ValAttr is a mandatory field", path)
	}
	if err := m.ValAttr.ValidateWithPath(path + "/ValAttr"); err != nil {
		return err
	}
	if err := m.LeaderAttr.ValidateWithPath(path + "/LeaderAttr"); err != nil {
		return err
	}
	if err := m.PosAttr.ValidateWithPath(path + "/PosAttr"); err != nil {
		return err
	}
	return nil
}
