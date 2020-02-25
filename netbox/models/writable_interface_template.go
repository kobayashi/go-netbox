// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableInterfaceTemplate writable interface template
// swagger:model WritableInterfaceTemplate
type WritableInterfaceTemplate struct {

	// Device type
	// Required: true
	DeviceType *int64 `json:"device_type"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Management only
	MgmtOnly bool `json:"mgmt_only,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// Type
	// Required: true
	// Enum: [virtual lag 100base-tx 1000base-t 2.5gbase-t 5gbase-t 10gbase-t 10gbase-cx4 1000base-x-gbic 1000base-x-sfp 10gbase-x-sfpp 10gbase-x-xfp 10gbase-x-xenpak 10gbase-x-x2 25gbase-x-sfp28 40gbase-x-qsfpp 50gbase-x-sfp28 100gbase-x-cfp 100gbase-x-cfp2 200gbase-x-cfp2 100gbase-x-cfp4 100gbase-x-cpak 100gbase-x-qsfp28 200gbase-x-qsfp56 400gbase-x-qsfpdd 400gbase-x-osfp ieee802.11a ieee802.11g ieee802.11n ieee802.11ac ieee802.11ad ieee802.11ax gsm cdma lte sonet-oc3 sonet-oc12 sonet-oc48 sonet-oc192 sonet-oc768 sonet-oc1920 sonet-oc3840 1gfc-sfp 2gfc-sfp 4gfc-sfp 8gfc-sfpp 16gfc-sfpp 32gfc-sfp28 128gfc-sfp28 inifiband-sdr inifiband-ddr inifiband-qdr inifiband-fdr10 inifiband-fdr inifiband-edr inifiband-hdr inifiband-ndr inifiband-xdr t1 e1 t3 e3 cisco-stackwise cisco-stackwise-plus cisco-flexstack cisco-flexstack-plus juniper-vcp extreme-summitstack extreme-summitstack-128 extreme-summitstack-256 extreme-summitstack-512 other]
	Type *string `json:"type"`
}

// Validate validates this writable interface template
func (m *WritableInterfaceTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableInterfaceTemplate) validateDeviceType(formats strfmt.Registry) error {

	if err := validate.Required("device_type", "body", m.DeviceType); err != nil {
		return err
	}

	return nil
}

func (m *WritableInterfaceTemplate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 64); err != nil {
		return err
	}

	return nil
}

var writableInterfaceTemplateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["virtual","lag","100base-tx","1000base-t","2.5gbase-t","5gbase-t","10gbase-t","10gbase-cx4","1000base-x-gbic","1000base-x-sfp","10gbase-x-sfpp","10gbase-x-xfp","10gbase-x-xenpak","10gbase-x-x2","25gbase-x-sfp28","40gbase-x-qsfpp","50gbase-x-sfp28","100gbase-x-cfp","100gbase-x-cfp2","200gbase-x-cfp2","100gbase-x-cfp4","100gbase-x-cpak","100gbase-x-qsfp28","200gbase-x-qsfp56","400gbase-x-qsfpdd","400gbase-x-osfp","ieee802.11a","ieee802.11g","ieee802.11n","ieee802.11ac","ieee802.11ad","ieee802.11ax","gsm","cdma","lte","sonet-oc3","sonet-oc12","sonet-oc48","sonet-oc192","sonet-oc768","sonet-oc1920","sonet-oc3840","1gfc-sfp","2gfc-sfp","4gfc-sfp","8gfc-sfpp","16gfc-sfpp","32gfc-sfp28","128gfc-sfp28","inifiband-sdr","inifiband-ddr","inifiband-qdr","inifiband-fdr10","inifiband-fdr","inifiband-edr","inifiband-hdr","inifiband-ndr","inifiband-xdr","t1","e1","t3","e3","cisco-stackwise","cisco-stackwise-plus","cisco-flexstack","cisco-flexstack-plus","juniper-vcp","extreme-summitstack","extreme-summitstack-128","extreme-summitstack-256","extreme-summitstack-512","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableInterfaceTemplateTypeTypePropEnum = append(writableInterfaceTemplateTypeTypePropEnum, v)
	}
}

