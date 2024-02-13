/*
OpenDota API

# Introduction The OpenDota API provides Dota 2 related data including advanced match data extracted from match replays.  You can find data that can be used to convert hero and ability IDs and other information provided by the API from the [dotaconstants](https://github.com/odota/dotaconstants) repository.  Without a key, you can make 2,000 free calls per day at a rate limit of 60 requests/minute. We also offer a Premium Tier with unlimited API calls and higher rate limits. Check out the [API page](https://www.opendota.com/api-keys) to learn more.     

API version: 25.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PlayerWinLossResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlayerWinLossResponse{}

// PlayerWinLossResponse struct for PlayerWinLossResponse
type PlayerWinLossResponse struct {
	// Number of wins
	Win *int32 `json:"win,omitempty"`
	// Number of loses
	Lose *int32 `json:"lose,omitempty"`
}

// NewPlayerWinLossResponse instantiates a new PlayerWinLossResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlayerWinLossResponse() *PlayerWinLossResponse {
	this := PlayerWinLossResponse{}
	return &this
}

// NewPlayerWinLossResponseWithDefaults instantiates a new PlayerWinLossResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlayerWinLossResponseWithDefaults() *PlayerWinLossResponse {
	this := PlayerWinLossResponse{}
	return &this
}

// GetWin returns the Win field value if set, zero value otherwise.
func (o *PlayerWinLossResponse) GetWin() int32 {
	if o == nil || IsNil(o.Win) {
		var ret int32
		return ret
	}
	return *o.Win
}

// GetWinOk returns a tuple with the Win field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerWinLossResponse) GetWinOk() (*int32, bool) {
	if o == nil || IsNil(o.Win) {
		return nil, false
	}
	return o.Win, true
}

// HasWin returns a boolean if a field has been set.
func (o *PlayerWinLossResponse) HasWin() bool {
	if o != nil && !IsNil(o.Win) {
		return true
	}

	return false
}

// SetWin gets a reference to the given int32 and assigns it to the Win field.
func (o *PlayerWinLossResponse) SetWin(v int32) {
	o.Win = &v
}

// GetLose returns the Lose field value if set, zero value otherwise.
func (o *PlayerWinLossResponse) GetLose() int32 {
	if o == nil || IsNil(o.Lose) {
		var ret int32
		return ret
	}
	return *o.Lose
}

// GetLoseOk returns a tuple with the Lose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerWinLossResponse) GetLoseOk() (*int32, bool) {
	if o == nil || IsNil(o.Lose) {
		return nil, false
	}
	return o.Lose, true
}

// HasLose returns a boolean if a field has been set.
func (o *PlayerWinLossResponse) HasLose() bool {
	if o != nil && !IsNil(o.Lose) {
		return true
	}

	return false
}

// SetLose gets a reference to the given int32 and assigns it to the Lose field.
func (o *PlayerWinLossResponse) SetLose(v int32) {
	o.Lose = &v
}

func (o PlayerWinLossResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlayerWinLossResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Win) {
		toSerialize["win"] = o.Win
	}
	if !IsNil(o.Lose) {
		toSerialize["lose"] = o.Lose
	}
	return toSerialize, nil
}

type NullablePlayerWinLossResponse struct {
	value *PlayerWinLossResponse
	isSet bool
}

func (v NullablePlayerWinLossResponse) Get() *PlayerWinLossResponse {
	return v.value
}

func (v *NullablePlayerWinLossResponse) Set(val *PlayerWinLossResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePlayerWinLossResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePlayerWinLossResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlayerWinLossResponse(val *PlayerWinLossResponse) *NullablePlayerWinLossResponse {
	return &NullablePlayerWinLossResponse{value: val, isSet: true}
}

func (v NullablePlayerWinLossResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlayerWinLossResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

