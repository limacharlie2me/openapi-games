# MatchResponsePlayersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchId** | Pointer to **int32** | The ID number of the match assigned by Valve | [optional] 
**PlayerSlot** | Pointer to **NullableInt32** | Which slot the player is in. 0-127 are Radiant, 128-255 are Dire | [optional] 
**AbilityUpgradesArr** | Pointer to **[]int32** | An array describing how abilities were upgraded | [optional] 
**AbilityUses** | Pointer to **map[string]interface{}** | Object containing information on how many times the played used their abilities | [optional] 
**AbilityTargets** | Pointer to **map[string]interface{}** | Object containing information on who the player used their abilities on | [optional] 
**DamageTargets** | Pointer to **map[string]interface{}** | Object containing information on how and how much damage the player dealt to other heroes | [optional] 
**AccountId** | Pointer to **int32** | The player account ID | [optional] 
**Actions** | Pointer to **map[string]interface{}** | Object containing information on how many and what type of actions the player issued to their hero | [optional] 
**AdditionalUnits** | Pointer to **[]map[string]interface{}** | Object containing information on additional units the player had under their control | [optional] 
**Assists** | Pointer to **int32** | Number of assists the player had | [optional] 
**Backpack0** | Pointer to **int32** | Item in backpack slot 0 | [optional] 
**Backpack1** | Pointer to **int32** | Item in backpack slot 1 | [optional] 
**Backpack2** | Pointer to **int32** | Item in backpack slot 2 | [optional] 
**BuybackLog** | Pointer to [**[]MatchResponsePlayersInnerBuybackLogInner**](MatchResponsePlayersInnerBuybackLogInner.md) | Array containing information about buybacks | [optional] 
**CampsStacked** | Pointer to **int32** | Number of camps stacked | [optional] 
**ConnectionLog** | Pointer to [**[]MatchResponsePlayersInnerConnectionLogInner**](MatchResponsePlayersInnerConnectionLogInner.md) | Array containing information about the player&#39;s disconnections and reconnections | [optional] 
**CreepsStacked** | Pointer to **int32** | Number of creeps stacked | [optional] 
**Damage** | Pointer to **map[string]interface{}** | Object containing information about damage dealt by the player to different units | [optional] 
**DamageInflictor** | Pointer to **map[string]interface{}** | Object containing information about about the sources of this player&#39;s damage to heroes | [optional] 
**DamageInflictorReceived** | Pointer to **map[string]interface{}** | Object containing information about the sources of damage received by this player from heroes | [optional] 
**DamageTaken** | Pointer to **map[string]interface{}** | Object containing information about from whom the player took damage | [optional] 
**Deaths** | Pointer to **int32** | Number of deaths | [optional] 
**Denies** | Pointer to **int32** | Number of denies | [optional] 
**DnT** | Pointer to **[]int32** | Array containing number of denies at different times of the match | [optional] 
**Gold** | Pointer to **int32** | Gold at the end of the game | [optional] 
**GoldPerMin** | Pointer to **int32** | Gold Per Minute obtained by this player | [optional] 
**GoldReasons** | Pointer to **map[string]interface{}** | Object containing information on how the player gainined gold over the course of the match | [optional] 
**GoldSpent** | Pointer to **int32** | How much gold the player spent | [optional] 
**GoldT** | Pointer to **[]int32** | Array containing total gold at different times of the match | [optional] 
**HeroDamage** | Pointer to **int32** | Hero Damage Dealt | [optional] 
**HeroHealing** | Pointer to **int32** | Hero Healing Done | [optional] 
**HeroHits** | Pointer to **map[string]interface{}** | Object containing information on how many ticks of damages the hero inflicted with different spells and damage inflictors | [optional] 
**HeroId** | Pointer to **int32** | The ID value of the hero played | [optional] 
**Item0** | Pointer to **int32** | Item in the player&#39;s first slot | [optional] 
**Item1** | Pointer to **int32** | Item in the player&#39;s second slot | [optional] 
**Item2** | Pointer to **int32** | Item in the player&#39;s third slot | [optional] 
**Item3** | Pointer to **int32** | Item in the player&#39;s fourth slot | [optional] 
**Item4** | Pointer to **int32** | Item in the player&#39;s fifth slot | [optional] 
**Item5** | Pointer to **int32** | Item in the player&#39;s sixth slot | [optional] 
**ItemUses** | Pointer to **map[string]interface{}** | Object containing information about how many times a player used items | [optional] 
**KillStreaks** | Pointer to **map[string]interface{}** | Object containing information about the player&#39;s killstreaks | [optional] 
**Killed** | Pointer to **map[string]interface{}** | Object containing information about what units the player killed | [optional] 
**KilledBy** | Pointer to **map[string]interface{}** | Object containing information about who killed the player | [optional] 
**Kills** | Pointer to **int32** | Number of kills | [optional] 
**KillsLog** | Pointer to [**[]MatchResponsePlayersInnerKillsLogInner**](MatchResponsePlayersInnerKillsLogInner.md) | Array containing information on which hero the player killed at what time | [optional] 
**LanePos** | Pointer to **map[string]interface{}** | Object containing information on lane position | [optional] 
**LastHits** | Pointer to **int32** | Number of last hits | [optional] 
**LeaverStatus** | Pointer to **int32** | Integer describing whether or not the player left the game. 0: didn&#39;t leave. 1: left safely. 2+: Abandoned | [optional] 
**Level** | Pointer to **int32** | Level at the end of the game | [optional] 
**LhT** | Pointer to **[]int32** | Array describing last hits at each minute in the game | [optional] 
**LifeState** | Pointer to **map[string]interface{}** | life_state | [optional] 
**MaxHeroHit** | Pointer to **map[string]interface{}** | Object with information on the highest damage instance the player inflicted | [optional] 
**MultiKills** | Pointer to **map[string]interface{}** | Object with information on the number of the number of multikills the player had | [optional] 
**Obs** | Pointer to **map[string]interface{}** | Object with information on where the player placed observer wards. The location takes the form (outer number, inner number) and are from ~64-192. | [optional] 
**ObsLeftLog** | Pointer to **[]map[string]interface{}** | obs_left_log | [optional] 
**ObsLog** | Pointer to **[]map[string]interface{}** | Object containing information on when and where the player placed observer wards | [optional] 
**ObsPlaced** | Pointer to **int32** | Total number of observer wards placed | [optional] 
**PartyId** | Pointer to **int32** | party_id | [optional] 
**PermanentBuffs** | Pointer to **[]map[string]interface{}** | Array describing permanent buffs the player had at the end of the game. List of constants can be found here: https://github.com/odota/dotaconstants/blob/master/json/permanent_buffs.json | [optional] 
**Pings** | Pointer to **int32** | Total number of pings | [optional] 
**Purchase** | Pointer to **map[string]interface{}** | Object containing information on the items the player purchased | [optional] 
**PurchaseLog** | Pointer to [**[]MatchResponsePlayersInnerPurchaseLogInner**](MatchResponsePlayersInnerPurchaseLogInner.md) | Object containing information on when items were purchased | [optional] 
**RunePickups** | Pointer to **int32** | Number of runes picked up | [optional] 
**Runes** | Pointer to **map[string]int32** | Object with information about which runes the player picked up | [optional] 
**RunesLog** | Pointer to [**[]MatchResponsePlayersInnerRunesLogInner**](MatchResponsePlayersInnerRunesLogInner.md) | Array with information on when runes were picked up | [optional] 
**Sen** | Pointer to **map[string]interface{}** | Object with information on where sentries were placed. The location takes the form (outer number, inner number) and are from ~64-192. | [optional] 
**SenLeftLog** | Pointer to **[]map[string]interface{}** | Array containing information on when and where the player placed sentries | [optional] 
**SenLog** | Pointer to **[]map[string]interface{}** | Array with information on when and where sentries were placed by the player | [optional] 
**SenPlaced** | Pointer to **int32** | How many sentries were placed by the player | [optional] 
**Stuns** | Pointer to **float32** | Total stun duration of all stuns by the player | [optional] 
**Times** | Pointer to **[]int32** | Time in seconds corresponding to the time of entries of other arrays in the match. | [optional] 
**TowerDamage** | Pointer to **int32** | Total tower damage done by the player | [optional] 
**XpPerMin** | Pointer to **int32** | Experience Per Minute obtained by the player | [optional] 
**XpReasons** | Pointer to **map[string]interface{}** | Object containing information on the sources of this player&#39;s experience | [optional] 
**XpT** | Pointer to **[]int32** | Experience at each minute of the game | [optional] 
**Personaname** | Pointer to **NullableString** | Player&#39;s Steam name | [optional] 
**Name** | Pointer to **NullableString** | name | [optional] 
**LastLogin** | Pointer to **NullableTime** | Time of player&#39;s last login | [optional] 
**RadiantWin** | Pointer to **NullableBool** | Boolean indicating whether Radiant won the match | [optional] 
**StartTime** | Pointer to **int32** | The Unix timestamp at which the game started | [optional] 
**Duration** | Pointer to **int32** | Duration of the game in seconds | [optional] 
**Cluster** | Pointer to **int32** | cluster | [optional] 
**LobbyType** | Pointer to **int32** | Integer corresponding to lobby type of match. List of constants can be found here: https://github.com/odota/dotaconstants/blob/master/json/lobby_type.json | [optional] 
**GameMode** | Pointer to **int32** | Integer corresponding to game mode played. List of constants can be found here: https://github.com/odota/dotaconstants/blob/master/json/game_mode.json | [optional] 
**Patch** | Pointer to **int32** | Patch ID, from dotaconstants | [optional] 
**Region** | Pointer to **int32** | Integer corresponding to the region the game was played on | [optional] 
**IsRadiant** | Pointer to **bool** | Boolean for whether or not the player is on Radiant | [optional] 
**Win** | Pointer to **int32** | Binary integer representing whether or not the player won | [optional] 
**Lose** | Pointer to **int32** | Binary integer representing whether or not the player lost | [optional] 
**TotalGold** | Pointer to **int32** | Total gold at the end of the game | [optional] 
**TotalXp** | Pointer to **int32** | Total experience at the end of the game | [optional] 
**KillsPerMin** | Pointer to **float32** | Number of kills per minute | [optional] 
**Kda** | Pointer to **float32** | kda | [optional] 
**Abandons** | Pointer to **int32** | abandons | [optional] 
**NeutralKills** | Pointer to **int32** | Total number of neutral creeps killed | [optional] 
**TowerKills** | Pointer to **int32** | Total number of tower kills the player had | [optional] 
**CourierKills** | Pointer to **int32** | Total number of courier kills the player had | [optional] 
**LaneKills** | Pointer to **int32** | Total number of lane creeps killed by the player | [optional] 
**HeroKills** | Pointer to **int32** | Total number of heroes killed by the player | [optional] 
**ObserverKills** | Pointer to **int32** | Total number of observer wards killed by the player | [optional] 
**SentryKills** | Pointer to **int32** | Total number of sentry wards killed by the player | [optional] 
**RoshanKills** | Pointer to **int32** | Total number of roshan kills (last hit on roshan) the player had | [optional] 
**NecronomiconKills** | Pointer to **int32** | Total number of Necronomicon creeps killed by the player | [optional] 
**AncientKills** | Pointer to **int32** | Total number of Ancient creeps killed by the player | [optional] 
**BuybackCount** | Pointer to **int32** | Total number of buyback the player used | [optional] 
**ObserverUses** | Pointer to **int32** | Number of observer wards used | [optional] 
**SentryUses** | Pointer to **int32** | Number of sentry wards used | [optional] 
**LaneEfficiency** | Pointer to **float32** | lane_efficiency | [optional] 
**LaneEfficiencyPct** | Pointer to **float32** | lane_efficiency_pct | [optional] 
**Lane** | Pointer to **NullableInt32** | Integer referring to which lane the hero laned in | [optional] 
**LaneRole** | Pointer to **NullableInt32** | lane_role | [optional] 
**IsRoaming** | Pointer to **NullableBool** | Boolean referring to whether or not the player roamed | [optional] 
**PurchaseTime** | Pointer to **map[string]interface{}** | Object with information on when the player last purchased an item | [optional] 
**FirstPurchaseTime** | Pointer to **map[string]interface{}** | Object with information on when the player first puchased an item | [optional] 
**ItemWin** | Pointer to **map[string]interface{}** | Object with information on whether or not the item won | [optional] 
**ItemUsage** | Pointer to **map[string]interface{}** | Object containing binary integers the tell whether the item was purchased by the player (note: this is always 1) | [optional] 
**PurchaseTpscroll** | Pointer to **int32** | Total number of TP scrolls purchased by the player | [optional] 
**ActionsPerMin** | Pointer to **int32** | Actions per minute | [optional] 
**LifeStateDead** | Pointer to **int32** | life_state_dead | [optional] 
**RankTier** | Pointer to **int32** | The rank tier of the player. Tens place indicates rank, ones place indicates stars. | [optional] 
**Cosmetics** | Pointer to [**[]MatchResponsePlayersInnerCosmeticsInner**](MatchResponsePlayersInnerCosmeticsInner.md) | cosmetics | [optional] 
**Benchmarks** | Pointer to **map[string]interface{}** | Object containing information on certain benchmarks like GPM, XPM, KDA, tower damage, etc | [optional] 

