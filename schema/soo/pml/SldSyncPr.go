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

	"github.com/unidoc/unioffice"
)

type SldSyncPr struct {
	CT_SlideSyncProperties
}

func NewSldSyncPr() *SldSyncPr {
	ret := &SldSyncPr{}
	ret.CT_SlideSyncProperties = *NewCT_SlideSyncProperties()
	return ret
}

func (m *SldSyncPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/presentationml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:p"}, Value: "http://schemas.openxmlformats.org/presentationml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "p:sldSyncPr"
	return m.CT_SlideSyncProperties.MarshalXML(e, start)
}

func (m *SldSyncPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_SlideSyncProperties = *NewCT_SlideSyncProperties()
	for _, attr := range start.Attr {
		if attr.Name.Local == "serverSldId" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ServerSldIdAttr = parsed
			continue
		}
		if attr.Name.Local == "serverSldModifiedTime" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.ServerSldModifiedTimeAttr = parsed
			continue
		}
		if attr.Name.Local == "clientInsertedTime" {
			parsed, err := ParseStdlibTime(attr.Value)
			if err != nil {
				return err
			}
			m.ClientInsertedTimeAttr = parsed
			continue
		}
	}
lSldSyncPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on SldSyncPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lSldSyncPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the SldSyncPr and its children
func (m *SldSyncPr) Validate() error {
	return m.ValidateWithPath("SldSyncPr")
}

// ValidateWithPath validates the SldSyncPr and its children, prefixing error messages with path
func (m *SldSyncPr) ValidateWithPath(path string) error {
	if err := m.CT_SlideSyncProperties.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
