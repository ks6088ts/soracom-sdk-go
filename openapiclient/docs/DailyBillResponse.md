# DailyBillResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillList** | Pointer to [**[]DailyBill**](DailyBill.md) | 日ごとの利用料金 | [optional] 

## Methods

### NewDailyBillResponse

`func NewDailyBillResponse() *DailyBillResponse`

NewDailyBillResponse instantiates a new DailyBillResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyBillResponseWithDefaults

`func NewDailyBillResponseWithDefaults() *DailyBillResponse`

NewDailyBillResponseWithDefaults instantiates a new DailyBillResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillList

`func (o *DailyBillResponse) GetBillList() []DailyBill`

GetBillList returns the BillList field if non-nil, zero value otherwise.

### GetBillListOk

`func (o *DailyBillResponse) GetBillListOk() (*[]DailyBill, bool)`

GetBillListOk returns a tuple with the BillList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillList

`func (o *DailyBillResponse) SetBillList(v []DailyBill)`

SetBillList sets BillList field to given value.

### HasBillList

`func (o *DailyBillResponse) HasBillList() bool`

HasBillList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


