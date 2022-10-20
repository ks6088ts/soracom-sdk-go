# LagoonDashboardPermissionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DashboardId** | Pointer to **float32** |  | [optional] 
**DashboardTitle** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**[]LagoonDashboardPermissionsResponsePermissionsInner**](LagoonDashboardPermissionsResponsePermissionsInner.md) |  | [optional] 

## Methods

### NewLagoonDashboardPermissionsResponse

`func NewLagoonDashboardPermissionsResponse() *LagoonDashboardPermissionsResponse`

NewLagoonDashboardPermissionsResponse instantiates a new LagoonDashboardPermissionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLagoonDashboardPermissionsResponseWithDefaults

`func NewLagoonDashboardPermissionsResponseWithDefaults() *LagoonDashboardPermissionsResponse`

NewLagoonDashboardPermissionsResponseWithDefaults instantiates a new LagoonDashboardPermissionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDashboardId

`func (o *LagoonDashboardPermissionsResponse) GetDashboardId() float32`

GetDashboardId returns the DashboardId field if non-nil, zero value otherwise.

### GetDashboardIdOk

`func (o *LagoonDashboardPermissionsResponse) GetDashboardIdOk() (*float32, bool)`

GetDashboardIdOk returns a tuple with the DashboardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardId

`func (o *LagoonDashboardPermissionsResponse) SetDashboardId(v float32)`

SetDashboardId sets DashboardId field to given value.

### HasDashboardId

`func (o *LagoonDashboardPermissionsResponse) HasDashboardId() bool`

HasDashboardId returns a boolean if a field has been set.

### GetDashboardTitle

`func (o *LagoonDashboardPermissionsResponse) GetDashboardTitle() string`

GetDashboardTitle returns the DashboardTitle field if non-nil, zero value otherwise.

### GetDashboardTitleOk

`func (o *LagoonDashboardPermissionsResponse) GetDashboardTitleOk() (*string, bool)`

GetDashboardTitleOk returns a tuple with the DashboardTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardTitle

`func (o *LagoonDashboardPermissionsResponse) SetDashboardTitle(v string)`

SetDashboardTitle sets DashboardTitle field to given value.

### HasDashboardTitle

`func (o *LagoonDashboardPermissionsResponse) HasDashboardTitle() bool`

HasDashboardTitle returns a boolean if a field has been set.

### GetPermissions

`func (o *LagoonDashboardPermissionsResponse) GetPermissions() []LagoonDashboardPermissionsResponsePermissionsInner`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *LagoonDashboardPermissionsResponse) GetPermissionsOk() (*[]LagoonDashboardPermissionsResponsePermissionsInner, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *LagoonDashboardPermissionsResponse) SetPermissions(v []LagoonDashboardPermissionsResponsePermissionsInner)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *LagoonDashboardPermissionsResponse) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


