// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// commercial license can be purchased via https://unidoc.io website.

package diagram_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/dml/diagram"
)

func TestCT_CTDescriptionConstructor(t *testing.T) {
	v := diagram.NewCT_CTDescription()
	if v == nil {
		t.Errorf("diagram.NewCT_CTDescription must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_CTDescription should validate: %s", err)
	}
}

func TestCT_CTDescriptionMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_CTDescription()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_CTDescription()
	xml.Unmarshal(buf, v2)
}
