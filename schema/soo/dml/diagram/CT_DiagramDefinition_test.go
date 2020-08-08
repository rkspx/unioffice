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

func TestCT_DiagramDefinitionConstructor(t *testing.T) {
	v := diagram.NewCT_DiagramDefinition()
	if v == nil {
		t.Errorf("diagram.NewCT_DiagramDefinition must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_DiagramDefinition should validate: %s", err)
	}
}

func TestCT_DiagramDefinitionMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_DiagramDefinition()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_DiagramDefinition()
	xml.Unmarshal(buf, v2)
}
