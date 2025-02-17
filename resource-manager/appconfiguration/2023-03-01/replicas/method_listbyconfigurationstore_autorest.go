package replicas

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByConfigurationStoreOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Replica

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByConfigurationStoreOperationResponse, error)
}

type ListByConfigurationStoreCompleteResult struct {
	Items []Replica
}

func (r ListByConfigurationStoreOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByConfigurationStoreOperationResponse) LoadMore(ctx context.Context) (resp ListByConfigurationStoreOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByConfigurationStore ...
func (c ReplicasClient) ListByConfigurationStore(ctx context.Context, id ConfigurationStoreId) (resp ListByConfigurationStoreOperationResponse, err error) {
	req, err := c.preparerForListByConfigurationStore(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "replicas.ReplicasClient", "ListByConfigurationStore", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "replicas.ReplicasClient", "ListByConfigurationStore", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByConfigurationStore(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "replicas.ReplicasClient", "ListByConfigurationStore", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByConfigurationStore prepares the ListByConfigurationStore request.
func (c ReplicasClient) preparerForListByConfigurationStore(ctx context.Context, id ConfigurationStoreId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/replicas", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByConfigurationStoreWithNextLink prepares the ListByConfigurationStore request with the given nextLink token.
func (c ReplicasClient) preparerForListByConfigurationStoreWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
	uri, err := url.Parse(nextLink)
	if err != nil {
		return nil, fmt.Errorf("parsing nextLink %q: %+v", nextLink, err)
	}
	queryParameters := map[string]interface{}{}
	for k, v := range uri.Query() {
		if len(v) == 0 {
			continue
		}
		val := v[0]
		val = autorest.Encode("query", val)
		queryParameters[k] = val
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByConfigurationStore handles the response to the ListByConfigurationStore request. The method always
// closes the http.Response Body.
func (c ReplicasClient) responderForListByConfigurationStore(resp *http.Response) (result ListByConfigurationStoreOperationResponse, err error) {
	type page struct {
		Values   []Replica `json:"value"`
		NextLink *string   `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	result.Model = &respObj.Values
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByConfigurationStoreOperationResponse, err error) {
			req, err := c.preparerForListByConfigurationStoreWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "replicas.ReplicasClient", "ListByConfigurationStore", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "replicas.ReplicasClient", "ListByConfigurationStore", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByConfigurationStore(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "replicas.ReplicasClient", "ListByConfigurationStore", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByConfigurationStoreComplete retrieves all of the results into a single object
func (c ReplicasClient) ListByConfigurationStoreComplete(ctx context.Context, id ConfigurationStoreId) (ListByConfigurationStoreCompleteResult, error) {
	return c.ListByConfigurationStoreCompleteMatchingPredicate(ctx, id, ReplicaOperationPredicate{})
}

// ListByConfigurationStoreCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c ReplicasClient) ListByConfigurationStoreCompleteMatchingPredicate(ctx context.Context, id ConfigurationStoreId, predicate ReplicaOperationPredicate) (resp ListByConfigurationStoreCompleteResult, err error) {
	items := make([]Replica, 0)

	page, err := c.ListByConfigurationStore(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading the initial page: %+v", err)
		return
	}
	if page.Model != nil {
		for _, v := range *page.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	for page.HasMore() {
		page, err = page.LoadMore(ctx)
		if err != nil {
			err = fmt.Errorf("loading the next page: %+v", err)
			return
		}

		if page.Model != nil {
			for _, v := range *page.Model {
				if predicate.Matches(v) {
					items = append(items, v)
				}
			}
		}
	}

	out := ListByConfigurationStoreCompleteResult{
		Items: items,
	}
	return out, nil
}
