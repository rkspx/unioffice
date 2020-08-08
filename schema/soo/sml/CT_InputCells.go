// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_InputCells struct {
	// Reference
	RAttr string
	// Deleted
	DeletedAttr *bool
	// Undone
	UndoneAttr *bool
	// Value
	ValAttr string
	// Number Format Id
	NumFmtIdAttr *uint32
}

func NewCT_InputCells() *CT_InputCells {
	ret := &CT_InputCells{}
	return ret
}

func (m *CT_InputCells) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r"},
		Value: fmt.Sprintf("%v", m.RAttr)})
	if m.DeletedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "deleted"},
			Value: fmt.Sprintf("%d", b2i(*m.DeletedAttr))})
	}
	if m.UndoneAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "undone"},
			Value: fmt.Sprintf("%d", b2i(*m.UndoneAttr))})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "val"},
		Value: fmt.Sprintf("%v", m.ValAttr)})
	if m.NumFmtIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "numFmtId"},
			Value: fmt.Sprintf("%v", *m.NumFmtIdAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_InputCells) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "r" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RAttr = parsed
			continue
		}
		if attr.Name.Local == "deleted" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DeletedAttr = &parsed
			continue
		}
		if attr.Name.Local == "undone" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UndoneAttr = &parsed
			continue
		}
		if attr.Name.Local == "val" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ValAttr = parsed
			continue
		}
		if attr.Name.Local == "numFmtId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.NumFmtIdAttr = &pt
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_InputCells: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_InputCells and its children
func (m *CT_InputCells) Validate() error {
	return m.ValidateWithPath("CT_InputCells")
}

// ValidateWithPath validates the CT_InputCells and its children, prefixing error messages with path
func (m *CT_InputCells) ValidateWithPath(path string) error {
	return nil
}
