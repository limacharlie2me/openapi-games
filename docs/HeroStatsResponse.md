# HeroStatsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID value of the hero played | [optional] 
**Name** | Pointer to **string** | Dota hero command name | [optional] 
**LocalizedName** | Pointer to **string** | Hero name | [optional] 
**PrimaryAttr** | Pointer to **string** | primary_attr | [optional] 
**AttackType** | Pointer to **string** | attack_type | [optional] 
**Roles** | Pointer to **[]string** | roles | [optional] 
**Img** | Pointer to **string** | img | [optional] 
**Icon** | Pointer to **string** | icon | [optional] 
**BaseHealth** | Pointer to **int32** | base_health | [optional] 
**BaseHealthRegen** | Pointer to **float32** | base_health_regen | [optional] 
**BaseMana** | Pointer to **int32** | base_mana | [optional] 
**BaseManaRegen** | Pointer to **int32** | base_mana_regen | [optional] 
**BaseArmor** | Pointer to **int32** | base_armor | [optional] 
**BaseMr** | Pointer to **int32** | base_mr | [optional] 
**BaseAttackMin** | Pointer to **int32** | base_attack_min | [optional] 
**BaseAttackMax** | Pointer to **int32** | base_attack_max | [optional] 
**BaseStr** | Pointer to **int32** | base_str | [optional] 
**BaseAgi** | Pointer to **int32** | base_agi | [optional] 
**BaseInt** | Pointer to **int32** | base_int | [optional] 
**StrGain** | Pointer to **float32** | str_gain | [optional] 
**AgiGain** | Pointer to **float32** | agi_gain | [optional] 
**IntGain** | Pointer to **float32** | int_gain | [optional] 
**AttackRange** | Pointer to **int32** | attack_range | [optional] 
**ProjectileSpeed** | Pointer to **int32** | projectile_speed | [optional] 
**AttackRate** | Pointer to **float32** | attack_rate | [optional] 
**BaseAttackTime** | Pointer to **int32** | base_attack_time | [optional] 
**AttackPoint** | Pointer to **float32** | attack_point | [optional] 
**MoveSpeed** | Pointer to **int32** | move_speed | [optional] 
**TurnRate** | Pointer to **float32** | turn_rate | [optional] 
**CmEnabled** | Pointer to **bool** | cm_enabled | [optional] 
**Legs** | Pointer to **int32** | legs | [optional] 
**DayVision** | Pointer to **int32** | day_vision | [optional] 
**NightVision** | Pointer to **int32** | night_vision | [optional] 
**HeroId** | Pointer to **int32** | The ID value of the hero played | [optional] 
**TurboPicks** | Pointer to **int32** | Picks in Turbo mode this month | [optional] 
**TurboWins** | Pointer to **int32** | Wins in Turbo mode this month | [optional] 
**ProBan** | Pointer to **int32** | pro_ban | [optional] 
**ProWin** | Pointer to **int32** | pro_win | [optional] 
**ProPick** | Pointer to **int32** | pro_pick | [optional] 
**Var1Pick** | Pointer to **int32** | Herald picks | [optional] 
**Var1Win** | Pointer to **int32** | Herald wins | [optional] 
**Var2Pick** | Pointer to **int32** | Guardian picks | [optional] 
**Var2Win** | Pointer to **int32** | Guardian wins | [optional] 
**Var3Pick** | Pointer to **int32** | Crusader picks | [optional] 
**Var3Win** | Pointer to **int32** | Crusader wins | [optional] 
**Var4Pick** | Pointer to **int32** | Archon picks | [optional] 
**Var4Win** | Pointer to **int32** | Archon wins | [optional] 
**Var5Pick** | Pointer to **int32** | Legend picks | [optional] 
**Var5Win** | Pointer to **int32** | Legend wins | [optional] 
**Var6Pick** | Pointer to **int32** | Ancient picks | [optional] 
**Var6Win** | Pointer to **int32** | Ancient wins | [optional] 
**Var7Pick** | Pointer to **int32** | Divine picks | [optional] 
**Var7Win** | Pointer to **int32** | Divine wins | [optional] 
**Var8Pick** | Pointer to **int32** | Immortal picks | [optional] 
**Var8Win** | Pointer to **int32** | Immortal wins | [optional] 

## Methods

### NewHeroStatsResponse

`func NewHeroStatsResponse() *HeroStatsResponse`

NewHeroStatsResponse instantiates a new HeroStatsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeroStatsResponseWithDefaults

`func NewHeroStatsResponseWithDefaults() *HeroStatsResponse`

