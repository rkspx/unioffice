// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// commercial license can be purchased via https://unidoc.io website.

package math_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/ofc/math"
)

func TestEG_ScriptStyleConstructor(t *testing.T) {
	v := math.NewEG_ScriptStyle()
	if v == nil {
		t.Errorf("math.NewEG_ScriptStyle must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.EG_ScriptStyle should validate: %s", err)
	}
}

func TestEG_ScriptStyleMarshalUnmarshal(t *testing.T) {
	v := math.NewEG_ScriptStyle()
	buf, _ := xml.Marshal(v)
	v2 := math.NewEG_ScriptStyle()
	xml.Unmarshal(buf, v2)
}
