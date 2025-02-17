package virtualmachineimages

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdgeZoneListSkusOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]VirtualMachineImageResource
}

// EdgeZoneListSkus ...
func (c VirtualMachineImagesClient) EdgeZoneListSkus(ctx context.Context, id VMImageOfferId) (result EdgeZoneListSkusOperationResponse, err error) {
	req, err := c.preparerForEdgeZoneListSkus(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimages.VirtualMachineImagesClient", "EdgeZoneListSkus", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimages.VirtualMachineImagesClient", "EdgeZoneListSkus", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForEdgeZoneListSkus(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimages.VirtualMachineImagesClient", "EdgeZoneListSkus", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForEdgeZoneListSkus prepares the EdgeZoneListSkus request.
func (c VirtualMachineImagesClient) preparerForEdgeZoneListSkus(ctx context.Context, id VMImageOfferId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/skus", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForEdgeZoneListSkus handles the response to the EdgeZoneListSkus request. The method always
// closes the http.Response Body.
func (c VirtualMachineImagesClient) responderForEdgeZoneListSkus(resp *http.Response) (result EdgeZoneListSkusOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
