// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// commercial license can be purchased via https://unidoc.io website.

package vml

import (
	"encoding/xml"
	"fmt"

	"github.com/unidoc/unioffice/schema/soo/ofc/sharedTypes"
)

type OfcCT_SignatureLine struct {
	IssignaturelineAttr        sharedTypes.ST_TrueFalse
	IdAttr                     *string
	ProvidAttr                 *string
	SigninginstructionssetAttr sharedTypes.ST_TrueFalse
	AllowcommentsAttr          sharedTypes.ST_TrueFalse
	ShowsigndateAttr           sharedTypes.ST_TrueFalse
	SuggestedsignerAttr        *string
	Suggestedsigner2Attr       *string
	SuggestedsigneremailAttr   *string
	SigninginstructionsAttr    *string
	AddlxmlAttr                *string
	SigprovurlAttr             *string
	ExtAttr                    ST_Ext
}

func NewOfcCT_SignatureLine() *OfcCT_SignatureLine {
	ret := &OfcCT_SignatureLine{}
	return ret
}

func (m *OfcCT_SignatureLine) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.IssignaturelineAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.IssignaturelineAttr.MarshalXMLAttr(xml.Name{Local: "issignatureline"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	if m.ProvidAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "provid"},
			Value: fmt.Sprintf("%v", *m.ProvidAttr)})
	}
	if m.SigninginstructionssetAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.SigninginstructionssetAttr.MarshalXMLAttr(xml.Name{Local: "signinginstructionsset"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.AllowcommentsAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.AllowcommentsAttr.MarshalXMLAttr(xml.Name{Local: "allowcomments"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ShowsigndateAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.ShowsigndateAttr.MarshalXMLAttr(xml.Name{Local: "showsigndate"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.SuggestedsignerAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "suggestedsigner"},
			Value: fmt.Sprintf("%v", *m.SuggestedsignerAttr)})
	}
	if m.Suggestedsigner2Attr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "suggestedsigner2"},
			Value: fmt.Sprintf("%v", *m.Suggestedsigner2Attr)})
	}
	if m.SuggestedsigneremailAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "suggestedsigneremail"},
			Value: fmt.Sprintf("%v", *m.SuggestedsigneremailAttr)})
	}
	if m.SigninginstructionsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "signinginstructions"},
			Value: fmt.Sprintf("%v", *m.SigninginstructionsAttr)})
	}
	if m.AddlxmlAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "addlxml"},
			Value: fmt.Sprintf("%v", *m.AddlxmlAttr)})
	}
	if m.SigprovurlAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sigprovurl"},
			Value: fmt.Sprintf("%v", *m.SigprovurlAttr)})
	}
	if m.ExtAttr != ST_ExtUnset {
		attr, err := m.ExtAttr.MarshalXMLAttr(xml.Name{Local: "ext"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *OfcCT_SignatureLine) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "suggestedsigner" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SuggestedsignerAttr = &parsed
			continue
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
		if attr.Name.Local == "provid" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ProvidAttr = &parsed
			continue
		}
		if attr.Name.Local == "signinginstructionsset" {
			m.SigninginstructionssetAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "allowcomments" {
			m.AllowcommentsAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "showsigndate" {
			m.ShowsigndateAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "issignatureline" {
			m.IssignaturelineAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "suggestedsigner2" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.Suggestedsigner2Attr = &parsed
			continue
		}
		if attr.Name.Local == "suggestedsigneremail" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SuggestedsigneremailAttr = &parsed
			continue
		}
		if attr.Name.Local == "signinginstructions" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SigninginstructionsAttr = &parsed
			continue
		}
		if attr.Name.Local == "addlxml" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AddlxmlAttr = &parsed
			continue
		}
		if attr.Name.Local == "sigprovurl" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SigprovurlAttr = &parsed
			continue
		}
		if attr.Name.Local == "ext" {
			m.ExtAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing OfcCT_SignatureLine: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the OfcCT_SignatureLine and its children
func (m *OfcCT_SignatureLine) Validate() error {
	return m.ValidateWithPath("OfcCT_SignatureLine")
}

// ValidateWithPath validates the OfcCT_SignatureLine and its children, prefixing error messages with path
func (m *OfcCT_SignatureLine) ValidateWithPath(path string) error {
	if err := m.IssignaturelineAttr.ValidateWithPath(path + "/IssignaturelineAttr"); err != nil {
		return err
	}
	if m.IdAttr != nil {
		if !sharedTypes.ST_GuidPatternRe.MatchString(*m.IdAttr) {
			return fmt.Errorf(`%s/m.IdAttr must match '%s' (have %v)`, path, sharedTypes.ST_GuidPatternRe, *m.IdAttr)
		}
	}
	if m.ProvidAttr != nil {
		if !sharedTypes.ST_GuidPatternRe.MatchString(*m.ProvidAttr) {
			return fmt.Errorf(`%s/m.ProvidAttr must match '%s' (have %v)`, path, sharedTypes.ST_GuidPatternRe, *m.ProvidAttr)
		}
	}
	if err := m.SigninginstructionssetAttr.ValidateWithPath(path + "/SigninginstructionssetAttr"); err != nil {
		return err
	}
	if err := m.AllowcommentsAttr.ValidateWithPath(path + "/AllowcommentsAttr"); err != nil {
		return err
	}
	if err := m.ShowsigndateAttr.ValidateWithPath(path + "/ShowsigndateAttr"); err != nil {
		return err
	}
	if err := m.ExtAttr.ValidateWithPath(path + "/ExtAttr"); err != nil {
		return err
	}
	return nil
}
