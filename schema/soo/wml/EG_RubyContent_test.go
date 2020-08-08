// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// commercial license can be purchased via https://unidoc.io website.

package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/wml"
)

func TestEG_RubyContentConstructor(t *testing.T) {
	v := wml.NewEG_RubyContent()
	if v == nil {
		t.Errorf("wml.NewEG_RubyContent must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_RubyContent should validate: %s", err)
	}
}

func TestEG_RubyContentMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_RubyContent()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_RubyContent()
	xml.Unmarshal(buf, v2)
}
