// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// commercial license can be purchased via https://unidoc.io website.

package picture_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/dml/picture"
)

func TestCT_PictureConstructor(t *testing.T) {
	v := picture.NewCT_Picture()
	if v == nil {
		t.Errorf("picture.NewCT_Picture must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed picture.CT_Picture should validate: %s", err)
	}
}

func TestCT_PictureMarshalUnmarshal(t *testing.T) {
	v := picture.NewCT_Picture()
	buf, _ := xml.Marshal(v)
	v2 := picture.NewCT_Picture()
	xml.Unmarshal(buf, v2)
}