## Methods

### NewMatchResponsePlayersInner

`func NewMatchResponsePlayersInner() *MatchResponsePlayersInner`

NewMatchResponsePlayersInner instantiates a new MatchResponsePlayersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchResponsePlayersInnerWithDefaults

`func NewMatchResponsePlayersInnerWithDefaults() *MatchResponsePlayersInner`

NewMatchResponsePlayersInnerWithDefaults instantiates a new MatchResponsePlayersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchId

`func (o *MatchResponsePlayersInner) GetMatchId() int32`

GetMatchId returns the MatchId field if non-nil, zero value otherwise.

### GetMatchIdOk

`func (o *MatchResponsePlayersInner) GetMatchIdOk() (*int32, bool)`

GetMatchIdOk returns a tuple with the MatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchId

`func (o *MatchResponsePlayersInner) SetMatchId(v int32)`

SetMatchId sets MatchId field to given value.

### HasMatchId

`func (o *MatchResponsePlayersInner) HasMatchId() bool`

HasMatchId returns a boolean if a field has been set.

### GetPlayerSlot

`func (o *MatchResponsePlayersInner) GetPlayerSlot() int32`

GetPlayerSlot returns the PlayerSlot field if non-nil, zero value otherwise.

### GetPlayerSlotOk

`func (o *MatchResponsePlayersInner) GetPlayerSlotOk() (*int32, bool)`

GetPlayerSlotOk returns a tuple with the PlayerSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerSlot

`func (o *MatchResponsePlayersInner) SetPlayerSlot(v int32)`

SetPlayerSlot sets PlayerSlot field to given value.

### HasPlayerSlot

`func (o *MatchResponsePlayersInner) HasPlayerSlot() bool`

HasPlayerSlot returns a boolean if a field has been set.

### SetPlayerSlotNil

`func (o *MatchResponsePlayersInner) SetPlayerSlotNil(b bool)`

 SetPlayerSlotNil sets the value for PlayerSlot to be an explicit nil

### UnsetPlayerSlot
`func (o *MatchResponsePlayersInner) UnsetPlayerSlot()`

UnsetPlayerSlot ensures that no value is present for PlayerSlot, not even an explicit nil
### GetAbilityUpgradesArr

`func (o *MatchResponsePlayersInner) GetAbilityUpgradesArr() []int32`

GetAbilityUpgradesArr returns the AbilityUpgradesArr field if non-nil, zero value otherwise.

### GetAbilityUpgradesArrOk

`func (o *MatchResponsePlayersInner) GetAbilityUpgradesArrOk() (*[]int32, bool)`

GetAbilityUpgradesArrOk returns a tuple with the AbilityUpgradesArr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbilityUpgradesArr

`func (o *MatchResponsePlayersInner) SetAbilityUpgradesArr(v []int32)`

SetAbilityUpgradesArr sets AbilityUpgradesArr field to given value.

### HasAbilityUpgradesArr

`func (o *MatchResponsePlayersInner) HasAbilityUpgradesArr() bool`

HasAbilityUpgradesArr returns a boolean if a field has been set.

### GetAbilityUses

`func (o *MatchResponsePlayersInner) GetAbilityUses() map[string]interface{}`

GetAbilityUses returns the AbilityUses field if non-nil, zero value otherwise.

### GetAbilityUsesOk

`func (o *MatchResponsePlayersInner) GetAbilityUsesOk() (*map[string]interface{}, bool)`

GetAbilityUsesOk returns a tuple with the AbilityUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbilityUses

`func (o *MatchResponsePlayersInner) SetAbilityUses(v map[string]interface{})`

SetAbilityUses sets AbilityUses field to given value.

### HasAbilityUses

`func (o *MatchResponsePlayersInner) HasAbilityUses() bool`

HasAbilityUses returns a boolean if a field has been set.

### GetAbilityTargets

`func (o *MatchResponsePlayersInner) GetAbilityTargets() map[string]interface{}`

GetAbilityTargets returns the AbilityTargets field if non-nil, zero value otherwise.

### GetAbilityTargetsOk

`func (o *MatchResponsePlayersInner) GetAbilityTargetsOk() (*map[string]interface{}, bool)`

GetAbilityTargetsOk returns a tuple with the AbilityTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbilityTargets

`func (o *MatchResponsePlayersInner) SetAbilityTargets(v map[string]interface{})`

SetAbilityTargets sets AbilityTargets field to given value.

### HasAbilityTargets

`func (o *MatchResponsePlayersInner) HasAbilityTargets() bool`

HasAbilityTargets returns a boolean if a field has been set.

### GetDamageTargets

`func (o *MatchResponsePlayersInner) GetDamageTargets() map[string]interface{}`

GetDamageTargets returns the DamageTargets field if non-nil, zero value otherwise.

### GetDamageTargetsOk

`func (o *MatchResponsePlayersInner) GetDamageTargetsOk() (*map[string]interface{}, bool)`

GetDamageTargetsOk returns a tuple with the DamageTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamageTargets

`func (o *MatchResponsePlayersInner) SetDamageTargets(v map[string]interface{})`

SetDamageTargets sets DamageTargets field to given value.

### HasDamageTargets

`func (o *MatchResponsePlayersInner) HasDamageTargets() bool`

HasDamageTargets returns a boolean if a field has been set.

### GetAccountId

`func (o *MatchResponsePlayersInner) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *MatchResponsePlayersInner) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *MatchResponsePlayersInner) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *MatchResponsePlayersInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetActions

`func (o *MatchResponsePlayersInner) GetActions() map[string]interface{}`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *MatchResponsePlayersInner) GetActionsOk() (*map[string]interface{}, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *MatchResponsePlayersInner) SetActions(v map[string]interface{})`

SetActions sets Actions field to given value.

### HasActions

