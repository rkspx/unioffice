// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// commercial license can be purchased via https://unidoc.io website.

package vml_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/urn/schemas_microsoft_com/vml"
)

func TestCT_HConstructor(t *testing.T) {
	v := vml.NewCT_H()
	if v == nil {
		t.Errorf("vml.NewCT_H must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.CT_H should validate: %s", err)
	}
}

func TestCT_HMarshalUnmarshal(t *testing.T) {
	v := vml.NewCT_H()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCT_H()
	xml.Unmarshal(buf, v2)
}
