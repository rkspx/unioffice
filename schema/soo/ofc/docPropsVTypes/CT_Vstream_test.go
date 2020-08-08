// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// commercial license can be purchased via https://unidoc.io website.

package docPropsVTypes_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/ofc/docPropsVTypes"
)

func TestCT_VstreamConstructor(t *testing.T) {
	v := docPropsVTypes.NewCT_Vstream()
	if v == nil {
		t.Errorf("docPropsVTypes.NewCT_Vstream must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed docPropsVTypes.CT_Vstream should validate: %s", err)
	}
}

func TestCT_VstreamMarshalUnmarshal(t *testing.T) {
	v := docPropsVTypes.NewCT_Vstream()
	buf, _ := xml.Marshal(v)
	v2 := docPropsVTypes.NewCT_Vstream()
	xml.Unmarshal(buf, v2)
}
