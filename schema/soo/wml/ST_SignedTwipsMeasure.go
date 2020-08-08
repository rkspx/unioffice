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
)

// ST_SignedTwipsMeasure is a union type
type ST_SignedTwipsMeasure struct {
	Int64               *int64
	ST_UniversalMeasure *string
}

func (m *ST_SignedTwipsMeasure) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_SignedTwipsMeasure) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Int64 != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.Int64)))
	}
	if m.ST_UniversalMeasure != nil {
		e.EncodeToken(xml.CharData(*m.ST_UniversalMeasure))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_SignedTwipsMeasure) ValidateWithPath(path string) error {
	mems := []string{}
	if m.Int64 != nil {
		mems = append(mems, "Int64")
	}
	if m.ST_UniversalMeasure != nil {
		mems = append(mems, "ST_UniversalMeasure")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_SignedTwipsMeasure) String() string {
	if m.Int64 != nil {
		return fmt.Sprintf("%v", *m.Int64)
	}
	if m.ST_UniversalMeasure != nil {
		return fmt.Sprintf("%v", *m.ST_UniversalMeasure)
	}
	return ""
}
