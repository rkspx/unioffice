// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// commercial license can be purchased via https://unidoc.io website.

package chartDrawing_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/dml/chartDrawing"
)

func TestCT_PictureNonVisualConstructor(t *testing.T) {
	v := chartDrawing.NewCT_PictureNonVisual()
	if v == nil {
		t.Errorf("chartDrawing.NewCT_PictureNonVisual must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chartDrawing.CT_PictureNonVisual should validate: %s", err)
	}
}

func TestCT_PictureNonVisualMarshalUnmarshal(t *testing.T) {
	v := chartDrawing.NewCT_PictureNonVisual()
	buf, _ := xml.Marshal(v)
	v2 := chartDrawing.NewCT_PictureNonVisual()
	xml.Unmarshal(buf, v2)
}
