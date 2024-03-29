# Go API client for fw

BloxOne Threat Defense Cloud is an extension of the BloxOne Suite that provides visibility into infected and compromised off-premises devices, roaming users, remote sites, and branch offices. You can subscribe to BloxOne Cloud and use its functionality to mitigate and control malware as well as provide unprecedented insight into your network security posture and enable timely action. BloxOne Cloud also offers unified policy management, reporting, and threat analytics across the entire spectrum. Using automated and high-quality threat intelligence feeds and unique behavioral analytics, it automatically stops device communications with C&Cs/botnets and prevents DNS based data exfiltration.

The mission-critical DNS infrastructure can become a vulnerable component in your network when it is inadequately protected by traditional security solutions and consequently used as an attack surface. Compromised DNS services can result in catastrophic network and system failures. To fully protect your network in today’s cyber security threat environment, Infoblox sets a new DNS security standard by offering scalable, enterprise-grade, and integrated protection for your DNS infrastructure.

Through the Infoblox Cloud Services Portal, you can view the status of your subscription and threat intelligence feeds, manage your network scope and roaming end users, and learn more about threats on your networks through the Infoblox Threat Lookup tool and predefined reports.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import fw "github.com/infobloxopen/bloxone-go-client"
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
ctx := context.WithValue(context.Background(), fw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), fw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), fw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), fw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost/api/atcfw/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccessCodesAPI* | [**AccessCodesCreateAccessCode**](docs/AccessCodesAPI.md#accesscodescreateaccesscode) | **Post** /access_codes | Create Access Codes
*AccessCodesAPI* | [**AccessCodesDeleteAccessCodes**](docs/AccessCodesAPI.md#accesscodesdeleteaccesscodes) | **Delete** /access_codes | Delete Access Codes
*AccessCodesAPI* | [**AccessCodesDeleteSingleAccessCodes**](docs/AccessCodesAPI.md#accesscodesdeletesingleaccesscodes) | **Delete** /access_codes/{access_key} | Delete Access Code By ID
*AccessCodesAPI* | [**AccessCodesListAccessCodes**](docs/AccessCodesAPI.md#accesscodeslistaccesscodes) | **Get** /access_codes | List Access Codes
*AccessCodesAPI* | [**AccessCodesReadAccessCode**](docs/AccessCodesAPI.md#accesscodesreadaccesscode) | **Get** /access_codes/{access_key} | Read Access Codes
*AccessCodesAPI* | [**AccessCodesUpdateAccessCode**](docs/AccessCodesAPI.md#accesscodesupdateaccesscode) | **Put** /access_codes/{payload.access_key} | Update Access Codes
*AppApprovalsAPI* | [**AppApprovalsListAppApprovals**](docs/AppApprovalsAPI.md#appapprovalslistappapprovals) | **Get** /app_approvals | 
*AppApprovalsAPI* | [**AppApprovalsReplaceAppApprovals**](docs/AppApprovalsAPI.md#appapprovalsreplaceappapprovals) | **Put** /app_approvals | Update Application Approval.
*AppApprovalsAPI* | [**AppApprovalsUpdateAppApprovals**](docs/AppApprovalsAPI.md#appapprovalsupdateappapprovals) | **Patch** /app_approvals | 
*ApplicationFiltersAPI* | [**ApplicationFiltersCreateApplicationFilter**](docs/ApplicationFiltersAPI.md#applicationfilterscreateapplicationfilter) | **Post** /application_filters | Create Application Filter.
*ApplicationFiltersAPI* | [**ApplicationFiltersDeleteApplicationFilters**](docs/ApplicationFiltersAPI.md#applicationfiltersdeleteapplicationfilters) | **Delete** /application_filters | Delete Application Filters.
*ApplicationFiltersAPI* | [**ApplicationFiltersDeleteSingleApplicationFilters**](docs/ApplicationFiltersAPI.md#applicationfiltersdeletesingleapplicationfilters) | **Delete** /application_filters/{id} | Delete Application Filter Object by ID.
*ApplicationFiltersAPI* | [**ApplicationFiltersListApplicationFilters**](docs/ApplicationFiltersAPI.md#applicationfilterslistapplicationfilters) | **Get** /application_filters | List Application Filters.
*ApplicationFiltersAPI* | [**ApplicationFiltersReadApplicationFilter**](docs/ApplicationFiltersAPI.md#applicationfiltersreadapplicationfilter) | **Get** /application_filters/{id} | Read Application Filter.
*ApplicationFiltersAPI* | [**ApplicationFiltersUpdateApplicationFilter**](docs/ApplicationFiltersAPI.md#applicationfiltersupdateapplicationfilter) | **Put** /application_filters/{id} | Update Application Filter.
*CategoryFiltersAPI* | [**CategoryFiltersCreateCategoryFilter**](docs/CategoryFiltersAPI.md#categoryfilterscreatecategoryfilter) | **Post** /category_filters | Create Category Filter.
*CategoryFiltersAPI* | [**CategoryFiltersDeleteCategoryFilters**](docs/CategoryFiltersAPI.md#categoryfiltersdeletecategoryfilters) | **Delete** /category_filters | Delete Category Filters By ID.
*CategoryFiltersAPI* | [**CategoryFiltersDeleteSingleCategoryFilters**](docs/CategoryFiltersAPI.md#categoryfiltersdeletesinglecategoryfilters) | **Delete** /category_filters/{id} | Delete Category Filters.
*CategoryFiltersAPI* | [**CategoryFiltersListCategoryFilters**](docs/CategoryFiltersAPI.md#categoryfilterslistcategoryfilters) | **Get** /category_filters | List Category Filters.
*CategoryFiltersAPI* | [**CategoryFiltersReadCategoryFilter**](docs/CategoryFiltersAPI.md#categoryfiltersreadcategoryfilter) | **Get** /category_filters/{id} | Read Category Filter.
*CategoryFiltersAPI* | [**CategoryFiltersUpdateCategoryFilter**](docs/CategoryFiltersAPI.md#categoryfiltersupdatecategoryfilter) | **Put** /category_filters/{id} | Update Category Filter.
*ContentCategoriesAPI* | [**ContentCategoriesListContentCategories**](docs/ContentCategoriesAPI.md#contentcategorieslistcontentcategories) | **Get** /content_categories | List Content Categories.
*InternalDomainListsAPI* | [**InternalDomainListsCreateInternalDomains**](docs/InternalDomainListsAPI.md#internaldomainlistscreateinternaldomains) | **Post** /internal_domain_lists | Create Internal Domains.
*InternalDomainListsAPI* | [**InternalDomainListsDeleteInternalDomains**](docs/InternalDomainListsAPI.md#internaldomainlistsdeleteinternaldomains) | **Delete** /internal_domain_lists | Delete Internal Domains.
*InternalDomainListsAPI* | [**InternalDomainListsDeleteSingleInternalDomains**](docs/InternalDomainListsAPI.md#internaldomainlistsdeletesingleinternaldomains) | **Delete** /internal_domain_lists/{id} | Delete Internal Domains.
*InternalDomainListsAPI* | [**InternalDomainListsInternalDomainsItemsPartialUpdate**](docs/InternalDomainListsAPI.md#internaldomainlistsinternaldomainsitemspartialupdate) | **Patch** /internal_domain_lists/{id}/items | Patch Internal Domains.
*InternalDomainListsAPI* | [**InternalDomainListsListInternalDomains**](docs/InternalDomainListsAPI.md#internaldomainlistslistinternaldomains) | **Get** /internal_domain_lists | List Internal Domains.
*InternalDomainListsAPI* | [**InternalDomainListsReadInternalDomains**](docs/InternalDomainListsAPI.md#internaldomainlistsreadinternaldomains) | **Get** /internal_domain_lists/{id} | Read Internal Domains.
*InternalDomainListsAPI* | [**InternalDomainListsUpdateInternalDomains**](docs/InternalDomainListsAPI.md#internaldomainlistsupdateinternaldomains) | **Put** /internal_domain_lists/{id} | Update Internal Domains.
*NamedListItemsAPI* | [**NamedListItemsDeleteNamedListItems**](docs/NamedListItemsAPI.md#namedlistitemsdeletenamedlistitems) | **Delete** /named_lists/{id}/items | Delete Named List Items.
*NamedListItemsAPI* | [**NamedListItemsInsertOrReplaceNamedListItems**](docs/NamedListItemsAPI.md#namedlistitemsinsertorreplacenamedlistitems) | **Post** /named_lists/{id}/items | Insert Named List Items.
*NamedListItemsAPI* | [**NamedListItemsNamedListItemsPartialUpdate**](docs/NamedListItemsAPI.md#namedlistitemsnamedlistitemspartialupdate) | **Patch** /named_lists/{id}/items | Partial Update Named List Items.
*NamedListsAPI* | [**NamedListsCreateNamedList**](docs/NamedListsAPI.md#namedlistscreatenamedlist) | **Post** /named_lists | Create Named List.
*NamedListsAPI* | [**NamedListsDeleteNamedLists**](docs/NamedListsAPI.md#namedlistsdeletenamedlists) | **Delete** /named_lists | Delete Named Lists.
*NamedListsAPI* | [**NamedListsDeleteSingleNamedLists**](docs/NamedListsAPI.md#namedlistsdeletesinglenamedlists) | **Delete** /named_lists/{id} | Delete Named Lists.
*NamedListsAPI* | [**NamedListsListNamedLists**](docs/NamedListsAPI.md#namedlistslistnamedlists) | **Get** /named_lists | List Named Lists.
*NamedListsAPI* | [**NamedListsListNamedListsCSV**](docs/NamedListsAPI.md#namedlistslistnamedlistscsv) | **Get** /named_lists_download | List Named Lists in CSV format.
*NamedListsAPI* | [**NamedListsMultiListUpdate**](docs/NamedListsAPI.md#namedlistsmultilistupdate) | **Patch** /named_lists | Patch Multiple Named Lists.
*NamedListsAPI* | [**NamedListsReadNamedList**](docs/NamedListsAPI.md#namedlistsreadnamedlist) | **Get** /named_lists/{id} | Read Named List.
*NamedListsAPI* | [**NamedListsUpdateNamedList**](docs/NamedListsAPI.md#namedlistsupdatenamedlist) | **Put** /named_lists/{id} | Update Named List.
*NamedListsAPI* | [**NamedListsUpdateNamedListPartial**](docs/NamedListsAPI.md#namedlistsupdatenamedlistpartial) | **Patch** /named_lists/{id} | Patch TI List.
*NetworkListsAPI* | [**NetworkListsCreateNetworkList**](docs/NetworkListsAPI.md#networklistscreatenetworklist) | **Post** /network_lists | Create Network List.
*NetworkListsAPI* | [**NetworkListsDeleteNetworkLists**](docs/NetworkListsAPI.md#networklistsdeletenetworklists) | **Delete** /network_lists | Delete Network Lists.
*NetworkListsAPI* | [**NetworkListsDeleteSingleNetworkLists**](docs/NetworkListsAPI.md#networklistsdeletesinglenetworklists) | **Delete** /network_lists/{id} | Delete Network Lists.
*NetworkListsAPI* | [**NetworkListsListNetworkLists**](docs/NetworkListsAPI.md#networklistslistnetworklists) | **Get** /network_lists | List Network Lists.
*NetworkListsAPI* | [**NetworkListsReadNetworkList**](docs/NetworkListsAPI.md#networklistsreadnetworklist) | **Get** /network_lists/{id} | Read Network List.
*NetworkListsAPI* | [**NetworkListsUpdateNetworkList**](docs/NetworkListsAPI.md#networklistsupdatenetworklist) | **Put** /network_lists/{id} | Update Network List.
*PopRegionsAPI* | [**PopRegionsListPoPRegions**](docs/PopRegionsAPI.md#popregionslistpopregions) | **Get** /pop_regions | List PoP Regions.
*PopRegionsAPI* | [**PopRegionsReadPoPRegion**](docs/PopRegionsAPI.md#popregionsreadpopregion) | **Get** /pop_regions/{id} | Read PoP Region.
*SecurityPoliciesAPI* | [**SecurityPoliciesCreateSecurityPolicy**](docs/SecurityPoliciesAPI.md#securitypoliciescreatesecuritypolicy) | **Post** /security_policies | Create Security Policy.
*SecurityPoliciesAPI* | [**SecurityPoliciesDeleteSecurityPolicy**](docs/SecurityPoliciesAPI.md#securitypoliciesdeletesecuritypolicy) | **Delete** /security_policies | Delete Security Policies.
*SecurityPoliciesAPI* | [**SecurityPoliciesDeleteSingleSecurityPolicy**](docs/SecurityPoliciesAPI.md#securitypoliciesdeletesinglesecuritypolicy) | **Delete** /security_policies/{id} | Delete Security Policy.
*SecurityPoliciesAPI* | [**SecurityPoliciesListSecurityPolicies**](docs/SecurityPoliciesAPI.md#securitypolicieslistsecuritypolicies) | **Get** /security_policies | List Security Policies.
*SecurityPoliciesAPI* | [**SecurityPoliciesReadSecurityPolicy**](docs/SecurityPoliciesAPI.md#securitypoliciesreadsecuritypolicy) | **Get** /security_policies/{id} | Read Security Policy.
*SecurityPoliciesAPI* | [**SecurityPoliciesUpdateSecurityPolicy**](docs/SecurityPoliciesAPI.md#securitypoliciesupdatesecuritypolicy) | **Put** /security_policies/{id} | Update Security Policy.
*SecurityPolicyRulesAPI* | [**SecurityPolicyRulesListSecurityPolicyRules**](docs/SecurityPolicyRulesAPI.md#securitypolicyruleslistsecuritypolicyrules) | **Get** /security_policy_rules | List Security Policy Rules.
*ThreatFeedsAPI* | [**ThreatFeedsListThreatFeeds**](docs/ThreatFeedsAPI.md#threatfeedslistthreatfeeds) | **Get** /threat_feeds | List Threat Feeds.


## Documentation For Models

 - [AccessCodesCreateAccessCode400Response](docs/AccessCodesCreateAccessCode400Response.md)
 - [AccessCodesCreateAccessCode400ResponseError](docs/AccessCodesCreateAccessCode400ResponseError.md)
 - [AccessCodesCreateAccessCode404Response](docs/AccessCodesCreateAccessCode404Response.md)
 - [AccessCodesCreateAccessCode404ResponseError](docs/AccessCodesCreateAccessCode404ResponseError.md)
 - [AccessCodesCreateAccessCode409Response](docs/AccessCodesCreateAccessCode409Response.md)
 - [AccessCodesCreateAccessCode409ResponseError](docs/AccessCodesCreateAccessCode409ResponseError.md)
 - [AccessCodesDeleteAccessCodes400Response](docs/AccessCodesDeleteAccessCodes400Response.md)
 - [AccessCodesDeleteAccessCodes400ResponseError](docs/AccessCodesDeleteAccessCodes400ResponseError.md)
 - [AccessCodesDeleteSingleAccessCodes400Response](docs/AccessCodesDeleteSingleAccessCodes400Response.md)
 - [AccessCodesDeleteSingleAccessCodes400ResponseError](docs/AccessCodesDeleteSingleAccessCodes400ResponseError.md)
 - [AccessCodesListAccessCodes500Response](docs/AccessCodesListAccessCodes500Response.md)
 - [AccessCodesListAccessCodes500ResponseError](docs/AccessCodesListAccessCodes500ResponseError.md)
 - [AccessCodesReadAccessCode404Response](docs/AccessCodesReadAccessCode404Response.md)
 - [AccessCodesReadAccessCode404ResponseError](docs/AccessCodesReadAccessCode404ResponseError.md)
 - [ApplicationFiltersDeleteApplicationFilters400Response](docs/ApplicationFiltersDeleteApplicationFilters400Response.md)
 - [ApplicationFiltersDeleteApplicationFilters400ResponseError](docs/ApplicationFiltersDeleteApplicationFilters400ResponseError.md)
 - [ApplicationFiltersDeleteSingleApplicationFilters400Response](docs/ApplicationFiltersDeleteSingleApplicationFilters400Response.md)
 - [ApplicationFiltersDeleteSingleApplicationFilters400ResponseError](docs/ApplicationFiltersDeleteSingleApplicationFilters400ResponseError.md)
 - [AtcfwAccessCode](docs/AtcfwAccessCode.md)
 - [AtcfwAccessCodeCreateResponse](docs/AtcfwAccessCodeCreateResponse.md)
 - [AtcfwAccessCodeDeleteRequest](docs/AtcfwAccessCodeDeleteRequest.md)
 - [AtcfwAccessCodeMultiResponse](docs/AtcfwAccessCodeMultiResponse.md)
 - [AtcfwAccessCodeReadResponse](docs/AtcfwAccessCodeReadResponse.md)
 - [AtcfwAccessCodeRule](docs/AtcfwAccessCodeRule.md)
 - [AtcfwAccessCodeUpdateResponse](docs/AtcfwAccessCodeUpdateResponse.md)
 - [AtcfwAppApproval](docs/AtcfwAppApproval.md)
 - [AtcfwAppApprovalMultiResponse](docs/AtcfwAppApprovalMultiResponse.md)
 - [AtcfwAppApprovalRemovalRequest](docs/AtcfwAppApprovalRemovalRequest.md)
 - [AtcfwAppApprovalsReplaceRequest](docs/AtcfwAppApprovalsReplaceRequest.md)
 - [AtcfwAppApprovalsUpdateRequest](docs/AtcfwAppApprovalsUpdateRequest.md)
 - [AtcfwApplicationCriterion](docs/AtcfwApplicationCriterion.md)
 - [AtcfwApplicationFilter](docs/AtcfwApplicationFilter.md)
 - [AtcfwApplicationFilterCreateResponse](docs/AtcfwApplicationFilterCreateResponse.md)
 - [AtcfwApplicationFilterMultiResponse](docs/AtcfwApplicationFilterMultiResponse.md)
 - [AtcfwApplicationFilterReadResponse](docs/AtcfwApplicationFilterReadResponse.md)
 - [AtcfwApplicationFilterUpdateResponse](docs/AtcfwApplicationFilterUpdateResponse.md)
 - [AtcfwApplicationFiltersDeleteRequest](docs/AtcfwApplicationFiltersDeleteRequest.md)
 - [AtcfwCategoryFilter](docs/AtcfwCategoryFilter.md)
 - [AtcfwCategoryFilterCreateResponse](docs/AtcfwCategoryFilterCreateResponse.md)
 - [AtcfwCategoryFilterMultiResponse](docs/AtcfwCategoryFilterMultiResponse.md)
 - [AtcfwCategoryFilterReadResponse](docs/AtcfwCategoryFilterReadResponse.md)
 - [AtcfwCategoryFilterUpdateResponse](docs/AtcfwCategoryFilterUpdateResponse.md)
 - [AtcfwCategoryFiltersDeleteRequest](docs/AtcfwCategoryFiltersDeleteRequest.md)
 - [AtcfwContentCategory](docs/AtcfwContentCategory.md)
 - [AtcfwContentCategoryMultiResponse](docs/AtcfwContentCategoryMultiResponse.md)
 - [AtcfwInternalDomains](docs/AtcfwInternalDomains.md)
 - [AtcfwInternalDomainsCreateResponse](docs/AtcfwInternalDomainsCreateResponse.md)
 - [AtcfwInternalDomainsDeleteRequest](docs/AtcfwInternalDomainsDeleteRequest.md)
 - [AtcfwInternalDomainsItems](docs/AtcfwInternalDomainsItems.md)
 - [AtcfwInternalDomainsMultiResponse](docs/AtcfwInternalDomainsMultiResponse.md)
 - [AtcfwInternalDomainsReadResponse](docs/AtcfwInternalDomainsReadResponse.md)
 - [AtcfwInternalDomainsUpdateResponse](docs/AtcfwInternalDomainsUpdateResponse.md)
 - [AtcfwItemStructs](docs/AtcfwItemStructs.md)
 - [AtcfwListPoPRegionsResponse](docs/AtcfwListPoPRegionsResponse.md)
 - [AtcfwListSeverityLevels](docs/AtcfwListSeverityLevels.md)
 - [AtcfwMultiListUpdate](docs/AtcfwMultiListUpdate.md)
 - [AtcfwNamedList](docs/AtcfwNamedList.md)
 - [AtcfwNamedListCSVListResponse](docs/AtcfwNamedListCSVListResponse.md)
 - [AtcfwNamedListCreateResponse](docs/AtcfwNamedListCreateResponse.md)
 - [AtcfwNamedListItemsDeleteRequest](docs/AtcfwNamedListItemsDeleteRequest.md)
 - [AtcfwNamedListItemsInsertOrUpdate](docs/AtcfwNamedListItemsInsertOrUpdate.md)
 - [AtcfwNamedListItemsInsertOrUpdateResponse](docs/AtcfwNamedListItemsInsertOrUpdateResponse.md)
 - [AtcfwNamedListItemsInsertOrUpdateResponseSuccess](docs/AtcfwNamedListItemsInsertOrUpdateResponseSuccess.md)
 - [AtcfwNamedListItemsPartialUpdate](docs/AtcfwNamedListItemsPartialUpdate.md)
 - [AtcfwNamedListRead](docs/AtcfwNamedListRead.md)
 - [AtcfwNamedListReadMultiResponse](docs/AtcfwNamedListReadMultiResponse.md)
 - [AtcfwNamedListReadResponse](docs/AtcfwNamedListReadResponse.md)
 - [AtcfwNamedListUpdateResponse](docs/AtcfwNamedListUpdateResponse.md)
 - [AtcfwNamedListsDeleteRequest](docs/AtcfwNamedListsDeleteRequest.md)
 - [AtcfwNetAddrDfpAssignment](docs/AtcfwNetAddrDfpAssignment.md)
 - [AtcfwNetworkList](docs/AtcfwNetworkList.md)
 - [AtcfwNetworkListCreateResponse](docs/AtcfwNetworkListCreateResponse.md)
 - [AtcfwNetworkListMultiResponse](docs/AtcfwNetworkListMultiResponse.md)
 - [AtcfwNetworkListReadResponse](docs/AtcfwNetworkListReadResponse.md)
 - [AtcfwNetworkListUpdateResponse](docs/AtcfwNetworkListUpdateResponse.md)
 - [AtcfwNetworkListsDeleteRequest](docs/AtcfwNetworkListsDeleteRequest.md)
 - [AtcfwPoPRegion](docs/AtcfwPoPRegion.md)
 - [AtcfwReadPoPRegionResponse](docs/AtcfwReadPoPRegionResponse.md)
 - [AtcfwSecurityPolicy](docs/AtcfwSecurityPolicy.md)
 - [AtcfwSecurityPolicyCreateResponse](docs/AtcfwSecurityPolicyCreateResponse.md)
 - [AtcfwSecurityPolicyDeleteRequest](docs/AtcfwSecurityPolicyDeleteRequest.md)
 - [AtcfwSecurityPolicyMultiResponse](docs/AtcfwSecurityPolicyMultiResponse.md)
 - [AtcfwSecurityPolicyReadResponse](docs/AtcfwSecurityPolicyReadResponse.md)
 - [AtcfwSecurityPolicyRule](docs/AtcfwSecurityPolicyRule.md)
 - [AtcfwSecurityPolicyRuleMultiResponse](docs/AtcfwSecurityPolicyRuleMultiResponse.md)
 - [AtcfwSecurityPolicyUpdateResponse](docs/AtcfwSecurityPolicyUpdateResponse.md)
 - [AtcfwThreatFeed](docs/AtcfwThreatFeed.md)
 - [AtcfwThreatFeedMultiResponse](docs/AtcfwThreatFeedMultiResponse.md)
 - [CategoryFiltersCreateCategoryFilter400Response](docs/CategoryFiltersCreateCategoryFilter400Response.md)
 - [CategoryFiltersCreateCategoryFilter400ResponseError](docs/CategoryFiltersCreateCategoryFilter400ResponseError.md)
 - [CategoryFiltersCreateCategoryFilter409Response](docs/CategoryFiltersCreateCategoryFilter409Response.md)
 - [CategoryFiltersCreateCategoryFilter409ResponseError](docs/CategoryFiltersCreateCategoryFilter409ResponseError.md)
 - [CategoryFiltersDeleteCategoryFilters400Response](docs/CategoryFiltersDeleteCategoryFilters400Response.md)
 - [CategoryFiltersDeleteCategoryFilters400ResponseError](docs/CategoryFiltersDeleteCategoryFilters400ResponseError.md)
 - [CategoryFiltersReadCategoryFilter404Response](docs/CategoryFiltersReadCategoryFilter404Response.md)
 - [CategoryFiltersReadCategoryFilter404ResponseError](docs/CategoryFiltersReadCategoryFilter404ResponseError.md)
 - [InternalDomainListsCreateInternalDomains400Response](docs/InternalDomainListsCreateInternalDomains400Response.md)
 - [InternalDomainListsCreateInternalDomains400ResponseError](docs/InternalDomainListsCreateInternalDomains400ResponseError.md)
 - [InternalDomainListsCreateInternalDomains409Response](docs/InternalDomainListsCreateInternalDomains409Response.md)
 - [InternalDomainListsCreateInternalDomains409ResponseError](docs/InternalDomainListsCreateInternalDomains409ResponseError.md)
 - [InternalDomainListsDeleteInternalDomains400Response](docs/InternalDomainListsDeleteInternalDomains400Response.md)
 - [InternalDomainListsDeleteInternalDomains400ResponseError](docs/InternalDomainListsDeleteInternalDomains400ResponseError.md)
 - [InternalDomainListsDeleteInternalDomains404Response](docs/InternalDomainListsDeleteInternalDomains404Response.md)
 - [InternalDomainListsDeleteInternalDomains404ResponseError](docs/InternalDomainListsDeleteInternalDomains404ResponseError.md)
 - [InternalDomainListsDeleteSingleInternalDomains400Response](docs/InternalDomainListsDeleteSingleInternalDomains400Response.md)
 - [InternalDomainListsDeleteSingleInternalDomains400ResponseError](docs/InternalDomainListsDeleteSingleInternalDomains400ResponseError.md)
 - [InternalDomainListsInternalDomainsItemsPartialUpdate400Response](docs/InternalDomainListsInternalDomainsItemsPartialUpdate400Response.md)
 - [InternalDomainListsInternalDomainsItemsPartialUpdate400ResponseError](docs/InternalDomainListsInternalDomainsItemsPartialUpdate400ResponseError.md)
 - [InternalDomainListsInternalDomainsItemsPartialUpdate404Response](docs/InternalDomainListsInternalDomainsItemsPartialUpdate404Response.md)
 - [InternalDomainListsInternalDomainsItemsPartialUpdate404ResponseError](docs/InternalDomainListsInternalDomainsItemsPartialUpdate404ResponseError.md)
 - [InternalDomainListsReadInternalDomains404Response](docs/InternalDomainListsReadInternalDomains404Response.md)
 - [InternalDomainListsReadInternalDomains404ResponseError](docs/InternalDomainListsReadInternalDomains404ResponseError.md)
 - [InternalDomainListsUpdateInternalDomains400Response](docs/InternalDomainListsUpdateInternalDomains400Response.md)
 - [InternalDomainListsUpdateInternalDomains400ResponseError](docs/InternalDomainListsUpdateInternalDomains400ResponseError.md)
 - [InternalDomainListsUpdateInternalDomains404Response](docs/InternalDomainListsUpdateInternalDomains404Response.md)
 - [InternalDomainListsUpdateInternalDomains404ResponseError](docs/InternalDomainListsUpdateInternalDomains404ResponseError.md)
 - [NamedListItemsDeleteNamedListItems400Response](docs/NamedListItemsDeleteNamedListItems400Response.md)
 - [NamedListItemsDeleteNamedListItems400ResponseError](docs/NamedListItemsDeleteNamedListItems400ResponseError.md)
 - [NamedListItemsInsertOrReplaceNamedListItems400Response](docs/NamedListItemsInsertOrReplaceNamedListItems400Response.md)
 - [NamedListItemsInsertOrReplaceNamedListItems400ResponseError](docs/NamedListItemsInsertOrReplaceNamedListItems400ResponseError.md)
 - [NamedListItemsNamedListItemsPartialUpdate400Response](docs/NamedListItemsNamedListItemsPartialUpdate400Response.md)
 - [NamedListItemsNamedListItemsPartialUpdate400ResponseError](docs/NamedListItemsNamedListItemsPartialUpdate400ResponseError.md)
 - [NamedListsCreateNamedList409Response](docs/NamedListsCreateNamedList409Response.md)
 - [NamedListsCreateNamedList409ResponseError](docs/NamedListsCreateNamedList409ResponseError.md)
 - [NamedListsDeleteNamedLists400Response](docs/NamedListsDeleteNamedLists400Response.md)
 - [NamedListsDeleteNamedLists400ResponseError](docs/NamedListsDeleteNamedLists400ResponseError.md)
 - [NamedListsDeleteNamedLists404Response](docs/NamedListsDeleteNamedLists404Response.md)
 - [NamedListsDeleteNamedLists404ResponseError](docs/NamedListsDeleteNamedLists404ResponseError.md)
 - [NamedListsDeleteSingleNamedLists404Response](docs/NamedListsDeleteSingleNamedLists404Response.md)
 - [NamedListsDeleteSingleNamedLists404ResponseError](docs/NamedListsDeleteSingleNamedLists404ResponseError.md)
 - [NamedListsMultiListUpdate404Response](docs/NamedListsMultiListUpdate404Response.md)
 - [NamedListsMultiListUpdate404ResponseError](docs/NamedListsMultiListUpdate404ResponseError.md)
 - [NamedListsUpdateNamedListPartial400Response](docs/NamedListsUpdateNamedListPartial400Response.md)
 - [NamedListsUpdateNamedListPartial400ResponseError](docs/NamedListsUpdateNamedListPartial400ResponseError.md)
 - [NamedListsUpdateNamedListPartial404Response](docs/NamedListsUpdateNamedListPartial404Response.md)
 - [NamedListsUpdateNamedListPartial404ResponseError](docs/NamedListsUpdateNamedListPartial404ResponseError.md)
 - [NamedListsUpdateNamedListPartial405Response](docs/NamedListsUpdateNamedListPartial405Response.md)
 - [NamedListsUpdateNamedListPartial405ResponseError](docs/NamedListsUpdateNamedListPartial405ResponseError.md)
 - [NetAddrDfpAssignmentScopeType](docs/NetAddrDfpAssignmentScopeType.md)
 - [NetworkListsCreateNetworkList409Response](docs/NetworkListsCreateNetworkList409Response.md)
 - [NetworkListsCreateNetworkList409ResponseError](docs/NetworkListsCreateNetworkList409ResponseError.md)
 - [NetworkListsDeleteNetworkLists400Response](docs/NetworkListsDeleteNetworkLists400Response.md)
 - [NetworkListsDeleteNetworkLists400ResponseError](docs/NetworkListsDeleteNetworkLists400ResponseError.md)
 - [NetworkListsDeleteNetworkLists404Response](docs/NetworkListsDeleteNetworkLists404Response.md)
 - [NetworkListsDeleteNetworkLists404ResponseError](docs/NetworkListsDeleteNetworkLists404ResponseError.md)
 - [PopRegionsReadPoPRegion404Response](docs/PopRegionsReadPoPRegion404Response.md)
 - [PopRegionsReadPoPRegion404ResponseError](docs/PopRegionsReadPoPRegion404ResponseError.md)
 - [ProtobufFieldMask](docs/ProtobufFieldMask.md)
 - [SecurityPoliciesCreateSecurityPolicy400Response](docs/SecurityPoliciesCreateSecurityPolicy400Response.md)
 - [SecurityPoliciesCreateSecurityPolicy400ResponseError](docs/SecurityPoliciesCreateSecurityPolicy400ResponseError.md)
 - [SecurityPoliciesCreateSecurityPolicy409Response](docs/SecurityPoliciesCreateSecurityPolicy409Response.md)
 - [SecurityPoliciesCreateSecurityPolicy409ResponseError](docs/SecurityPoliciesCreateSecurityPolicy409ResponseError.md)
 - [SecurityPoliciesReadSecurityPolicy404Response](docs/SecurityPoliciesReadSecurityPolicy404Response.md)
 - [SecurityPoliciesReadSecurityPolicy404ResponseError](docs/SecurityPoliciesReadSecurityPolicy404ResponseError.md)
 - [ThreatFeedSource](docs/ThreatFeedSource.md)


## Documentation For Authorization

Endpoints do not require authorization.


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



