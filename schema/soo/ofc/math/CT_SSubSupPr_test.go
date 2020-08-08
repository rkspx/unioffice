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

func TestCT_SSubSupPrConstructor(t *testing.T) {
	v := math.NewCT_SSubSupPr()
	if v == nil {
		t.Errorf("math.NewCT_SSubSupPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_SSubSupPr should validate: %s", err)
	}
}

func TestCT_SSubSupPrMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_SSubSupPr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_SSubSupPr()
	xml.Unmarshal(buf, v2)
}
