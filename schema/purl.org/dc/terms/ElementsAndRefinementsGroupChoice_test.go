// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// commercial license can be purchased via https://unidoc.io website.

package terms_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/purl.org/dc/terms"
)

func TestElementsAndRefinementsGroupChoiceConstructor(t *testing.T) {
	v := terms.NewElementsAndRefinementsGroupChoice()
	if v == nil {
		t.Errorf("terms.NewElementsAndRefinementsGroupChoice must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.ElementsAndRefinementsGroupChoice should validate: %s", err)
	}
}

func TestElementsAndRefinementsGroupChoiceMarshalUnmarshal(t *testing.T) {
	v := terms.NewElementsAndRefinementsGroupChoice()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewElementsAndRefinementsGroupChoice()
	xml.Unmarshal(buf, v2)
}