NewHeroStatsResponseWithDefaults instantiates a new HeroStatsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HeroStatsResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HeroStatsResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HeroStatsResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *HeroStatsResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *HeroStatsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HeroStatsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HeroStatsResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HeroStatsResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocalizedName

`func (o *HeroStatsResponse) GetLocalizedName() string`

GetLocalizedName returns the LocalizedName field if non-nil, zero value otherwise.

### GetLocalizedNameOk

`func (o *HeroStatsResponse) GetLocalizedNameOk() (*string, bool)`

GetLocalizedNameOk returns a tuple with the LocalizedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedName

`func (o *HeroStatsResponse) SetLocalizedName(v string)`

SetLocalizedName sets LocalizedName field to given value.

### HasLocalizedName

`func (o *HeroStatsResponse) HasLocalizedName() bool`

HasLocalizedName returns a boolean if a field has been set.

### GetPrimaryAttr

`func (o *HeroStatsResponse) GetPrimaryAttr() string`

GetPrimaryAttr returns the PrimaryAttr field if non-nil, zero value otherwise.

### GetPrimaryAttrOk

`func (o *HeroStatsResponse) GetPrimaryAttrOk() (*string, bool)`

GetPrimaryAttrOk returns a tuple with the PrimaryAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAttr

`func (o *HeroStatsResponse) SetPrimaryAttr(v string)`

SetPrimaryAttr sets PrimaryAttr field to given value.

### HasPrimaryAttr

`func (o *HeroStatsResponse) HasPrimaryAttr() bool`

HasPrimaryAttr returns a boolean if a field has been set.

### GetAttackType

`func (o *HeroStatsResponse) GetAttackType() string`

GetAttackType returns the AttackType field if non-nil, zero value otherwise.

### GetAttackTypeOk

`func (o *HeroStatsResponse) GetAttackTypeOk() (*string, bool)`

GetAttackTypeOk returns a tuple with the AttackType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackType

`func (o *HeroStatsResponse) SetAttackType(v string)`

SetAttackType sets AttackType field to given value.

### HasAttackType

`func (o *HeroStatsResponse) HasAttackType() bool`

HasAttackType returns a boolean if a field has been set.

### GetRoles

`func (o *HeroStatsResponse) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *HeroStatsResponse) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *HeroStatsResponse) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *HeroStatsResponse) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetImg

`func (o *HeroStatsResponse) GetImg() string`

GetImg returns the Img field if non-nil, zero value otherwise.

### GetImgOk

`func (o *HeroStatsResponse) GetImgOk() (*string, bool)`

GetImgOk returns a tuple with the Img field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImg

`func (o *HeroStatsResponse) SetImg(v string)`

SetImg sets Img field to given value.

### HasImg

`func (o *HeroStatsResponse) HasImg() bool`

HasImg returns a boolean if a field has been set.

### GetIcon

`func (o *HeroStatsResponse) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *HeroStatsResponse) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *HeroStatsResponse) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *HeroStatsResponse) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetBaseHealth

`func (o *HeroStatsResponse) GetBaseHealth() int32`

GetBaseHealth returns the BaseHealth field if non-nil, zero value otherwise.

### GetBaseHealthOk

`func (o *HeroStatsResponse) GetBaseHealthOk() (*int32, bool)`

GetBaseHealthOk returns a tuple with the BaseHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseHealth

`func (o *HeroStatsResponse) SetBaseHealth(v int32)`

SetBaseHealth sets BaseHealth field to given value.

### HasBaseHealth

`func (o *HeroStatsResponse) HasBaseHealth() bool`

HasBaseHealth returns a boolean if a field has been set.

### GetBaseHealthRegen

`func (o *HeroStatsResponse) GetBaseHealthRegen() float32`

GetBaseHealthRegen returns the BaseHealthRegen field if non-nil, zero value otherwise.

### GetBaseHealthRegenOk

`func (o *HeroStatsResponse) GetBaseHealthRegenOk() (*float32, bool)`

GetBaseHealthRegenOk returns a tuple with the BaseHealthRegen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseHealthRegen

`func (o *HeroStatsResponse) SetBaseHealthRegen(v float32)`

SetBaseHealthRegen sets BaseHealthRegen field to given value.

### HasBaseHealthRegen

`func (o *HeroStatsResponse) HasBaseHealthRegen() bool`

HasBaseHealthRegen returns a boolean if a field has been set.

### GetBaseMana

`func (o *HeroStatsResponse) GetBaseMana() int32`

GetBaseMana returns the BaseMana field if non-nil, zero value otherwise.

### GetBaseManaOk

`func (o *HeroStatsResponse) GetBaseManaOk() (*int32, bool)`

GetBaseManaOk returns a tuple with the BaseMana field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseMana

