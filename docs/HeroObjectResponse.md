# HeroObjectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID value of the hero played | 
**Name** | Pointer to **string** | Dota hero command name | [optional] 
**LocalizedName** | Pointer to **string** | Hero name | [optional] 
**PrimaryAttr** | Pointer to **string** | Hero primary shorthand attribute name, e.g. &#39;agi&#39; | [optional] 
**AttackType** | Pointer to **string** | Hero attack type, either &#39;Melee&#39; or &#39;Ranged&#39; | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 

## Methods

### NewHeroObjectResponse

`func NewHeroObjectResponse(id int32, ) *HeroObjectResponse`

NewHeroObjectResponse instantiates a new HeroObjectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeroObjectResponseWithDefaults

`func NewHeroObjectResponseWithDefaults() *HeroObjectResponse`

NewHeroObjectResponseWithDefaults instantiates a new HeroObjectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HeroObjectResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HeroObjectResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HeroObjectResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *HeroObjectResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HeroObjectResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HeroObjectResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HeroObjectResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocalizedName

`func (o *HeroObjectResponse) GetLocalizedName() string`

GetLocalizedName returns the LocalizedName field if non-nil, zero value otherwise.

### GetLocalizedNameOk

`func (o *HeroObjectResponse) GetLocalizedNameOk() (*string, bool)`

GetLocalizedNameOk returns a tuple with the LocalizedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedName

`func (o *HeroObjectResponse) SetLocalizedName(v string)`

SetLocalizedName sets LocalizedName field to given value.

### HasLocalizedName

`func (o *HeroObjectResponse) HasLocalizedName() bool`

HasLocalizedName returns a boolean if a field has been set.

### GetPrimaryAttr

`func (o *HeroObjectResponse) GetPrimaryAttr() string`

GetPrimaryAttr returns the PrimaryAttr field if non-nil, zero value otherwise.

### GetPrimaryAttrOk

`func (o *HeroObjectResponse) GetPrimaryAttrOk() (*string, bool)`

GetPrimaryAttrOk returns a tuple with the PrimaryAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAttr

`func (o *HeroObjectResponse) SetPrimaryAttr(v string)`

SetPrimaryAttr sets PrimaryAttr field to given value.

### HasPrimaryAttr

`func (o *HeroObjectResponse) HasPrimaryAttr() bool`

HasPrimaryAttr returns a boolean if a field has been set.

### GetAttackType

`func (o *HeroObjectResponse) GetAttackType() string`

GetAttackType returns the AttackType field if non-nil, zero value otherwise.

### GetAttackTypeOk

`func (o *HeroObjectResponse) GetAttackTypeOk() (*string, bool)`

GetAttackTypeOk returns a tuple with the AttackType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackType

`func (o *HeroObjectResponse) SetAttackType(v string)`

SetAttackType sets AttackType field to given value.

### HasAttackType

`func (o *HeroObjectResponse) HasAttackType() bool`

HasAttackType returns a boolean if a field has been set.

### GetRoles

`func (o *HeroObjectResponse) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *HeroObjectResponse) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *HeroObjectResponse) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *HeroObjectResponse) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


