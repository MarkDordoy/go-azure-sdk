package openshiftclusters

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// Update ...
func (c OpenShiftClustersClient) Update(ctx context.Context, id ProviderOpenShiftClusterId, input OpenShiftClusterUpdate) (result UpdateOperationResponse, err error) {
	req, err := c.preparerForUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "openshiftclusters.OpenShiftClustersClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "openshiftclusters.OpenShiftClustersClient", "Update", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// UpdateThenPoll performs Update then polls until it's completed
func (c OpenShiftClustersClient) UpdateThenPoll(ctx context.Context, id ProviderOpenShiftClusterId, input OpenShiftClusterUpdate) error {
	result, err := c.Update(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing Update: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Update: %+v", err)
	}

	return nil
}

// preparerForUpdate prepares the Update request.
func (c OpenShiftClustersClient) preparerForUpdate(ctx context.Context, id ProviderOpenShiftClusterId, input OpenShiftClusterUpdate) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForUpdate sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (c OpenShiftClustersClient) senderForUpdate(ctx context.Context, req *http.Request) (future UpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