`func (o *MatchResponsePlayersInner) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetAdditionalUnits

`func (o *MatchResponsePlayersInner) GetAdditionalUnits() []map[string]interface{}`

GetAdditionalUnits returns the AdditionalUnits field if non-nil, zero value otherwise.

### GetAdditionalUnitsOk

`func (o *MatchResponsePlayersInner) GetAdditionalUnitsOk() (*[]map[string]interface{}, bool)`

GetAdditionalUnitsOk returns a tuple with the AdditionalUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalUnits

`func (o *MatchResponsePlayersInner) SetAdditionalUnits(v []map[string]interface{})`

SetAdditionalUnits sets AdditionalUnits field to given value.

### HasAdditionalUnits

`func (o *MatchResponsePlayersInner) HasAdditionalUnits() bool`

HasAdditionalUnits returns a boolean if a field has been set.

### SetAdditionalUnitsNil

`func (o *MatchResponsePlayersInner) SetAdditionalUnitsNil(b bool)`

 SetAdditionalUnitsNil sets the value for AdditionalUnits to be an explicit nil

### UnsetAdditionalUnits
`func (o *MatchResponsePlayersInner) UnsetAdditionalUnits()`

UnsetAdditionalUnits ensures that no value is present for AdditionalUnits, not even an explicit nil
### GetAssists

`func (o *MatchResponsePlayersInner) GetAssists() int32`

GetAssists returns the Assists field if non-nil, zero value otherwise.

### GetAssistsOk

`func (o *MatchResponsePlayersInner) GetAssistsOk() (*int32, bool)`

GetAssistsOk returns a tuple with the Assists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssists

`func (o *MatchResponsePlayersInner) SetAssists(v int32)`

SetAssists sets Assists field to given value.

### HasAssists

`func (o *MatchResponsePlayersInner) HasAssists() bool`

HasAssists returns a boolean if a field has been set.

### GetBackpack0

`func (o *MatchResponsePlayersInner) GetBackpack0() int32`

GetBackpack0 returns the Backpack0 field if non-nil, zero value otherwise.

### GetBackpack0Ok

`func (o *MatchResponsePlayersInner) GetBackpack0Ok() (*int32, bool)`

GetBackpack0Ok returns a tuple with the Backpack0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackpack0

`func (o *MatchResponsePlayersInner) SetBackpack0(v int32)`

SetBackpack0 sets Backpack0 field to given value.

### HasBackpack0

`func (o *MatchResponsePlayersInner) HasBackpack0() bool`

HasBackpack0 returns a boolean if a field has been set.

### GetBackpack1

`func (o *MatchResponsePlayersInner) GetBackpack1() int32`

GetBackpack1 returns the Backpack1 field if non-nil, zero value otherwise.

### GetBackpack1Ok

`func (o *MatchResponsePlayersInner) GetBackpack1Ok() (*int32, bool)`

GetBackpack1Ok returns a tuple with the Backpack1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackpack1

`func (o *MatchResponsePlayersInner) SetBackpack1(v int32)`

SetBackpack1 sets Backpack1 field to given value.

### HasBackpack1

`func (o *MatchResponsePlayersInner) HasBackpack1() bool`

HasBackpack1 returns a boolean if a field has been set.

### GetBackpack2

`func (o *MatchResponsePlayersInner) GetBackpack2() int32`

GetBackpack2 returns the Backpack2 field if non-nil, zero value otherwise.

### GetBackpack2Ok

`func (o *MatchResponsePlayersInner) GetBackpack2Ok() (*int32, bool)`

GetBackpack2Ok returns a tuple with the Backpack2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackpack2

`func (o *MatchResponsePlayersInner) SetBackpack2(v int32)`

SetBackpack2 sets Backpack2 field to given value.

### HasBackpack2

`func (o *MatchResponsePlayersInner) HasBackpack2() bool`

HasBackpack2 returns a boolean if a field has been set.

### GetBuybackLog

`func (o *MatchResponsePlayersInner) GetBuybackLog() []MatchResponsePlayersInnerBuybackLogInner`

GetBuybackLog returns the BuybackLog field if non-nil, zero value otherwise.

### GetBuybackLogOk

`func (o *MatchResponsePlayersInner) GetBuybackLogOk() (*[]MatchResponsePlayersInnerBuybackLogInner, bool)`

GetBuybackLogOk returns a tuple with the BuybackLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuybackLog

`func (o *MatchResponsePlayersInner) SetBuybackLog(v []MatchResponsePlayersInnerBuybackLogInner)`

SetBuybackLog sets BuybackLog field to given value.

### HasBuybackLog

`func (o *MatchResponsePlayersInner) HasBuybackLog() bool`

HasBuybackLog returns a boolean if a field has been set.

### GetCampsStacked

`func (o *MatchResponsePlayersInner) GetCampsStacked() int32`

GetCampsStacked returns the CampsStacked field if non-nil, zero value otherwise.

### GetCampsStackedOk

`func (o *MatchResponsePlayersInner) GetCampsStackedOk() (*int32, bool)`

GetCampsStackedOk returns a tuple with the CampsStacked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampsStacked

`func (o *MatchResponsePlayersInner) SetCampsStacked(v int32)`

SetCampsStacked sets CampsStacked field to given value.

### HasCampsStacked

`func (o *MatchResponsePlayersInner) HasCampsStacked() bool`

HasCampsStacked returns a boolean if a field has been set.

### GetConnectionLog

`func (o *MatchResponsePlayersInner) GetConnectionLog() []MatchResponsePlayersInnerConnectionLogInner`

GetConnectionLog returns the ConnectionLog field if non-nil, zero value otherwise.

### GetConnectionLogOk

`func (o *MatchResponsePlayersInner) GetConnectionLogOk() (*[]MatchResponsePlayersInnerConnectionLogInner, bool)`

GetConnectionLogOk returns a tuple with the ConnectionLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionLog

`func (o *MatchResponsePlayersInner) SetConnectionLog(v []MatchResponsePlayersInnerConnectionLogInner)`

SetConnectionLog sets ConnectionLog field to given value.

### HasConnectionLog

`func (o *MatchResponsePlayersInner) HasConnectionLog() bool`

HasConnectionLog returns a boolean if a field has been set.

### GetCreepsStacked

`func (o *MatchResponsePlayersInner) GetCreepsStacked() int32`

GetCreepsStacked returns the CreepsStacked field if non-nil, zero value otherwise.

### GetCreepsStackedOk

`func (o *MatchResponsePlayersInner) GetCreepsStackedOk() (*int32, bool)`

GetCreepsStackedOk returns a tuple with the CreepsStacked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreepsStacked

`func (o *MatchResponsePlayersInner) SetCreepsStacked(v int32)`

SetCreepsStacked sets CreepsStacked field to given value.

### HasCreepsStacked

`func (o *MatchResponsePlayersInner) HasCreepsStacked() bool`

HasCreepsStacked returns a boolean if a field has been set.

### GetDamage

`func (o *MatchResponsePlayersInner) GetDamage() map[string]interface{}`

GetDamage returns the Damage field if non-nil, zero value otherwise.

### GetDamageOk

`func (o *MatchResponsePlayersInner) GetDamageOk() (*map[string]interface{}, bool)`

GetDamageOk returns a tuple with the Damage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamage

`func (o *MatchResponsePlayersInner) SetDamage(v map[string]interface{})`

SetDamage sets Damage field to given value.

### HasDamage

`func (o *MatchResponsePlayersInner) HasDamage() bool`

HasDamage returns a boolean if a field has been set.

### GetDamageInflictor

`func (o *MatchResponsePlayersInner) GetDamageInflictor() map[string]interface{}`

GetDamageInflictor returns the DamageInflictor field if non-nil, zero value otherwise.

### GetDamageInflictorOk

`func (o *MatchResponsePlayersInner) GetDamageInflictorOk() (*map[string]interface{}, bool)`

GetDamageInflictorOk returns a tuple with the DamageInflictor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamageInflictor

`func (o *MatchResponsePlayersInner) SetDamageInflictor(v map[string]interface{})`

SetDamageInflictor sets DamageInflictor field to given value.

### HasDamageInflictor

`func (o *MatchResponsePlayersInner) HasDamageInflictor() bool`

HasDamageInflictor returns a boolean if a field has been set.

### GetDamageInflictorReceived

`func (o *MatchResponsePlayersInner) GetDamageInflictorReceived() map[string]interface{}`

GetDamageInflictorReceived returns the DamageInflictorReceived field if non-nil, zero value otherwise.

### GetDamageInflictorReceivedOk

`func (o *MatchResponsePlayersInner) GetDamageInflictorReceivedOk() (*map[string]interface{}, bool)`

GetDamageInflictorReceivedOk returns a tuple with the DamageInflictorReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamageInflictorReceived

`func (o *MatchResponsePlayersInner) SetDamageInflictorReceived(v map[string]interface{})`

SetDamageInflictorReceived sets DamageInflictorReceived field to given value.

### HasDamageInflictorReceived

`func (o *MatchResponsePlayersInner) HasDamageInflictorReceived() bool`

HasDamageInflictorReceived returns a boolean if a field has been set.

### GetDamageTaken

`func (o *MatchResponsePlayersInner) GetDamageTaken() map[string]interface{}`

GetDamageTaken returns the DamageTaken field if non-nil, zero value otherwise.

### GetDamageTakenOk

`func (o *MatchResponsePlayersInner) GetDamageTakenOk() (*map[string]interface{}, bool)`

GetDamageTakenOk returns a tuple with the DamageTaken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamageTaken

`func (o *MatchResponsePlayersInner) SetDamageTaken(v map[string]interface{})`

SetDamageTaken sets DamageTaken field to given value.

### HasDamageTaken

`func (o *MatchResponsePlayersInner) HasDamageTaken() bool`

HasDamageTaken returns a boolean if a field has been set.

### GetDeaths

`func (o *MatchResponsePlayersInner) GetDeaths() int32`

GetDeaths returns the Deaths field if non-nil, zero value otherwise.

### GetDeathsOk

`func (o *MatchResponsePlayersInner) GetDeathsOk() (*int32, bool)`

GetDeathsOk returns a tuple with the Deaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeaths

`func (o *MatchResponsePlayersInner) SetDeaths(v int32)`

SetDeaths sets Deaths field to given value.

### HasDeaths

`func (o *MatchResponsePlayersInner) HasDeaths() bool`

HasDeaths returns a boolean if a field has been set.

### GetDenies

`func (o *MatchResponsePlayersInner) GetDenies() int32`

GetDenies returns the Denies field if non-nil, zero value otherwise.

### GetDeniesOk

`func (o *MatchResponsePlayersInner) GetDeniesOk() (*int32, bool)`

GetDeniesOk returns a tuple with the Denies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenies

`func (o *MatchResponsePlayersInner) SetDenies(v int32)`

SetDenies sets Denies field to given value.

### HasDenies

`func (o *MatchResponsePlayersInner) HasDenies() bool`

HasDenies returns a boolean if a field has been set.

### GetDnT

`func (o *MatchResponsePlayersInner) GetDnT() []int32`

GetDnT returns the DnT field if non-nil, zero value otherwise.

### GetDnTOk

`func (o *MatchResponsePlayersInner) GetDnTOk() (*[]int32, bool)`

GetDnTOk returns a tuple with the DnT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnT

`func (o *MatchResponsePlayersInner) SetDnT(v []int32)`

SetDnT sets DnT field to given value.

### HasDnT

`func (o *MatchResponsePlayersInner) HasDnT() bool`

HasDnT returns a boolean if a field has been set.

### GetGold

`func (o *MatchResponsePlayersInner) GetGold() int32`

GetGold returns the Gold field if non-nil, zero value otherwise.

### GetGoldOk

`func (o *MatchResponsePlayersInner) GetGoldOk() (*int32, bool)`

GetGoldOk returns a tuple with the Gold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGold

`func (o *MatchResponsePlayersInner) SetGold(v int32)`

SetGold sets Gold field to given value.

### HasGold

`func (o *MatchResponsePlayersInner) HasGold() bool`

HasGold returns a boolean if a field has been set.

### GetGoldPerMin

`func (o *MatchResponsePlayersInner) GetGoldPerMin() int32`

GetGoldPerMin returns the GoldPerMin field if non-nil, zero value otherwise.

### GetGoldPerMinOk

`func (o *MatchResponsePlayersInner) GetGoldPerMinOk() (*int32, bool)`

GetGoldPerMinOk returns a tuple with the GoldPerMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoldPerMin

`func (o *MatchResponsePlayersInner) SetGoldPerMin(v int32)`

SetGoldPerMin sets GoldPerMin field to given value.

### HasGoldPerMin

`func (o *MatchResponsePlayersInner) HasGoldPerMin() bool`

HasGoldPerMin returns a boolean if a field has been set.

### GetGoldReasons

`func (o *MatchResponsePlayersInner) GetGoldReasons() map[string]interface{}`

GetGoldReasons returns the GoldReasons field if non-nil, zero value otherwise.

### GetGoldReasonsOk

`func (o *MatchResponsePlayersInner) GetGoldReasonsOk() (*map[string]interface{}, bool)`

GetGoldReasonsOk returns a tuple with the GoldReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoldReasons

`func (o *MatchResponsePlayersInner) SetGoldReasons(v map[string]interface{})`

SetGoldReasons sets GoldReasons field to given value.

### HasGoldReasons

`func (o *MatchResponsePlayersInner) HasGoldReasons() bool`

HasGoldReasons returns a boolean if a field has been set.

### GetGoldSpent

`func (o *MatchResponsePlayersInner) GetGoldSpent() int32`

GetGoldSpent returns the GoldSpent field if non-nil, zero value otherwise.

### GetGoldSpentOk

`func (o *MatchResponsePlayersInner) GetGoldSpentOk() (*int32, bool)`

GetGoldSpentOk returns a tuple with the GoldSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoldSpent

`func (o *MatchResponsePlayersInner) SetGoldSpent(v int32)`

SetGoldSpent sets GoldSpent field to given value.

### HasGoldSpent

`func (o *MatchResponsePlayersInner) HasGoldSpent() bool`

HasGoldSpent returns a boolean if a field has been set.

### GetGoldT

`func (o *MatchResponsePlayersInner) GetGoldT() []int32`

GetGoldT returns the GoldT field if non-nil, zero value otherwise.

### GetGoldTOk

`func (o *MatchResponsePlayersInner) GetGoldTOk() (*[]int32, bool)`

GetGoldTOk returns a tuple with the GoldT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoldT

`func (o *MatchResponsePlayersInner) SetGoldT(v []int32)`

SetGoldT sets GoldT field to given value.

### HasGoldT

`func (o *MatchResponsePlayersInner) HasGoldT() bool`

HasGoldT returns a boolean if a field has been set.

### GetHeroDamage

`func (o *MatchResponsePlayersInner) GetHeroDamage() int32`

GetHeroDamage returns the HeroDamage field if non-nil, zero value otherwise.

### GetHeroDamageOk

`func (o *MatchResponsePlayersInner) GetHeroDamageOk() (*int32, bool)`

GetHeroDamageOk returns a tuple with the HeroDamage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeroDamage

`func (o *MatchResponsePlayersInner) SetHeroDamage(v int32)`

SetHeroDamage sets HeroDamage field to given value.

### HasHeroDamage

`func (o *MatchResponsePlayersInner) HasHeroDamage() bool`

HasHeroDamage returns a boolean if a field has been set.

### GetHeroHealing

`func (o *MatchResponsePlayersInner) GetHeroHealing() int32`

GetHeroHealing returns the HeroHealing field if non-nil, zero value otherwise.

### GetHeroHealingOk

`func (o *MatchResponsePlayersInner) GetHeroHealingOk() (*int32, bool)`

GetHeroHealingOk returns a tuple with the HeroHealing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeroHealing

`func (o *MatchResponsePlayersInner) SetHeroHealing(v int32)`

SetHeroHealing sets HeroHealing field to given value.

### HasHeroHealing

`func (o *MatchResponsePlayersInner) HasHeroHealing() bool`

HasHeroHealing returns a boolean if a field has been set.

### GetHeroHits

`func (o *MatchResponsePlayersInner) GetHeroHits() map[string]interface{}`

GetHeroHits returns the HeroHits field if non-nil, zero value otherwise.

### GetHeroHitsOk

`func (o *MatchResponsePlayersInner) GetHeroHitsOk() (*map[string]interface{}, bool)`

GetHeroHitsOk returns a tuple with the HeroHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeroHits

`func (o *MatchResponsePlayersInner) SetHeroHits(v map[string]interface{})`

SetHeroHits sets HeroHits field to given value.

### HasHeroHits

`func (o *MatchResponsePlayersInner) HasHeroHits() bool`

HasHeroHits returns a boolean if a field has been set.

### GetHeroId

`func (o *MatchResponsePlayersInner) GetHeroId() int32`

GetHeroId returns the HeroId field if non-nil, zero value otherwise.

### GetHeroIdOk

`func (o *MatchResponsePlayersInner) GetHeroIdOk() (*int32, bool)`

GetHeroIdOk returns a tuple with the HeroId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeroId

`func (o *MatchResponsePlayersInner) SetHeroId(v int32)`

SetHeroId sets HeroId field to given value.

### HasHeroId

`func (o *MatchResponsePlayersInner) HasHeroId() bool`

HasHeroId returns a boolean if a field has been set.

### GetItem0

`func (o *MatchResponsePlayersInner) GetItem0() int32`

GetItem0 returns the Item0 field if non-nil, zero value otherwise.

### GetItem0Ok

`func (o *MatchResponsePlayersInner) GetItem0Ok() (*int32, bool)`

GetItem0Ok returns a tuple with the Item0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem0

`func (o *MatchResponsePlayersInner) SetItem0(v int32)`

SetItem0 sets Item0 field to given value.

### HasItem0

`func (o *MatchResponsePlayersInner) HasItem0() bool`

HasItem0 returns a boolean if a field has been set.

### GetItem1

`func (o *MatchResponsePlayersInner) GetItem1() int32`

GetItem1 returns the Item1 field if non-nil, zero value otherwise.

### GetItem1Ok

`func (o *MatchResponsePlayersInner) GetItem1Ok() (*int32, bool)`

GetItem1Ok returns a tuple with the Item1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem1

`func (o *MatchResponsePlayersInner) SetItem1(v int32)`

SetItem1 sets Item1 field to given value.

### HasItem1

`func (o *MatchResponsePlayersInner) HasItem1() bool`

HasItem1 returns a boolean if a field has been set.

### GetItem2

`func (o *MatchResponsePlayersInner) GetItem2() int32`

GetItem2 returns the Item2 field if non-nil, zero value otherwise.

### GetItem2Ok

`func (o *MatchResponsePlayersInner) GetItem2Ok() (*int32, bool)`

GetItem2Ok returns a tuple with the Item2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem2

`func (o *MatchResponsePlayersInner) SetItem2(v int32)`

SetItem2 sets Item2 field to given value.

### HasItem2

`func (o *MatchResponsePlayersInner) HasItem2() bool`

HasItem2 returns a boolean if a field has been set.

### GetItem3

`func (o *MatchResponsePlayersInner) GetItem3() int32`

GetItem3 returns the Item3 field if non-nil, zero value otherwise.

### GetItem3Ok

`func (o *MatchResponsePlayersInner) GetItem3Ok() (*int32, bool)`

GetItem3Ok returns a tuple with the Item3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem3

`func (o *MatchResponsePlayersInner) SetItem3(v int32)`

SetItem3 sets Item3 field to given value.

### HasItem3

`func (o *MatchResponsePlayersInner) HasItem3() bool`

HasItem3 returns a boolean if a field has been set.

### GetItem4

`func (o *MatchResponsePlayersInner) GetItem4() int32`

GetItem4 returns the Item4 field if non-nil, zero value otherwise.

### GetItem4Ok

`func (o *MatchResponsePlayersInner) GetItem4Ok() (*int32, bool)`

GetItem4Ok returns a tuple with the Item4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem4

`func (o *MatchResponsePlayersInner) SetItem4(v int32)`

SetItem4 sets Item4 field to given value.

### HasItem4

`func (o *MatchResponsePlayersInner) HasItem4() bool`

HasItem4 returns a boolean if a field has been set.

### GetItem5

`func (o *MatchResponsePlayersInner) GetItem5() int32`

GetItem5 returns the Item5 field if non-nil, zero value otherwise.

### GetItem5Ok

`func (o *MatchResponsePlayersInner) GetItem5Ok() (*int32, bool)`

GetItem5Ok returns a tuple with the Item5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem5

`func (o *MatchResponsePlayersInner) SetItem5(v int32)`

SetItem5 sets Item5 field to given value.

### HasItem5

`func (o *MatchResponsePlayersInner) HasItem5() bool`

HasItem5 returns a boolean if a field has been set.

### GetItemUses

`func (o *MatchResponsePlayersInner) GetItemUses() map[string]interface{}`

GetItemUses returns the ItemUses field if non-nil, zero value otherwise.

### GetItemUsesOk

`func (o *MatchResponsePlayersInner) GetItemUsesOk() (*map[string]interface{}, bool)`

GetItemUsesOk returns a tuple with the ItemUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemUses

`func (o *MatchResponsePlayersInner) SetItemUses(v map[string]interface{})`

SetItemUses sets ItemUses field to given value.

### HasItemUses

`func (o *MatchResponsePlayersInner) HasItemUses() bool`

HasItemUses returns a boolean if a field has been set.

### GetKillStreaks

`func (o *MatchResponsePlayersInner) GetKillStreaks() map[string]interface{}`

GetKillStreaks returns the KillStreaks field if non-nil, zero value otherwise.

### GetKillStreaksOk

`func (o *MatchResponsePlayersInner) GetKillStreaksOk() (*map[string]interface{}, bool)`

GetKillStreaksOk returns a tuple with the KillStreaks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillStreaks

`func (o *MatchResponsePlayersInner) SetKillStreaks(v map[string]interface{})`

SetKillStreaks sets KillStreaks field to given value.

### HasKillStreaks

`func (o *MatchResponsePlayersInner) HasKillStreaks() bool`

HasKillStreaks returns a boolean if a field has been set.

### GetKilled

`func (o *MatchResponsePlayersInner) GetKilled() map[string]interface{}`

GetKilled returns the Killed field if non-nil, zero value otherwise.

### GetKilledOk

`func (o *MatchResponsePlayersInner) GetKilledOk() (*map[string]interface{}, bool)`

GetKilledOk returns a tuple with the Killed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKilled

`func (o *MatchResponsePlayersInner) SetKilled(v map[string]interface{})`

SetKilled sets Killed field to given value.

### HasKilled

`func (o *MatchResponsePlayersInner) HasKilled() bool`

HasKilled returns a boolean if a field has been set.

### GetKilledBy

`func (o *MatchResponsePlayersInner) GetKilledBy() map[string]interface{}`

GetKilledBy returns the KilledBy field if non-nil, zero value otherwise.

### GetKilledByOk

`func (o *MatchResponsePlayersInner) GetKilledByOk() (*map[string]interface{}, bool)`

GetKilledByOk returns a tuple with the KilledBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKilledBy

`func (o *MatchResponsePlayersInner) SetKilledBy(v map[string]interface{})`

SetKilledBy sets KilledBy field to given value.

### HasKilledBy

`func (o *MatchResponsePlayersInner) HasKilledBy() bool`

HasKilledBy returns a boolean if a field has been set.

### GetKills

`func (o *MatchResponsePlayersInner) GetKills() int32`

GetKills returns the Kills field if non-nil, zero value otherwise.

### GetKillsOk

`func (o *MatchResponsePlayersInner) GetKillsOk() (*int32, bool)`

GetKillsOk returns a tuple with the Kills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKills

`func (o *MatchResponsePlayersInner) SetKills(v int32)`

SetKills sets Kills field to given value.

### HasKills

`func (o *MatchResponsePlayersInner) HasKills() bool`

HasKills returns a boolean if a field has been set.

### GetKillsLog

`func (o *MatchResponsePlayersInner) GetKillsLog() []MatchResponsePlayersInnerKillsLogInner`

GetKillsLog returns the KillsLog field if non-nil, zero value otherwise.

### GetKillsLogOk

`func (o *MatchResponsePlayersInner) GetKillsLogOk() (*[]MatchResponsePlayersInnerKillsLogInner, bool)`

GetKillsLogOk returns a tuple with the KillsLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillsLog

`func (o *MatchResponsePlayersInner) SetKillsLog(v []MatchResponsePlayersInnerKillsLogInner)`

SetKillsLog sets KillsLog field to given value.

### HasKillsLog

`func (o *MatchResponsePlayersInner) HasKillsLog() bool`

HasKillsLog returns a boolean if a field has been set.

### GetLanePos

`func (o *MatchResponsePlayersInner) GetLanePos() map[string]interface{}`

GetLanePos returns the LanePos field if non-nil, zero value otherwise.

### GetLanePosOk

`func (o *MatchResponsePlayersInner) GetLanePosOk() (*map[string]interface{}, bool)`

GetLanePosOk returns a tuple with the LanePos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanePos

`func (o *MatchResponsePlayersInner) SetLanePos(v map[string]interface{})`

SetLanePos sets LanePos field to given value.

### HasLanePos

`func (o *MatchResponsePlayersInner) HasLanePos() bool`

HasLanePos returns a boolean if a field has been set.

### GetLastHits

`func (o *MatchResponsePlayersInner) GetLastHits() int32`

GetLastHits returns the LastHits field if non-nil, zero value otherwise.

### GetLastHitsOk

`func (o *MatchResponsePlayersInner) GetLastHitsOk() (*int32, bool)`

GetLastHitsOk returns a tuple with the LastHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHits

`func (o *MatchResponsePlayersInner) SetLastHits(v int32)`

SetLastHits sets LastHits field to given value.

### HasLastHits

`func (o *MatchResponsePlayersInner) HasLastHits() bool`

HasLastHits returns a boolean if a field has been set.

### GetLeaverStatus

`func (o *MatchResponsePlayersInner) GetLeaverStatus() int32`

GetLeaverStatus returns the LeaverStatus field if non-nil, zero value otherwise.

### GetLeaverStatusOk

`func (o *MatchResponsePlayersInner) GetLeaverStatusOk() (*int32, bool)`

GetLeaverStatusOk returns a tuple with the LeaverStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaverStatus

`func (o *MatchResponsePlayersInner) SetLeaverStatus(v int32)`

SetLeaverStatus sets LeaverStatus field to given value.

### HasLeaverStatus

`func (o *MatchResponsePlayersInner) HasLeaverStatus() bool`

HasLeaverStatus returns a boolean if a field has been set.

### GetLevel

`func (o *MatchResponsePlayersInner) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *MatchResponsePlayersInner) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *MatchResponsePlayersInner) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *MatchResponsePlayersInner) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetLhT

