// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// commercial license can be purchased via https://unidoc.io website.

package chart

import (
	"encoding/xml"
	"fmt"
)

// ST_HoleSize is a union type
type ST_HoleSize struct {
	ST_HoleSizePercent *string
	ST_HoleSizeUByte   *uint8
}

func (m *ST_HoleSize) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_HoleSize) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_HoleSizePercent != nil {
		e.EncodeToken(xml.CharData(*m.ST_HoleSizePercent))
	}
	if m.ST_HoleSizeUByte != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.ST_HoleSizeUByte)))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_HoleSize) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_HoleSizePercent != nil {
		mems = append(mems, "ST_HoleSizePercent")
	}
	if m.ST_HoleSizeUByte != nil {
		mems = append(mems, "ST_HoleSizeUByte")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_HoleSize) String() string {
	if m.ST_HoleSizePercent != nil {
		return fmt.Sprintf("%v", *m.ST_HoleSizePercent)
	}
	if m.ST_HoleSizeUByte != nil {
		return fmt.Sprintf("%v", *m.ST_HoleSizeUByte)
	}
	return ""
}
