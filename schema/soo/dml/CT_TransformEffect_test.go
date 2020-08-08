// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// commercial license can be purchased via https://unidoc.io website.

package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/dml"
)

func TestCT_TransformEffectConstructor(t *testing.T) {
	v := dml.NewCT_TransformEffect()
	if v == nil {
		t.Errorf("dml.NewCT_TransformEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TransformEffect should validate: %s", err)
	}
}

func TestCT_TransformEffectMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TransformEffect()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TransformEffect()
	xml.Unmarshal(buf, v2)
}