`func (o *MatchResponsePlayersInner) GetLhT() []int32`

GetLhT returns the LhT field if non-nil, zero value otherwise.

### GetLhTOk

`func (o *MatchResponsePlayersInner) GetLhTOk() (*[]int32, bool)`

GetLhTOk returns a tuple with the LhT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLhT

`func (o *MatchResponsePlayersInner) SetLhT(v []int32)`

SetLhT sets LhT field to given value.

### HasLhT

`func (o *MatchResponsePlayersInner) HasLhT() bool`

HasLhT returns a boolean if a field has been set.

### GetLifeState

`func (o *MatchResponsePlayersInner) GetLifeState() map[string]interface{}`

GetLifeState returns the LifeState field if non-nil, zero value otherwise.

### GetLifeStateOk

`func (o *MatchResponsePlayersInner) GetLifeStateOk() (*map[string]interface{}, bool)`

GetLifeStateOk returns a tuple with the LifeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeState

`func (o *MatchResponsePlayersInner) SetLifeState(v map[string]interface{})`

SetLifeState sets LifeState field to given value.

### HasLifeState

`func (o *MatchResponsePlayersInner) HasLifeState() bool`

HasLifeState returns a boolean if a field has been set.

### GetMaxHeroHit

`func (o *MatchResponsePlayersInner) GetMaxHeroHit() map[string]interface{}`