`func (o *HeroStatsResponse) SetBaseMana(v int32)`

SetBaseMana sets BaseMana field to given value.

### HasBaseMana

`func (o *HeroStatsResponse) HasBaseMana() bool`

HasBaseMana returns a boolean if a field has been set.

### GetBaseManaRegen

`func (o *HeroStatsResponse) GetBaseManaRegen() int32`

GetBaseManaRegen returns the BaseManaRegen field if non-nil, zero value otherwise.

### GetBaseManaRegenOk

`func (o *HeroStatsResponse) GetBaseManaRegenOk() (*int32, bool)`

GetBaseManaRegenOk returns a tuple with the BaseManaRegen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseManaRegen

`func (o *HeroStatsResponse) SetBaseManaRegen(v int32)`

SetBaseManaRegen sets BaseManaRegen field to given value.

### HasBaseManaRegen

`func (o *HeroStatsResponse) HasBaseManaRegen() bool`

HasBaseManaRegen returns a boolean if a field has been set.

### GetBaseArmor

`func (o *HeroStatsResponse) GetBaseArmor() int32`

GetBaseArmor returns the BaseArmor field if non-nil, zero value otherwise.

### GetBaseArmorOk

`func (o *HeroStatsResponse) GetBaseArmorOk() (*int32, bool)`

GetBaseArmorOk returns a tuple with the BaseArmor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseArmor

`func (o *HeroStatsResponse) SetBaseArmor(v int32)`

SetBaseArmor sets BaseArmor field to given value.

### HasBaseArmor

`func (o *HeroStatsResponse) HasBaseArmor() bool`

HasBaseArmor returns a boolean if a field has been set.

### GetBaseMr

`func (o *HeroStatsResponse) GetBaseMr() int32`

GetBaseMr returns the BaseMr field if non-nil, zero value otherwise.

### GetBaseMrOk

`func (o *HeroStatsResponse) GetBaseMrOk() (*int32, bool)`

GetBaseMrOk returns a tuple with the BaseMr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseMr

`func (o *HeroStatsResponse) SetBaseMr(v int32)`

SetBaseMr sets BaseMr field to given value.

### HasBaseMr

`func (o *HeroStatsResponse) HasBaseMr() bool`

HasBaseMr returns a boolean if a field has been set.

### GetBaseAttackMin

`func (o *HeroStatsResponse) GetBaseAttackMin() int32`

GetBaseAttackMin returns the BaseAttackMin field if non-nil, zero value otherwise.

### GetBaseAttackMinOk

`func (o *HeroStatsResponse) GetBaseAttackMinOk() (*int32, bool)`

GetBaseAttackMinOk returns a tuple with the BaseAttackMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAttackMin

`func (o *HeroStatsResponse) SetBaseAttackMin(v int32)`

SetBaseAttackMin sets BaseAttackMin field to given value.

### HasBaseAttackMin

`func (o *HeroStatsResponse) HasBaseAttackMin() bool`

HasBaseAttackMin returns a boolean if a field has been set.

### GetBaseAttackMax

`func (o *HeroStatsResponse) GetBaseAttackMax() int32`

GetBaseAttackMax returns the BaseAttackMax field if non-nil, zero value otherwise.

### GetBaseAttackMaxOk

`func (o *HeroStatsResponse) GetBaseAttackMaxOk() (*int32, bool)`

GetBaseAttackMaxOk returns a tuple with the BaseAttackMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAttackMax

`func (o *HeroStatsResponse) SetBaseAttackMax(v int32)`

SetBaseAttackMax sets BaseAttackMax field to given value.

### HasBaseAttackMax

`func (o *HeroStatsResponse) HasBaseAttackMax() bool`

HasBaseAttackMax returns a boolean if a field has been set.

### GetBaseStr

`func (o *HeroStatsResponse) GetBaseStr() int32`

GetBaseStr returns the BaseStr field if non-nil, zero value otherwise.

### GetBaseStrOk

`func (o *HeroStatsResponse) GetBaseStrOk() (*int32, bool)`

GetBaseStrOk returns a tuple with the BaseStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseStr

`func (o *HeroStatsResponse) SetBaseStr(v int32)`

SetBaseStr sets BaseStr field to given value.

### HasBaseStr

`func (o *HeroStatsResponse) HasBaseStr() bool`

HasBaseStr returns a boolean if a field has been set.

### GetBaseAgi

`func (o *HeroStatsResponse) GetBaseAgi() int32`

GetBaseAgi returns the BaseAgi field if non-nil, zero value otherwise.

### GetBaseAgiOk

`func (o *HeroStatsResponse) GetBaseAgiOk() (*int32, bool)`

GetBaseAgiOk returns a tuple with the BaseAgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAgi

