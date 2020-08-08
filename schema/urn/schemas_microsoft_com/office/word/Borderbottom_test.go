// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// commercial license can be purchased via https://unidoc.io website.

package word_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/urn/schemas_microsoft_com/office/word"
)

func TestBorderbottomConstructor(t *testing.T) {
	v := word.NewBorderbottom()
	if v == nil {
		t.Errorf("word.NewBorderbottom must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed word.Borderbottom should validate: %s", err)
	}
}

func TestBorderbottomMarshalUnmarshal(t *testing.T) {
	v := word.NewBorderbottom()
	buf, _ := xml.Marshal(v)
	v2 := word.NewBorderbottom()
	xml.Unmarshal(buf, v2)
}