GetMaxHeroHit returns the MaxHeroHit field if non-nil, zero value otherwise.

### GetMaxHeroHitOk

`func (o *MatchResponsePlayersInner) GetMaxHeroHitOk() (*map[string]interface{}, bool)`

GetMaxHeroHitOk returns a tuple with the MaxHeroHit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHeroHit

`func (o *MatchResponsePlayersInner) SetMaxHeroHit(v map[string]interface{})`

SetMaxHeroHit sets MaxHeroHit field to given value.

### HasMaxHeroHit

`func (o *MatchResponsePlayersInner) HasMaxHeroHit() bool`

HasMaxHeroHit returns a boolean if a field has been set.

### GetMultiKills

`func (o *MatchResponsePlayersInner) GetMultiKills() map[string]interface{}`

GetMultiKills returns the MultiKills field if non-nil, zero value otherwise.

### GetMultiKillsOk

`func (o *MatchResponsePlayersInner) GetMultiKillsOk() (*map[string]interface{}, bool)`

GetMultiKillsOk returns a tuple with the MultiKills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiKills

`func (o *MatchResponsePlayersInner) SetMultiKills(v map[string]interface{})`

SetMultiKills sets MultiKills field to given value.

### HasMultiKills

`func (o *MatchResponsePlayersInner) HasMultiKills() bool`

HasMultiKills returns a boolean if a field has been set.

### GetObs

`func (o *MatchResponsePlayersInner) GetObs() map[string]interface{}`

GetObs returns the Obs field if non-nil, zero value otherwise.

### GetObsOk

`func (o *MatchResponsePlayersInner) GetObsOk() (*map[string]interface{}, bool)`

GetObsOk returns a tuple with the Obs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObs

`func (o *MatchResponsePlayersInner) SetObs(v map[string]interface{})`

SetObs sets Obs field to given value.

### HasObs

`func (o *MatchResponsePlayersInner) HasObs() bool`

HasObs returns a boolean if a field has been set.

### GetObsLeftLog

`func (o *MatchResponsePlayersInner) GetObsLeftLog() []map[string]interface{}`

GetObsLeftLog returns the ObsLeftLog field if non-nil, zero value otherwise.

### GetObsLeftLogOk

`func (o *MatchResponsePlayersInner) GetObsLeftLogOk() (*[]map[string]interface{}, bool)`

GetObsLeftLogOk returns a tuple with the ObsLeftLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObsLeftLog

`func (o *MatchResponsePlayersInner) SetObsLeftLog(v []map[string]interface{})`

SetObsLeftLog sets ObsLeftLog field to given value.

### HasObsLeftLog

`func (o *MatchResponsePlayersInner) HasObsLeftLog() bool`

HasObsLeftLog returns a boolean if a field has been set.

### GetObsLog

`func (o *MatchResponsePlayersInner) GetObsLog() []map[string]interface{}`

GetObsLog returns the ObsLog field if non-nil, zero value otherwise.

### GetObsLogOk

`func (o *MatchResponsePlayersInner) GetObsLogOk() (*[]map[string]interface{}, bool)`

GetObsLogOk returns a tuple with the ObsLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObsLog

`func (o *MatchResponsePlayersInner) SetObsLog(v []map[string]interface{})`

SetObsLog sets ObsLog field to given value.

### HasObsLog

`func (o *MatchResponsePlayersInner) HasObsLog() bool`

HasObsLog returns a boolean if a field has been set.

### GetObsPlaced

`func (o *MatchResponsePlayersInner) GetObsPlaced() int32`

GetObsPlaced returns the ObsPlaced field if non-nil, zero value otherwise.

### GetObsPlacedOk

`func (o *MatchResponsePlayersInner) GetObsPlacedOk() (*int32, bool)`

GetObsPlacedOk returns a tuple with the ObsPlaced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObsPlaced

`func (o *MatchResponsePlayersInner) SetObsPlaced(v int32)`

SetObsPlaced sets ObsPlaced field to given value.

### HasObsPlaced

`func (o *MatchResponsePlayersInner) HasObsPlaced() bool`

HasObsPlaced returns a boolean if a field has been set.

### GetPartyId

`func (o *MatchResponsePlayersInner) GetPartyId() int32`

GetPartyId returns the PartyId field if non-nil, zero value otherwise.

### GetPartyIdOk

`func (o *MatchResponsePlayersInner) GetPartyIdOk() (*int32, bool)`

GetPartyIdOk returns a tuple with the PartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyId

`func (o *MatchResponsePlayersInner) SetPartyId(v int32)`

SetPartyId sets PartyId field to given value.

### HasPartyId

`func (o *MatchResponsePlayersInner) HasPartyId() bool`

HasPartyId returns a boolean if a field has been set.

### GetPermanentBuffs

`func (o *MatchResponsePlayersInner) GetPermanentBuffs() []map[string]interface{}`

GetPermanentBuffs returns the PermanentBuffs field if non-nil, zero value otherwise.

### GetPermanentBuffsOk

`func (o *MatchResponsePlayersInner) GetPermanentBuffsOk() (*[]map[string]interface{}, bool)`

GetPermanentBuffsOk returns a tuple with the PermanentBuffs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanentBuffs

`func (o *MatchResponsePlayersInner) SetPermanentBuffs(v []map[string]interface{})`

SetPermanentBuffs sets PermanentBuffs field to given value.

### HasPermanentBuffs

`func (o *MatchResponsePlayersInner) HasPermanentBuffs() bool`

HasPermanentBuffs returns a boolean if a field has been set.

### GetPings

`func (o *MatchResponsePlayersInner) GetPings() int32`

GetPings returns the Pings field if non-nil, zero value otherwise.

### GetPingsOk

`func (o *MatchResponsePlayersInner) GetPingsOk() (*int32, bool)`

GetPingsOk returns a tuple with the Pings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPings

`func (o *MatchResponsePlayersInner) SetPings(v int32)`

SetPings sets Pings field to given value.

### HasPings

`func (o *MatchResponsePlayersInner) HasPings() bool`

HasPings returns a boolean if a field has been set.

### GetPurchase

`func (o *MatchResponsePlayersInner) GetPurchase() map[string]interface{}`

GetPurchase returns the Purchase field if non-nil, zero value otherwise.

### GetPurchaseOk

`func (o *MatchResponsePlayersInner) GetPurchaseOk() (*map[string]interface{}, bool)`

GetPurchaseOk returns a tuple with the Purchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchase

`func (o *MatchResponsePlayersInner) SetPurchase(v map[string]interface{})`

SetPurchase sets Purchase field to given value.

### HasPurchase

`func (o *MatchResponsePlayersInner) HasPurchase() bool`

HasPurchase returns a boolean if a field has been set.

### GetPurchaseLog

`func (o *MatchResponsePlayersInner) GetPurchaseLog() []MatchResponsePlayersInnerPurchaseLogInner`

GetPurchaseLog returns the PurchaseLog field if non-nil, zero value otherwise.

### GetPurchaseLogOk

`func (o *MatchResponsePlayersInner) GetPurchaseLogOk() (*[]MatchResponsePlayersInnerPurchaseLogInner, bool)`

GetPurchaseLogOk returns a tuple with the PurchaseLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseLog

`func (o *MatchResponsePlayersInner) SetPurchaseLog(v []MatchResponsePlayersInnerPurchaseLogInner)`

SetPurchaseLog sets PurchaseLog field to given value.

### HasPurchaseLog

`func (o *MatchResponsePlayersInner) HasPurchaseLog() bool`

HasPurchaseLog returns a boolean if a field has been set.

### GetRunePickups

`func (o *MatchResponsePlayersInner) GetRunePickups() int32`

GetRunePickups returns the RunePickups field if non-nil, zero value otherwise.

### GetRunePickupsOk

`func (o *MatchResponsePlayersInner) GetRunePickupsOk() (*int32, bool)`

GetRunePickupsOk returns a tuple with the RunePickups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunePickups

`func (o *MatchResponsePlayersInner) SetRunePickups(v int32)`

SetRunePickups sets RunePickups field to given value.

### HasRunePickups

`func (o *MatchResponsePlayersInner) HasRunePickups() bool`

HasRunePickups returns a boolean if a field has been set.

### GetRunes

`func (o *MatchResponsePlayersInner) GetRunes() map[string]int32`

GetRunes returns the Runes field if non-nil, zero value otherwise.

### GetRunesOk

`func (o *MatchResponsePlayersInner) GetRunesOk() (*map[string]int32, bool)`

GetRunesOk returns a tuple with the Runes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunes

`func (o *MatchResponsePlayersInner) SetRunes(v map[string]int32)`

SetRunes sets Runes field to given value.

### HasRunes

`func (o *MatchResponsePlayersInner) HasRunes() bool`

HasRunes returns a boolean if a field has been set.

### GetRunesLog

`func (o *MatchResponsePlayersInner) GetRunesLog() []MatchResponsePlayersInnerRunesLogInner`

GetRunesLog returns the RunesLog field if non-nil, zero value otherwise.

### GetRunesLogOk

`func (o *MatchResponsePlayersInner) GetRunesLogOk() (*[]MatchResponsePlayersInnerRunesLogInner, bool)`

GetRunesLogOk returns a tuple with the RunesLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunesLog

`func (o *MatchResponsePlayersInner) SetRunesLog(v []MatchResponsePlayersInnerRunesLogInner)`

SetRunesLog sets RunesLog field to given value.

### HasRunesLog

`func (o *MatchResponsePlayersInner) HasRunesLog() bool`

HasRunesLog returns a boolean if a field has been set.

### GetSen

`func (o *MatchResponsePlayersInner) GetSen() map[string]interface{}`

GetSen returns the Sen field if non-nil, zero value otherwise.

### GetSenOk

`func (o *MatchResponsePlayersInner) GetSenOk() (*map[string]interface{}, bool)`

GetSenOk returns a tuple with the Sen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSen

`func (o *MatchResponsePlayersInner) SetSen(v map[string]interface{})`

SetSen sets Sen field to given value.

### HasSen

`func (o *MatchResponsePlayersInner) HasSen() bool`

HasSen returns a boolean if a field has been set.

### GetSenLeftLog

`func (o *MatchResponsePlayersInner) GetSenLeftLog() []map[string]interface{}`

GetSenLeftLog returns the SenLeftLog field if non-nil, zero value otherwise.

### GetSenLeftLogOk

`func (o *MatchResponsePlayersInner) GetSenLeftLogOk() (*[]map[string]interface{}, bool)`

GetSenLeftLogOk returns a tuple with the SenLeftLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenLeftLog

`func (o *MatchResponsePlayersInner) SetSenLeftLog(v []map[string]interface{})`

SetSenLeftLog sets SenLeftLog field to given value.

### HasSenLeftLog

`func (o *MatchResponsePlayersInner) HasSenLeftLog() bool`

HasSenLeftLog returns a boolean if a field has been set.

### GetSenLog

`func (o *MatchResponsePlayersInner) GetSenLog() []map[string]interface{}`

GetSenLog returns the SenLog field if non-nil, zero value otherwise.

### GetSenLogOk

`func (o *MatchResponsePlayersInner) GetSenLogOk() (*[]map[string]interface{}, bool)`

GetSenLogOk returns a tuple with the SenLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenLog

`func (o *MatchResponsePlayersInner) SetSenLog(v []map[string]interface{})`

SetSenLog sets SenLog field to given value.

### HasSenLog