`func (o *HeroStatsResponse) SetBaseAgi(v int32)`

SetBaseAgi sets BaseAgi field to given value.

### HasBaseAgi

`func (o *HeroStatsResponse) HasBaseAgi() bool`

HasBaseAgi returns a boolean if a field has been set.

### GetBaseInt

`func (o *HeroStatsResponse) GetBaseInt() int32`

GetBaseInt returns the BaseInt field if non-nil, zero value otherwise.

### GetBaseIntOk

`func (o *HeroStatsResponse) GetBaseIntOk() (*int32, bool)`

GetBaseIntOk returns a tuple with the BaseInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseInt

`func (o *HeroStatsResponse) SetBaseInt(v int32)`

SetBaseInt sets BaseInt field to given value.

### HasBaseInt

`func (o *HeroStatsResponse) HasBaseInt() bool`

HasBaseInt returns a boolean if a field has been set.

### GetStrGain

`func (o *HeroStatsResponse) GetStrGain() float32`

GetStrGain returns the StrGain field if non-nil, zero value otherwise.

### GetStrGainOk

`func (o *HeroStatsResponse) GetStrGainOk() (*float32, bool)`

GetStrGainOk returns a tuple with the StrGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrGain

`func (o *HeroStatsResponse) SetStrGain(v float32)`

SetStrGain sets StrGain field to given value.

### HasStrGain

`func (o *HeroStatsResponse) HasStrGain() bool`

HasStrGain returns a boolean if a field has been set.

### GetAgiGain

`func (o *HeroStatsResponse) GetAgiGain() float32`

GetAgiGain returns the AgiGain field if non-nil, zero value otherwise.

### GetAgiGainOk

`func (o *HeroStatsResponse) GetAgiGainOk() (*float32, bool)`

GetAgiGainOk returns a tuple with the AgiGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgiGain

`func (o *HeroStatsResponse) SetAgiGain(v float32)`

SetAgiGain sets AgiGain field to given value.

### HasAgiGain

`func (o *HeroStatsResponse) HasAgiGain() bool`

HasAgiGain returns a boolean if a field has been set.

### GetIntGain

`func (o *HeroStatsResponse) GetIntGain() float32`

GetIntGain returns the IntGain field if non-nil, zero value otherwise.

### GetIntGainOk

`func (o *HeroStatsResponse) GetIntGainOk() (*float32, bool)`

GetIntGainOk returns a tuple with the IntGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntGain

`func (o *HeroStatsResponse) SetIntGain(v float32)`

SetIntGain sets IntGain field to given value.

### HasIntGain

`func (o *HeroStatsResponse) HasIntGain() bool`

HasIntGain returns a boolean if a field has been set.

### GetAttackRange

`func (o *HeroStatsResponse) GetAttackRange() int32`

GetAttackRange returns the AttackRange field if non-nil, zero value otherwise.

### GetAttackRangeOk

`func (o *HeroStatsResponse) GetAttackRangeOk() (*int32, bool)`

GetAttackRangeOk returns a tuple with the AttackRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackRange

`func (o *HeroStatsResponse) SetAttackRange(v int32)`

SetAttackRange sets AttackRange field to given value.

### HasAttackRange

`func (o *HeroStatsResponse) HasAttackRange() bool`

HasAttackRange returns a boolean if a field has been set.

### GetProjectileSpeed

`func (o *HeroStatsResponse) GetProjectileSpeed() int32`

GetProjectileSpeed returns the ProjectileSpeed field if non-nil, zero value otherwise.

### GetProjectileSpeedOk

`func (o *HeroStatsResponse) GetProjectileSpeedOk() (*int32, bool)`

GetProjectileSpeedOk returns a tuple with the ProjectileSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectileSpeed

`func (o *HeroStatsResponse) SetProjectileSpeed(v int32)`

SetProjectileSpeed sets ProjectileSpeed field to given value.

### HasProjectileSpeed

`func (o *HeroStatsResponse) HasProjectileSpeed() bool`

HasProjectileSpeed returns a boolean if a field has been set.

### GetAttackRate

`func (o *HeroStatsResponse) GetAttackRate() float32`

GetAttackRate returns the AttackRate field if non-nil, zero value otherwise.

### GetAttackRateOk

`func (o *HeroStatsResponse) GetAttackRateOk() (*float32, bool)`

GetAttackRateOk returns a tuple with the AttackRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackRate

`func (o *HeroStatsResponse) SetAttackRate(v float32)`

SetAttackRate sets AttackRate field to given value.

### HasAttackRate

`func (o *HeroStatsResponse) HasAttackRate() bool`

HasAttackRate returns a boolean if a field has been set.

