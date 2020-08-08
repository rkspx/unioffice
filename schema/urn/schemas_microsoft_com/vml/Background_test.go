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

func TestBackgroundConstructor(t *testing.T) {
	v := vml.NewBackground()
	if v == nil {
		t.Errorf("vml.NewBackground must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed vml.Background should validate: %s", err)
	}
}

func TestBackgroundMarshalUnmarshal(t *testing.T) {
	v := vml.NewBackground()
	buf, _ := xml.Marshal(v)
	v2 := vml.NewBackground()
	xml.Unmarshal(buf, v2)
}
