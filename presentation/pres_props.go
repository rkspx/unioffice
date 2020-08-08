// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// A trial license code for evaluation can be obtained at https://unidoc.io website.

package presentation

import (
	"github.com/unidoc/unioffice/schema/soo/dml"
	"github.com/unidoc/unioffice/schema/soo/pml"
)

// PresentationProperties contains document specific properties.
type PresentationProperties struct {
	x *pml.PresentationPr
}

// NewPresentationProperties constructs a new PresentationProperties.
func NewPresentationProperties() PresentationProperties {
	return PresentationProperties{x: pml.NewPresentationPr()}
}

// X returns the inner wrapped XML type.
func (p PresentationProperties) X() *pml.PresentationPr {
	return p.x
}

// HtmlPubPr returns the HtmlPubPr property.
func (p PresentationProperties) HtmlPubPr() *pml.CT_HtmlPublishProperties {
	return p.x.HtmlPubPr
}

// WebPr returns the WebPr property.
func (p PresentationProperties) WebPr() *pml.CT_WebProperties {
	return p.x.WebPr
}

// PrnPr returns the PrnPr property.
func (p PresentationProperties) PrnPr() *pml.CT_PrintProperties {
	return p.x.PrnPr
}

// ShowPr returns the ShowPr property.
func (p PresentationProperties) ShowPr() *pml.CT_ShowProperties {
	return p.x.ShowPr
}

// ClrMru returns the ClrMru property.
func (p PresentationProperties) ClrMru() *dml.CT_ColorMRU {
	return p.x.ClrMru
}

// ExtLst returns the ExtLst property.
func (p PresentationProperties) ExtLst() *pml.CT_ExtensionList {
	return p.x.ExtLst
}
