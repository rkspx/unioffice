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
	"strconv"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/schema/soo/dml"
)

type WdCT_Anchor struct {
	DistTAttr          *uint32
	DistBAttr          *uint32
	DistLAttr          *uint32
	DistRAttr          *uint32
	SimplePosAttr      *bool
	RelativeHeightAttr uint32
	BehindDocAttr      bool
	LockedAttr         bool
	LayoutInCellAttr   bool
	HiddenAttr         *bool
	AllowOverlapAttr   bool
	SimplePos          *dml.CT_Point2D
	PositionH          *WdCT_PosH
	PositionV          *WdCT_PosV
	Extent             *dml.CT_PositiveSize2D
	EffectExtent       *WdCT_EffectExtent
	Choice             *WdEG_WrapTypeChoice
	DocPr              *dml.CT_NonVisualDrawingProps
	CNvGraphicFramePr  *dml.CT_NonVisualGraphicFrameProperties
	Graphic            *dml.Graphic
}

func NewWdCT_Anchor() *WdCT_Anchor {
	ret := &WdCT_Anchor{}
	ret.SimplePos = dml.NewCT_Point2D()
	ret.PositionH = NewWdCT_PosH()
	ret.PositionV = NewWdCT_PosV()
	ret.Extent = dml.NewCT_PositiveSize2D()
	ret.DocPr = dml.NewCT_NonVisualDrawingProps()
	ret.Graphic = dml.NewGraphic()
	return ret
}

func (m *WdCT_Anchor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.DistTAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distT"},
			Value: fmt.Sprintf("%v", *m.DistTAttr)})
	}
	if m.DistBAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distB"},
			Value: fmt.Sprintf("%v", *m.DistBAttr)})
	}
	if m.DistLAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distL"},
			Value: fmt.Sprintf("%v", *m.DistLAttr)})
	}
	if m.DistRAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distR"},
			Value: fmt.Sprintf("%v", *m.DistRAttr)})
	}
	if m.SimplePosAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "simplePos"},
			Value: fmt.Sprintf("%d", b2i(*m.SimplePosAttr))})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "relativeHeight"},
		Value: fmt.Sprintf("%v", m.RelativeHeightAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "behindDoc"},
		Value: fmt.Sprintf("%d", b2i(m.BehindDocAttr))})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "locked"},
		Value: fmt.Sprintf("%d", b2i(m.LockedAttr))})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "layoutInCell"},
		Value: fmt.Sprintf("%d", b2i(m.LayoutInCellAttr))})
	if m.HiddenAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hidden"},
			Value: fmt.Sprintf("%d", b2i(*m.HiddenAttr))})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "allowOverlap"},
		Value: fmt.Sprintf("%d", b2i(m.AllowOverlapAttr))})
	e.EncodeToken(start)
	sesimplePos := xml.StartElement{Name: xml.Name{Local: "wp:simplePos"}}
	e.EncodeElement(m.SimplePos, sesimplePos)
	sepositionH := xml.StartElement{Name: xml.Name{Local: "wp:positionH"}}
	e.EncodeElement(m.PositionH, sepositionH)
	sepositionV := xml.StartElement{Name: xml.Name{Local: "wp:positionV"}}
	e.EncodeElement(m.PositionV, sepositionV)
	seextent := xml.StartElement{Name: xml.Name{Local: "wp:extent"}}
	e.EncodeElement(m.Extent, seextent)
	if m.EffectExtent != nil {
		seeffectExtent := xml.StartElement{Name: xml.Name{Local: "wp:effectExtent"}}
		e.EncodeElement(m.EffectExtent, seeffectExtent)
	}
	if m.Choice != nil {
		m.Choice.MarshalXML(e, xml.StartElement{})
	}
	sedocPr := xml.StartElement{Name: xml.Name{Local: "wp:docPr"}}
	e.EncodeElement(m.DocPr, sedocPr)
	if m.CNvGraphicFramePr != nil {
		secNvGraphicFramePr := xml.StartElement{Name: xml.Name{Local: "wp:cNvGraphicFramePr"}}
		e.EncodeElement(m.CNvGraphicFramePr, secNvGraphicFramePr)
	}
	segraphic := xml.StartElement{Name: xml.Name{Local: "a:graphic"}}
	e.EncodeElement(m.Graphic, segraphic)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *WdCT_Anchor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.SimplePos = dml.NewCT_Point2D()
	m.PositionH = NewWdCT_PosH()
	m.PositionV = NewWdCT_PosV()
	m.Extent = dml.NewCT_PositiveSize2D()
	m.DocPr = dml.NewCT_NonVisualDrawingProps()
	m.Graphic = dml.NewGraphic()
	for _, attr := range start.Attr {
		if attr.Name.Local == "distT" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistTAttr = &pt
			continue
		}
		if attr.Name.Local == "distL" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistLAttr = &pt
			continue
		}
		if attr.Name.Local == "simplePos" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SimplePosAttr = &parsed
			continue
		}
		if attr.Name.Local == "behindDoc" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.BehindDocAttr = parsed
			continue
		}
		if attr.Name.Local == "layoutInCell" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LayoutInCellAttr = parsed
			continue
		}
		if attr.Name.Local == "hidden" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.HiddenAttr = &parsed
			continue
		}
		if attr.Name.Local == "distB" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistBAttr = &pt
			continue
		}
		if attr.Name.Local == "distR" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DistRAttr = &pt
			continue
		}
		if attr.Name.Local == "relativeHeight" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.RelativeHeightAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "locked" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LockedAttr = parsed
			continue
		}
		if attr.Name.Local == "allowOverlap" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AllowOverlapAttr = parsed
			continue
		}
	}
