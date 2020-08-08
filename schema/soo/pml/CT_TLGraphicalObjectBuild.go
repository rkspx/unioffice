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

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/schema/soo/dml"
)

type CT_TLGraphicalObjectBuild struct {
	// Build As One
	BldAsOne *CT_Empty
	// Build Sub Elements
	BldSub       *dml.CT_AnimationGraphicalObjectBuildProperties
	SpidAttr     *uint32
	GrpIdAttr    *uint32
	UiExpandAttr *bool
}

func NewCT_TLGraphicalObjectBuild() *CT_TLGraphicalObjectBuild {
	ret := &CT_TLGraphicalObjectBuild{}
	return ret
}

func (m *CT_TLGraphicalObjectBuild) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.SpidAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "spid"},
			Value: fmt.Sprintf("%v", *m.SpidAttr)})
	}
	if m.GrpIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "grpId"},
			Value: fmt.Sprintf("%v", *m.GrpIdAttr)})
	}
	if m.UiExpandAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uiExpand"},
			Value: fmt.Sprintf("%d", b2i(*m.UiExpandAttr))})
	}
	e.EncodeToken(start)
	if m.BldAsOne != nil {
		sebldAsOne := xml.StartElement{Name: xml.Name{Local: "p:bldAsOne"}}
		e.EncodeElement(m.BldAsOne, sebldAsOne)
	}
	if m.BldSub != nil {
		sebldSub := xml.StartElement{Name: xml.Name{Local: "p:bldSub"}}
		e.EncodeElement(m.BldSub, sebldSub)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLGraphicalObjectBuild) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "spid" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.SpidAttr = &pt
			continue
		}
		if attr.Name.Local == "grpId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.GrpIdAttr = &pt
			continue
		}
		if attr.Name.Local == "uiExpand" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UiExpandAttr = &parsed
			continue
		}
	}
lCT_TLGraphicalObjectBuild:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "bldAsOne"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "bldAsOne"}:
				m.BldAsOne = NewCT_Empty()
				if err := d.DecodeElement(m.BldAsOne, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "bldSub"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "bldSub"}:
				m.BldSub = dml.NewCT_AnimationGraphicalObjectBuildProperties()
				if err := d.DecodeElement(m.BldSub, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_TLGraphicalObjectBuild %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLGraphicalObjectBuild
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLGraphicalObjectBuild and its children
func (m *CT_TLGraphicalObjectBuild) Validate() error {
	return m.ValidateWithPath("CT_TLGraphicalObjectBuild")
}

// ValidateWithPath validates the CT_TLGraphicalObjectBuild and its children, prefixing error messages with path
func (m *CT_TLGraphicalObjectBuild) ValidateWithPath(path string) error {
	if m.BldAsOne != nil {
		if err := m.BldAsOne.ValidateWithPath(path + "/BldAsOne"); err != nil {
			return err
		}
	}
	if m.BldSub != nil {
		if err := m.BldSub.ValidateWithPath(path + "/BldSub"); err != nil {
			return err
		}
	}
	return nil
}
