// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package diagram_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/dml/diagram"
)

func TestCT_SampleDataConstructor(t *testing.T) {
	v := diagram.NewCT_SampleData()
	if v == nil {
		t.Errorf("diagram.NewCT_SampleData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_SampleData should validate: %s", err)
	}
}

func TestCT_SampleDataMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_SampleData()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_SampleData()
	xml.Unmarshal(buf, v2)
}
