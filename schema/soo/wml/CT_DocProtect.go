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

	"github.com/unidoc/unioffice/schema/soo/ofc/sharedTypes"
)

type CT_DocProtect struct {
	// Document Editing Restrictions
	EditAttr ST_DocProtect
	// Only Allow Formatting With Unlocked Styles
	FormattingAttr *sharedTypes.ST_OnOff
	// Enforce Document Protection Settings
	EnforcementAttr                *sharedTypes.ST_OnOff
	AlgorithmNameAttr              *string
	HashValueAttr                  *string
	SaltValueAttr                  *string
	SpinCountAttr                  *int64
	CryptProviderTypeAttr          sharedTypes.ST_CryptProv
	CryptAlgorithmClassAttr        sharedTypes.ST_AlgClass
	CryptAlgorithmTypeAttr         sharedTypes.ST_AlgType
	CryptAlgorithmSidAttr          *int64
	CryptSpinCountAttr             *int64
	CryptProviderAttr              *string
	AlgIdExtAttr                   *string
	AlgIdExtSourceAttr             *string
	CryptProviderTypeExtAttr       *string
	CryptProviderTypeExtSourceAttr *string
	HashAttr                       *string
	SaltAttr                       *string
}

func NewCT_DocProtect() *CT_DocProtect {
	ret := &CT_DocProtect{}
	return ret
}

func (m *CT_DocProtect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.EditAttr != ST_DocProtectUnset {
		attr, err := m.EditAttr.MarshalXMLAttr(xml.Name{Local: "w:edit"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.FormattingAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:formatting"},
			Value: fmt.Sprintf("%v", *m.FormattingAttr)})
	}
	if m.EnforcementAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:enforcement"},
			Value: fmt.Sprintf("%v", *m.EnforcementAttr)})
	}
	if m.AlgorithmNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:algorithmName"},
			Value: fmt.Sprintf("%v", *m.AlgorithmNameAttr)})
	}
	if m.HashValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:hashValue"},
			Value: fmt.Sprintf("%v", *m.HashValueAttr)})
	}
	if m.SaltValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:saltValue"},
			Value: fmt.Sprintf("%v", *m.SaltValueAttr)})
	}
	if m.SpinCountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:spinCount"},
			Value: fmt.Sprintf("%v", *m.SpinCountAttr)})
	}
	if m.CryptProviderTypeAttr != sharedTypes.ST_CryptProvUnset {
		attr, err := m.CryptProviderTypeAttr.MarshalXMLAttr(xml.Name{Local: "w:cryptProviderType"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.CryptAlgorithmClassAttr != sharedTypes.ST_AlgClassUnset {
		attr, err := m.CryptAlgorithmClassAttr.MarshalXMLAttr(xml.Name{Local: "w:cryptAlgorithmClass"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.CryptAlgorithmTypeAttr != sharedTypes.ST_AlgTypeUnset {
		attr, err := m.CryptAlgorithmTypeAttr.MarshalXMLAttr(xml.Name{Local: "w:cryptAlgorithmType"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.CryptAlgorithmSidAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:cryptAlgorithmSid"},
			Value: fmt.Sprintf("%v", *m.CryptAlgorithmSidAttr)})
	}
	if m.CryptSpinCountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:cryptSpinCount"},
			Value: fmt.Sprintf("%v", *m.CryptSpinCountAttr)})
	}
	if m.CryptProviderAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:cryptProvider"},
			Value: fmt.Sprintf("%v", *m.CryptProviderAttr)})
	}
	if m.AlgIdExtAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:algIdExt"},
			Value: fmt.Sprintf("%v", *m.AlgIdExtAttr)})
	}
	if m.AlgIdExtSourceAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:algIdExtSource"},
			Value: fmt.Sprintf("%v", *m.AlgIdExtSourceAttr)})
	}
	if m.CryptProviderTypeExtAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:cryptProviderTypeExt"},
			Value: fmt.Sprintf("%v", *m.CryptProviderTypeExtAttr)})
	}
	if m.CryptProviderTypeExtSourceAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:cryptProviderTypeExtSource"},
			Value: fmt.Sprintf("%v", *m.CryptProviderTypeExtSourceAttr)})
	}
	if m.HashAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:hash"},
			Value: fmt.Sprintf("%v", *m.HashAttr)})
	}
	if m.SaltAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:salt"},
			Value: fmt.Sprintf("%v", *m.SaltAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DocProtect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "cryptAlgorithmType" {
			m.CryptAlgorithmTypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "cryptAlgorithmSid" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.CryptAlgorithmSidAttr = &parsed
			continue
		}
		if attr.Name.Local == "formatting" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.FormattingAttr = &parsed
			continue
		}
		if attr.Name.Local == "cryptSpinCount" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.CryptSpinCountAttr = &parsed
			continue
		}
		if attr.Name.Local == "algorithmName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AlgorithmNameAttr = &parsed
			continue
		}
		if attr.Name.Local == "cryptProvider" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CryptProviderAttr = &parsed
			continue
		}
		if attr.Name.Local == "saltValue" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SaltValueAttr = &parsed
			continue
		}
		if attr.Name.Local == "cryptProviderType" {
			m.CryptProviderTypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "cryptAlgorithmClass" {
			m.CryptAlgorithmClassAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "edit" {
			m.EditAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "enforcement" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.EnforcementAttr = &parsed
			continue
		}
		if attr.Name.Local == "hashValue" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.HashValueAttr = &parsed
			continue
		}
		if attr.Name.Local == "spinCount" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.SpinCountAttr = &parsed
			continue
		}
		if attr.Name.Local == "algIdExt" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AlgIdExtAttr = &parsed
			continue
		}
		if attr.Name.Local == "algIdExtSource" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AlgIdExtSourceAttr = &parsed
			continue
		}
		if attr.Name.Local == "cryptProviderTypeExt" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CryptProviderTypeExtAttr = &parsed
			continue
		}
		if attr.Name.Local == "cryptProviderTypeExtSource" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CryptProviderTypeExtSourceAttr = &parsed
			continue
		}
		if attr.Name.Local == "hash" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.HashAttr = &parsed
			continue
		}
		if attr.Name.Local == "salt" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SaltAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_DocProtect: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_DocProtect and its children
func (m *CT_DocProtect) Validate() error {
	return m.ValidateWithPath("CT_DocProtect")
}

// ValidateWithPath validates the CT_DocProtect and its children, prefixing error messages with path
func (m *CT_DocProtect) ValidateWithPath(path string) error {
	if err := m.EditAttr.ValidateWithPath(path + "/EditAttr"); err != nil {
		return err
	}
	if m.FormattingAttr != nil {
		if err := m.FormattingAttr.ValidateWithPath(path + "/FormattingAttr"); err != nil {
			return err
		}
	}
	if m.EnforcementAttr != nil {
		if err := m.EnforcementAttr.ValidateWithPath(path + "/EnforcementAttr"); err != nil {
			return err
		}
	}
	if err := m.CryptProviderTypeAttr.ValidateWithPath(path + "/CryptProviderTypeAttr"); err != nil {
		return err
	}
	if err := m.CryptAlgorithmClassAttr.ValidateWithPath(path + "/CryptAlgorithmClassAttr"); err != nil {
		return err
	}
	if err := m.CryptAlgorithmTypeAttr.ValidateWithPath(path + "/CryptAlgorithmTypeAttr"); err != nil {
		return err
	}
	return nil
}
