package registeredserverresource

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegisteredServersGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *RegisteredServer
}

// RegisteredServersGet ...
func (c RegisteredServerResourceClient) RegisteredServersGet(ctx context.Context, id RegisteredServerId) (result RegisteredServersGetOperationResponse, err error) {
	req, err := c.preparerForRegisteredServersGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registeredserverresource.RegisteredServerResourceClient", "RegisteredServersGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "registeredserverresource.RegisteredServerResourceClient", "RegisteredServersGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRegisteredServersGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registeredserverresource.RegisteredServerResourceClient", "RegisteredServersGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRegisteredServersGet prepares the RegisteredServersGet request.
func (c RegisteredServerResourceClient) preparerForRegisteredServersGet(ctx context.Context, id RegisteredServerId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRegisteredServersGet handles the response to the RegisteredServersGet request. The method always
// closes the http.Response Body.
func (c RegisteredServerResourceClient) responderForRegisteredServersGet(resp *http.Response) (result RegisteredServersGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
