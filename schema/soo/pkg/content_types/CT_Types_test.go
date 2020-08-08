// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// commercial license can be purchased via https://unidoc.io website.

package content_types_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/pkg/content_types"
)

func TestCT_TypesConstructor(t *testing.T) {
	v := content_types.NewCT_Types()
	if v == nil {
		t.Errorf("content_types.NewCT_Types must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed content_types.CT_Types should validate: %s", err)
	}
}

func TestCT_TypesMarshalUnmarshal(t *testing.T) {
	v := content_types.NewCT_Types()
	buf, _ := xml.Marshal(v)
	v2 := content_types.NewCT_Types()
	xml.Unmarshal(buf, v2)
}
