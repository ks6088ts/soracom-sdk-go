# RoutingFilterEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Whether to allow or deny the outbound packets with a destination in the specified range | 
**IpRange** | **string** | IPv4 address range in CIDR format, e.g. a.b.c.d/x | 

## Methods

### NewRoutingFilterEntry

`func NewRoutingFilterEntry(action string, ipRange string, ) *RoutingFilterEntry`

NewRoutingFilterEntry instantiates a new RoutingFilterEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingFilterEntryWithDefaults

`func NewRoutingFilterEntryWithDefaults() *RoutingFilterEntry`

NewRoutingFilterEntryWithDefaults instantiates a new RoutingFilterEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *RoutingFilterEntry) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RoutingFilterEntry) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RoutingFilterEntry) SetAction(v string)`

SetAction sets Action field to given value.


### GetIpRange

`func (o *RoutingFilterEntry) GetIpRange() string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *RoutingFilterEntry) GetIpRangeOk() (*string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *RoutingFilterEntry) SetIpRange(v string)`

SetIpRange sets IpRange field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


