# IssueSubscriberTransferTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransferDestinationOperatorEmail** | **string** | 移管先オペレーターEmail | 
**TransferDestinationOperatorId** | **string** | 移管先オペレーターID | 
**TransferImsi** | **[]string** | 移管する IMSI リスト | 

## Methods

### NewIssueSubscriberTransferTokenRequest

`func NewIssueSubscriberTransferTokenRequest(transferDestinationOperatorEmail string, transferDestinationOperatorId string, transferImsi []string, ) *IssueSubscriberTransferTokenRequest`

NewIssueSubscriberTransferTokenRequest instantiates a new IssueSubscriberTransferTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueSubscriberTransferTokenRequestWithDefaults

`func NewIssueSubscriberTransferTokenRequestWithDefaults() *IssueSubscriberTransferTokenRequest`

NewIssueSubscriberTransferTokenRequestWithDefaults instantiates a new IssueSubscriberTransferTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransferDestinationOperatorEmail

`func (o *IssueSubscriberTransferTokenRequest) GetTransferDestinationOperatorEmail() string`

GetTransferDestinationOperatorEmail returns the TransferDestinationOperatorEmail field if non-nil, zero value otherwise.

### GetTransferDestinationOperatorEmailOk

`func (o *IssueSubscriberTransferTokenRequest) GetTransferDestinationOperatorEmailOk() (*string, bool)`

GetTransferDestinationOperatorEmailOk returns a tuple with the TransferDestinationOperatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferDestinationOperatorEmail

`func (o *IssueSubscriberTransferTokenRequest) SetTransferDestinationOperatorEmail(v string)`

SetTransferDestinationOperatorEmail sets TransferDestinationOperatorEmail field to given value.


### GetTransferDestinationOperatorId

`func (o *IssueSubscriberTransferTokenRequest) GetTransferDestinationOperatorId() string`

GetTransferDestinationOperatorId returns the TransferDestinationOperatorId field if non-nil, zero value otherwise.

### GetTransferDestinationOperatorIdOk

`func (o *IssueSubscriberTransferTokenRequest) GetTransferDestinationOperatorIdOk() (*string, bool)`

GetTransferDestinationOperatorIdOk returns a tuple with the TransferDestinationOperatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferDestinationOperatorId

`func (o *IssueSubscriberTransferTokenRequest) SetTransferDestinationOperatorId(v string)`

SetTransferDestinationOperatorId sets TransferDestinationOperatorId field to given value.


### GetTransferImsi

`func (o *IssueSubscriberTransferTokenRequest) GetTransferImsi() []string`

GetTransferImsi returns the TransferImsi field if non-nil, zero value otherwise.

### GetTransferImsiOk

`func (o *IssueSubscriberTransferTokenRequest) GetTransferImsiOk() (*[]string, bool)`

GetTransferImsiOk returns a tuple with the TransferImsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferImsi

`func (o *IssueSubscriberTransferTokenRequest) SetTransferImsi(v []string)`

SetTransferImsi sets TransferImsi field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