`func (o *MatchResponsePlayersInner) HasSenLog() bool`

HasSenLog returns a boolean if a field has been set.

### GetSenPlaced

`func (o *MatchResponsePlayersInner) GetSenPlaced() int32`

GetSenPlaced returns the SenPlaced field if non-nil, zero value otherwise.

### GetSenPlacedOk

`func (o *MatchResponsePlayersInner) GetSenPlacedOk() (*int32, bool)`

GetSenPlacedOk returns a tuple with the SenPlaced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenPlaced

`func (o *MatchResponsePlayersInner) SetSenPlaced(v int32)`

SetSenPlaced sets SenPlaced field to given value.

### HasSenPlaced

`func (o *MatchResponsePlayersInner) HasSenPlaced() bool`

HasSenPlaced returns a boolean if a field has been set.

### GetStuns

`func (o *MatchResponsePlayersInner) GetStuns() float32`

GetStuns returns the Stuns field if non-nil, zero value otherwise.

### GetStunsOk

`func (o *MatchResponsePlayersInner) GetStunsOk() (*float32, bool)`

GetStunsOk returns a tuple with the Stuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStuns

`func (o *MatchResponsePlayersInner) SetStuns(v float32)`

SetStuns sets Stuns field to given value.

### HasStuns

`func (o *MatchResponsePlayersInner) HasStuns() bool`

HasStuns returns a boolean if a field has been set.

### GetTimes

`func (o *MatchResponsePlayersInner) GetTimes() []int32`

GetTimes returns the Times field if non-nil, zero value otherwise.

### GetTimesOk

`func (o *MatchResponsePlayersInner) GetTimesOk() (*[]int32, bool)`

GetTimesOk returns a tuple with the Times field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimes

`func (o *MatchResponsePlayersInner) SetTimes(v []int32)`

SetTimes sets Times field to given value.

### HasTimes

`func (o *MatchResponsePlayersInner) HasTimes() bool`

HasTimes returns a boolean if a field has been set.

### GetTowerDamage

`func (o *MatchResponsePlayersInner) GetTowerDamage() int32`

GetTowerDamage returns the TowerDamage field if non-nil, zero value otherwise.

### GetTowerDamageOk

`func (o *MatchResponsePlayersInner) GetTowerDamageOk() (*int32, bool)`

GetTowerDamageOk returns a tuple with the TowerDamage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTowerDamage

`func (o *MatchResponsePlayersInner) SetTowerDamage(v int32)`

SetTowerDamage sets TowerDamage field to given value.

### HasTowerDamage

`func (o *MatchResponsePlayersInner) HasTowerDamage() bool`

HasTowerDamage returns a boolean if a field has been set.

### GetXpPerMin

`func (o *MatchResponsePlayersInner) GetXpPerMin() int32`

GetXpPerMin returns the XpPerMin field if non-nil, zero value otherwise.

### GetXpPerMinOk

`func (o *MatchResponsePlayersInner) GetXpPerMinOk() (*int32, bool)`

GetXpPerMinOk returns a tuple with the XpPerMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXpPerMin

`func (o *MatchResponsePlayersInner) SetXpPerMin(v int32)`

SetXpPerMin sets XpPerMin field to given value.

### HasXpPerMin

`func (o *MatchResponsePlayersInner) HasXpPerMin() bool`

HasXpPerMin returns a boolean if a field has been set.

### GetXpReasons

`func (o *MatchResponsePlayersInner) GetXpReasons() map[string]interface{}`

GetXpReasons returns the XpReasons field if non-nil, zero value otherwise.

### GetXpReasonsOk

`func (o *MatchResponsePlayersInner) GetXpReasonsOk() (*map[string]interface{}, bool)`

GetXpReasonsOk returns a tuple with the XpReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXpReasons

`func (o *MatchResponsePlayersInner) SetXpReasons(v map[string]interface{})`

SetXpReasons sets XpReasons field to given value.

### HasXpReasons

`func (o *MatchResponsePlayersInner) HasXpReasons() bool`

HasXpReasons returns a boolean if a field has been set.

### GetXpT

`func (o *MatchResponsePlayersInner) GetXpT() []int32`

GetXpT returns the XpT field if non-nil, zero value otherwise.

### GetXpTOk

`func (o *MatchResponsePlayersInner) GetXpTOk() (*[]int32, bool)`

GetXpTOk returns a tuple with the XpT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXpT

`func (o *MatchResponsePlayersInner) SetXpT(v []int32)`

SetXpT sets XpT field to given value.

### HasXpT

`func (o *MatchResponsePlayersInner) HasXpT() bool`

HasXpT returns a boolean if a field has been set.

### GetPersonaname

`func (o *MatchResponsePlayersInner) GetPersonaname() string`

GetPersonaname returns the Personaname field if non-nil, zero value otherwise.

### GetPersonanameOk

`func (o *MatchResponsePlayersInner) GetPersonanameOk() (*string, bool)`

GetPersonanameOk returns a tuple with the Personaname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonaname

`func (o *MatchResponsePlayersInner) SetPersonaname(v string)`

SetPersonaname sets Personaname field to given value.

### HasPersonaname

`func (o *MatchResponsePlayersInner) HasPersonaname() bool`

HasPersonaname returns a boolean if a field has been set.

### SetPersonanameNil

`func (o *MatchResponsePlayersInner) SetPersonanameNil(b bool)`

 SetPersonanameNil sets the value for Personaname to be an explicit nil

### UnsetPersonaname
`func (o *MatchResponsePlayersInner) UnsetPersonaname()`

UnsetPersonaname ensures that no value is present for Personaname, not even an explicit nil
### GetName

`func (o *MatchResponsePlayersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MatchResponsePlayersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MatchResponsePlayersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MatchResponsePlayersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MatchResponsePlayersInner) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MatchResponsePlayersInner) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLastLogin

`func (o *MatchResponsePlayersInner) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *MatchResponsePlayersInner) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *MatchResponsePlayersInner) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *MatchResponsePlayersInner) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### SetLastLoginNil

`func (o *MatchResponsePlayersInner) SetLastLoginNil(b bool)`

 SetLastLoginNil sets the value for LastLogin to be an explicit nil

### UnsetLastLogin
`func (o *MatchResponsePlayersInner) UnsetLastLogin()`

UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
### GetRadiantWin

`func (o *MatchResponsePlayersInner) GetRadiantWin() bool`

GetRadiantWin returns the RadiantWin field if non-nil, zero value otherwise.

### GetRadiantWinOk

`func (o *MatchResponsePlayersInner) GetRadiantWinOk() (*bool, bool)`

GetRadiantWinOk returns a tuple with the RadiantWin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiantWin

`func (o *MatchResponsePlayersInner) SetRadiantWin(v bool)`

SetRadiantWin sets RadiantWin field to given value.

### HasRadiantWin

`func (o *MatchResponsePlayersInner) HasRadiantWin() bool`

HasRadiantWin returns a boolean if a field has been set.

### SetRadiantWinNil

`func (o *MatchResponsePlayersInner) SetRadiantWinNil(b bool)`

 SetRadiantWinNil sets the value for RadiantWin to be an explicit nil

### UnsetRadiantWin
`func (o *MatchResponsePlayersInner) UnsetRadiantWin()`

UnsetRadiantWin ensures that no value is present for RadiantWin, not even an explicit nil
### GetStartTime

`func (o *MatchResponsePlayersInner) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MatchResponsePlayersInner) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MatchResponsePlayersInner) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MatchResponsePlayersInner) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetDuration

`func (o *MatchResponsePlayersInner) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *MatchResponsePlayersInner) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *MatchResponsePlayersInner) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *MatchResponsePlayersInner) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetCluster

`func (o *MatchResponsePlayersInner) GetCluster() int32`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *MatchResponsePlayersInner) GetClusterOk() (*int32, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *MatchResponsePlayersInner) SetCluster(v int32)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *MatchResponsePlayersInner) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetLobbyType

`func (o *MatchResponsePlayersInner) GetLobbyType() int32`

GetLobbyType returns the LobbyType field if non-nil, zero value otherwise.

### GetLobbyTypeOk

`func (o *MatchResponsePlayersInner) GetLobbyTypeOk() (*int32, bool)`

GetLobbyTypeOk returns a tuple with the LobbyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLobbyType

`func (o *MatchResponsePlayersInner) SetLobbyType(v int32)`

SetLobbyType sets LobbyType field to given value.

### HasLobbyType

`func (o *MatchResponsePlayersInner) HasLobbyType() bool`

HasLobbyType returns a boolean if a field has been set.

### GetGameMode

`func (o *MatchResponsePlayersInner) GetGameMode() int32`

GetGameMode returns the GameMode field if non-nil, zero value otherwise.

### GetGameModeOk

`func (o *MatchResponsePlayersInner) GetGameModeOk() (*int32, bool)`

GetGameModeOk returns a tuple with the GameMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameMode

`func (o *MatchResponsePlayersInner) SetGameMode(v int32)`

SetGameMode sets GameMode field to given value.

### HasGameMode

`func (o *MatchResponsePlayersInner) HasGameMode() bool`

HasGameMode returns a boolean if a field has been set.

### GetPatch

`func (o *MatchResponsePlayersInner) GetPatch() int32`

GetPatch returns the Patch field if non-nil, zero value otherwise.

### GetPatchOk

`func (o *MatchResponsePlayersInner) GetPatchOk() (*int32, bool)`

GetPatchOk returns a tuple with the Patch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatch

`func (o *MatchResponsePlayersInner) SetPatch(v int32)`

SetPatch sets Patch field to given value.

### HasPatch

`func (o *MatchResponsePlayersInner) HasPatch() bool`

HasPatch returns a boolean if a field has been set.

### GetRegion

`func (o *MatchResponsePlayersInner) GetRegion() int32`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *MatchResponsePlayersInner) GetRegionOk() (*int32, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *MatchResponsePlayersInner) SetRegion(v int32)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *MatchResponsePlayersInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetIsRadiant

`func (o *MatchResponsePlayersInner) GetIsRadiant() bool`

GetIsRadiant returns the IsRadiant field if non-nil, zero value otherwise.

### GetIsRadiantOk

`func (o *MatchResponsePlayersInner) GetIsRadiantOk() (*bool, bool)`

GetIsRadiantOk returns a tuple with the IsRadiant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRadiant

`func (o *MatchResponsePlayersInner) SetIsRadiant(v bool)`

SetIsRadiant sets IsRadiant field to given value.

### HasIsRadiant

`func (o *MatchResponsePlayersInner) HasIsRadiant() bool`

HasIsRadiant returns a boolean if a field has been set.

### GetWin

`func (o *MatchResponsePlayersInner) GetWin() int32`

GetWin returns the Win field if non-nil, zero value otherwise.

### GetWinOk

`func (o *MatchResponsePlayersInner) GetWinOk() (*int32, bool)`

GetWinOk returns a tuple with the Win field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWin

`func (o *MatchResponsePlayersInner) SetWin(v int32)`

SetWin sets Win field to given value.

### HasWin

`func (o *MatchResponsePlayersInner) HasWin() bool`

HasWin returns a boolean if a field has been set.

### GetLose

`func (o *MatchResponsePlayersInner) GetLose() int32`

GetLose returns the Lose field if non-nil, zero value otherwise.

### GetLoseOk

`func (o *MatchResponsePlayersInner) GetLoseOk() (*int32, bool)`

GetLoseOk returns a tuple with the Lose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLose

`func (o *MatchResponsePlayersInner) SetLose(v int32)`

SetLose sets Lose field to given value.

### HasLose

`func (o *MatchResponsePlayersInner) HasLose() bool`

HasLose returns a boolean if a field has been set.

### GetTotalGold

`func (o *MatchResponsePlayersInner) GetTotalGold() int32`

