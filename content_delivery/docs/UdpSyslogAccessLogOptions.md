# UdpSyslogAccessLogOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultLogFormat** | **string** | The name of the log format to use as the default.  The log formats are defined by the operator. | 
**Facility** | Pointer to **string** | The syslog facility to use when transmitting logs. Possible values: &#x60;LOCAL0&#x60;, &#x60;LOCAL1&#x60;, &#x60;LOCAL2&#x60;, &#x60;LOCAL3&#x60;, &#x60;LOCAL4&#x60;, &#x60;LOCAL5&#x60;, &#x60;LOCAL6&#x60;, &#x60;LOCAL7&#x60;, &#x60;KERN&#x60;, &#x60;USER&#x60;, &#x60;MAIL&#x60;, &#x60;DAEMON&#x60;, &#x60;AUTH&#x60;, &#x60;INTERN&#x60;, &#x60;LPR&#x60;, &#x60;NEWS&#x60;, &#x60;UUCP&#x60;, &#x60;CLOCK&#x60;, &#x60;AUTHPRIV&#x60;, &#x60;FTP&#x60;, &#x60;NTP&#x60;, &#x60;AUDIT&#x60;, &#x60;ALERT&#x60;, &#x60;CRON&#x60;. Default value: &#x60;LOCAL7&#x60;. | [optional] [default to "LOCAL7"]
**Severity** | Pointer to **string** | The syslog severity to use when transmitting logs. Possible values: &#x60;DEBUG&#x60;, &#x60;INFO&#x60;, &#x60;NOTICE&#x60;, &#x60;WARNING&#x60;, &#x60;ERROR&#x60;, &#x60;CRITICAL&#x60;, &#x60;ALERT&#x60;, &#x60;EMERGENCY&#x60;. Default value: &#x60;INFO&#x60;. | [optional] [default to "INFO"]
**Servers** | Pointer to [**[]UdpSyslogAccessLogOptionsServers**](UdpSyslogAccessLogOptionsServers.md) | List of servers to send logs to.  Maximum of 6. | [optional] 

## Methods

### NewUdpSyslogAccessLogOptions

`func NewUdpSyslogAccessLogOptions(defaultLogFormat string, ) *UdpSyslogAccessLogOptions`

NewUdpSyslogAccessLogOptions instantiates a new UdpSyslogAccessLogOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUdpSyslogAccessLogOptionsWithDefaults

`func NewUdpSyslogAccessLogOptionsWithDefaults() *UdpSyslogAccessLogOptions`

NewUdpSyslogAccessLogOptionsWithDefaults instantiates a new UdpSyslogAccessLogOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultLogFormat

`func (o *UdpSyslogAccessLogOptions) GetDefaultLogFormat() string`

GetDefaultLogFormat returns the DefaultLogFormat field if non-nil, zero value otherwise.

### GetDefaultLogFormatOk

`func (o *UdpSyslogAccessLogOptions) GetDefaultLogFormatOk() (*string, bool)`

GetDefaultLogFormatOk returns a tuple with the DefaultLogFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLogFormat

`func (o *UdpSyslogAccessLogOptions) SetDefaultLogFormat(v string)`

SetDefaultLogFormat sets DefaultLogFormat field to given value.


### GetFacility

`func (o *UdpSyslogAccessLogOptions) GetFacility() string`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *UdpSyslogAccessLogOptions) GetFacilityOk() (*string, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *UdpSyslogAccessLogOptions) SetFacility(v string)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *UdpSyslogAccessLogOptions) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetSeverity

`func (o *UdpSyslogAccessLogOptions) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *UdpSyslogAccessLogOptions) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *UdpSyslogAccessLogOptions) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *UdpSyslogAccessLogOptions) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetServers

`func (o *UdpSyslogAccessLogOptions) GetServers() []UdpSyslogAccessLogOptionsServers`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *UdpSyslogAccessLogOptions) GetServersOk() (*[]UdpSyslogAccessLogOptionsServers, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *UdpSyslogAccessLogOptions) SetServers(v []UdpSyslogAccessLogOptionsServers)`

SetServers sets Servers field to given value.

### HasServers

`func (o *UdpSyslogAccessLogOptions) HasServers() bool`

HasServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


