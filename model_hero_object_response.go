/*
OpenDota API

# Introduction The OpenDota API provides Dota 2 related data including advanced match data extracted from match replays.  You can find data that can be used to convert hero and ability IDs and other information provided by the API from the [dotaconstants](https://github.com/odota/dotaconstants) repository.  Without a key, you can make 2,000 free calls per day at a rate limit of 60 requests/minute. We also offer a Premium Tier with unlimited API calls and higher rate limits. Check out the [API page](https://www.opendota.com/api-keys) to learn more.     

API version: 25.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the HeroObjectResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HeroObjectResponse{}

// HeroObjectResponse struct for HeroObjectResponse
type HeroObjectResponse struct {
	// The ID value of the hero played
	Id int32 `json:"id"`
	// Dota hero command name
	Name *string `json:"name,omitempty"`
	// Hero name
	LocalizedName *string `json:"localized_name,omitempty"`
	// Hero primary shorthand attribute name, e.g. 'agi'
	PrimaryAttr *string `json:"primary_attr,omitempty"`
	// Hero attack type, either 'Melee' or 'Ranged'
	AttackType *string `json:"attack_type,omitempty"`
	Roles []string `json:"roles,omitempty"`
}

type _HeroObjectResponse HeroObjectResponse

// NewHeroObjectResponse instantiates a new HeroObjectResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeroObjectResponse(id int32) *HeroObjectResponse {
	this := HeroObjectResponse{}
	this.Id = id
	return &this
}

// NewHeroObjectResponseWithDefaults instantiates a new HeroObjectResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeroObjectResponseWithDefaults() *HeroObjectResponse {
	this := HeroObjectResponse{}
	return &this
}

// GetId returns the Id field value
func (o *HeroObjectResponse) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HeroObjectResponse) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *HeroObjectResponse) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HeroObjectResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeroObjectResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HeroObjectResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HeroObjectResponse) SetName(v string) {
	o.Name = &v
}

// GetLocalizedName returns the LocalizedName field value if set, zero value otherwise.
func (o *HeroObjectResponse) GetLocalizedName() string {
	if o == nil || IsNil(o.LocalizedName) {
		var ret string
		return ret
	}
	return *o.LocalizedName
}

// GetLocalizedNameOk returns a tuple with the LocalizedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeroObjectResponse) GetLocalizedNameOk() (*string, bool) {
	if o == nil || IsNil(o.LocalizedName) {
		return nil, false
	}
	return o.LocalizedName, true
}

// HasLocalizedName returns a boolean if a field has been set.
func (o *HeroObjectResponse) HasLocalizedName() bool {
	if o != nil && !IsNil(o.LocalizedName) {
		return true
	}

	return false
}

// SetLocalizedName gets a reference to the given string and assigns it to the LocalizedName field.
func (o *HeroObjectResponse) SetLocalizedName(v string) {
	o.LocalizedName = &v
}

// GetPrimaryAttr returns the PrimaryAttr field value if set, zero value otherwise.
func (o *HeroObjectResponse) GetPrimaryAttr() string {
	if o == nil || IsNil(o.PrimaryAttr) {
		var ret string
		return ret
	}
	return *o.PrimaryAttr
}

// GetPrimaryAttrOk returns a tuple with the PrimaryAttr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeroObjectResponse) GetPrimaryAttrOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryAttr) {
		return nil, false
	}
	return o.PrimaryAttr, true
}

// HasPrimaryAttr returns a boolean if a field has been set.
func (o *HeroObjectResponse) HasPrimaryAttr() bool {
	if o != nil && !IsNil(o.PrimaryAttr) {
		return true
	}

	return false
}

// SetPrimaryAttr gets a reference to the given string and assigns it to the PrimaryAttr field.
func (o *HeroObjectResponse) SetPrimaryAttr(v string) {
	o.PrimaryAttr = &v
}

// GetAttackType returns the AttackType field value if set, zero value otherwise.
func (o *HeroObjectResponse) GetAttackType() string {
	if o == nil || IsNil(o.AttackType) {
		var ret string
		return ret
	}
	return *o.AttackType
}

// GetAttackTypeOk returns a tuple with the AttackType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeroObjectResponse) GetAttackTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AttackType) {
		return nil, false
	}
	return o.AttackType, true
}

// HasAttackType returns a boolean if a field has been set.
func (o *HeroObjectResponse) HasAttackType() bool {
	if o != nil && !IsNil(o.AttackType) {
		return true
	}

	return false
}

// SetAttackType gets a reference to the given string and assigns it to the AttackType field.
func (o *HeroObjectResponse) SetAttackType(v string) {
	o.AttackType = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *HeroObjectResponse) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeroObjectResponse) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *HeroObjectResponse) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *HeroObjectResponse) SetRoles(v []string) {
	o.Roles = v
}

func (o HeroObjectResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HeroObjectResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.LocalizedName) {
		toSerialize["localized_name"] = o.LocalizedName
	}
	if !IsNil(o.PrimaryAttr) {
		toSerialize["primary_attr"] = o.PrimaryAttr
	}
	if !IsNil(o.AttackType) {
		toSerialize["attack_type"] = o.AttackType
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

func (o *HeroObjectResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varHeroObjectResponse := _HeroObjectResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHeroObjectResponse)

	if err != nil {
		return err
	}

	*o = HeroObjectResponse(varHeroObjectResponse)

	return err
}

type NullableHeroObjectResponse struct {
	value *HeroObjectResponse
	isSet bool
}

func (v NullableHeroObjectResponse) Get() *HeroObjectResponse {
	return v.value
}

func (v *NullableHeroObjectResponse) Set(val *HeroObjectResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHeroObjectResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHeroObjectResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeroObjectResponse(val *HeroObjectResponse) *NullableHeroObjectResponse {
	return &NullableHeroObjectResponse{value: val, isSet: true}
}

func (v NullableHeroObjectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeroObjectResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


