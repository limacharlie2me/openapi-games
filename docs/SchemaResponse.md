# SchemaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TableName** | Pointer to **string** | table_name | [optional] 
**ColumnName** | Pointer to **string** | column_name | [optional] 
**DataType** | Pointer to **string** | data_type | [optional] 

## Methods

### NewSchemaResponse

`func NewSchemaResponse() *SchemaResponse`

NewSchemaResponse instantiates a new SchemaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaResponseWithDefaults

`func NewSchemaResponseWithDefaults() *SchemaResponse`

NewSchemaResponseWithDefaults instantiates a new SchemaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTableName

`func (o *SchemaResponse) GetTableName() string`

GetTableName returns the TableName field if non-nil, zero value otherwise.

### GetTableNameOk

`func (o *SchemaResponse) GetTableNameOk() (*string, bool)`

GetTableNameOk returns a tuple with the TableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableName

`func (o *SchemaResponse) SetTableName(v string)`

SetTableName sets TableName field to given value.

### HasTableName

`func (o *SchemaResponse) HasTableName() bool`

HasTableName returns a boolean if a field has been set.

### GetColumnName

`func (o *SchemaResponse) GetColumnName() string`

GetColumnName returns the ColumnName field if non-nil, zero value otherwise.

### GetColumnNameOk

`func (o *SchemaResponse) GetColumnNameOk() (*string, bool)`

GetColumnNameOk returns a tuple with the ColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnName

`func (o *SchemaResponse) SetColumnName(v string)`

SetColumnName sets ColumnName field to given value.

### HasColumnName

`func (o *SchemaResponse) HasColumnName() bool`

HasColumnName returns a boolean if a field has been set.

### GetDataType

`func (o *SchemaResponse) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *SchemaResponse) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *SchemaResponse) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *SchemaResponse) HasDataType() bool`

HasDataType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


