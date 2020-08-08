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

func TestCT_TextPathConstructor(t *testing.T) {
	v := vml.NewCT_TextPath()
	if v == nil {
		t.Errorf("vml.NewCT_TextPath must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.CT_TextPath should validate: %s", err)
	}
}

func TestCT_TextPathMarshalUnmarshal(t *testing.T) {
	v := vml.NewCT_TextPath()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewCT_TextPath()
	xml.Unmarshal(buf, v2)
}
