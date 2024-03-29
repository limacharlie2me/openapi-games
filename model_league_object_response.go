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

// checks if the LeagueObjectResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LeagueObjectResponse{}

// LeagueObjectResponse struct for LeagueObjectResponse
type LeagueObjectResponse struct {
	// leagueid
	Leagueid *int32 `json:"leagueid,omitempty"`
	// ticket
	Ticket *string `json:"ticket,omitempty"`
	// banner
	Banner *string `json:"banner,omitempty"`
	// tier
	Tier *string `json:"tier,omitempty"`
	// League name
	Name *string `json:"name,omitempty"`
}

// NewLeagueObjectResponse instantiates a new LeagueObjectResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLeagueObjectResponse() *LeagueObjectResponse {
	this := LeagueObjectResponse{}
	return &this
}

// NewLeagueObjectResponseWithDefaults instantiates a new LeagueObjectResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLeagueObjectResponseWithDefaults() *LeagueObjectResponse {
	this := LeagueObjectResponse{}
	return &this
}

// GetLeagueid returns the Leagueid field value if set, zero value otherwise.
func (o *LeagueObjectResponse) GetLeagueid() int32 {
	if o == nil || IsNil(o.Leagueid) {
		var ret int32
		return ret
	}
	return *o.Leagueid
}

// GetLeagueidOk returns a tuple with the Leagueid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LeagueObjectResponse) GetLeagueidOk() (*int32, bool) {
	if o == nil || IsNil(o.Leagueid) {
		return nil, false
	}
	return o.Leagueid, true
}

// HasLeagueid returns a boolean if a field has been set.
func (o *LeagueObjectResponse) HasLeagueid() bool {
	if o != nil && !IsNil(o.Leagueid) {
		return true
	}

	return false
}

// SetLeagueid gets a reference to the given int32 and assigns it to the Leagueid field.
func (o *LeagueObjectResponse) SetLeagueid(v int32) {
	o.Leagueid = &v
}

// GetTicket returns the Ticket field value if set, zero value otherwise.
func (o *LeagueObjectResponse) GetTicket() string {
	if o == nil || IsNil(o.Ticket) {
		var ret string
		return ret
	}
	return *o.Ticket
}

// GetTicketOk returns a tuple with the Ticket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LeagueObjectResponse) GetTicketOk() (*string, bool) {
	if o == nil || IsNil(o.Ticket) {
		return nil, false
	}
	return o.Ticket, true
}

// HasTicket returns a boolean if a field has been set.
func (o *LeagueObjectResponse) HasTicket() bool {
	if o != nil && !IsNil(o.Ticket) {
		return true
	}

	return false
}

// SetTicket gets a reference to the given string and assigns it to the Ticket field.
func (o *LeagueObjectResponse) SetTicket(v string) {
	o.Ticket = &v
}

// GetBanner returns the Banner field value if set, zero value otherwise.
func (o *LeagueObjectResponse) GetBanner() string {
	if o == nil || IsNil(o.Banner) {
		var ret string
		return ret
	}
	return *o.Banner
}

// GetBannerOk returns a tuple with the Banner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LeagueObjectResponse) GetBannerOk() (*string, bool) {
	if o == nil || IsNil(o.Banner) {
		return nil, false
	}
	return o.Banner, true
}

// HasBanner returns a boolean if a field has been set.
func (o *LeagueObjectResponse) HasBanner() bool {
	if o != nil && !IsNil(o.Banner) {
		return true
	}

	return false
}

// SetBanner gets a reference to the given string and assigns it to the Banner field.
func (o *LeagueObjectResponse) SetBanner(v string) {
	o.Banner = &v
}

// GetTier returns the Tier field value if set, zero value otherwise.
func (o *LeagueObjectResponse) GetTier() string {
	if o == nil || IsNil(o.Tier) {
		var ret string
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LeagueObjectResponse) GetTierOk() (*string, bool) {
	if o == nil || IsNil(o.Tier) {
		return nil, false
	}
	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *LeagueObjectResponse) HasTier() bool {
	if o != nil && !IsNil(o.Tier) {
		return true
	}

	return false
}

// SetTier gets a reference to the given string and assigns it to the Tier field.
func (o *LeagueObjectResponse) SetTier(v string) {
	o.Tier = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LeagueObjectResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LeagueObjectResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LeagueObjectResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LeagueObjectResponse) SetName(v string) {
	o.Name = &v
}

func (o LeagueObjectResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LeagueObjectResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Leagueid) {
		toSerialize["leagueid"] = o.Leagueid
	}
	if !IsNil(o.Ticket) {
		toSerialize["ticket"] = o.Ticket
	}
	if !IsNil(o.Banner) {
		toSerialize["banner"] = o.Banner
	}
	if !IsNil(o.Tier) {
		toSerialize["tier"] = o.Tier
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableLeagueObjectResponse struct {
	value *LeagueObjectResponse
	isSet bool
}

func (v NullableLeagueObjectResponse) Get() *LeagueObjectResponse {
	return v.value
}

func (v *NullableLeagueObjectResponse) Set(val *LeagueObjectResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLeagueObjectResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLeagueObjectResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLeagueObjectResponse(val *LeagueObjectResponse) *NullableLeagueObjectResponse {
	return &NullableLeagueObjectResponse{value: val, isSet: true}
}

func (v NullableLeagueObjectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLeagueObjectResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


