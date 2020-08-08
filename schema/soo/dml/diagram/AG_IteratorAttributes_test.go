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

func TestAG_IteratorAttributesConstructor(t *testing.T) {
	v := diagram.NewAG_IteratorAttributes()
	if v == nil {
		t.Errorf("diagram.NewAG_IteratorAttributes must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.AG_IteratorAttributes should validate: %s", err)
	}
}

func TestAG_IteratorAttributesMarshalUnmarshal(t *testing.T) {
	v := diagram.NewAG_IteratorAttributes()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewAG_IteratorAttributes()
	xml.Unmarshal(buf, v2)
}
