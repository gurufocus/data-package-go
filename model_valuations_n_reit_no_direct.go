/*
Gurufocus Data Package API

API for accessing Gurufocus data packages, please go to [https://www.gurufocus.com/user/me?tab=account&subtab=api-token](https://www.gurufocus.com/user/me?tab=account&subtab=api-token) to view or generate authorization keys.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ValuationsNREITNODIRECT type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValuationsNREITNODIRECT{}

// ValuationsNREITNODIRECT struct for ValuationsNREITNODIRECT
type ValuationsNREITNODIRECT struct {
	BasicInformation *ValuationsNREITNODIRECTBasicInformation `json:"basic_information,omitempty"`
	PerShareData *ValuationsNREITNODIRECTPerShareData `json:"per_share_data,omitempty"`
	Ratios *ValuationsNREITNODIRECTRatios `json:"ratios,omitempty"`
	ValuationRatios *ValuationsNREITNODIRECTValuationRatios `json:"valuation_ratios,omitempty"`
	ValuationandQuality *ValuationsNREITNODIRECTValuationandQuality `json:"valuationand_quality,omitempty"`
}

// NewValuationsNREITNODIRECT instantiates a new ValuationsNREITNODIRECT object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValuationsNREITNODIRECT() *ValuationsNREITNODIRECT {
	this := ValuationsNREITNODIRECT{}
	return &this
}

// NewValuationsNREITNODIRECTWithDefaults instantiates a new ValuationsNREITNODIRECT object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValuationsNREITNODIRECTWithDefaults() *ValuationsNREITNODIRECT {
	this := ValuationsNREITNODIRECT{}
	return &this
}

// GetBasicInformation returns the BasicInformation field value if set, zero value otherwise.
func (o *ValuationsNREITNODIRECT) GetBasicInformation() ValuationsNREITNODIRECTBasicInformation {
	if o == nil || IsNil(o.BasicInformation) {
		var ret ValuationsNREITNODIRECTBasicInformation
		return ret
	}
	return *o.BasicInformation
}

// GetBasicInformationOk returns a tuple with the BasicInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsNREITNODIRECT) GetBasicInformationOk() (*ValuationsNREITNODIRECTBasicInformation, bool) {
	if o == nil || IsNil(o.BasicInformation) {
		return nil, false
	}
	return o.BasicInformation, true
}

// HasBasicInformation returns a boolean if a field has been set.
func (o *ValuationsNREITNODIRECT) HasBasicInformation() bool {
	if o != nil && !IsNil(o.BasicInformation) {
		return true
	}

	return false
}

// SetBasicInformation gets a reference to the given ValuationsNREITNODIRECTBasicInformation and assigns it to the BasicInformation field.
func (o *ValuationsNREITNODIRECT) SetBasicInformation(v ValuationsNREITNODIRECTBasicInformation) {
	o.BasicInformation = &v
}

// GetPerShareData returns the PerShareData field value if set, zero value otherwise.
func (o *ValuationsNREITNODIRECT) GetPerShareData() ValuationsNREITNODIRECTPerShareData {
	if o == nil || IsNil(o.PerShareData) {
		var ret ValuationsNREITNODIRECTPerShareData
		return ret
	}
	return *o.PerShareData
}

// GetPerShareDataOk returns a tuple with the PerShareData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsNREITNODIRECT) GetPerShareDataOk() (*ValuationsNREITNODIRECTPerShareData, bool) {
	if o == nil || IsNil(o.PerShareData) {
		return nil, false
	}
	return o.PerShareData, true
}

// HasPerShareData returns a boolean if a field has been set.
func (o *ValuationsNREITNODIRECT) HasPerShareData() bool {
	if o != nil && !IsNil(o.PerShareData) {
		return true
	}

	return false
}

// SetPerShareData gets a reference to the given ValuationsNREITNODIRECTPerShareData and assigns it to the PerShareData field.
func (o *ValuationsNREITNODIRECT) SetPerShareData(v ValuationsNREITNODIRECTPerShareData) {
	o.PerShareData = &v
}

// GetRatios returns the Ratios field value if set, zero value otherwise.
func (o *ValuationsNREITNODIRECT) GetRatios() ValuationsNREITNODIRECTRatios {
	if o == nil || IsNil(o.Ratios) {
		var ret ValuationsNREITNODIRECTRatios
		return ret
	}
	return *o.Ratios
}

// GetRatiosOk returns a tuple with the Ratios field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsNREITNODIRECT) GetRatiosOk() (*ValuationsNREITNODIRECTRatios, bool) {
	if o == nil || IsNil(o.Ratios) {
		return nil, false
	}
	return o.Ratios, true
}

// HasRatios returns a boolean if a field has been set.
func (o *ValuationsNREITNODIRECT) HasRatios() bool {
	if o != nil && !IsNil(o.Ratios) {
		return true
	}

	return false
}

// SetRatios gets a reference to the given ValuationsNREITNODIRECTRatios and assigns it to the Ratios field.
func (o *ValuationsNREITNODIRECT) SetRatios(v ValuationsNREITNODIRECTRatios) {
	o.Ratios = &v
}

// GetValuationRatios returns the ValuationRatios field value if set, zero value otherwise.
func (o *ValuationsNREITNODIRECT) GetValuationRatios() ValuationsNREITNODIRECTValuationRatios {
	if o == nil || IsNil(o.ValuationRatios) {
		var ret ValuationsNREITNODIRECTValuationRatios
		return ret
	}
	return *o.ValuationRatios
}

// GetValuationRatiosOk returns a tuple with the ValuationRatios field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsNREITNODIRECT) GetValuationRatiosOk() (*ValuationsNREITNODIRECTValuationRatios, bool) {
	if o == nil || IsNil(o.ValuationRatios) {
		return nil, false
	}
	return o.ValuationRatios, true
}

// HasValuationRatios returns a boolean if a field has been set.
func (o *ValuationsNREITNODIRECT) HasValuationRatios() bool {
	if o != nil && !IsNil(o.ValuationRatios) {
		return true
	}

	return false
}

// SetValuationRatios gets a reference to the given ValuationsNREITNODIRECTValuationRatios and assigns it to the ValuationRatios field.
func (o *ValuationsNREITNODIRECT) SetValuationRatios(v ValuationsNREITNODIRECTValuationRatios) {
	o.ValuationRatios = &v
}

// GetValuationandQuality returns the ValuationandQuality field value if set, zero value otherwise.
func (o *ValuationsNREITNODIRECT) GetValuationandQuality() ValuationsNREITNODIRECTValuationandQuality {
	if o == nil || IsNil(o.ValuationandQuality) {
		var ret ValuationsNREITNODIRECTValuationandQuality
		return ret
	}
	return *o.ValuationandQuality
}

// GetValuationandQualityOk returns a tuple with the ValuationandQuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuationsNREITNODIRECT) GetValuationandQualityOk() (*ValuationsNREITNODIRECTValuationandQuality, bool) {
	if o == nil || IsNil(o.ValuationandQuality) {
		return nil, false
	}
	return o.ValuationandQuality, true
}

// HasValuationandQuality returns a boolean if a field has been set.
func (o *ValuationsNREITNODIRECT) HasValuationandQuality() bool {
	if o != nil && !IsNil(o.ValuationandQuality) {
		return true
	}

	return false
}

// SetValuationandQuality gets a reference to the given ValuationsNREITNODIRECTValuationandQuality and assigns it to the ValuationandQuality field.
func (o *ValuationsNREITNODIRECT) SetValuationandQuality(v ValuationsNREITNODIRECTValuationandQuality) {
	o.ValuationandQuality = &v
}

func (o ValuationsNREITNODIRECT) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValuationsNREITNODIRECT) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BasicInformation) {
		toSerialize["basic_information"] = o.BasicInformation
	}
	if !IsNil(o.PerShareData) {
		toSerialize["per_share_data"] = o.PerShareData
	}
	if !IsNil(o.Ratios) {
		toSerialize["ratios"] = o.Ratios
	}
	if !IsNil(o.ValuationRatios) {
		toSerialize["valuation_ratios"] = o.ValuationRatios
	}
	if !IsNil(o.ValuationandQuality) {
		toSerialize["valuationand_quality"] = o.ValuationandQuality
	}
	return toSerialize, nil
}

type NullableValuationsNREITNODIRECT struct {
	value *ValuationsNREITNODIRECT
	isSet bool
}

func (v NullableValuationsNREITNODIRECT) Get() *ValuationsNREITNODIRECT {
	return v.value
}

func (v *NullableValuationsNREITNODIRECT) Set(val *ValuationsNREITNODIRECT) {
	v.value = val
	v.isSet = true
}

func (v NullableValuationsNREITNODIRECT) IsSet() bool {
	return v.isSet
}

func (v *NullableValuationsNREITNODIRECT) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValuationsNREITNODIRECT(val *ValuationsNREITNODIRECT) *NullableValuationsNREITNODIRECT {
	return &NullableValuationsNREITNODIRECT{value: val, isSet: true}
}

func (v NullableValuationsNREITNODIRECT) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValuationsNREITNODIRECT) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


