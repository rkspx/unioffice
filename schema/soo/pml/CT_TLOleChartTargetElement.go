// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// commercial license can be purchased via https://unidoc.io website.

package pml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_TLOleChartTargetElement struct {
	// Type
	TypeAttr ST_TLChartSubelementType
	// Level
	LvlAttr *uint32
}

func NewCT_TLOleChartTargetElement() *CT_TLOleChartTargetElement {
	ret := &CT_TLOleChartTargetElement{}
	ret.TypeAttr = ST_TLChartSubelementType(1)
	return ret
}

func (m *CT_TLOleChartTargetElement) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	if m.LvlAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lvl"},
			Value: fmt.Sprintf("%v", *m.LvlAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLOleChartTargetElement) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TypeAttr = ST_TLChartSubelementType(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "lvl" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.LvlAttr = &pt
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TLOleChartTargetElement: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TLOleChartTargetElement and its children
func (m *CT_TLOleChartTargetElement) Validate() error {
	return m.ValidateWithPath("CT_TLOleChartTargetElement")
}

// ValidateWithPath validates the CT_TLOleChartTargetElement and its children, prefixing error messages with path
func (m *CT_TLOleChartTargetElement) ValidateWithPath(path string) error {
	if m.TypeAttr == ST_TLChartSubelementTypeUnset {
		return fmt.Errorf("%s/TypeAttr is a mandatory field", path)
	}
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	return nil
}