const (

	// WritableInterfaceTemplateTypeVirtual captures enum value "virtual"
	WritableInterfaceTemplateTypeVirtual string = "virtual"

	// WritableInterfaceTemplateTypeLag captures enum value "lag"
	WritableInterfaceTemplateTypeLag string = "lag"

	// WritableInterfaceTemplateTypeNr100baseTx captures enum value "100base-tx"
	WritableInterfaceTemplateTypeNr100baseTx string = "100base-tx"

	// WritableInterfaceTemplateTypeNr1000baset captures enum value "1000base-t"
	WritableInterfaceTemplateTypeNr1000baset string = "1000base-t"

	// WritableInterfaceTemplateTypeNr25gbaset captures enum value "2.5gbase-t"
	WritableInterfaceTemplateTypeNr25gbaset string = "2.5gbase-t"

	// WritableInterfaceTemplateTypeNr5gbaset captures enum value "5gbase-t"
	WritableInterfaceTemplateTypeNr5gbaset string = "5gbase-t"

	// WritableInterfaceTemplateTypeNr10gbaset captures enum value "10gbase-t"
	WritableInterfaceTemplateTypeNr10gbaset string = "10gbase-t"

	// WritableInterfaceTemplateTypeNr10gbaseCx4 captures enum value "10gbase-cx4"
	WritableInterfaceTemplateTypeNr10gbaseCx4 string = "10gbase-cx4"

	// WritableInterfaceTemplateTypeNr1000basexGbic captures enum value "1000base-x-gbic"
	WritableInterfaceTemplateTypeNr1000basexGbic string = "1000base-x-gbic"

	// WritableInterfaceTemplateTypeNr1000basexSfp captures enum value "1000base-x-sfp"
	WritableInterfaceTemplateTypeNr1000basexSfp string = "1000base-x-sfp"

	// WritableInterfaceTemplateTypeNr10gbasexSfpp captures enum value "10gbase-x-sfpp"
	WritableInterfaceTemplateTypeNr10gbasexSfpp string = "10gbase-x-sfpp"

	// WritableInterfaceTemplateTypeNr10gbasexXfp captures enum value "10gbase-x-xfp"
	WritableInterfaceTemplateTypeNr10gbasexXfp string = "10gbase-x-xfp"

	// WritableInterfaceTemplateTypeNr10gbasexXenpak captures enum value "10gbase-x-xenpak"
	WritableInterfaceTemplateTypeNr10gbasexXenpak string = "10gbase-x-xenpak"

	// WritableInterfaceTemplateTypeNr10gbasexX2 captures enum value "10gbase-x-x2"
	WritableInterfaceTemplateTypeNr10gbasexX2 string = "10gbase-x-x2"

	// WritableInterfaceTemplateTypeNr25gbasexSfp28 captures enum value "25gbase-x-sfp28"
	WritableInterfaceTemplateTypeNr25gbasexSfp28 string = "25gbase-x-sfp28"

	// WritableInterfaceTemplateTypeNr40gbasexQsfpp captures enum value "40gbase-x-qsfpp"
	WritableInterfaceTemplateTypeNr40gbasexQsfpp string = "40gbase-x-qsfpp"

	// WritableInterfaceTemplateTypeNr50gbasexSfp28 captures enum value "50gbase-x-sfp28"
	WritableInterfaceTemplateTypeNr50gbasexSfp28 string = "50gbase-x-sfp28"

	// WritableInterfaceTemplateTypeNr100gbasexCfp captures enum value "100gbase-x-cfp"
	WritableInterfaceTemplateTypeNr100gbasexCfp string = "100gbase-x-cfp"

	// WritableInterfaceTemplateTypeNr100gbasexCfp2 captures enum value "100gbase-x-cfp2"
	WritableInterfaceTemplateTypeNr100gbasexCfp2 string = "100gbase-x-cfp2"

	// WritableInterfaceTemplateTypeNr200gbasexCfp2 captures enum value "200gbase-x-cfp2"
	WritableInterfaceTemplateTypeNr200gbasexCfp2 string = "200gbase-x-cfp2"

	// WritableInterfaceTemplateTypeNr100gbasexCfp4 captures enum value "100gbase-x-cfp4"
	WritableInterfaceTemplateTypeNr100gbasexCfp4 string = "100gbase-x-cfp4"

	// WritableInterfaceTemplateTypeNr100gbasexCpak captures enum value "100gbase-x-cpak"
	WritableInterfaceTemplateTypeNr100gbasexCpak string = "100gbase-x-cpak"

	// WritableInterfaceTemplateTypeNr100gbasexQsfp28 captures enum value "100gbase-x-qsfp28"
	WritableInterfaceTemplateTypeNr100gbasexQsfp28 string = "100gbase-x-qsfp28"

	// WritableInterfaceTemplateTypeNr200gbasexQsfp56 captures enum value "200gbase-x-qsfp56"
	WritableInterfaceTemplateTypeNr200gbasexQsfp56 string = "200gbase-x-qsfp56"

	// WritableInterfaceTemplateTypeNr400gbasexQsfpdd captures enum value "400gbase-x-qsfpdd"
	WritableInterfaceTemplateTypeNr400gbasexQsfpdd string = "400gbase-x-qsfpdd"

	// WritableInterfaceTemplateTypeNr400gbasexOsfp captures enum value "400gbase-x-osfp"
	WritableInterfaceTemplateTypeNr400gbasexOsfp string = "400gbase-x-osfp"

	// WritableInterfaceTemplateTypeIeee80211a captures enum value "ieee802.11a"
	WritableInterfaceTemplateTypeIeee80211a string = "ieee802.11a"

	// WritableInterfaceTemplateTypeIeee80211g captures enum value "ieee802.11g"
	WritableInterfaceTemplateTypeIeee80211g string = "ieee802.11g"

	// WritableInterfaceTemplateTypeIeee80211n captures enum value "ieee802.11n"
	WritableInterfaceTemplateTypeIeee80211n string = "ieee802.11n"

	// WritableInterfaceTemplateTypeIeee80211ac captures enum value "ieee802.11ac"
	WritableInterfaceTemplateTypeIeee80211ac string = "ieee802.11ac"

	// WritableInterfaceTemplateTypeIeee80211ad captures enum value "ieee802.11ad"
	WritableInterfaceTemplateTypeIeee80211ad string = "ieee802.11ad"

	// WritableInterfaceTemplateTypeIeee80211ax captures enum value "ieee802.11ax"
	WritableInterfaceTemplateTypeIeee80211ax string = "ieee802.11ax"

	// WritableInterfaceTemplateTypeGsm captures enum value "gsm"
	WritableInterfaceTemplateTypeGsm string = "gsm"

	// WritableInterfaceTemplateTypeCdma captures enum value "cdma"
	WritableInterfaceTemplateTypeCdma string = "cdma"

	// WritableInterfaceTemplateTypeLte captures enum value "lte"
	WritableInterfaceTemplateTypeLte string = "lte"

	// WritableInterfaceTemplateTypeSonetOc3 captures enum value "sonet-oc3"
	WritableInterfaceTemplateTypeSonetOc3 string = "sonet-oc3"

	// WritableInterfaceTemplateTypeSonetOc12 captures enum value "sonet-oc12"
	WritableInterfaceTemplateTypeSonetOc12 string = "sonet-oc12"

	// WritableInterfaceTemplateTypeSonetOc48 captures enum value "sonet-oc48"
	WritableInterfaceTemplateTypeSonetOc48 string = "sonet-oc48"

	// WritableInterfaceTemplateTypeSonetOc192 captures enum value "sonet-oc192"
	WritableInterfaceTemplateTypeSonetOc192 string = "sonet-oc192"

	// WritableInterfaceTemplateTypeSonetOc768 captures enum value "sonet-oc768"
	WritableInterfaceTemplateTypeSonetOc768 string = "sonet-oc768"

	// WritableInterfaceTemplateTypeSonetOc1920 captures enum value "sonet-oc1920"
	WritableInterfaceTemplateTypeSonetOc1920 string = "sonet-oc1920"

	// WritableInterfaceTemplateTypeSonetOc3840 captures enum value "sonet-oc3840"
	WritableInterfaceTemplateTypeSonetOc3840 string = "sonet-oc3840"

	// WritableInterfaceTemplateTypeNr1gfcSfp captures enum value "1gfc-sfp"
	WritableInterfaceTemplateTypeNr1gfcSfp string = "1gfc-sfp"

	// WritableInterfaceTemplateTypeNr2gfcSfp captures enum value "2gfc-sfp"
	WritableInterfaceTemplateTypeNr2gfcSfp string = "2gfc-sfp"

	// WritableInterfaceTemplateTypeNr4gfcSfp captures enum value "4gfc-sfp"
	WritableInterfaceTemplateTypeNr4gfcSfp string = "4gfc-sfp"

	// WritableInterfaceTemplateTypeNr8gfcSfpp captures enum value "8gfc-sfpp"
	WritableInterfaceTemplateTypeNr8gfcSfpp string = "8gfc-sfpp"

	// WritableInterfaceTemplateTypeNr16gfcSfpp captures enum value "16gfc-sfpp"
	WritableInterfaceTemplateTypeNr16gfcSfpp string = "16gfc-sfpp"

	// WritableInterfaceTemplateTypeNr32gfcSfp28 captures enum value "32gfc-sfp28"
	WritableInterfaceTemplateTypeNr32gfcSfp28 string = "32gfc-sfp28"

	// WritableInterfaceTemplateTypeNr128gfcSfp28 captures enum value "128gfc-sfp28"
	WritableInterfaceTemplateTypeNr128gfcSfp28 string = "128gfc-sfp28"

	// WritableInterfaceTemplateTypeInifibandSdr captures enum value "inifiband-sdr"
	WritableInterfaceTemplateTypeInifibandSdr string = "inifiband-sdr"

	// WritableInterfaceTemplateTypeInifibandDdr captures enum value "inifiband-ddr"
	WritableInterfaceTemplateTypeInifibandDdr string = "inifiband-ddr"

	// WritableInterfaceTemplateTypeInifibandQdr captures enum value "inifiband-qdr"
	WritableInterfaceTemplateTypeInifibandQdr string = "inifiband-qdr"

	// WritableInterfaceTemplateTypeInifibandFdr10 captures enum value "inifiband-fdr10"
	WritableInterfaceTemplateTypeInifibandFdr10 string = "inifiband-fdr10"

	// WritableInterfaceTemplateTypeInifibandFdr captures enum value "inifiband-fdr"
	WritableInterfaceTemplateTypeInifibandFdr string = "inifiband-fdr"

	// WritableInterfaceTemplateTypeInifibandEdr captures enum value "inifiband-edr"
	WritableInterfaceTemplateTypeInifibandEdr string = "inifiband-edr"

	// WritableInterfaceTemplateTypeInifibandHdr captures enum value "inifiband-hdr"
	WritableInterfaceTemplateTypeInifibandHdr string = "inifiband-hdr"

	// WritableInterfaceTemplateTypeInifibandNdr captures enum value "inifiband-ndr"
	WritableInterfaceTemplateTypeInifibandNdr string = "inifiband-ndr"

	// WritableInterfaceTemplateTypeInifibandXdr captures enum value "inifiband-xdr"
	WritableInterfaceTemplateTypeInifibandXdr string = "inifiband-xdr"

	// WritableInterfaceTemplateTypeT1 captures enum value "t1"
	WritableInterfaceTemplateTypeT1 string = "t1"

	// WritableInterfaceTemplateTypeE1 captures enum value "e1"
	WritableInterfaceTemplateTypeE1 string = "e1"

	// WritableInterfaceTemplateTypeT3 captures enum value "t3"
	WritableInterfaceTemplateTypeT3 string = "t3"

	// WritableInterfaceTemplateTypeE3 captures enum value "e3"
	WritableInterfaceTemplateTypeE3 string = "e3"

	// WritableInterfaceTemplateTypeCiscoStackwise captures enum value "cisco-stackwise"
	WritableInterfaceTemplateTypeCiscoStackwise string = "cisco-stackwise"

	// WritableInterfaceTemplateTypeCiscoStackwisePlus captures enum value "cisco-stackwise-plus"
	WritableInterfaceTemplateTypeCiscoStackwisePlus string = "cisco-stackwise-plus"

	// WritableInterfaceTemplateTypeCiscoFlexstack captures enum value "cisco-flexstack"
	WritableInterfaceTemplateTypeCiscoFlexstack string = "cisco-flexstack"

	// WritableInterfaceTemplateTypeCiscoFlexstackPlus captures enum value "cisco-flexstack-plus"
	WritableInterfaceTemplateTypeCiscoFlexstackPlus string = "cisco-flexstack-plus"

	// WritableInterfaceTemplateTypeJuniperVcp captures enum value "juniper-vcp"
	WritableInterfaceTemplateTypeJuniperVcp string = "juniper-vcp"

	// WritableInterfaceTemplateTypeExtremeSummitstack captures enum value "extreme-summitstack"
	WritableInterfaceTemplateTypeExtremeSummitstack string = "extreme-summitstack"

	// WritableInterfaceTemplateTypeExtremeSummitstack128 captures enum value "extreme-summitstack-128"
	WritableInterfaceTemplateTypeExtremeSummitstack128 string = "extreme-summitstack-128"

	// WritableInterfaceTemplateTypeExtremeSummitstack256 captures enum value "extreme-summitstack-256"
	WritableInterfaceTemplateTypeExtremeSummitstack256 string = "extreme-summitstack-256"

	// WritableInterfaceTemplateTypeExtremeSummitstack512 captures enum value "extreme-summitstack-512"
	WritableInterfaceTemplateTypeExtremeSummitstack512 string = "extreme-summitstack-512"

	// WritableInterfaceTemplateTypeOther captures enum value "other"
	WritableInterfaceTemplateTypeOther string = "other"
)

// prop value enum
func (m *WritableInterfaceTemplate) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, writableInterfaceTemplateTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WritableInterfaceTemplate) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableInterfaceTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableInterfaceTemplate) UnmarshalBinary(b []byte) error {
	var res WritableInterfaceTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
