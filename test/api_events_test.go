/*
authentik

Testing EventsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_EventsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EventsAPIService EventsEventsActionsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EventsAPI.EventsEventsActionsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsEventsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EventsAPI.EventsEventsCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsEventsDestroy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var eventUuid string

		httpRes, err := apiClient.EventsAPI.EventsEventsDestroy(context.Background(), eventUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsEventsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EventsAPI.EventsEventsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsEventsPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var eventUuid string

		resp, httpRes, err := apiClient.EventsAPI.EventsEventsPartialUpdate(context.Background(), eventUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsEventsPerMonthList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EventsAPI.EventsEventsPerMonthList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsEventsRetrieve", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var eventUuid string

		resp, httpRes, err := apiClient.EventsAPI.EventsEventsRetrieve(context.Background(), eventUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsEventsTopPerUserList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EventsAPI.EventsEventsTopPerUserList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsEventsUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var eventUuid string

		resp, httpRes, err := apiClient.EventsAPI.EventsEventsUpdate(context.Background(), eventUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsNotificationsDestroy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		httpRes, err := apiClient.EventsAPI.EventsNotificationsDestroy(context.Background(), uuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsNotificationsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EventsAPI.EventsNotificationsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsNotificationsMarkAllSeenCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.EventsAPI.EventsNotificationsMarkAllSeenCreate(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsNotificationsPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		resp, httpRes, err := apiClient.EventsAPI.EventsNotificationsPartialUpdate(context.Background(), uuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsNotificationsRetrieve", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		resp, httpRes, err := apiClient.EventsAPI.EventsNotificationsRetrieve(context.Background(), uuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsNotificationsUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		resp, httpRes, err := apiClient.EventsAPI.EventsNotificationsUpdate(context.Background(), uuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsNotificationsUsedByList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		resp, httpRes, err := apiClient.EventsAPI.EventsNotificationsUsedByList(context.Background(), uuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsRulesCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EventsAPI.EventsRulesCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsRulesDestroy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pbmUuid string

		httpRes, err := apiClient.EventsAPI.EventsRulesDestroy(context.Background(), pbmUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsRulesList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EventsAPI.EventsRulesList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsRulesPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pbmUuid string

		resp, httpRes, err := apiClient.EventsAPI.EventsRulesPartialUpdate(context.Background(), pbmUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsRulesRetrieve", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pbmUuid string

		resp, httpRes, err := apiClient.EventsAPI.EventsRulesRetrieve(context.Background(), pbmUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsRulesUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pbmUuid string

		resp, httpRes, err := apiClient.EventsAPI.EventsRulesUpdate(context.Background(), pbmUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsRulesUsedByList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pbmUuid string

		resp, httpRes, err := apiClient.EventsAPI.EventsRulesUsedByList(context.Background(), pbmUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsTransportsCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EventsAPI.EventsTransportsCreate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsTransportsDestroy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		httpRes, err := apiClient.EventsAPI.EventsTransportsDestroy(context.Background(), uuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsTransportsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EventsAPI.EventsTransportsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsTransportsPartialUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		resp, httpRes, err := apiClient.EventsAPI.EventsTransportsPartialUpdate(context.Background(), uuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsTransportsRetrieve", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		resp, httpRes, err := apiClient.EventsAPI.EventsTransportsRetrieve(context.Background(), uuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsTransportsTestCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		resp, httpRes, err := apiClient.EventsAPI.EventsTransportsTestCreate(context.Background(), uuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsTransportsUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		resp, httpRes, err := apiClient.EventsAPI.EventsTransportsUpdate(context.Background(), uuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EventsAPIService EventsTransportsUsedByList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuid string

		resp, httpRes, err := apiClient.EventsAPI.EventsTransportsUsedByList(context.Background(), uuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
