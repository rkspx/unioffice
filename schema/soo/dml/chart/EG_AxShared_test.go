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

func TestEG_AxSharedConstructor(t *testing.T) {
	v := chart.NewEG_AxShared()
	if v == nil {
		t.Errorf("chart.NewEG_AxShared must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed chart.EG_AxShared should validate: %s", err)
	}
}

func TestEG_AxSharedMarshalUnmarshal(t *testing.T) {
	v := chart.NewEG_AxShared()
	buf, _ := xml.Marshal(v)
	v2 := chart.NewEG_AxShared()
	xml.Unmarshal(buf, v2)
}
