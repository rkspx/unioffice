// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// commercial license can be purchased via https://unidoc.io website.

package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/sml"
)

func TestCT_RevisionConflictConstructor(t *testing.T) {
	v := sml.NewCT_RevisionConflict()
	if v == nil {
		t.Errorf("sml.NewCT_RevisionConflict must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_RevisionConflict should validate: %s", err)
	}
}

func TestCT_RevisionConflictMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_RevisionConflict()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_RevisionConflict()
	xml.Unmarshal(buf, v2)
}