### GetBaseAttackTime

`func (o *HeroStatsResponse) GetBaseAttackTime() int32`

GetBaseAttackTime returns the BaseAttackTime field if non-nil, zero value otherwise.

### GetBaseAttackTimeOk

`func (o *HeroStatsResponse) GetBaseAttackTimeOk() (*int32, bool)`

GetBaseAttackTimeOk returns a tuple with the BaseAttackTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAttackTime

`func (o *HeroStatsResponse) SetBaseAttackTime(v int32)`

SetBaseAttackTime sets BaseAttackTime field to given value.

### HasBaseAttackTime

`func (o *HeroStatsResponse) HasBaseAttackTime() bool`

HasBaseAttackTime returns a boolean if a field has been set.

### GetAttackPoint

`func (o *HeroStatsResponse) GetAttackPoint() float32`

GetAttackPoint returns the AttackPoint field if non-nil, zero value otherwise.

### GetAttackPointOk

`func (o *HeroStatsResponse) GetAttackPointOk() (*float32, bool)`

GetAttackPointOk returns a tuple with the AttackPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackPoint

`func (o *HeroStatsResponse) SetAttackPoint(v float32)`

SetAttackPoint sets AttackPoint field to given value.

### HasAttackPoint

`func (o *HeroStatsResponse) HasAttackPoint() bool`

HasAttackPoint returns a boolean if a field has been set.

### GetMoveSpeed

`func (o *HeroStatsResponse) GetMoveSpeed() int32`

GetMoveSpeed returns the MoveSpeed field if non-nil, zero value otherwise.

### GetMoveSpeedOk

`func (o *HeroStatsResponse) GetMoveSpeedOk() (*int32, bool)`

GetMoveSpeedOk returns a tuple with the MoveSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveSpeed

`func (o *HeroStatsResponse) SetMoveSpeed(v int32)`

SetMoveSpeed sets MoveSpeed field to given value.

### HasMoveSpeed

`func (o *HeroStatsResponse) HasMoveSpeed() bool`

HasMoveSpeed returns a boolean if a field has been set.

### GetTurnRate

`func (o *HeroStatsResponse) GetTurnRate() float32`

GetTurnRate returns the TurnRate field if non-nil, zero value otherwise.

### GetTurnRateOk

`func (o *HeroStatsResponse) GetTurnRateOk() (*float32, bool)`

GetTurnRateOk returns a tuple with the TurnRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurnRate

`func (o *HeroStatsResponse) SetTurnRate(v float32)`

SetTurnRate sets TurnRate field to given value.

### HasTurnRate

`func (o *HeroStatsResponse) HasTurnRate() bool`

HasTurnRate returns a boolean if a field has been set.

### GetCmEnabled

`func (o *HeroStatsResponse) GetCmEnabled() bool`

GetCmEnabled returns the CmEnabled field if non-nil, zero value otherwise.

### GetCmEnabledOk

`func (o *HeroStatsResponse) GetCmEnabledOk() (*bool, bool)`

GetCmEnabledOk returns a tuple with the CmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmEnabled

`func (o *HeroStatsResponse) SetCmEnabled(v bool)`

SetCmEnabled sets CmEnabled field to given value.

### HasCmEnabled

`func (o *HeroStatsResponse) HasCmEnabled() bool`

HasCmEnabled returns a boolean if a field has been set.

### GetLegs

