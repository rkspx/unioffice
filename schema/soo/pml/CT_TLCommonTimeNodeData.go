// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by unioffice ECMA-376 generator
//
// Use of this software package and source code is governed by the terms of the
// UniDoc End User License Agreement (EULA) that is available at:
// https://unidoc.io/eula/
// commercial license can be purchased via https://unidoc.io website.

package pml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/schema/soo/dml"
)

type CT_TLCommonTimeNodeData struct {
	// ID
	IdAttr *uint32
	// Preset ID
	PresetIDAttr *int32
	// Preset Types
	PresetClassAttr ST_TLTimeNodePresetClassType
	// Preset SubType
	PresetSubtypeAttr *int32
	// Duration
	DurAttr *ST_TLTime
	// Repeat Count
	RepeatCountAttr *ST_TLTime
	// Repeat Duration
	RepeatDurAttr *ST_TLTime
	// Speed
	SpdAttr *dml.ST_Percentage
	// Acceleration
	AccelAttr *dml.ST_PositiveFixedPercentage
	// Deceleration
	DecelAttr *dml.ST_PositiveFixedPercentage
	// Auto Reverse
	AutoRevAttr *bool
	// Restart
	RestartAttr ST_TLTimeNodeRestartType
	// Fill
	FillAttr ST_TLTimeNodeFillType
	// Synchronization Behavior
	SyncBehaviorAttr ST_TLTimeNodeSyncType
	// Time Filter
	TmFilterAttr *string
	// Event Filter
	EvtFilterAttr *string
	// Display
	DisplayAttr *bool
	// Master Relation
	MasterRelAttr ST_TLTimeNodeMasterRelation
	// Build level
	BldLvlAttr *int32
	// Group ID
	GrpIdAttr *uint32
	// After Effect
	AfterEffectAttr *bool
	// Node Type
	NodeTypeAttr ST_TLTimeNodeType
	// Node Placeholder
	NodePhAttr *bool
	// Start Conditions List
	StCondLst *CT_TLTimeConditionList
	// End Conditions List
	EndCondLst *CT_TLTimeConditionList
	// EndSync
	EndSync *CT_TLTimeCondition
	// Iterate
	Iterate *CT_TLIterateData
	// Children Time Node List
	ChildTnLst *CT_TimeNodeList
	// Sub-TimeNodes List
	SubTnLst *CT_TimeNodeList
}

func NewCT_TLCommonTimeNodeData() *CT_TLCommonTimeNodeData {
	ret := &CT_TLCommonTimeNodeData{}
	return ret
}

