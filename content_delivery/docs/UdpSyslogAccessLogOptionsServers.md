# UdpSyslogAccessLogOptionsServers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogFormat** | Pointer to **string** | The name of the log format to use as the default.  The log formats are defined by the operator. | [optional] 
**Server** | **string** | Server:\\&lt;port\\&gt; of syslog server. | 

## Methods

### NewUdpSyslogAccessLogOptionsServers

`func NewUdpSyslogAccessLogOptionsServers(server string, ) *UdpSyslogAccessLogOptionsServers`

NewUdpSyslogAccessLogOptionsServers instantiates a new UdpSyslogAccessLogOptionsServers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUdpSyslogAccessLogOptionsServersWithDefaults

`func NewUdpSyslogAccessLogOptionsServersWithDefaults() *UdpSyslogAccessLogOptionsServers`

NewUdpSyslogAccessLogOptionsServersWithDefaults instantiates a new UdpSyslogAccessLogOptionsServers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogFormat

`func (o *UdpSyslogAccessLogOptionsServers) GetLogFormat() string`

GetLogFormat returns the LogFormat field if non-nil, zero value otherwise.

### GetLogFormatOk

`func (o *UdpSyslogAccessLogOptionsServers) GetLogFormatOk() (*string, bool)`

GetLogFormatOk returns a tuple with the LogFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFormat

`func (o *UdpSyslogAccessLogOptionsServers) SetLogFormat(v string)`

SetLogFormat sets LogFormat field to given value.

### HasLogFormat

`func (o *UdpSyslogAccessLogOptionsServers) HasLogFormat() bool`

HasLogFormat returns a boolean if a field has been set.

### GetServer

`func (o *UdpSyslogAccessLogOptionsServers) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *UdpSyslogAccessLogOptionsServers) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *UdpSyslogAccessLogOptionsServers) SetServer(v string)`

SetServer sets Server field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


