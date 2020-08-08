// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// commercial license can be purchased via https://unidoc.io website.

package pml_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/pml"
)

func TestCT_TLAnimateScaleBehaviorConstructor(t *testing.T) {
	v := pml.NewCT_TLAnimateScaleBehavior()
	if v == nil {
		t.Errorf("pml.NewCT_TLAnimateScaleBehavior must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLAnimateScaleBehavior should validate: %s", err)
	}
}

func TestCT_TLAnimateScaleBehaviorMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLAnimateScaleBehavior()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLAnimateScaleBehavior()
	xml.Unmarshal(buf, v2)
}