`func (o *HeroStatsResponse) GetLegs() int32`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *HeroStatsResponse) GetLegsOk() (*int32, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *HeroStatsResponse) SetLegs(v int32)`

SetLegs sets Legs field to given value.

### HasLegs

`func (o *HeroStatsResponse) HasLegs() bool`

HasLegs returns a boolean if a field has been set.

### GetDayVision

`func (o *HeroStatsResponse) GetDayVision() int32`

GetDayVision returns the DayVision field if non-nil, zero value otherwise.

### GetDayVisionOk

`func (o *HeroStatsResponse) GetDayVisionOk() (*int32, bool)`

GetDayVisionOk returns a tuple with the DayVision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayVision

`func (o *HeroStatsResponse) SetDayVision(v int32)`

SetDayVision sets DayVision field to given value.

### HasDayVision

`func (o *HeroStatsResponse) HasDayVision() bool`

HasDayVision returns a boolean if a field has been set.

### GetNightVision

`func (o *HeroStatsResponse) GetNightVision() int32`

GetNightVision returns the NightVision field if non-nil, zero value otherwise.

### GetNightVisionOk

`func (o *HeroStatsResponse) GetNightVisionOk() (*int32, bool)`

GetNightVisionOk returns a tuple with the NightVision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNightVision

`func (o *HeroStatsResponse) SetNightVision(v int32)`

SetNightVision sets NightVision field to given value.

### HasNightVision

`func (o *HeroStatsResponse) HasNightVision() bool`

HasNightVision returns a boolean if a field has been set.

### GetHeroId

`func (o *HeroStatsResponse) GetHeroId() int32`

GetHeroId returns the HeroId field if non-nil, zero value otherwise.

### GetHeroIdOk

`func (o *HeroStatsResponse) GetHeroIdOk() (*int32, bool)`

GetHeroIdOk returns a tuple with the HeroId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeroId

`func (o *HeroStatsResponse) SetHeroId(v int32)`

SetHeroId sets HeroId field to given value.

### HasHeroId

`func (o *HeroStatsResponse) HasHeroId() bool`

HasHeroId returns a boolean if a field has been set.

### GetTurboPicks

`func (o *HeroStatsResponse) GetTurboPicks() int32`

GetTurboPicks returns the TurboPicks field if non-nil, zero value otherwise.

### GetTurboPicksOk

`func (o *HeroStatsResponse) GetTurboPicksOk() (*int32, bool)`

GetTurboPicksOk returns a tuple with the TurboPicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurboPicks

`func (o *HeroStatsResponse) SetTurboPicks(v int32)`

SetTurboPicks sets TurboPicks field to given value.

### HasTurboPicks

`func (o *HeroStatsResponse) HasTurboPicks() bool`

HasTurboPicks returns a boolean if a field has been set.

### GetTurboWins

`func (o *HeroStatsResponse) GetTurboWins() int32`

GetTurboWins returns the TurboWins field if non-nil, zero value otherwise.

### GetTurboWinsOk

`func (o *HeroStatsResponse) GetTurboWinsOk() (*int32, bool)`

GetTurboWinsOk returns a tuple with the TurboWins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurboWins

`func (o *HeroStatsResponse) SetTurboWins(v int32)`

SetTurboWins sets TurboWins field to given value.

### HasTurboWins

`func (o *HeroStatsResponse) HasTurboWins() bool`

HasTurboWins returns a boolean if a field has been set.

### GetProBan

`func (o *HeroStatsResponse) GetProBan() int32`

GetProBan returns the ProBan field if non-nil, zero value otherwise.

### GetProBanOk

`func (o *HeroStatsResponse) GetProBanOk() (*int32, bool)`

GetProBanOk returns a tuple with the ProBan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProBan

`func (o *HeroStatsResponse) SetProBan(v int32)`

SetProBan sets ProBan field to given value.

### HasProBan

`func (o *HeroStatsResponse) HasProBan() bool`

HasProBan returns a boolean if a field has been set.

### GetProWin

`func (o *HeroStatsResponse) GetProWin() int32`

GetProWin returns the ProWin field if non-nil, zero value otherwise.

### GetProWinOk

`func (o *HeroStatsResponse) GetProWinOk() (*int32, bool)`

GetProWinOk returns a tuple with the ProWin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProWin

`func (o *HeroStatsResponse) SetProWin(v int32)`

SetProWin sets ProWin field to given value.

### HasProWin

`func (o *HeroStatsResponse) HasProWin() bool`

HasProWin returns a boolean if a field has been set.

### GetProPick

`func (o *HeroStatsResponse) GetProPick() int32`

GetProPick returns the ProPick field if non-nil, zero value otherwise.

### GetProPickOk

`func (o *HeroStatsResponse) GetProPickOk() (*int32, bool)`

GetProPickOk returns a tuple with the ProPick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProPick

`func (o *HeroStatsResponse) SetProPick(v int32)`

SetProPick sets ProPick field to given value.

### HasProPick

`func (o *HeroStatsResponse) HasProPick() bool`

HasProPick returns a boolean if a field has been set.

### GetVar1Pick

`func (o *HeroStatsResponse) GetVar1Pick() int32`

GetVar1Pick returns the Var1Pick field if non-nil, zero value otherwise.

### GetVar1PickOk

`func (o *HeroStatsResponse) GetVar1PickOk() (*int32, bool)`

GetVar1PickOk returns a tuple with the Var1Pick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar1Pick

`func (o *HeroStatsResponse) SetVar1Pick(v int32)`

SetVar1Pick sets Var1Pick field to given value.

### HasVar1Pick

`func (o *HeroStatsResponse) HasVar1Pick() bool`

HasVar1Pick returns a boolean if a field has been set.

### GetVar1Win

`func (o *HeroStatsResponse) GetVar1Win() int32`

GetVar1Win returns the Var1Win field if non-nil, zero value otherwise.

### GetVar1WinOk

`func (o *HeroStatsResponse) GetVar1WinOk() (*int32, bool)`

GetVar1WinOk returns a tuple with the Var1Win field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar1Win

`func (o *HeroStatsResponse) SetVar1Win(v int32)`

SetVar1Win sets Var1Win field to given value.

### HasVar1Win

`func (o *HeroStatsResponse) HasVar1Win() bool`

HasVar1Win returns a boolean if a field has been set.

### GetVar2Pick

`func (o *HeroStatsResponse) GetVar2Pick() int32`

GetVar2Pick returns the Var2Pick field if non-nil, zero value otherwise.

### GetVar2PickOk

`func (o *HeroStatsResponse) GetVar2PickOk() (*int32, bool)`

GetVar2PickOk returns a tuple with the Var2Pick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar2Pick

`func (o *HeroStatsResponse) SetVar2Pick(v int32)`

SetVar2Pick sets Var2Pick field to given value.

### HasVar2Pick

`func (o *HeroStatsResponse) HasVar2Pick() bool`

HasVar2Pick returns a boolean if a field has been set.

### GetVar2Win

`func (o *HeroStatsResponse) GetVar2Win() int32`

GetVar2Win returns the Var2Win field if non-nil, zero value otherwise.

### GetVar2WinOk

`func (o *HeroStatsResponse) GetVar2WinOk() (*int32, bool)`

GetVar2WinOk returns a tuple with the Var2Win field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar2Win

`func (o *HeroStatsResponse) SetVar2Win(v int32)`

SetVar2Win sets Var2Win field to given value.

### HasVar2Win

`func (o *HeroStatsResponse) HasVar2Win() bool`

HasVar2Win returns a boolean if a field has been set.

### GetVar3Pick

`func (o *HeroStatsResponse) GetVar3Pick() int32`

GetVar3Pick returns the Var3Pick field if non-nil, zero value otherwise.

### GetVar3PickOk

`func (o *HeroStatsResponse) GetVar3PickOk() (*int32, bool)`

GetVar3PickOk returns a tuple with the Var3Pick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar3Pick

`func (o *HeroStatsResponse) SetVar3Pick(v int32)`

SetVar3Pick sets Var3Pick field to given value.

### HasVar3Pick

`func (o *HeroStatsResponse) HasVar3Pick() bool`

HasVar3Pick returns a boolean if a field has been set.

### GetVar3Win

`func (o *HeroStatsResponse) GetVar3Win() int32`

GetVar3Win returns the Var3Win field if non-nil, zero value otherwise.

### GetVar3WinOk

`func (o *HeroStatsResponse) GetVar3WinOk() (*int32, bool)`

GetVar3WinOk returns a tuple with the Var3Win field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar3Win

`func (o *HeroStatsResponse) SetVar3Win(v int32)`

SetVar3Win sets Var3Win field to given value.

### HasVar3Win

`func (o *HeroStatsResponse) HasVar3Win() bool`

HasVar3Win returns a boolean if a field has been set.

### GetVar4Pick

`func (o *HeroStatsResponse) GetVar4Pick() int32`

GetVar4Pick returns the Var4Pick field if non-nil, zero value otherwise.

### GetVar4PickOk

`func (o *HeroStatsResponse) GetVar4PickOk() (*int32, bool)`

GetVar4PickOk returns a tuple with the Var4Pick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar4Pick

`func (o *HeroStatsResponse) SetVar4Pick(v int32)`

SetVar4Pick sets Var4Pick field to given value.

### HasVar4Pick

`func (o *HeroStatsResponse) HasVar4Pick() bool`

HasVar4Pick returns a boolean if a field has been set.

### GetVar4Win

`func (o *HeroStatsResponse) GetVar4Win() int32`

GetVar4Win returns the Var4Win field if non-nil, zero value otherwise.

### GetVar4WinOk

`func (o *HeroStatsResponse) GetVar4WinOk() (*int32, bool)`

GetVar4WinOk returns a tuple with the Var4Win field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar4Win

`func (o *HeroStatsResponse) SetVar4Win(v int32)`

SetVar4Win sets Var4Win field to given value.

### HasVar4Win

`func (o *HeroStatsResponse) HasVar4Win() bool`

HasVar4Win returns a boolean if a field has been set.

### GetVar5Pick

`func (o *HeroStatsResponse) GetVar5Pick() int32`

GetVar5Pick returns the Var5Pick field if non-nil, zero value otherwise.

### GetVar5PickOk

`func (o *HeroStatsResponse) GetVar5PickOk() (*int32, bool)`

GetVar5PickOk returns a tuple with the Var5Pick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5Pick

`func (o *HeroStatsResponse) SetVar5Pick(v int32)`

SetVar5Pick sets Var5Pick field to given value.

### HasVar5Pick

`func (o *HeroStatsResponse) HasVar5Pick() bool`

HasVar5Pick returns a boolean if a field has been set.

### GetVar5Win

`func (o *HeroStatsResponse) GetVar5Win() int32`

GetVar5Win returns the Var5Win field if non-nil, zero value otherwise.

### GetVar5WinOk

`func (o *HeroStatsResponse) GetVar5WinOk() (*int32, bool)`

GetVar5WinOk returns a tuple with the Var5Win field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5Win

`func (o *HeroStatsResponse) SetVar5Win(v int32)`

SetVar5Win sets Var5Win field to given value.

### HasVar5Win

`func (o *HeroStatsResponse) HasVar5Win() bool`

HasVar5Win returns a boolean if a field has been set.

### GetVar6Pick

`func (o *HeroStatsResponse) GetVar6Pick() int32`

GetVar6Pick returns the Var6Pick field if non-nil, zero value otherwise.

### GetVar6PickOk

`func (o *HeroStatsResponse) GetVar6PickOk() (*int32, bool)`

GetVar6PickOk returns a tuple with the Var6Pick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar6Pick

`func (o *HeroStatsResponse) SetVar6Pick(v int32)`

SetVar6Pick sets Var6Pick field to given value.

### HasVar6Pick

`func (o *HeroStatsResponse) HasVar6Pick() bool`

HasVar6Pick returns a boolean if a field has been set.

### GetVar6Win

`func (o *HeroStatsResponse) GetVar6Win() int32`

GetVar6Win returns the Var6Win field if non-nil, zero value otherwise.

### GetVar6WinOk

`func (o *HeroStatsResponse) GetVar6WinOk() (*int32, bool)`

GetVar6WinOk returns a tuple with the Var6Win field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar6Win

`func (o *HeroStatsResponse) SetVar6Win(v int32)`

SetVar6Win sets Var6Win field to given value.

### HasVar6Win

`func (o *HeroStatsResponse) HasVar6Win() bool`

HasVar6Win returns a boolean if a field has been set.

### GetVar7Pick

`func (o *HeroStatsResponse) GetVar7Pick() int32`

GetVar7Pick returns the Var7Pick field if non-nil, zero value otherwise.

### GetVar7PickOk

`func (o *HeroStatsResponse) GetVar7PickOk() (*int32, bool)`

GetVar7PickOk returns a tuple with the Var7Pick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7Pick

`func (o *HeroStatsResponse) SetVar7Pick(v int32)`

SetVar7Pick sets Var7Pick field to given value.

### HasVar7Pick

`func (o *HeroStatsResponse) HasVar7Pick() bool`

HasVar7Pick returns a boolean if a field has been set.

### GetVar7Win

`func (o *HeroStatsResponse) GetVar7Win() int32`

GetVar7Win returns the Var7Win field if non-nil, zero value otherwise.

### GetVar7WinOk

`func (o *HeroStatsResponse) GetVar7WinOk() (*int32, bool)`

GetVar7WinOk returns a tuple with the Var7Win field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7Win

`func (o *HeroStatsResponse) SetVar7Win(v int32)`

SetVar7Win sets Var7Win field to given value.

### HasVar7Win

`func (o *HeroStatsResponse) HasVar7Win() bool`

HasVar7Win returns a boolean if a field has been set.

### GetVar8Pick

`func (o *HeroStatsResponse) GetVar8Pick() int32`

GetVar8Pick returns the Var8Pick field if non-nil, zero value otherwise.

### GetVar8PickOk

`func (o *HeroStatsResponse) GetVar8PickOk() (*int32, bool)`

GetVar8PickOk returns a tuple with the Var8Pick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar8Pick

`func (o *HeroStatsResponse) SetVar8Pick(v int32)`

SetVar8Pick sets Var8Pick field to given value.

### HasVar8Pick

`func (o *HeroStatsResponse) HasVar8Pick() bool`

HasVar8Pick returns a boolean if a field has been set.

### GetVar8Win

`func (o *HeroStatsResponse) GetVar8Win() int32`

GetVar8Win returns the Var8Win field if non-nil, zero value otherwise.

### GetVar8WinOk

`func (o *HeroStatsResponse) GetVar8WinOk() (*int32, bool)`

GetVar8WinOk returns a tuple with the Var8Win field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar8Win

`func (o *HeroStatsResponse) SetVar8Win(v int32)`

SetVar8Win sets Var8Win field to given value.

### HasVar8Win

`func (o *HeroStatsResponse) HasVar8Win() bool`

HasVar8Win returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


