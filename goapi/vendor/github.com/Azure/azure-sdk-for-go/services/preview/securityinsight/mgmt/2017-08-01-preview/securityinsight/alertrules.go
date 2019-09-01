package securityinsight

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// AlertRulesClient is the API spec for Microsoft.SecurityInsights (Azure Security Insights) resource provider
type AlertRulesClient struct {
	BaseClient
}

// NewAlertRulesClient creates an instance of the AlertRulesClient client.
func NewAlertRulesClient(subscriptionID string) AlertRulesClient {
	return NewAlertRulesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAlertRulesClientWithBaseURI creates an instance of the AlertRulesClient client.
func NewAlertRulesClientWithBaseURI(baseURI string, subscriptionID string) AlertRulesClient {
	return AlertRulesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates the alert rule.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// operationalInsightsResourceProvider - the namespace of workspaces resource provider-
// Microsoft.OperationalInsights.
// workspaceName - the name of the workspace.
// ruleID - alert rule ID
// alertRule - the alert rule
func (client AlertRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string, alertRule BasicAlertRule) (result AlertRuleModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlertRulesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.AlertRulesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, ruleID, alertRule)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.AlertRulesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.AlertRulesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.AlertRulesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client AlertRulesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string, alertRule BasicAlertRule) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationalInsightsResourceProvider": autorest.Encode("path", operationalInsightsResourceProvider),
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"ruleId":                              autorest.Encode("path", ruleID),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRules/{ruleId}", pathParameters),
		autorest.WithJSON(alertRule),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client AlertRulesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client AlertRulesClient) CreateOrUpdateResponder(resp *http.Response) (result AlertRuleModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdateAction creates or updates the action of alert rule.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// operationalInsightsResourceProvider - the namespace of workspaces resource provider-
// Microsoft.OperationalInsights.
// workspaceName - the name of the workspace.
// ruleID - alert rule ID
// actionID - action ID
// action - the action
func (client AlertRulesClient) CreateOrUpdateAction(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string, actionID string, action Action) (result Action, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlertRulesClient.CreateOrUpdateAction")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.AlertRulesClient", "CreateOrUpdateAction", err.Error())
	}

	req, err := client.CreateOrUpdateActionPreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, ruleID, actionID, action)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.AlertRulesClient", "CreateOrUpdateAction", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateActionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.AlertRulesClient", "CreateOrUpdateAction", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateActionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.AlertRulesClient", "CreateOrUpdateAction", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdateActionPreparer prepares the CreateOrUpdateAction request.
func (client AlertRulesClient) CreateOrUpdateActionPreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string, actionID string, action Action) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"actionId":                            autorest.Encode("path", actionID),
		"operationalInsightsResourceProvider": autorest.Encode("path", operationalInsightsResourceProvider),
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"ruleId":                              autorest.Encode("path", ruleID),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRules/{ruleId}/actions/{actionId}", pathParameters),
		autorest.WithJSON(action),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateActionSender sends the CreateOrUpdateAction request. The method will close the
// http.Response Body if it receives an error.
func (client AlertRulesClient) CreateOrUpdateActionSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// CreateOrUpdateActionResponder handles the response to the CreateOrUpdateAction request. The method always
// closes the http.Response Body.
func (client AlertRulesClient) CreateOrUpdateActionResponder(resp *http.Response) (result Action, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete the alert rule.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// operationalInsightsResourceProvider - the namespace of workspaces resource provider-
// Microsoft.OperationalInsights.
// workspaceName - the name of the workspace.
// ruleID - alert rule ID
func (client AlertRulesClient) Delete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlertRulesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.AlertRulesClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, ruleID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.AlertRulesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "securityinsight.AlertRulesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.AlertRulesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AlertRulesClient) DeletePreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationalInsightsResourceProvider": autorest.Encode("path", operationalInsightsResourceProvider),
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"ruleId":                              autorest.Encode("path", ruleID),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRules/{ruleId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AlertRulesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AlertRulesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteAction delete the action of alert rule.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// operationalInsightsResourceProvider - the namespace of workspaces resource provider-
// Microsoft.OperationalInsights.
// workspaceName - the name of the workspace.
// ruleID - alert rule ID
// actionID - action ID
func (client AlertRulesClient) DeleteAction(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string, actionID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlertRulesClient.DeleteAction")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.AlertRulesClient", "DeleteAction", err.Error())
	}

	req, err := client.DeleteActionPreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, ruleID, actionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.AlertRulesClient", "DeleteAction", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteActionSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "securityinsight.AlertRulesClient", "DeleteAction", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteActionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.AlertRulesClient", "DeleteAction", resp, "Failure responding to request")
	}

	return
}

