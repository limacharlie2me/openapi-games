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

// checks if the MetadataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataResponse{}

// MetadataResponse struct for MetadataResponse
type MetadataResponse struct {
	// banner
	Banner map[string]interface{} `json:"banner,omitempty"`
}

// NewMetadataResponse instantiates a new MetadataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataResponse() *MetadataResponse {
	this := MetadataResponse{}
	return &this
}

// NewMetadataResponseWithDefaults instantiates a new MetadataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataResponseWithDefaults() *MetadataResponse {
	this := MetadataResponse{}
	return &this
}

// GetBanner returns the Banner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataResponse) GetBanner() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Banner
}

// GetBannerOk returns a tuple with the Banner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataResponse) GetBannerOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Banner) {
		return map[string]interface{}{}, false
	}
	return o.Banner, true
}

// HasBanner returns a boolean if a field has been set.
func (o *MetadataResponse) HasBanner() bool {
	if o != nil && IsNil(o.Banner) {
		return true
	}

	return false
}

// SetBanner gets a reference to the given map[string]interface{} and assigns it to the Banner field.
func (o *MetadataResponse) SetBanner(v map[string]interface{}) {
	o.Banner = v
}

func (o MetadataResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Banner != nil {
		toSerialize["banner"] = o.Banner
	}
	return toSerialize, nil
}

type NullableMetadataResponse struct {
	value *MetadataResponse
	isSet bool
}

func (v NullableMetadataResponse) Get() *MetadataResponse {
	return v.value
}

func (v *NullableMetadataResponse) Set(val *MetadataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataResponse(val *MetadataResponse) *NullableMetadataResponse {
	return &NullableMetadataResponse{value: val, isSet: true}
}

func (v NullableMetadataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


