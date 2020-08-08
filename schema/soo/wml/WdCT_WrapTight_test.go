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

func TestWdCT_WrapTightConstructor(t *testing.T) {
	v := wml.NewWdCT_WrapTight()
	if v == nil {
		t.Errorf("wml.NewWdCT_WrapTight must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_WrapTight should validate: %s", err)
	}
}

func TestWdCT_WrapTightMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_WrapTight()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_WrapTight()
	xml.Unmarshal(buf, v2)
}
