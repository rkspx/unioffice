// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package vml_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/urn/schemas_microsoft_com/vml"
)

func TestOfcSignaturelineConstructor(t *testing.T) {
	v := vml.NewOfcSignatureline()
	if v == nil {
		t.Errorf("vml.NewOfcSignatureline must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.OfcSignatureline should validate: %s", err)
	}
}

func TestOfcSignaturelineMarshalUnmarshal(t *testing.T) {
	v := vml.NewOfcSignatureline()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewOfcSignatureline()
	xml.Unmarshal(buf, v2)
}
