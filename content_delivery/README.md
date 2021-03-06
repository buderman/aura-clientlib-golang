# Go API client for content_delivery

Aura LCDN Content Delivery API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./content_delivery"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://aura.akamai.com:443/api/content-delivery/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**CreateCdnPrefix**](docs/DefaultApi.md#createcdnprefix) | **Post** /content-providers/{contentProviderId}/cdn-prefixes | cdn prefixes
*DefaultApi* | [**CreateOrigin**](docs/DefaultApi.md#createorigin) | **Post** /content-providers/{contentProviderId}/origins | origins
*DefaultApi* | [**CreateSharedSecretSet**](docs/DefaultApi.md#createsharedsecretset) | **Post** /content-providers/{contentProviderId}/shared-secret-sets | shared secret sets
*DefaultApi* | [**CreateTlsDeliveryProfile**](docs/DefaultApi.md#createtlsdeliveryprofile) | **Post** /content-providers/{contentProviderId}/tls-delivery-profiles | tls delivery profiles
*DefaultApi* | [**CreateTlsIngestProfiles**](docs/DefaultApi.md#createtlsingestprofiles) | **Post** /content-providers/{contentProviderId}/tls-ingest-profiles | tls ingest profiles
*DefaultApi* | [**DeleteCdnPrefix**](docs/DefaultApi.md#deletecdnprefix) | **Delete** /content-providers/{contentProviderId}/cdn-prefixes/{cdnPrefixId} | cdn prefix
*DefaultApi* | [**DeleteOrigin**](docs/DefaultApi.md#deleteorigin) | **Delete** /content-providers/{contentProviderId}/origins/{originId} | origin
*DefaultApi* | [**DeleteSharedSecretSet**](docs/DefaultApi.md#deletesharedsecretset) | **Delete** /content-providers/{contentProviderId}/shared-secret-sets/{sharedSecretSetId} | shared secret set
*DefaultApi* | [**DeleteTlsDeliveryProfile**](docs/DefaultApi.md#deletetlsdeliveryprofile) | **Delete** /content-providers/{contentProviderId}/tls-delivery-profiles/{tlsDeliveryProfileId} | tls delivery profile
*DefaultApi* | [**DeleteTlsIngestProfile**](docs/DefaultApi.md#deletetlsingestprofile) | **Delete** /content-providers/{contentProviderId}/tls-ingest-profiles/{tlsIngestProfileId} | tls ingest profile
*DefaultApi* | [**ListCdnPrefixes**](docs/DefaultApi.md#listcdnprefixes) | **Get** /content-providers/{contentProviderId}/cdn-prefixes | cdn prefixes
*DefaultApi* | [**ListContentProviders**](docs/DefaultApi.md#listcontentproviders) | **Get** /content-providers | content providers
*DefaultApi* | [**ListOrigins**](docs/DefaultApi.md#listorigins) | **Get** /content-providers/{contentProviderId}/origins | origins
*DefaultApi* | [**ListSharedSecretSets**](docs/DefaultApi.md#listsharedsecretsets) | **Get** /content-providers/{contentProviderId}/shared-secret-sets | shared secret sets
*DefaultApi* | [**ListTlsDeliveryProfileCdnPrefixes**](docs/DefaultApi.md#listtlsdeliveryprofilecdnprefixes) | **Get** /content-providers/{contentProviderId}/tls-delivery-profiles/{tlsDeliveryProfileId}/cdn-prefixes | cdn prefixes per tls
*DefaultApi* | [**ListTlsDeliveryProfiles**](docs/DefaultApi.md#listtlsdeliveryprofiles) | **Get** /content-providers/{contentProviderId}/tls-delivery-profiles | tls delivery profiles
*DefaultApi* | [**ListTlsIngestProfileOrigins**](docs/DefaultApi.md#listtlsingestprofileorigins) | **Get** /content-providers/{contentProviderId}/tls-ingest-profiles/{tlsIngestProfileId}/origins | origin per tls ingest profile
*DefaultApi* | [**ListTlsIngestProfiles**](docs/DefaultApi.md#listtlsingestprofiles) | **Get** /content-providers/{contentProviderId}/tls-ingest-profiles | tls ingest profiles
*DefaultApi* | [**ReadCdnPrefix**](docs/DefaultApi.md#readcdnprefix) | **Get** /content-providers/{contentProviderId}/cdn-prefixes/{cdnPrefixId} | cdn prefix
*DefaultApi* | [**ReadCdnPrefixRules**](docs/DefaultApi.md#readcdnprefixrules) | **Get** /content-providers/{contentProviderId}/cdn-prefixes/{cdnPrefixId}/rules | rules
*DefaultApi* | [**ReadContentProvider**](docs/DefaultApi.md#readcontentprovider) | **Get** /content-providers/{contentProviderId} | content provider
*DefaultApi* | [**ReadOrigin**](docs/DefaultApi.md#readorigin) | **Get** /content-providers/{contentProviderId}/origins/{originId} | origin
*DefaultApi* | [**ReadSharedSecretSet**](docs/DefaultApi.md#readsharedsecretset) | **Get** /content-providers/{contentProviderId}/shared-secret-sets/{sharedSecretSetId} | shared secret set
*DefaultApi* | [**ReadSharedSecretSetSecrets**](docs/DefaultApi.md#readsharedsecretsetsecrets) | **Get** /content-providers/{contentProviderId}/shared-secret-sets/{sharedSecretSetId}/secrets | secrets
*DefaultApi* | [**ReadTlsDeliveryProfile**](docs/DefaultApi.md#readtlsdeliveryprofile) | **Get** /content-providers/{contentProviderId}/tls-delivery-profiles/{tlsDeliveryProfileId} | tls delivery profile
*DefaultApi* | [**ReadTlsIngestProfile**](docs/DefaultApi.md#readtlsingestprofile) | **Get** /content-providers/{contentProviderId}/tls-ingest-profiles/{tlsIngestProfileId} | tls ingest profile
*DefaultApi* | [**UpdateCdnPrefix**](docs/DefaultApi.md#updatecdnprefix) | **Put** /content-providers/{contentProviderId}/cdn-prefixes/{cdnPrefixId} | cdn prefix
*DefaultApi* | [**UpdateCdnPrefixRules**](docs/DefaultApi.md#updatecdnprefixrules) | **Put** /content-providers/{contentProviderId}/cdn-prefixes/{cdnPrefixId}/rules | rules
*DefaultApi* | [**UpdateOrigin**](docs/DefaultApi.md#updateorigin) | **Put** /content-providers/{contentProviderId}/origins/{originId} | origin
*DefaultApi* | [**UpdateSharedSecretSetSecrets**](docs/DefaultApi.md#updatesharedsecretsetsecrets) | **Put** /content-providers/{contentProviderId}/shared-secret-sets/{sharedSecretSetId}/secrets | secrets
*DefaultApi* | [**UpdateTlsDeliveryProfile**](docs/DefaultApi.md#updatetlsdeliveryprofile) | **Put** /content-providers/{contentProviderId}/tls-delivery-profiles/{tlsDeliveryProfileId} | tls delivery profile
*DefaultApi* | [**UpdateTlsIngestProfile**](docs/DefaultApi.md#updatetlsingestprofile) | **Put** /content-providers/{contentProviderId}/tls-ingest-profiles/{tlsIngestProfileId} | tls ingest profile


## Documentation For Models

 - [BehaviorType](docs/BehaviorType.md)
 - [CacheControlOverride](docs/CacheControlOverride.md)
 - [CacheControlOverrideOptions](docs/CacheControlOverrideOptions.md)
 - [CacheKeyUri](docs/CacheKeyUri.md)
 - [CacheKeyUriOptions](docs/CacheKeyUriOptions.md)
 - [CachingType](docs/CachingType.md)
 - [CdnPrefix](docs/CdnPrefix.md)
 - [CdnPrefix1](docs/CdnPrefix1.md)
 - [CdnPrefix2](docs/CdnPrefix2.md)
 - [CdnPrefix3](docs/CdnPrefix3.md)
 - [CdnPrefixCollection](docs/CdnPrefixCollection.md)
 - [CdnPrefixCreate](docs/CdnPrefixCreate.md)
 - [CdnPrefixReadDetailed](docs/CdnPrefixReadDetailed.md)
 - [CdnPrefixRules](docs/CdnPrefixRules.md)
 - [CdnPrefixRulesChildren](docs/CdnPrefixRulesChildren.md)
 - [CdnPrefixUpdate](docs/CdnPrefixUpdate.md)
 - [ChunkSize](docs/ChunkSize.md)
 - [ChunkSizeOptions](docs/ChunkSizeOptions.md)
 - [ContentProvider](docs/ContentProvider.md)
 - [ContentProvider1](docs/ContentProvider1.md)
 - [ContentProvider2](docs/ContentProvider2.md)
 - [ContentProviderCollection](docs/ContentProviderCollection.md)
 - [ContentProviderReadDetailed](docs/ContentProviderReadDetailed.md)
 - [Cors](docs/Cors.md)
 - [CorsOptions](docs/CorsOptions.md)
 - [CorsOptionsAccessControlAllowOrigin](docs/CorsOptionsAccessControlAllowOrigin.md)
 - [Criterion](docs/Criterion.md)
 - [CriterionOptions](docs/CriterionOptions.md)
 - [DefaultBehaviorType](docs/DefaultBehaviorType.md)
 - [DefaultFrontEndCache](docs/DefaultFrontEndCache.md)
 - [DefaultFrontEndCacheOptions](docs/DefaultFrontEndCacheOptions.md)
 - [DefaultHttp2Delivery](docs/DefaultHttp2Delivery.md)
 - [DefaultHttp2DeliveryOptions](docs/DefaultHttp2DeliveryOptions.md)
 - [DefaultHttpDelivery](docs/DefaultHttpDelivery.md)
 - [DefaultHttpDeliveryOptions](docs/DefaultHttpDeliveryOptions.md)
 - [DefaultHttpsDelivery](docs/DefaultHttpsDelivery.md)
 - [DefaultHttpsDeliveryOptions](docs/DefaultHttpsDeliveryOptions.md)
 - [DefaultOriginServer](docs/DefaultOriginServer.md)
 - [DefaultOriginServerOptions](docs/DefaultOriginServerOptions.md)
 - [DefaultRamOnlyCaching](docs/DefaultRamOnlyCaching.md)
 - [DefaultRamOnlyCachingOptions](docs/DefaultRamOnlyCachingOptions.md)
 - [DefaultStoragePartition](docs/DefaultStoragePartition.md)
 - [DefaultStoragePartitionOptions](docs/DefaultStoragePartitionOptions.md)
 - [DeliveryOverrideType](docs/DeliveryOverrideType.md)
 - [Dscp](docs/Dscp.md)
 - [DscpOptions](docs/DscpOptions.md)
 - [DynamicHierarchyType](docs/DynamicHierarchyType.md)
 - [EcdhCurveType](docs/EcdhCurveType.md)
 - [EdgeHostType](docs/EdgeHostType.md)
 - [ExcludedQueryArguments](docs/ExcludedQueryArguments.md)
 - [ExcludedQueryArgumentsOptions](docs/ExcludedQueryArgumentsOptions.md)
 - [ExternalOriginAuth](docs/ExternalOriginAuth.md)
 - [ExternalOriginAuthOptions](docs/ExternalOriginAuthOptions.md)
 - [FastRerouteType](docs/FastRerouteType.md)
 - [FastcgiAuth](docs/FastcgiAuth.md)
 - [FastcgiAuthOptions](docs/FastcgiAuthOptions.md)
 - [FollowRedirect](docs/FollowRedirect.md)
 - [FollowRedirectOptions](docs/FollowRedirectOptions.md)
 - [FrontEndCache](docs/FrontEndCache.md)
 - [Http2Delivery](docs/Http2Delivery.md)
 - [HttpDelivery](docs/HttpDelivery.md)
 - [HttpDeliveryOptions](docs/HttpDeliveryOptions.md)
 - [HttpsDelivery](docs/HttpsDelivery.md)
 - [LastMileAcceleration](docs/LastMileAcceleration.md)
 - [LastMileAccelerationOptions](docs/LastMileAccelerationOptions.md)
 - [LimitRate](docs/LimitRate.md)
 - [LimitRateOptions](docs/LimitRateOptions.md)
 - [Origin](docs/Origin.md)
 - [Origin1](docs/Origin1.md)
 - [Origin2](docs/Origin2.md)
 - [Origin3](docs/Origin3.md)
 - [OriginAffinityUri](docs/OriginAffinityUri.md)
 - [OriginAffinityUriOptions](docs/OriginAffinityUriOptions.md)
 - [OriginCollection](docs/OriginCollection.md)
 - [OriginCreate](docs/OriginCreate.md)
 - [OriginRangeRequest](docs/OriginRangeRequest.md)
 - [OriginRangeRequestOptions](docs/OriginRangeRequestOptions.md)
 - [OriginReadDetailed](docs/OriginReadDetailed.md)
 - [OriginServer](docs/OriginServer.md)
 - [OriginServerUri](docs/OriginServerUri.md)
 - [OriginServerUriOptions](docs/OriginServerUriOptions.md)
 - [OriginUpdate](docs/OriginUpdate.md)
 - [PageType](docs/PageType.md)
 - [PrefixPrioritizationType](docs/PrefixPrioritizationType.md)
 - [ProtocolType](docs/ProtocolType.md)
 - [PurgeKeyUri](docs/PurgeKeyUri.md)
 - [PurgeKeyUriOptions](docs/PurgeKeyUriOptions.md)
 - [RamOnlyCaching](docs/RamOnlyCaching.md)
 - [Rules](docs/Rules.md)
 - [SharedSecretSet](docs/SharedSecretSet.md)
 - [SharedSecretSet1](docs/SharedSecretSet1.md)
 - [SharedSecretSet2](docs/SharedSecretSet2.md)
 - [SharedSecretSet3](docs/SharedSecretSet3.md)
 - [SharedSecretSetCollection](docs/SharedSecretSetCollection.md)
 - [SharedSecretSetCreate](docs/SharedSecretSetCreate.md)
 - [SharedSecretSetReadDetailed](docs/SharedSecretSetReadDetailed.md)
 - [SharedSecretSetSecrets](docs/SharedSecretSetSecrets.md)
 - [SharedSecretSetUpdate](docs/SharedSecretSetUpdate.md)
 - [SiteRedirectMode](docs/SiteRedirectMode.md)
 - [SiteRedirectModeOptions](docs/SiteRedirectModeOptions.md)
 - [SiteRedirectType](docs/SiteRedirectType.md)
 - [SiteRedirectTypeOptions](docs/SiteRedirectTypeOptions.md)
 - [StoragePartition](docs/StoragePartition.md)
 - [TlsBufferSizeType](docs/TlsBufferSizeType.md)
 - [TlsDeliveryProfile](docs/TlsDeliveryProfile.md)
 - [TlsDeliveryProfile1](docs/TlsDeliveryProfile1.md)
 - [TlsDeliveryProfile2](docs/TlsDeliveryProfile2.md)
 - [TlsDeliveryProfile3](docs/TlsDeliveryProfile3.md)
 - [TlsDeliveryProfileCollection](docs/TlsDeliveryProfileCollection.md)
 - [TlsDeliveryProfileCreate](docs/TlsDeliveryProfileCreate.md)
 - [TlsDeliveryProfileReadDetailed](docs/TlsDeliveryProfileReadDetailed.md)
 - [TlsDeliveryProfileUpdate](docs/TlsDeliveryProfileUpdate.md)
 - [TlsIngestProfile](docs/TlsIngestProfile.md)
 - [TlsIngestProfile1](docs/TlsIngestProfile1.md)
 - [TlsIngestProfile2](docs/TlsIngestProfile2.md)
 - [TlsIngestProfile3](docs/TlsIngestProfile3.md)
 - [TlsIngestProfileCollection](docs/TlsIngestProfileCollection.md)
 - [TlsIngestProfileCreate](docs/TlsIngestProfileCreate.md)
 - [TlsIngestProfileReadDetailed](docs/TlsIngestProfileReadDetailed.md)
 - [TlsIngestProfileUpdate](docs/TlsIngestProfileUpdate.md)
 - [UdpSyslogAccessLog](docs/UdpSyslogAccessLog.md)
 - [UdpSyslogAccessLogOptions](docs/UdpSyslogAccessLogOptions.md)
 - [UdpSyslogAccessLogOptionsServers](docs/UdpSyslogAccessLogOptionsServers.md)
 - [UriSignature](docs/UriSignature.md)
 - [UriSignatureOptions](docs/UriSignatureOptions.md)
 - [X509CertificateSummaryType](docs/X509CertificateSummaryType.md)
 - [X509CertificateSummaryTypeFingerprints](docs/X509CertificateSummaryTypeFingerprints.md)
 - [X509CertificateSummaryTypeIssuedBy](docs/X509CertificateSummaryTypeIssuedBy.md)
 - [X509CertificateSummaryTypeIssuedTo](docs/X509CertificateSummaryTypeIssuedTo.md)
 - [X509CertificateSummaryTypeValidity](docs/X509CertificateSummaryTypeValidity.md)
 - [X509CrlSummaryType](docs/X509CrlSummaryType.md)
 - [X509CrlSummaryTypeIssuedBy](docs/X509CrlSummaryTypeIssuedBy.md)


## Documentation For Authorization



### oauth2


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: N/A

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



