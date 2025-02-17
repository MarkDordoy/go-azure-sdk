package addons

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByRoleOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]Addon

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByRoleOperationResponse, error)
}

type ListByRoleCompleteResult struct {
	Items []Addon
}

func (r ListByRoleOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByRoleOperationResponse) LoadMore(ctx context.Context) (resp ListByRoleOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByRole ...
func (c AddonsClient) ListByRole(ctx context.Context, id RoleId) (resp ListByRoleOperationResponse, err error) {
	req, err := c.preparerForListByRole(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "addons.AddonsClient", "ListByRole", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "addons.AddonsClient", "ListByRole", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByRole(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "addons.AddonsClient", "ListByRole", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByRole prepares the ListByRole request.
func (c AddonsClient) preparerForListByRole(ctx context.Context, id RoleId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/addons", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByRoleWithNextLink prepares the ListByRole request with the given nextLink token.
func (c AddonsClient) preparerForListByRoleWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByRole handles the response to the ListByRole request. The method always
// closes the http.Response Body.
func (c AddonsClient) responderForListByRole(resp *http.Response) (result ListByRoleOperationResponse, err error) {
	type page struct {
		Values   []json.RawMessage `json:"value"`
		NextLink *string           `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	temp := make([]Addon, 0)
	for i, v := range respObj.Values {
		val, err := unmarshalAddonImplementation(v)
		if err != nil {
			err = fmt.Errorf("unmarshalling item %d for Addon (%q): %+v", i, v, err)
			return result, err
		}
		temp = append(temp, val)
	}
	result.Model = &temp
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByRoleOperationResponse, err error) {
			req, err := c.preparerForListByRoleWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "addons.AddonsClient", "ListByRole", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "addons.AddonsClient", "ListByRole", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByRole(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "addons.AddonsClient", "ListByRole", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByRoleComplete retrieves all of the results into a single object
func (c AddonsClient) ListByRoleComplete(ctx context.Context, id RoleId) (ListByRoleCompleteResult, error) {
	return c.ListByRoleCompleteMatchingPredicate(ctx, id, AddonOperationPredicate{})
}

// ListByRoleCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c AddonsClient) ListByRoleCompleteMatchingPredicate(ctx context.Context, id RoleId, predicate AddonOperationPredicate) (resp ListByRoleCompleteResult, err error) {
	items := make([]Addon, 0)

	page, err := c.ListByRole(ctx, id)
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

	out := ListByRoleCompleteResult{
		Items: items,
	}
	return out, nil
}