lWdCT_Anchor:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "simplePos"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "simplePos"}:
				if err := d.DecodeElement(m.SimplePos, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "positionH"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "positionH"}:
				if err := d.DecodeElement(m.PositionH, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "positionV"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "positionV"}:
				if err := d.DecodeElement(m.PositionV, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "extent"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "extent"}:
				if err := d.DecodeElement(m.Extent, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "effectExtent"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "effectExtent"}:
				m.EffectExtent = NewWdCT_EffectExtent()
				if err := d.DecodeElement(m.EffectExtent, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "wrapNone"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "wrapNone"}:
				m.Choice = NewWdEG_WrapTypeChoice()
				if err := d.DecodeElement(&m.Choice.WrapNone, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "wrapSquare"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "wrapSquare"}:
				m.Choice = NewWdEG_WrapTypeChoice()
				if err := d.DecodeElement(&m.Choice.WrapSquare, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "wrapTight"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "wrapTight"}:
				m.Choice = NewWdEG_WrapTypeChoice()
				if err := d.DecodeElement(&m.Choice.WrapTight, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "wrapThrough"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "wrapThrough"}:
				m.Choice = NewWdEG_WrapTypeChoice()
				if err := d.DecodeElement(&m.Choice.WrapThrough, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "wrapTopAndBottom"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "wrapTopAndBottom"}:
				m.Choice = NewWdEG_WrapTypeChoice()
				if err := d.DecodeElement(&m.Choice.WrapTopAndBottom, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "docPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "docPr"}:
				if err := d.DecodeElement(m.DocPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "cNvGraphicFramePr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "cNvGraphicFramePr"}:
				m.CNvGraphicFramePr = dml.NewCT_NonVisualGraphicFrameProperties()
				if err := d.DecodeElement(m.CNvGraphicFramePr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "graphic"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "graphic"}:
				if err := d.DecodeElement(m.Graphic, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on WdCT_Anchor %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWdCT_Anchor
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the WdCT_Anchor and its children
func (m *WdCT_Anchor) Validate() error {
	return m.ValidateWithPath("WdCT_Anchor")
}

// ValidateWithPath validates the WdCT_Anchor and its children, prefixing error messages with path
func (m *WdCT_Anchor) ValidateWithPath(path string) error {
	if err := m.SimplePos.ValidateWithPath(path + "/SimplePos"); err != nil {
		return err
	}
	if err := m.PositionH.ValidateWithPath(path + "/PositionH"); err != nil {
		return err
	}
	if err := m.PositionV.ValidateWithPath(path + "/PositionV"); err != nil {
		return err
	}
	if err := m.Extent.ValidateWithPath(path + "/Extent"); err != nil {
		return err
	}
	if m.EffectExtent != nil {
		if err := m.EffectExtent.ValidateWithPath(path + "/EffectExtent"); err != nil {
			return err
		}
	}
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	if err := m.DocPr.ValidateWithPath(path + "/DocPr"); err != nil {
		return err
	}
	if m.CNvGraphicFramePr != nil {
		if err := m.CNvGraphicFramePr.ValidateWithPath(path + "/CNvGraphicFramePr"); err != nil {
			return err
		}
	}
	if err := m.Graphic.ValidateWithPath(path + "/Graphic"); err != nil {
		return err
	}
	return nil
}