// DeleteActionPreparer prepares the DeleteAction request.
func (client AlertRulesClient) DeleteActionPreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string, actionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"actionId":                            autorest.Encode("path", actionID),
		"operationalInsightsResourceProvider": autorest.Encode("path", operationalInsightsResourceProvider),
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"ruleId":                              autorest.Encode("path", ruleID),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRules/{ruleId}/actions/{actionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteActionSender sends the DeleteAction request. The method will close the
// http.Response Body if it receives an error.
func (client AlertRulesClient) DeleteActionSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteActionResponder handles the response to the DeleteAction request. The method always
// closes the http.Response Body.
func (client AlertRulesClient) DeleteActionResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the alert rule.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// operationalInsightsResourceProvider - the namespace of workspaces resource provider-
// Microsoft.OperationalInsights.
// workspaceName - the name of the workspace.
// ruleID - alert rule ID
func (client AlertRulesClient) Get(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string) (result AlertRuleModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlertRulesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.AlertRulesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, ruleID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.AlertRulesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.AlertRulesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.AlertRulesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client AlertRulesClient) GetPreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationalInsightsResourceProvider": autorest.Encode("path", operationalInsightsResourceProvider),
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"ruleId":                              autorest.Encode("path", ruleID),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRules/{ruleId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AlertRulesClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AlertRulesClient) GetResponder(resp *http.Response) (result AlertRuleModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAction gets the action of alert rule.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// operationalInsightsResourceProvider - the namespace of workspaces resource provider-
// Microsoft.OperationalInsights.
// workspaceName - the name of the workspace.
// ruleID - alert rule ID
// actionID - action ID
func (client AlertRulesClient) GetAction(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string, actionID string) (result Action, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlertRulesClient.GetAction")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.AlertRulesClient", "GetAction", err.Error())
	}

	req, err := client.GetActionPreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName, ruleID, actionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.AlertRulesClient", "GetAction", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetActionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.AlertRulesClient", "GetAction", resp, "Failure sending request")
		return
	}

	result, err = client.GetActionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.AlertRulesClient", "GetAction", resp, "Failure responding to request")
	}

	return
}

// GetActionPreparer prepares the GetAction request.
func (client AlertRulesClient) GetActionPreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string, ruleID string, actionID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"actionId":                            autorest.Encode("path", actionID),
		"operationalInsightsResourceProvider": autorest.Encode("path", operationalInsightsResourceProvider),
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"ruleId":                              autorest.Encode("path", ruleID),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRules/{ruleId}/actions/{actionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetActionSender sends the GetAction request. The method will close the
// http.Response Body if it receives an error.
func (client AlertRulesClient) GetActionSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetActionResponder handles the response to the GetAction request. The method always
// closes the http.Response Body.
func (client AlertRulesClient) GetActionResponder(resp *http.Response) (result Action, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all alert rules.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// operationalInsightsResourceProvider - the namespace of workspaces resource provider-
// Microsoft.OperationalInsights.
// workspaceName - the name of the workspace.
func (client AlertRulesClient) List(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (result AlertRulesListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlertRulesClient.List")
		defer func() {
			sc := -1
			if result.arl.Response.Response != nil {
				sc = result.arl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("securityinsight.AlertRulesClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.AlertRulesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.arl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "securityinsight.AlertRulesClient", "List", resp, "Failure sending request")
		return
	}

	result.arl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.AlertRulesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client AlertRulesClient) ListPreparer(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationalInsightsResourceProvider": autorest.Encode("path", operationalInsightsResourceProvider),
		"resourceGroupName":                   autorest.Encode("path", resourceGroupName),
		"subscriptionId":                      autorest.Encode("path", client.SubscriptionID),
		"workspaceName":                       autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2019-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{operationalInsightsResourceProvider}/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/alertRules", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AlertRulesClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AlertRulesClient) ListResponder(resp *http.Response) (result AlertRulesList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client AlertRulesClient) listNextResults(ctx context.Context, lastResults AlertRulesList) (result AlertRulesList, err error) {
	req, err := lastResults.alertRulesListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "securityinsight.AlertRulesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "securityinsight.AlertRulesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securityinsight.AlertRulesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client AlertRulesClient) ListComplete(ctx context.Context, resourceGroupName string, operationalInsightsResourceProvider string, workspaceName string) (result AlertRulesListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AlertRulesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, operationalInsightsResourceProvider, workspaceName)
	return
}