func (m *CT_TLCommonTimeNodeData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	if m.PresetIDAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "presetID"},
			Value: fmt.Sprintf("%v", *m.PresetIDAttr)})
	}
	if m.PresetClassAttr != ST_TLTimeNodePresetClassTypeUnset {
		attr, err := m.PresetClassAttr.MarshalXMLAttr(xml.Name{Local: "presetClass"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.PresetSubtypeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "presetSubtype"},
			Value: fmt.Sprintf("%v", *m.PresetSubtypeAttr)})
	}
	if m.DurAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dur"},
			Value: fmt.Sprintf("%v", *m.DurAttr)})
	}
	if m.RepeatCountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "repeatCount"},
			Value: fmt.Sprintf("%v", *m.RepeatCountAttr)})
	}
	if m.RepeatDurAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "repeatDur"},
			Value: fmt.Sprintf("%v", *m.RepeatDurAttr)})
	}
	if m.SpdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "spd"},
			Value: fmt.Sprintf("%v", *m.SpdAttr)})
	}
	if m.AccelAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "accel"},
			Value: fmt.Sprintf("%v", *m.AccelAttr)})
	}
	if m.DecelAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "decel"},
			Value: fmt.Sprintf("%v", *m.DecelAttr)})
	}
	if m.AutoRevAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "autoRev"},
			Value: fmt.Sprintf("%d", b2i(*m.AutoRevAttr))})
	}
	if m.RestartAttr != ST_TLTimeNodeRestartTypeUnset {
		attr, err := m.RestartAttr.MarshalXMLAttr(xml.Name{Local: "restart"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.FillAttr != ST_TLTimeNodeFillTypeUnset {
		attr, err := m.FillAttr.MarshalXMLAttr(xml.Name{Local: "fill"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.SyncBehaviorAttr != ST_TLTimeNodeSyncTypeUnset {
		attr, err := m.SyncBehaviorAttr.MarshalXMLAttr(xml.Name{Local: "syncBehavior"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.TmFilterAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "tmFilter"},
			Value: fmt.Sprintf("%v", *m.TmFilterAttr)})
	}
	if m.EvtFilterAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "evtFilter"},
			Value: fmt.Sprintf("%v", *m.EvtFilterAttr)})
	}
	if m.DisplayAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "display"},
			Value: fmt.Sprintf("%d", b2i(*m.DisplayAttr))})
	}
	if m.MasterRelAttr != ST_TLTimeNodeMasterRelationUnset {
		attr, err := m.MasterRelAttr.MarshalXMLAttr(xml.Name{Local: "masterRel"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.BldLvlAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "bldLvl"},
			Value: fmt.Sprintf("%v", *m.BldLvlAttr)})
	}
	if m.GrpIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "grpId"},
			Value: fmt.Sprintf("%v", *m.GrpIdAttr)})
	}
	if m.AfterEffectAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "afterEffect"},
			Value: fmt.Sprintf("%d", b2i(*m.AfterEffectAttr))})
	}
	if m.NodeTypeAttr != ST_TLTimeNodeTypeUnset {
		attr, err := m.NodeTypeAttr.MarshalXMLAttr(xml.Name{Local: "nodeType"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.NodePhAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "nodePh"},
			Value: fmt.Sprintf("%d", b2i(*m.NodePhAttr))})
	}
	e.EncodeToken(start)
	if m.StCondLst != nil {
		sestCondLst := xml.StartElement{Name: xml.Name{Local: "p:stCondLst"}}
		e.EncodeElement(m.StCondLst, sestCondLst)
	}
	if m.EndCondLst != nil {
		seendCondLst := xml.StartElement{Name: xml.Name{Local: "p:endCondLst"}}
		e.EncodeElement(m.EndCondLst, seendCondLst)
	}
	if m.EndSync != nil {
		seendSync := xml.StartElement{Name: xml.Name{Local: "p:endSync"}}
		e.EncodeElement(m.EndSync, seendSync)
	}
	if m.Iterate != nil {
		seiterate := xml.StartElement{Name: xml.Name{Local: "p:iterate"}}
		e.EncodeElement(m.Iterate, seiterate)
	}
	if m.ChildTnLst != nil {
		sechildTnLst := xml.StartElement{Name: xml.Name{Local: "p:childTnLst"}}
		e.EncodeElement(m.ChildTnLst, sechildTnLst)
	}
	if m.SubTnLst != nil {
		sesubTnLst := xml.StartElement{Name: xml.Name{Local: "p:subTnLst"}}
		e.EncodeElement(m.SubTnLst, sesubTnLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLCommonTimeNodeData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "tmFilter" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TmFilterAttr = &parsed
			continue
		}
		if attr.Name.Local == "evtFilter" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.EvtFilterAttr = &parsed
			continue
		}
		if attr.Name.Local == "presetID" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.PresetIDAttr = &pt
			continue
		}
		if attr.Name.Local == "display" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.DisplayAttr = &parsed
			continue
		}
		if attr.Name.Local == "presetSubtype" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.PresetSubtypeAttr = &pt
			continue
		}
		if attr.Name.Local == "masterRel" {
			m.MasterRelAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "repeatCount" {
			parsed, err := ParseUnionST_TLTime(attr.Value)
			if err != nil {
				return err
			}
			m.RepeatCountAttr = &parsed
			continue
		}
		if attr.Name.Local == "bldLvl" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.BldLvlAttr = &pt
			continue
		}
		if attr.Name.Local == "spd" {
			parsed, err := ParseUnionST_Percentage(attr.Value)
			if err != nil {
				return err
			}
			m.SpdAttr = &parsed
			continue
		}
		if attr.Name.Local == "grpId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.GrpIdAttr = &pt
			continue
		}
		if attr.Name.Local == "afterEffect" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AfterEffectAttr = &parsed
			continue
		}
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.IdAttr = &pt
			continue
		}
		if attr.Name.Local == "nodePh" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NodePhAttr = &parsed
			continue
		}
		if attr.Name.Local == "accel" {
			parsed, err := ParseUnionST_PositiveFixedPercentage(attr.Value)
			if err != nil {
				return err
			}
			m.AccelAttr = &parsed
			continue
		}
		if attr.Name.Local == "decel" {
			parsed, err := ParseUnionST_PositiveFixedPercentage(attr.Value)
			if err != nil {
				return err
			}
			m.DecelAttr = &parsed
			continue
		}
		if attr.Name.Local == "restart" {
			m.RestartAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "nodeType" {
			m.NodeTypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "presetClass" {
			m.PresetClassAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "repeatDur" {
			parsed, err := ParseUnionST_TLTime(attr.Value)
			if err != nil {
				return err
			}
			m.RepeatDurAttr = &parsed
			continue
		}
		if attr.Name.Local == "syncBehavior" {
			m.SyncBehaviorAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "autoRev" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AutoRevAttr = &parsed
			continue
		}
		if attr.Name.Local == "fill" {
			m.FillAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "dur" {
			parsed, err := ParseUnionST_TLTime(attr.Value)
			if err != nil {
				return err
			}
			m.DurAttr = &parsed
			continue
		}
	}
lCT_TLCommonTimeNodeData:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "stCondLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "stCondLst"}:
				m.StCondLst = NewCT_TLTimeConditionList()
				if err := d.DecodeElement(m.StCondLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "endCondLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "endCondLst"}:
				m.EndCondLst = NewCT_TLTimeConditionList()
				if err := d.DecodeElement(m.EndCondLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "endSync"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "endSync"}:
				m.EndSync = NewCT_TLTimeCondition()
				if err := d.DecodeElement(m.EndSync, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "iterate"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "iterate"}:
				m.Iterate = NewCT_TLIterateData()
				if err := d.DecodeElement(m.Iterate, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "childTnLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "childTnLst"}:
				m.ChildTnLst = NewCT_TimeNodeList()
				if err := d.DecodeElement(m.ChildTnLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "subTnLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "subTnLst"}:
				m.SubTnLst = NewCT_TimeNodeList()
				if err := d.DecodeElement(m.SubTnLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_TLCommonTimeNodeData %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLCommonTimeNodeData
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLCommonTimeNodeData and its children
func (m *CT_TLCommonTimeNodeData) Validate() error {
	return m.ValidateWithPath("CT_TLCommonTimeNodeData")
}

// ValidateWithPath validates the CT_TLCommonTimeNodeData and its children, prefixing error messages with path
func (m *CT_TLCommonTimeNodeData) ValidateWithPath(path string) error {
	if err := m.PresetClassAttr.ValidateWithPath(path + "/PresetClassAttr"); err != nil {
		return err
	}
	if m.DurAttr != nil {
		if err := m.DurAttr.ValidateWithPath(path + "/DurAttr"); err != nil {
			return err
		}
	}
	if m.RepeatCountAttr != nil {
		if err := m.RepeatCountAttr.ValidateWithPath(path + "/RepeatCountAttr"); err != nil {
			return err
		}
	}
	if m.RepeatDurAttr != nil {
		if err := m.RepeatDurAttr.ValidateWithPath(path + "/RepeatDurAttr"); err != nil {
			return err
		}
	}
	if m.SpdAttr != nil {
		if err := m.SpdAttr.ValidateWithPath(path + "/SpdAttr"); err != nil {
			return err
		}
	}
	if m.AccelAttr != nil {
		if err := m.AccelAttr.ValidateWithPath(path + "/AccelAttr"); err != nil {
			return err
		}
	}
	if m.DecelAttr != nil {
		if err := m.DecelAttr.ValidateWithPath(path + "/DecelAttr"); err != nil {
			return err
		}
	}
	if err := m.RestartAttr.ValidateWithPath(path + "/RestartAttr"); err != nil {
		return err
	}
	if err := m.FillAttr.ValidateWithPath(path + "/FillAttr"); err != nil {
		return err
	}
	if err := m.SyncBehaviorAttr.ValidateWithPath(path + "/SyncBehaviorAttr"); err != nil {
		return err
	}
	if err := m.MasterRelAttr.ValidateWithPath(path + "/MasterRelAttr"); err != nil {
		return err
	}
	if err := m.NodeTypeAttr.ValidateWithPath(path + "/NodeTypeAttr"); err != nil {
		return err
	}
	if m.StCondLst != nil {
		if err := m.StCondLst.ValidateWithPath(path + "/StCondLst"); err != nil {
			return err
		}
	}
	if m.EndCondLst != nil {
		if err := m.EndCondLst.ValidateWithPath(path + "/EndCondLst"); err != nil {
			return err
		}
	}
	if m.EndSync != nil {
		if err := m.EndSync.ValidateWithPath(path + "/EndSync"); err != nil {
			return err
		}
	}
	if m.Iterate != nil {
		if err := m.Iterate.ValidateWithPath(path + "/Iterate"); err != nil {
			return err
		}
	}
	if m.ChildTnLst != nil {
		if err := m.ChildTnLst.ValidateWithPath(path + "/ChildTnLst"); err != nil {
			return err
		}
	}
	if m.SubTnLst != nil {
		if err := m.SubTnLst.ValidateWithPath(path + "/SubTnLst"); err != nil {
			return err
		}
	}
	return nil
}