GetTotalGold returns the TotalGold field if non-nil, zero value otherwise.

### GetTotalGoldOk

`func (o *MatchResponsePlayersInner) GetTotalGoldOk() (*int32, bool)`

GetTotalGoldOk returns a tuple with the TotalGold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGold

`func (o *MatchResponsePlayersInner) SetTotalGold(v int32)`

SetTotalGold sets TotalGold field to given value.

### HasTotalGold

`func (o *MatchResponsePlayersInner) HasTotalGold() bool`

HasTotalGold returns a boolean if a field has been set.

### GetTotalXp

`func (o *MatchResponsePlayersInner) GetTotalXp() int32`

GetTotalXp returns the TotalXp field if non-nil, zero value otherwise.

### GetTotalXpOk

`func (o *MatchResponsePlayersInner) GetTotalXpOk() (*int32, bool)`

GetTotalXpOk returns a tuple with the TotalXp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalXp

`func (o *MatchResponsePlayersInner) SetTotalXp(v int32)`

SetTotalXp sets TotalXp field to given value.

### HasTotalXp

`func (o *MatchResponsePlayersInner) HasTotalXp() bool`

HasTotalXp returns a boolean if a field has been set.

### GetKillsPerMin

`func (o *MatchResponsePlayersInner) GetKillsPerMin() float32`

GetKillsPerMin returns the KillsPerMin field if non-nil, zero value otherwise.

### GetKillsPerMinOk

`func (o *MatchResponsePlayersInner) GetKillsPerMinOk() (*float32, bool)`

GetKillsPerMinOk returns a tuple with the KillsPerMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKillsPerMin

`func (o *MatchResponsePlayersInner) SetKillsPerMin(v float32)`

SetKillsPerMin sets KillsPerMin field to given value.

### HasKillsPerMin

`func (o *MatchResponsePlayersInner) HasKillsPerMin() bool`

HasKillsPerMin returns a boolean if a field has been set.

### GetKda

`func (o *MatchResponsePlayersInner) GetKda() float32`

GetKda returns the Kda field if non-nil, zero value otherwise.

### GetKdaOk

`func (o *MatchResponsePlayersInner) GetKdaOk() (*float32, bool)`

GetKdaOk returns a tuple with the Kda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKda

`func (o *MatchResponsePlayersInner) SetKda(v float32)`

SetKda sets Kda field to given value.

### HasKda

`func (o *MatchResponsePlayersInner) HasKda() bool`

HasKda returns a boolean if a field has been set.

### GetAbandons

`func (o *MatchResponsePlayersInner) GetAbandons() int32`

GetAbandons returns the Abandons field if non-nil, zero value otherwise.

### GetAbandonsOk

`func (o *MatchResponsePlayersInner) GetAbandonsOk() (*int32, bool)`

GetAbandonsOk returns a tuple with the Abandons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbandons

`func (o *MatchResponsePlayersInner) SetAbandons(v int32)`

SetAbandons sets Abandons field to given value.

### HasAbandons

`func (o *MatchResponsePlayersInner) HasAbandons() bool`

HasAbandons returns a boolean if a field has been set.

### GetNeutralKills

`func (o *MatchResponsePlayersInner) GetNeutralKills() int32`

GetNeutralKills returns the NeutralKills field if non-nil, zero value otherwise.

### GetNeutralKillsOk

`func (o *MatchResponsePlayersInner) GetNeutralKillsOk() (*int32, bool)`

GetNeutralKillsOk returns a tuple with the NeutralKills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeutralKills

`func (o *MatchResponsePlayersInner) SetNeutralKills(v int32)`

SetNeutralKills sets NeutralKills field to given value.

### HasNeutralKills

`func (o *MatchResponsePlayersInner) HasNeutralKills() bool`

HasNeutralKills returns a boolean if a field has been set.

### GetTowerKills

`func (o *MatchResponsePlayersInner) GetTowerKills() int32`

GetTowerKills returns the TowerKills field if non-nil, zero value otherwise.

### GetTowerKillsOk

`func (o *MatchResponsePlayersInner) GetTowerKillsOk() (*int32, bool)`

GetTowerKillsOk returns a tuple with the TowerKills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTowerKills

`func (o *MatchResponsePlayersInner) SetTowerKills(v int32)`

SetTowerKills sets TowerKills field to given value.

### HasTowerKills

`func (o *MatchResponsePlayersInner) HasTowerKills() bool`

HasTowerKills returns a boolean if a field has been set.

### GetCourierKills

`func (o *MatchResponsePlayersInner) GetCourierKills() int32`

GetCourierKills returns the CourierKills field if non-nil, zero value otherwise.

### GetCourierKillsOk

`func (o *MatchResponsePlayersInner) GetCourierKillsOk() (*int32, bool)`

GetCourierKillsOk returns a tuple with the CourierKills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourierKills

`func (o *MatchResponsePlayersInner) SetCourierKills(v int32)`

SetCourierKills sets CourierKills field to given value.

### HasCourierKills

`func (o *MatchResponsePlayersInner) HasCourierKills() bool`

HasCourierKills returns a boolean if a field has been set.

### GetLaneKills

`func (o *MatchResponsePlayersInner) GetLaneKills() int32`

GetLaneKills returns the LaneKills field if non-nil, zero value otherwise.

### GetLaneKillsOk

`func (o *MatchResponsePlayersInner) GetLaneKillsOk() (*int32, bool)`

GetLaneKillsOk returns a tuple with the LaneKills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaneKills

`func (o *MatchResponsePlayersInner) SetLaneKills(v int32)`

SetLaneKills sets LaneKills field to given value.

### HasLaneKills

`func (o *MatchResponsePlayersInner) HasLaneKills() bool`

HasLaneKills returns a boolean if a field has been set.

### GetHeroKills

`func (o *MatchResponsePlayersInner) GetHeroKills() int32`

GetHeroKills returns the HeroKills field if non-nil, zero value otherwise.

### GetHeroKillsOk

`func (o *MatchResponsePlayersInner) GetHeroKillsOk() (*int32, bool)`

GetHeroKillsOk returns a tuple with the HeroKills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeroKills

`func (o *MatchResponsePlayersInner) SetHeroKills(v int32)`

SetHeroKills sets HeroKills field to given value.

### HasHeroKills

`func (o *MatchResponsePlayersInner) HasHeroKills() bool`

HasHeroKills returns a boolean if a field has been set.

### GetObserverKills

`func (o *MatchResponsePlayersInner) GetObserverKills() int32`

GetObserverKills returns the ObserverKills field if non-nil, zero value otherwise.

### GetObserverKillsOk

`func (o *MatchResponsePlayersInner) GetObserverKillsOk() (*int32, bool)`

GetObserverKillsOk returns a tuple with the ObserverKills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObserverKills

`func (o *MatchResponsePlayersInner) SetObserverKills(v int32)`

SetObserverKills sets ObserverKills field to given value.

### HasObserverKills

`func (o *MatchResponsePlayersInner) HasObserverKills() bool`

HasObserverKills returns a boolean if a field has been set.

### GetSentryKills

`func (o *MatchResponsePlayersInner) GetSentryKills() int32`

GetSentryKills returns the SentryKills field if non-nil, zero value otherwise.

### GetSentryKillsOk

`func (o *MatchResponsePlayersInner) GetSentryKillsOk() (*int32, bool)`

GetSentryKillsOk returns a tuple with the SentryKills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentryKills

`func (o *MatchResponsePlayersInner) SetSentryKills(v int32)`

SetSentryKills sets SentryKills field to given value.

### HasSentryKills

`func (o *MatchResponsePlayersInner) HasSentryKills() bool`

HasSentryKills returns a boolean if a field has been set.

### GetRoshanKills

`func (o *MatchResponsePlayersInner) GetRoshanKills() int32`

GetRoshanKills returns the RoshanKills field if non-nil, zero value otherwise.

### GetRoshanKillsOk

`func (o *MatchResponsePlayersInner) GetRoshanKillsOk() (*int32, bool)`

GetRoshanKillsOk returns a tuple with the RoshanKills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoshanKills

`func (o *MatchResponsePlayersInner) SetRoshanKills(v int32)`

SetRoshanKills sets RoshanKills field to given value.

### HasRoshanKills

`func (o *MatchResponsePlayersInner) HasRoshanKills() bool`

HasRoshanKills returns a boolean if a field has been set.

### GetNecronomiconKills

`func (o *MatchResponsePlayersInner) GetNecronomiconKills() int32`

GetNecronomiconKills returns the NecronomiconKills field if non-nil, zero value otherwise.

### GetNecronomiconKillsOk

`func (o *MatchResponsePlayersInner) GetNecronomiconKillsOk() (*int32, bool)`

GetNecronomiconKillsOk returns a tuple with the NecronomiconKills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNecronomiconKills

`func (o *MatchResponsePlayersInner) SetNecronomiconKills(v int32)`

SetNecronomiconKills sets NecronomiconKills field to given value.

### HasNecronomiconKills

`func (o *MatchResponsePlayersInner) HasNecronomiconKills() bool`

HasNecronomiconKills returns a boolean if a field has been set.

### GetAncientKills

`func (o *MatchResponsePlayersInner) GetAncientKills() int32`

GetAncientKills returns the AncientKills field if non-nil, zero value otherwise.

### GetAncientKillsOk

`func (o *MatchResponsePlayersInner) GetAncientKillsOk() (*int32, bool)`

GetAncientKillsOk returns a tuple with the AncientKills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncientKills

`func (o *MatchResponsePlayersInner) SetAncientKills(v int32)`

SetAncientKills sets AncientKills field to given value.

### HasAncientKills

`func (o *MatchResponsePlayersInner) HasAncientKills() bool`

HasAncientKills returns a boolean if a field has been set.

### GetBuybackCount

`func (o *MatchResponsePlayersInner) GetBuybackCount() int32`

GetBuybackCount returns the BuybackCount field if non-nil, zero value otherwise.

### GetBuybackCountOk

`func (o *MatchResponsePlayersInner) GetBuybackCountOk() (*int32, bool)`

GetBuybackCountOk returns a tuple with the BuybackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuybackCount

`func (o *MatchResponsePlayersInner) SetBuybackCount(v int32)`

SetBuybackCount sets BuybackCount field to given value.

### HasBuybackCount

`func (o *MatchResponsePlayersInner) HasBuybackCount() bool`

HasBuybackCount returns a boolean if a field has been set.

### GetObserverUses

`func (o *MatchResponsePlayersInner) GetObserverUses() int32`

GetObserverUses returns the ObserverUses field if non-nil, zero value otherwise.

### GetObserverUsesOk

`func (o *MatchResponsePlayersInner) GetObserverUsesOk() (*int32, bool)`

GetObserverUsesOk returns a tuple with the ObserverUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObserverUses

`func (o *MatchResponsePlayersInner) SetObserverUses(v int32)`

SetObserverUses sets ObserverUses field to given value.

### HasObserverUses

`func (o *MatchResponsePlayersInner) HasObserverUses() bool`

HasObserverUses returns a boolean if a field has been set.

### GetSentryUses

`func (o *MatchResponsePlayersInner) GetSentryUses() int32`

GetSentryUses returns the SentryUses field if non-nil, zero value otherwise.

### GetSentryUsesOk

`func (o *MatchResponsePlayersInner) GetSentryUsesOk() (*int32, bool)`

GetSentryUsesOk returns a tuple with the SentryUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentryUses

`func (o *MatchResponsePlayersInner) SetSentryUses(v int32)`

SetSentryUses sets SentryUses field to given value.

### HasSentryUses

`func (o *MatchResponsePlayersInner) HasSentryUses() bool`

HasSentryUses returns a boolean if a field has been set.

### GetLaneEfficiency

`func (o *MatchResponsePlayersInner) GetLaneEfficiency() float32`

