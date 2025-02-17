package appplatform

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicesRegenerateTestKeyOperationResponse struct {
	HttpResponse *http.Response
	Model        *TestKeys
}

// ServicesRegenerateTestKey ...
func (c AppPlatformClient) ServicesRegenerateTestKey(ctx context.Context, id SpringId, input RegenerateTestKeyRequestPayload) (result ServicesRegenerateTestKeyOperationResponse, err error) {
	req, err := c.preparerForServicesRegenerateTestKey(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ServicesRegenerateTestKey", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ServicesRegenerateTestKey", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForServicesRegenerateTestKey(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ServicesRegenerateTestKey", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForServicesRegenerateTestKey prepares the ServicesRegenerateTestKey request.
func (c AppPlatformClient) preparerForServicesRegenerateTestKey(ctx context.Context, id SpringId, input RegenerateTestKeyRequestPayload) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/regenerateTestKey", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForServicesRegenerateTestKey handles the response to the ServicesRegenerateTestKey request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForServicesRegenerateTestKey(resp *http.Response) (result ServicesRegenerateTestKeyOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
