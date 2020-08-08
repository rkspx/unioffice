// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/wml"
)

func TestCT_PageSzConstructor(t *testing.T) {
	v := wml.NewCT_PageSz()
	if v == nil {
		t.Errorf("wml.NewCT_PageSz must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_PageSz should validate: %s", err)
	}
}

func TestCT_PageSzMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_PageSz()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_PageSz()
	xml.Unmarshal(buf, v2)
}