GetLaneEfficiency returns the LaneEfficiency field if non-nil, zero value otherwise.

### GetLaneEfficiencyOk

`func (o *MatchResponsePlayersInner) GetLaneEfficiencyOk() (*float32, bool)`

GetLaneEfficiencyOk returns a tuple with the LaneEfficiency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaneEfficiency

`func (o *MatchResponsePlayersInner) SetLaneEfficiency(v float32)`

SetLaneEfficiency sets LaneEfficiency field to given value.

### HasLaneEfficiency

`func (o *MatchResponsePlayersInner) HasLaneEfficiency() bool`

HasLaneEfficiency returns a boolean if a field has been set.

### GetLaneEfficiencyPct

`func (o *MatchResponsePlayersInner) GetLaneEfficiencyPct() float32`

GetLaneEfficiencyPct returns the LaneEfficiencyPct field if non-nil, zero value otherwise.

### GetLaneEfficiencyPctOk

`func (o *MatchResponsePlayersInner) GetLaneEfficiencyPctOk() (*float32, bool)`

GetLaneEfficiencyPctOk returns a tuple with the LaneEfficiencyPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaneEfficiencyPct

`func (o *MatchResponsePlayersInner) SetLaneEfficiencyPct(v float32)`

SetLaneEfficiencyPct sets LaneEfficiencyPct field to given value.

### HasLaneEfficiencyPct

`func (o *MatchResponsePlayersInner) HasLaneEfficiencyPct() bool`

HasLaneEfficiencyPct returns a boolean if a field has been set.

### GetLane

`func (o *MatchResponsePlayersInner) GetLane() int32`

GetLane returns the Lane field if non-nil, zero value otherwise.

### GetLaneOk

`func (o *MatchResponsePlayersInner) GetLaneOk() (*int32, bool)`

GetLaneOk returns a tuple with the Lane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLane

`func (o *MatchResponsePlayersInner) SetLane(v int32)`

SetLane sets Lane field to given value.

### HasLane

`func (o *MatchResponsePlayersInner) HasLane() bool`

HasLane returns a boolean if a field has been set.

### SetLaneNil

`func (o *MatchResponsePlayersInner) SetLaneNil(b bool)`

 SetLaneNil sets the value for Lane to be an explicit nil

### UnsetLane
`func (o *MatchResponsePlayersInner) UnsetLane()`

UnsetLane ensures that no value is present for Lane, not even an explicit nil
### GetLaneRole

`func (o *MatchResponsePlayersInner) GetLaneRole() int32`

GetLaneRole returns the LaneRole field if non-nil, zero value otherwise.

### GetLaneRoleOk

`func (o *MatchResponsePlayersInner) GetLaneRoleOk() (*int32, bool)`

GetLaneRoleOk returns a tuple with the LaneRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaneRole

`func (o *MatchResponsePlayersInner) SetLaneRole(v int32)`

SetLaneRole sets LaneRole field to given value.

### HasLaneRole

`func (o *MatchResponsePlayersInner) HasLaneRole() bool`

HasLaneRole returns a boolean if a field has been set.

### SetLaneRoleNil

`func (o *MatchResponsePlayersInner) SetLaneRoleNil(b bool)`

 SetLaneRoleNil sets the value for LaneRole to be an explicit nil

### UnsetLaneRole
`func (o *MatchResponsePlayersInner) UnsetLaneRole()`

UnsetLaneRole ensures that no value is present for LaneRole, not even an explicit nil
### GetIsRoaming

`func (o *MatchResponsePlayersInner) GetIsRoaming() bool`

GetIsRoaming returns the IsRoaming field if non-nil, zero value otherwise.

### GetIsRoamingOk

`func (o *MatchResponsePlayersInner) GetIsRoamingOk() (*bool, bool)`

GetIsRoamingOk returns a tuple with the IsRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRoaming

`func (o *MatchResponsePlayersInner) SetIsRoaming(v bool)`

SetIsRoaming sets IsRoaming field to given value.

### HasIsRoaming

`func (o *MatchResponsePlayersInner) HasIsRoaming() bool`

HasIsRoaming returns a boolean if a field has been set.

### SetIsRoamingNil

`func (o *MatchResponsePlayersInner) SetIsRoamingNil(b bool)`

 SetIsRoamingNil sets the value for IsRoaming to be an explicit nil

### UnsetIsRoaming
`func (o *MatchResponsePlayersInner) UnsetIsRoaming()`

UnsetIsRoaming ensures that no value is present for IsRoaming, not even an explicit nil
### GetPurchaseTime

`func (o *MatchResponsePlayersInner) GetPurchaseTime() map[string]interface{}`

GetPurchaseTime returns the PurchaseTime field if non-nil, zero value otherwise.

### GetPurchaseTimeOk

`func (o *MatchResponsePlayersInner) GetPurchaseTimeOk() (*map[string]interface{}, bool)`

GetPurchaseTimeOk returns a tuple with the PurchaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseTime

`func (o *MatchResponsePlayersInner) SetPurchaseTime(v map[string]interface{})`

SetPurchaseTime sets PurchaseTime field to given value.

### HasPurchaseTime

`func (o *MatchResponsePlayersInner) HasPurchaseTime() bool`

HasPurchaseTime returns a boolean if a field has been set.

### GetFirstPurchaseTime

`func (o *MatchResponsePlayersInner) GetFirstPurchaseTime() map[string]interface{}`

GetFirstPurchaseTime returns the FirstPurchaseTime field if non-nil, zero value otherwise.

### GetFirstPurchaseTimeOk

`func (o *MatchResponsePlayersInner) GetFirstPurchaseTimeOk() (*map[string]interface{}, bool)`

GetFirstPurchaseTimeOk returns a tuple with the FirstPurchaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstPurchaseTime

`func (o *MatchResponsePlayersInner) SetFirstPurchaseTime(v map[string]interface{})`

SetFirstPurchaseTime sets FirstPurchaseTime field to given value.

### HasFirstPurchaseTime

`func (o *MatchResponsePlayersInner) HasFirstPurchaseTime() bool`

HasFirstPurchaseTime returns a boolean if a field has been set.

### GetItemWin

`func (o *MatchResponsePlayersInner) GetItemWin() map[string]interface{}`

GetItemWin returns the ItemWin field if non-nil, zero value otherwise.

### GetItemWinOk

`func (o *MatchResponsePlayersInner) GetItemWinOk() (*map[string]interface{}, bool)`

GetItemWinOk returns a tuple with the ItemWin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemWin

`func (o *MatchResponsePlayersInner) SetItemWin(v map[string]interface{})`

SetItemWin sets ItemWin field to given value.

### HasItemWin

`func (o *MatchResponsePlayersInner) HasItemWin() bool`

HasItemWin returns a boolean if a field has been set.

### GetItemUsage

`func (o *MatchResponsePlayersInner) GetItemUsage() map[string]interface{}`

GetItemUsage returns the ItemUsage field if non-nil, zero value otherwise.

### GetItemUsageOk

`func (o *MatchResponsePlayersInner) GetItemUsageOk() (*map[string]interface{}, bool)`

GetItemUsageOk returns a tuple with the ItemUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemUsage

`func (o *MatchResponsePlayersInner) SetItemUsage(v map[string]interface{})`

SetItemUsage sets ItemUsage field to given value.

### HasItemUsage

`func (o *MatchResponsePlayersInner) HasItemUsage() bool`

HasItemUsage returns a boolean if a field has been set.

### GetPurchaseTpscroll

`func (o *MatchResponsePlayersInner) GetPurchaseTpscroll() int32`

GetPurchaseTpscroll returns the PurchaseTpscroll field if non-nil, zero value otherwise.

### GetPurchaseTpscrollOk

`func (o *MatchResponsePlayersInner) GetPurchaseTpscrollOk() (*int32, bool)`

GetPurchaseTpscrollOk returns a tuple with the PurchaseTpscroll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseTpscroll

`func (o *MatchResponsePlayersInner) SetPurchaseTpscroll(v int32)`

SetPurchaseTpscroll sets PurchaseTpscroll field to given value.

### HasPurchaseTpscroll

`func (o *MatchResponsePlayersInner) HasPurchaseTpscroll() bool`

HasPurchaseTpscroll returns a boolean if a field has been set.

### GetActionsPerMin

`func (o *MatchResponsePlayersInner) GetActionsPerMin() int32`

GetActionsPerMin returns the ActionsPerMin field if non-nil, zero value otherwise.

### GetActionsPerMinOk

`func (o *MatchResponsePlayersInner) GetActionsPerMinOk() (*int32, bool)`

GetActionsPerMinOk returns a tuple with the ActionsPerMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionsPerMin

`func (o *MatchResponsePlayersInner) SetActionsPerMin(v int32)`

SetActionsPerMin sets ActionsPerMin field to given value.

### HasActionsPerMin

`func (o *MatchResponsePlayersInner) HasActionsPerMin() bool`

HasActionsPerMin returns a boolean if a field has been set.

### GetLifeStateDead

`func (o *MatchResponsePlayersInner) GetLifeStateDead() int32`

GetLifeStateDead returns the LifeStateDead field if non-nil, zero value otherwise.

### GetLifeStateDeadOk

`func (o *MatchResponsePlayersInner) GetLifeStateDeadOk() (*int32, bool)`

GetLifeStateDeadOk returns a tuple with the LifeStateDead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeStateDead

`func (o *MatchResponsePlayersInner) SetLifeStateDead(v int32)`

SetLifeStateDead sets LifeStateDead field to given value.

### HasLifeStateDead

`func (o *MatchResponsePlayersInner) HasLifeStateDead() bool`

HasLifeStateDead returns a boolean if a field has been set.

### GetRankTier

`func (o *MatchResponsePlayersInner) GetRankTier() int32`

GetRankTier returns the RankTier field if non-nil, zero value otherwise.

### GetRankTierOk

`func (o *MatchResponsePlayersInner) GetRankTierOk() (*int32, bool)`

GetRankTierOk returns a tuple with the RankTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankTier

`func (o *MatchResponsePlayersInner) SetRankTier(v int32)`

SetRankTier sets RankTier field to given value.

### HasRankTier

`func (o *MatchResponsePlayersInner) HasRankTier() bool`

HasRankTier returns a boolean if a field has been set.

### GetCosmetics

`func (o *MatchResponsePlayersInner) GetCosmetics() []MatchResponsePlayersInnerCosmeticsInner`

GetCosmetics returns the Cosmetics field if non-nil, zero value otherwise.

### GetCosmeticsOk

`func (o *MatchResponsePlayersInner) GetCosmeticsOk() (*[]MatchResponsePlayersInnerCosmeticsInner, bool)`

GetCosmeticsOk returns a tuple with the Cosmetics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosmetics

`func (o *MatchResponsePlayersInner) SetCosmetics(v []MatchResponsePlayersInnerCosmeticsInner)`

SetCosmetics sets Cosmetics field to given value.

### HasCosmetics

`func (o *MatchResponsePlayersInner) HasCosmetics() bool`

HasCosmetics returns a boolean if a field has been set.

### GetBenchmarks

`func (o *MatchResponsePlayersInner) GetBenchmarks() map[string]interface{}`

GetBenchmarks returns the Benchmarks field if non-nil, zero value otherwise.

### GetBenchmarksOk

`func (o *MatchResponsePlayersInner) GetBenchmarksOk() (*map[string]interface{}, bool)`

GetBenchmarksOk returns a tuple with the Benchmarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBenchmarks

`func (o *MatchResponsePlayersInner) SetBenchmarks(v map[string]interface{})`

SetBenchmarks sets Benchmarks field to given value.

### HasBenchmarks

`func (o *MatchResponsePlayersInner) HasBenchmarks() bool`

HasBenchmarks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


