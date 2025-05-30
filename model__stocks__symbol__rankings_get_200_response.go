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

// checks if the StocksSymbolRankingsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StocksSymbolRankingsGet200Response{}

// StocksSymbolRankingsGet200Response struct for StocksSymbolRankingsGet200Response
type StocksSymbolRankingsGet200Response struct {
	BasicInformation *StockRankingsBasicInformation `json:"basic_information,omitempty"`
	GuruFocusRankings *StockRankingsGuruFocusRankings `json:"guru_focus_rankings,omitempty"`
}

// NewStocksSymbolRankingsGet200Response instantiates a new StocksSymbolRankingsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStocksSymbolRankingsGet200Response() *StocksSymbolRankingsGet200Response {
	this := StocksSymbolRankingsGet200Response{}
	return &this
}

// NewStocksSymbolRankingsGet200ResponseWithDefaults instantiates a new StocksSymbolRankingsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStocksSymbolRankingsGet200ResponseWithDefaults() *StocksSymbolRankingsGet200Response {
	this := StocksSymbolRankingsGet200Response{}
	return &this
}

// GetBasicInformation returns the BasicInformation field value if set, zero value otherwise.
func (o *StocksSymbolRankingsGet200Response) GetBasicInformation() StockRankingsBasicInformation {
	if o == nil || IsNil(o.BasicInformation) {
		var ret StockRankingsBasicInformation
		return ret
	}
	return *o.BasicInformation
}

// GetBasicInformationOk returns a tuple with the BasicInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StocksSymbolRankingsGet200Response) GetBasicInformationOk() (*StockRankingsBasicInformation, bool) {
	if o == nil || IsNil(o.BasicInformation) {
		return nil, false
	}
	return o.BasicInformation, true
}

// HasBasicInformation returns a boolean if a field has been set.
func (o *StocksSymbolRankingsGet200Response) HasBasicInformation() bool {
	if o != nil && !IsNil(o.BasicInformation) {
		return true
	}

	return false
}

// SetBasicInformation gets a reference to the given StockRankingsBasicInformation and assigns it to the BasicInformation field.
func (o *StocksSymbolRankingsGet200Response) SetBasicInformation(v StockRankingsBasicInformation) {
	o.BasicInformation = &v
}

// GetGuruFocusRankings returns the GuruFocusRankings field value if set, zero value otherwise.
func (o *StocksSymbolRankingsGet200Response) GetGuruFocusRankings() StockRankingsGuruFocusRankings {
	if o == nil || IsNil(o.GuruFocusRankings) {
		var ret StockRankingsGuruFocusRankings
		return ret
	}
	return *o.GuruFocusRankings
}

// GetGuruFocusRankingsOk returns a tuple with the GuruFocusRankings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StocksSymbolRankingsGet200Response) GetGuruFocusRankingsOk() (*StockRankingsGuruFocusRankings, bool) {
	if o == nil || IsNil(o.GuruFocusRankings) {
		return nil, false
	}
	return o.GuruFocusRankings, true
}

// HasGuruFocusRankings returns a boolean if a field has been set.
func (o *StocksSymbolRankingsGet200Response) HasGuruFocusRankings() bool {
	if o != nil && !IsNil(o.GuruFocusRankings) {
		return true
	}

	return false
}

// SetGuruFocusRankings gets a reference to the given StockRankingsGuruFocusRankings and assigns it to the GuruFocusRankings field.
func (o *StocksSymbolRankingsGet200Response) SetGuruFocusRankings(v StockRankingsGuruFocusRankings) {
	o.GuruFocusRankings = &v
}

func (o StocksSymbolRankingsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StocksSymbolRankingsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BasicInformation) {
		toSerialize["basic_information"] = o.BasicInformation
	}
	if !IsNil(o.GuruFocusRankings) {
		toSerialize["guru_focus_rankings"] = o.GuruFocusRankings
	}
	return toSerialize, nil
}

type NullableStocksSymbolRankingsGet200Response struct {
	value *StocksSymbolRankingsGet200Response
	isSet bool
}

func (v NullableStocksSymbolRankingsGet200Response) Get() *StocksSymbolRankingsGet200Response {
	return v.value
}

func (v *NullableStocksSymbolRankingsGet200Response) Set(val *StocksSymbolRankingsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableStocksSymbolRankingsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableStocksSymbolRankingsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStocksSymbolRankingsGet200Response(val *StocksSymbolRankingsGet200Response) *NullableStocksSymbolRankingsGet200Response {
	return &NullableStocksSymbolRankingsGet200Response{value: val, isSet: true}
}

func (v NullableStocksSymbolRankingsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStocksSymbolRankingsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


