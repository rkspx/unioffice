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

func TestAG_ChromakeyConstructor(t *testing.T) {
	v := vml.NewAG_Chromakey()
	if v == nil {
		t.Errorf("vml.NewAG_Chromakey must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.AG_Chromakey should validate: %s", err)
	}
}

func TestAG_ChromakeyMarshalUnmarshal(t *testing.T) {
	v := vml.NewAG_Chromakey()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewAG_Chromakey()
	xml.Unmarshal(buf, v2)
}
