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

type CT_PrintOptions struct {
	// Horizontal Centered
	HorizontalCenteredAttr *bool
	// Vertical Centered
	VerticalCenteredAttr *bool
	// Print Headings
	HeadingsAttr *bool
	// Print Grid Lines
	GridLinesAttr *bool
	// Grid Lines Set
	GridLinesSetAttr *bool
}

func NewCT_PrintOptions() *CT_PrintOptions {
	ret := &CT_PrintOptions{}
	return ret
}

func (m *CT_PrintOptions) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.HorizontalCenteredAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "horizontalCentered"},
			Value: fmt.Sprintf("%d", b2i(*m.HorizontalCenteredAttr))})
	}
	if m.VerticalCenteredAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "verticalCentered"},
			Value: fmt.Sprintf("%d", b2i(*m.VerticalCenteredAttr))})
	}
	if m.HeadingsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "headings"},
			Value: fmt.Sprintf("%d", b2i(*m.HeadingsAttr))})
	}
	if m.GridLinesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "gridLines"},
			Value: fmt.Sprintf("%d", b2i(*m.GridLinesAttr))})
	}
	if m.GridLinesSetAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "gridLinesSet"},
			Value: fmt.Sprintf("%d", b2i(*m.GridLinesSetAttr))})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PrintOptions) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "horizontalCentered" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HorizontalCenteredAttr = &parsed
			continue
		}
		if attr.Name.Local == "verticalCentered" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.VerticalCenteredAttr = &parsed
			continue
		}
		if attr.Name.Local == "headings" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HeadingsAttr = &parsed
			continue
		}
		if attr.Name.Local == "gridLines" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.GridLinesAttr = &parsed
			continue
		}
		if attr.Name.Local == "gridLinesSet" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.GridLinesSetAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_PrintOptions: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_PrintOptions and its children
func (m *CT_PrintOptions) Validate() error {
	return m.ValidateWithPath("CT_PrintOptions")
}

// ValidateWithPath validates the CT_PrintOptions and its children, prefixing error messages with path
func (m *CT_PrintOptions) ValidateWithPath(path string) error {
	return nil
}
