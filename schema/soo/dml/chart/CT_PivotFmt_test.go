// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// commercial license can be purchased via https://unidoc.io website.

package chart_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/dml/chart"
)

func TestCT_PivotFmtConstructor(t *testing.T) {
	v := chart.NewCT_PivotFmt()
	if v == nil {
		t.Errorf("chart.NewCT_PivotFmt must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.CT_PivotFmt should validate: %s", err)
	}
}

func TestCT_PivotFmtMarshalUnmarshal(t *testing.T) {
	v := chart.NewCT_PivotFmt()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewCT_PivotFmt()
	xml.Unmarshal(buf, v2)
}
