package appplatform

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

type StoragesDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// StoragesDelete ...
func (c AppPlatformClient) StoragesDelete(ctx context.Context, id StorageId) (result StoragesDeleteOperationResponse, err error) {
	req, err := c.preparerForStoragesDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "StoragesDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForStoragesDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "StoragesDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// StoragesDeleteThenPoll performs StoragesDelete then polls until it's completed
func (c AppPlatformClient) StoragesDeleteThenPoll(ctx context.Context, id StorageId) error {
	result, err := c.StoragesDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing StoragesDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after StoragesDelete: %+v", err)
	}

	return nil
}

// preparerForStoragesDelete prepares the StoragesDelete request.
func (c AppPlatformClient) preparerForStoragesDelete(ctx context.Context, id StorageId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForStoragesDelete sends the StoragesDelete request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForStoragesDelete(ctx context.Context, req *http.Request) (future StoragesDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
