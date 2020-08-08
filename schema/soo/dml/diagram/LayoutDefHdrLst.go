// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package diagram

import (
	"encoding/xml"

	"github.com/unidoc/unioffice"
)

type LayoutDefHdrLst struct {
	CT_DiagramDefinitionHeaderLst
}

func NewLayoutDefHdrLst() *LayoutDefHdrLst {
	ret := &LayoutDefHdrLst{}
	ret.CT_DiagramDefinitionHeaderLst = *NewCT_DiagramDefinitionHeaderLst()
	return ret
}

func (m *LayoutDefHdrLst) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/diagram"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:di"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/diagram"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "layoutDefHdrLst"
	return m.CT_DiagramDefinitionHeaderLst.MarshalXML(e, start)
}

func (m *LayoutDefHdrLst) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_DiagramDefinitionHeaderLst = *NewCT_DiagramDefinitionHeaderLst()
lLayoutDefHdrLst:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "layoutDefHdr"}:
				tmp := NewCT_DiagramDefinitionHeader()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.LayoutDefHdr = append(m.LayoutDefHdr, tmp)
			default:
				unioffice.Log("skipping unsupported element on LayoutDefHdrLst %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lLayoutDefHdrLst
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the LayoutDefHdrLst and its children
func (m *LayoutDefHdrLst) Validate() error {
	return m.ValidateWithPath("LayoutDefHdrLst")
}

// ValidateWithPath validates the LayoutDefHdrLst and its children, prefixing error messages with path
func (m *LayoutDefHdrLst) ValidateWithPath(path string) error {
	if err := m.CT_DiagramDefinitionHeaderLst.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
