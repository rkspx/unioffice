// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package chartDrawing

import (
	"encoding/xml"
	"fmt"

	"github.com/unidoc/unioffice"
)

type CT_Drawing struct {
	EG_Anchor []*EG_Anchor
}

func NewCT_Drawing() *CT_Drawing {
	ret := &CT_Drawing{}
	return ret
}

func (m *CT_Drawing) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "CT_Drawing"
	e.EncodeToken(start)
	if m.EG_Anchor != nil {
		for _, c := range m.EG_Anchor {
			c.MarshalXML(e, xml.StartElement{})
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Drawing) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Drawing:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", Local: "relSizeAnchor"}:
				tmpanchor := NewEG_Anchor()
				tmpanchor.RelSizeAnchor = NewCT_RelSizeAnchor()
				if err := d.DecodeElement(tmpanchor.RelSizeAnchor, &el); err != nil {
					return err
				}
				m.EG_Anchor = append(m.EG_Anchor, tmpanchor)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chartDrawing", Local: "absSizeAnchor"}:
				tmpanchor := NewEG_Anchor()
				tmpanchor.AbsSizeAnchor = NewCT_AbsSizeAnchor()
				if err := d.DecodeElement(tmpanchor.AbsSizeAnchor, &el); err != nil {
					return err
				}
				m.EG_Anchor = append(m.EG_Anchor, tmpanchor)
			default:
				unioffice.Log("skipping unsupported element on CT_Drawing %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Drawing
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Drawing and its children
func (m *CT_Drawing) Validate() error {
	return m.ValidateWithPath("CT_Drawing")
}

// ValidateWithPath validates the CT_Drawing and its children, prefixing error messages with path
func (m *CT_Drawing) ValidateWithPath(path string) error {
	for i, v := range m.EG_Anchor {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_Anchor[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
