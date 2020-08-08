// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// commercial license can be purchased via https://unidoc.io website.

package terms

import (
	"encoding/xml"
	"fmt"

	"github.com/unidoc/unioffice"
)

type ElementsAndRefinementsGroup struct {
	Choice []*ElementsAndRefinementsGroupChoice
}

func NewElementsAndRefinementsGroup() *ElementsAndRefinementsGroup {
	ret := &ElementsAndRefinementsGroup{}
	return ret
}

func (m *ElementsAndRefinementsGroup) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Choice != nil {
		for _, c := range m.Choice {
			c.MarshalXML(e, xml.StartElement{})
		}
	}
	return nil
}

func (m *ElementsAndRefinementsGroup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lElementsAndRefinementsGroup:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://purl.org/dc/elements/1.1/", Local: "any"}:
				tmp := NewElementsAndRefinementsGroupChoice()
				if err := d.DecodeElement(&tmp.Any, &el); err != nil {
					return err
				}
				m.Choice = append(m.Choice, tmp)
			default:
				unioffice.Log("skipping unsupported element on ElementsAndRefinementsGroup %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lElementsAndRefinementsGroup
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the ElementsAndRefinementsGroup and its children
func (m *ElementsAndRefinementsGroup) Validate() error {
	return m.ValidateWithPath("ElementsAndRefinementsGroup")
}

// ValidateWithPath validates the ElementsAndRefinementsGroup and its children, prefixing error messages with path
func (m *ElementsAndRefinementsGroup) ValidateWithPath(path string) error {
	for i, v := range m.Choice {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Choice[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
