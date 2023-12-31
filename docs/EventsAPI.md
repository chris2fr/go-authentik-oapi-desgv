# \EventsAPI

All URIs are relative to *http://localhost/api/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventsEventsActionsList**](EventsAPI.md#EventsEventsActionsList) | **Get** /events/events/actions/ | 
[**EventsEventsCreate**](EventsAPI.md#EventsEventsCreate) | **Post** /events/events/ | 
[**EventsEventsDestroy**](EventsAPI.md#EventsEventsDestroy) | **Delete** /events/events/{event_uuid}/ | 
[**EventsEventsList**](EventsAPI.md#EventsEventsList) | **Get** /events/events/ | 
[**EventsEventsPartialUpdate**](EventsAPI.md#EventsEventsPartialUpdate) | **Patch** /events/events/{event_uuid}/ | 
[**EventsEventsPerMonthList**](EventsAPI.md#EventsEventsPerMonthList) | **Get** /events/events/per_month/ | 
[**EventsEventsRetrieve**](EventsAPI.md#EventsEventsRetrieve) | **Get** /events/events/{event_uuid}/ | 
[**EventsEventsTopPerUserList**](EventsAPI.md#EventsEventsTopPerUserList) | **Get** /events/events/top_per_user/ | 
[**EventsEventsUpdate**](EventsAPI.md#EventsEventsUpdate) | **Put** /events/events/{event_uuid}/ | 
[**EventsNotificationsDestroy**](EventsAPI.md#EventsNotificationsDestroy) | **Delete** /events/notifications/{uuid}/ | 
[**EventsNotificationsList**](EventsAPI.md#EventsNotificationsList) | **Get** /events/notifications/ | 
[**EventsNotificationsMarkAllSeenCreate**](EventsAPI.md#EventsNotificationsMarkAllSeenCreate) | **Post** /events/notifications/mark_all_seen/ | 
[**EventsNotificationsPartialUpdate**](EventsAPI.md#EventsNotificationsPartialUpdate) | **Patch** /events/notifications/{uuid}/ | 
[**EventsNotificationsRetrieve**](EventsAPI.md#EventsNotificationsRetrieve) | **Get** /events/notifications/{uuid}/ | 
[**EventsNotificationsUpdate**](EventsAPI.md#EventsNotificationsUpdate) | **Put** /events/notifications/{uuid}/ | 
[**EventsNotificationsUsedByList**](EventsAPI.md#EventsNotificationsUsedByList) | **Get** /events/notifications/{uuid}/used_by/ | 
[**EventsRulesCreate**](EventsAPI.md#EventsRulesCreate) | **Post** /events/rules/ | 
[**EventsRulesDestroy**](EventsAPI.md#EventsRulesDestroy) | **Delete** /events/rules/{pbm_uuid}/ | 
[**EventsRulesList**](EventsAPI.md#EventsRulesList) | **Get** /events/rules/ | 
[**EventsRulesPartialUpdate**](EventsAPI.md#EventsRulesPartialUpdate) | **Patch** /events/rules/{pbm_uuid}/ | 
[**EventsRulesRetrieve**](EventsAPI.md#EventsRulesRetrieve) | **Get** /events/rules/{pbm_uuid}/ | 
[**EventsRulesUpdate**](EventsAPI.md#EventsRulesUpdate) | **Put** /events/rules/{pbm_uuid}/ | 
[**EventsRulesUsedByList**](EventsAPI.md#EventsRulesUsedByList) | **Get** /events/rules/{pbm_uuid}/used_by/ | 
[**EventsTransportsCreate**](EventsAPI.md#EventsTransportsCreate) | **Post** /events/transports/ | 
[**EventsTransportsDestroy**](EventsAPI.md#EventsTransportsDestroy) | **Delete** /events/transports/{uuid}/ | 
[**EventsTransportsList**](EventsAPI.md#EventsTransportsList) | **Get** /events/transports/ | 
[**EventsTransportsPartialUpdate**](EventsAPI.md#EventsTransportsPartialUpdate) | **Patch** /events/transports/{uuid}/ | 
[**EventsTransportsRetrieve**](EventsAPI.md#EventsTransportsRetrieve) | **Get** /events/transports/{uuid}/ | 
[**EventsTransportsTestCreate**](EventsAPI.md#EventsTransportsTestCreate) | **Post** /events/transports/{uuid}/test/ | 
[**EventsTransportsUpdate**](EventsAPI.md#EventsTransportsUpdate) | **Put** /events/transports/{uuid}/ | 
[**EventsTransportsUsedByList**](EventsAPI.md#EventsTransportsUsedByList) | **Get** /events/transports/{uuid}/used_by/ | 



## EventsEventsActionsList

> []TypeCreate EventsEventsActionsList(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.EventsEventsActionsList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsEventsActionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsEventsActionsList`: []TypeCreate
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsEventsActionsList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEventsEventsActionsListRequest struct via the builder pattern


### Return type

[**[]TypeCreate**](TypeCreate.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsEventsCreate

> Event EventsEventsCreate(ctx).EventRequest(eventRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    eventRequest := *openapiclient.NewEventRequest(openapiclient.EventActions("login"), "App_example") // EventRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.EventsEventsCreate(context.Background()).EventRequest(eventRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsEventsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsEventsCreate`: Event
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsEventsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsEventsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventRequest** | [**EventRequest**](EventRequest.md) |  | 

### Return type

[**Event**](Event.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsEventsDestroy

> EventsEventsDestroy(ctx, eventUuid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    eventUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Event.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EventsAPI.EventsEventsDestroy(context.Background(), eventUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsEventsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventUuid** | **string** | A UUID string identifying this Event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsEventsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsEventsList

> PaginatedEventList EventsEventsList(ctx).Action(action).ClientIp(clientIp).ContextAuthorizedApp(contextAuthorizedApp).ContextModelApp(contextModelApp).ContextModelName(contextModelName).ContextModelPk(contextModelPk).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).TenantName(tenantName).Username(username).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    action := "action_example" // string |  (optional)
    clientIp := "clientIp_example" // string |  (optional)
    contextAuthorizedApp := "contextAuthorizedApp_example" // string | Context Authorized application (optional)
    contextModelApp := "contextModelApp_example" // string | Context Model App (optional)
    contextModelName := "contextModelName_example" // string | Context Model Name (optional)
    contextModelPk := "contextModelPk_example" // string | Context Model Primary Key (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)
    tenantName := "tenantName_example" // string | Tenant name (optional)
    username := "username_example" // string | Username (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.EventsEventsList(context.Background()).Action(action).ClientIp(clientIp).ContextAuthorizedApp(contextAuthorizedApp).ContextModelApp(contextModelApp).ContextModelName(contextModelName).ContextModelPk(contextModelPk).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).TenantName(tenantName).Username(username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsEventsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsEventsList`: PaginatedEventList
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsEventsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsEventsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** |  | 
 **clientIp** | **string** |  | 
 **contextAuthorizedApp** | **string** | Context Authorized application | 
 **contextModelApp** | **string** | Context Model App | 
 **contextModelName** | **string** | Context Model Name | 
 **contextModelPk** | **string** | Context Model Primary Key | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **tenantName** | **string** | Tenant name | 
 **username** | **string** | Username | 

### Return type

[**PaginatedEventList**](PaginatedEventList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsEventsPartialUpdate

> Event EventsEventsPartialUpdate(ctx, eventUuid).PatchedEventRequest(patchedEventRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    eventUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Event.
    patchedEventRequest := *openapiclient.NewPatchedEventRequest() // PatchedEventRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.EventsEventsPartialUpdate(context.Background(), eventUuid).PatchedEventRequest(patchedEventRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsEventsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsEventsPartialUpdate`: Event
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsEventsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventUuid** | **string** | A UUID string identifying this Event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsEventsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedEventRequest** | [**PatchedEventRequest**](PatchedEventRequest.md) |  | 

### Return type

[**Event**](Event.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsEventsPerMonthList

> []Coordinate EventsEventsPerMonthList(ctx).Action(action).Query(query).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    action := "action_example" // string |  (optional)
    query := "query_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.EventsEventsPerMonthList(context.Background()).Action(action).Query(query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsEventsPerMonthList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsEventsPerMonthList`: []Coordinate
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsEventsPerMonthList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsEventsPerMonthListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** |  | 
 **query** | **string** |  | 

### Return type

[**[]Coordinate**](Coordinate.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsEventsRetrieve

> Event EventsEventsRetrieve(ctx, eventUuid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    eventUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Event.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.EventsEventsRetrieve(context.Background(), eventUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsEventsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsEventsRetrieve`: Event
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsEventsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventUuid** | **string** | A UUID string identifying this Event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsEventsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Event**](Event.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsEventsTopPerUserList

> []EventTopPerUser EventsEventsTopPerUserList(ctx).Action(action).TopN(topN).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    action := "action_example" // string |  (optional)
    topN := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.EventsEventsTopPerUserList(context.Background()).Action(action).TopN(topN).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsEventsTopPerUserList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsEventsTopPerUserList`: []EventTopPerUser
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsEventsTopPerUserList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsEventsTopPerUserListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** |  | 
 **topN** | **int32** |  | 

### Return type

[**[]EventTopPerUser**](EventTopPerUser.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsEventsUpdate

> Event EventsEventsUpdate(ctx, eventUuid).EventRequest(eventRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    eventUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Event.
    eventRequest := *openapiclient.NewEventRequest(openapiclient.EventActions("login"), "App_example") // EventRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.EventsEventsUpdate(context.Background(), eventUuid).EventRequest(eventRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsEventsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsEventsUpdate`: Event
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsEventsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventUuid** | **string** | A UUID string identifying this Event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsEventsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventRequest** | [**EventRequest**](EventRequest.md) |  | 

### Return type

[**Event**](Event.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsNotificationsDestroy

> EventsNotificationsDestroy(ctx, uuid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EventsAPI.EventsNotificationsDestroy(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsNotificationsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Notification. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsNotificationsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsNotificationsList

> PaginatedNotificationList EventsNotificationsList(ctx).Body(body).Created(created).Event(event).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Seen(seen).Severity(severity).User(user).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    body := "body_example" // string |  (optional)
    created := time.Now() // time.Time |  (optional)
    event := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)
    seen := true // bool |  (optional)
    severity := "severity_example" // string | * `notice` - Notice * `warning` - Warning * `alert` - Alert (optional)
    user := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.EventsNotificationsList(context.Background()).Body(body).Created(created).Event(event).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Seen(seen).Severity(severity).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsNotificationsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsNotificationsList`: PaginatedNotificationList
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsNotificationsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsNotificationsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 
 **created** | **time.Time** |  | 
 **event** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **seen** | **bool** |  | 
 **severity** | **string** | * &#x60;notice&#x60; - Notice * &#x60;warning&#x60; - Warning * &#x60;alert&#x60; - Alert | 
 **user** | **int32** |  | 

### Return type

[**PaginatedNotificationList**](PaginatedNotificationList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsNotificationsMarkAllSeenCreate

> EventsNotificationsMarkAllSeenCreate(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EventsAPI.EventsNotificationsMarkAllSeenCreate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsNotificationsMarkAllSeenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEventsNotificationsMarkAllSeenCreateRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsNotificationsPartialUpdate

> Notification EventsNotificationsPartialUpdate(ctx, uuid).PatchedNotificationRequest(patchedNotificationRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification.
    patchedNotificationRequest := *openapiclient.NewPatchedNotificationRequest() // PatchedNotificationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.EventsNotificationsPartialUpdate(context.Background(), uuid).PatchedNotificationRequest(patchedNotificationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsNotificationsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsNotificationsPartialUpdate`: Notification
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsNotificationsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Notification. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsNotificationsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedNotificationRequest** | [**PatchedNotificationRequest**](PatchedNotificationRequest.md) |  | 

### Return type

[**Notification**](Notification.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsNotificationsRetrieve

> Notification EventsNotificationsRetrieve(ctx, uuid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.EventsNotificationsRetrieve(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsNotificationsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsNotificationsRetrieve`: Notification
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsNotificationsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Notification. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsNotificationsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Notification**](Notification.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsNotificationsUpdate

> Notification EventsNotificationsUpdate(ctx, uuid).NotificationRequest(notificationRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification.
    notificationRequest := *openapiclient.NewNotificationRequest() // NotificationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.EventsNotificationsUpdate(context.Background(), uuid).NotificationRequest(notificationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsNotificationsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsNotificationsUpdate`: Notification
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsNotificationsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Notification. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsNotificationsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationRequest** | [**NotificationRequest**](NotificationRequest.md) |  | 

### Return type

[**Notification**](Notification.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsNotificationsUsedByList

> []UsedBy EventsNotificationsUsedByList(ctx, uuid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.EventsNotificationsUsedByList(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsNotificationsUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsNotificationsUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsNotificationsUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Notification. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsNotificationsUsedByListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]UsedBy**](UsedBy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsRulesCreate

> NotificationRule EventsRulesCreate(ctx).NotificationRuleRequest(notificationRuleRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    notificationRuleRequest := *openapiclient.NewNotificationRuleRequest("Name_example") // NotificationRuleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.EventsRulesCreate(context.Background()).NotificationRuleRequest(notificationRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsRulesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsRulesCreate`: NotificationRule
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsRulesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsRulesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationRuleRequest** | [**NotificationRuleRequest**](NotificationRuleRequest.md) |  | 

### Return type

[**NotificationRule**](NotificationRule.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsRulesDestroy

> EventsRulesDestroy(ctx, pbmUuid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    pbmUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification Rule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EventsAPI.EventsRulesDestroy(context.Background(), pbmUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsRulesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pbmUuid** | **string** | A UUID string identifying this Notification Rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsRulesDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsRulesList

> PaginatedNotificationRuleList EventsRulesList(ctx).GroupName(groupName).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Severity(severity).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    groupName := "groupName_example" // string |  (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)
    severity := "severity_example" // string | Controls which severity level the created notifications will have.  * `notice` - Notice * `warning` - Warning * `alert` - Alert (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.EventsRulesList(context.Background()).GroupName(groupName).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).Severity(severity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsRulesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsRulesList`: PaginatedNotificationRuleList
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsRulesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsRulesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupName** | **string** |  | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **severity** | **string** | Controls which severity level the created notifications will have.  * &#x60;notice&#x60; - Notice * &#x60;warning&#x60; - Warning * &#x60;alert&#x60; - Alert | 

### Return type

[**PaginatedNotificationRuleList**](PaginatedNotificationRuleList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsRulesPartialUpdate

> NotificationRule EventsRulesPartialUpdate(ctx, pbmUuid).PatchedNotificationRuleRequest(patchedNotificationRuleRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    pbmUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification Rule.
    patchedNotificationRuleRequest := *openapiclient.NewPatchedNotificationRuleRequest() // PatchedNotificationRuleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.EventsRulesPartialUpdate(context.Background(), pbmUuid).PatchedNotificationRuleRequest(patchedNotificationRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsRulesPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsRulesPartialUpdate`: NotificationRule
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsRulesPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pbmUuid** | **string** | A UUID string identifying this Notification Rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsRulesPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedNotificationRuleRequest** | [**PatchedNotificationRuleRequest**](PatchedNotificationRuleRequest.md) |  | 

### Return type

[**NotificationRule**](NotificationRule.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsRulesRetrieve

> NotificationRule EventsRulesRetrieve(ctx, pbmUuid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    pbmUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification Rule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.EventsRulesRetrieve(context.Background(), pbmUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsRulesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsRulesRetrieve`: NotificationRule
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsRulesRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pbmUuid** | **string** | A UUID string identifying this Notification Rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsRulesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationRule**](NotificationRule.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsRulesUpdate

> NotificationRule EventsRulesUpdate(ctx, pbmUuid).NotificationRuleRequest(notificationRuleRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    pbmUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification Rule.
    notificationRuleRequest := *openapiclient.NewNotificationRuleRequest("Name_example") // NotificationRuleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.EventsRulesUpdate(context.Background(), pbmUuid).NotificationRuleRequest(notificationRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsRulesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsRulesUpdate`: NotificationRule
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsRulesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pbmUuid** | **string** | A UUID string identifying this Notification Rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsRulesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationRuleRequest** | [**NotificationRuleRequest**](NotificationRuleRequest.md) |  | 

### Return type

[**NotificationRule**](NotificationRule.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsRulesUsedByList

> []UsedBy EventsRulesUsedByList(ctx, pbmUuid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    pbmUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification Rule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.EventsRulesUsedByList(context.Background(), pbmUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsRulesUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsRulesUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsRulesUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pbmUuid** | **string** | A UUID string identifying this Notification Rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsRulesUsedByListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]UsedBy**](UsedBy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsTransportsCreate

> NotificationTransport EventsTransportsCreate(ctx).NotificationTransportRequest(notificationTransportRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    notificationTransportRequest := *openapiclient.NewNotificationTransportRequest("Name_example") // NotificationTransportRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.EventsTransportsCreate(context.Background()).NotificationTransportRequest(notificationTransportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsTransportsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsTransportsCreate`: NotificationTransport
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsTransportsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsTransportsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationTransportRequest** | [**NotificationTransportRequest**](NotificationTransportRequest.md) |  | 

### Return type

[**NotificationTransport**](NotificationTransport.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsTransportsDestroy

> EventsTransportsDestroy(ctx, uuid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification Transport.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EventsAPI.EventsTransportsDestroy(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsTransportsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Notification Transport. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsTransportsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsTransportsList

> PaginatedNotificationTransportList EventsTransportsList(ctx).Mode(mode).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).SendOnce(sendOnce).WebhookUrl(webhookUrl).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    mode := "mode_example" // string | * `local` - authentik inbuilt notifications * `webhook` - Generic Webhook * `webhook_slack` - Slack Webhook (Slack/Discord) * `email` - Email (optional)
    name := "name_example" // string |  (optional)
    ordering := "ordering_example" // string | Which field to use when ordering the results. (optional)
    page := int32(56) // int32 | A page number within the paginated result set. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    search := "search_example" // string | A search term. (optional)
    sendOnce := true // bool |  (optional)
    webhookUrl := "webhookUrl_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.EventsTransportsList(context.Background()).Mode(mode).Name(name).Ordering(ordering).Page(page).PageSize(pageSize).Search(search).SendOnce(sendOnce).WebhookUrl(webhookUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsTransportsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsTransportsList`: PaginatedNotificationTransportList
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsTransportsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEventsTransportsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mode** | **string** | * &#x60;local&#x60; - authentik inbuilt notifications * &#x60;webhook&#x60; - Generic Webhook * &#x60;webhook_slack&#x60; - Slack Webhook (Slack/Discord) * &#x60;email&#x60; - Email | 
 **name** | **string** |  | 
 **ordering** | **string** | Which field to use when ordering the results. | 
 **page** | **int32** | A page number within the paginated result set. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **search** | **string** | A search term. | 
 **sendOnce** | **bool** |  | 
 **webhookUrl** | **string** |  | 

### Return type

[**PaginatedNotificationTransportList**](PaginatedNotificationTransportList.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsTransportsPartialUpdate

> NotificationTransport EventsTransportsPartialUpdate(ctx, uuid).PatchedNotificationTransportRequest(patchedNotificationTransportRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification Transport.
    patchedNotificationTransportRequest := *openapiclient.NewPatchedNotificationTransportRequest() // PatchedNotificationTransportRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.EventsTransportsPartialUpdate(context.Background(), uuid).PatchedNotificationTransportRequest(patchedNotificationTransportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsTransportsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsTransportsPartialUpdate`: NotificationTransport
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsTransportsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Notification Transport. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsTransportsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchedNotificationTransportRequest** | [**PatchedNotificationTransportRequest**](PatchedNotificationTransportRequest.md) |  | 

### Return type

[**NotificationTransport**](NotificationTransport.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsTransportsRetrieve

> NotificationTransport EventsTransportsRetrieve(ctx, uuid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification Transport.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.EventsTransportsRetrieve(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsTransportsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsTransportsRetrieve`: NotificationTransport
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsTransportsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Notification Transport. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsTransportsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationTransport**](NotificationTransport.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsTransportsTestCreate

> NotificationTransportTest EventsTransportsTestCreate(ctx, uuid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification Transport.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.EventsTransportsTestCreate(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsTransportsTestCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsTransportsTestCreate`: NotificationTransportTest
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsTransportsTestCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Notification Transport. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsTransportsTestCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationTransportTest**](NotificationTransportTest.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsTransportsUpdate

> NotificationTransport EventsTransportsUpdate(ctx, uuid).NotificationTransportRequest(notificationTransportRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification Transport.
    notificationTransportRequest := *openapiclient.NewNotificationTransportRequest("Name_example") // NotificationTransportRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.EventsTransportsUpdate(context.Background(), uuid).NotificationTransportRequest(notificationTransportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsTransportsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsTransportsUpdate`: NotificationTransport
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsTransportsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Notification Transport. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsTransportsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationTransportRequest** | [**NotificationTransportRequest**](NotificationTransportRequest.md) |  | 

### Return type

[**NotificationTransport**](NotificationTransport.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EventsTransportsUsedByList

> []UsedBy EventsTransportsUsedByList(ctx, uuid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A UUID string identifying this Notification Transport.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.EventsTransportsUsedByList(context.Background(), uuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.EventsTransportsUsedByList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EventsTransportsUsedByList`: []UsedBy
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.EventsTransportsUsedByList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | A UUID string identifying this Notification Transport. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEventsTransportsUsedByListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]UsedBy**](UsedBy.md)

### Authorization

[authentik](../README.md#authentik)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

