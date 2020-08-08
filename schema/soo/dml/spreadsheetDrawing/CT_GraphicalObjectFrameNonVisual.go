// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// commercial license can be purchased via https://unidoc.io website.

package spreadsheetDrawing

import (
	"encoding/xml"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/schema/soo/dml"
)

type CT_GraphicalObjectFrameNonVisual struct {
	CNvPr             *dml.CT_NonVisualDrawingProps
	CNvGraphicFramePr *dml.CT_NonVisualGraphicFrameProperties
}

func NewCT_GraphicalObjectFrameNonVisual() *CT_GraphicalObjectFrameNonVisual {
	ret := &CT_GraphicalObjectFrameNonVisual{}
	ret.CNvPr = dml.NewCT_NonVisualDrawingProps()
	ret.CNvGraphicFramePr = dml.NewCT_NonVisualGraphicFrameProperties()
	return ret
}

func (m *CT_GraphicalObjectFrameNonVisual) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secNvPr := xml.StartElement{Name: xml.Name{Local: "xdr:cNvPr"}}
	e.EncodeElement(m.CNvPr, secNvPr)
	secNvGraphicFramePr := xml.StartElement{Name: xml.Name{Local: "xdr:cNvGraphicFramePr"}}
	e.EncodeElement(m.CNvGraphicFramePr, secNvGraphicFramePr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GraphicalObjectFrameNonVisual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CNvPr = dml.NewCT_NonVisualDrawingProps()
	m.CNvGraphicFramePr = dml.NewCT_NonVisualGraphicFrameProperties()
lCT_GraphicalObjectFrameNonVisual:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "cNvPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "cNvPr"}:
				if err := d.DecodeElement(m.CNvPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "cNvGraphicFramePr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "cNvGraphicFramePr"}:
				if err := d.DecodeElement(m.CNvGraphicFramePr, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_GraphicalObjectFrameNonVisual %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GraphicalObjectFrameNonVisual
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GraphicalObjectFrameNonVisual and its children
func (m *CT_GraphicalObjectFrameNonVisual) Validate() error {
	return m.ValidateWithPath("CT_GraphicalObjectFrameNonVisual")
}

// ValidateWithPath validates the CT_GraphicalObjectFrameNonVisual and its children, prefixing error messages with path
func (m *CT_GraphicalObjectFrameNonVisual) ValidateWithPath(path string) error {
	if err := m.CNvPr.ValidateWithPath(path + "/CNvPr"); err != nil {
		return err
	}
	if err := m.CNvGraphicFramePr.ValidateWithPath(path + "/CNvGraphicFramePr"); err != nil {
		return err
	}
	return nil
}
