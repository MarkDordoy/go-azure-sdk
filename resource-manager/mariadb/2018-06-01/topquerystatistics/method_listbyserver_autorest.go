package topquerystatistics

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

type ListByServerOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]QueryStatistic

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByServerOperationResponse, error)
}

type ListByServerCompleteResult struct {
	Items []QueryStatistic
}

func (r ListByServerOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByServerOperationResponse) LoadMore(ctx context.Context) (resp ListByServerOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ListByServer ...
func (c TopQueryStatisticsClient) ListByServer(ctx context.Context, id ServerId, input TopQueryStatisticsInput) (resp ListByServerOperationResponse, err error) {
	req, err := c.preparerForListByServer(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "topquerystatistics.TopQueryStatisticsClient", "ListByServer", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "topquerystatistics.TopQueryStatisticsClient", "ListByServer", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByServer(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "topquerystatistics.TopQueryStatisticsClient", "ListByServer", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByServer prepares the ListByServer request.
func (c TopQueryStatisticsClient) preparerForListByServer(ctx context.Context, id ServerId, input TopQueryStatisticsInput) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/topQueryStatistics", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByServerWithNextLink prepares the ListByServer request with the given nextLink token.
func (c TopQueryStatisticsClient) preparerForListByServerWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByServer handles the response to the ListByServer request. The method always
// closes the http.Response Body.
func (c TopQueryStatisticsClient) responderForListByServer(resp *http.Response) (result ListByServerOperationResponse, err error) {
	type page struct {
		Values   []QueryStatistic `json:"value"`
		NextLink *string          `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByServerOperationResponse, err error) {
			req, err := c.preparerForListByServerWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "topquerystatistics.TopQueryStatisticsClient", "ListByServer", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "topquerystatistics.TopQueryStatisticsClient", "ListByServer", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByServer(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "topquerystatistics.TopQueryStatisticsClient", "ListByServer", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByServerComplete retrieves all of the results into a single object
func (c TopQueryStatisticsClient) ListByServerComplete(ctx context.Context, id ServerId, input TopQueryStatisticsInput) (ListByServerCompleteResult, error) {
	return c.ListByServerCompleteMatchingPredicate(ctx, id, input, QueryStatisticOperationPredicate{})
}

// ListByServerCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c TopQueryStatisticsClient) ListByServerCompleteMatchingPredicate(ctx context.Context, id ServerId, input TopQueryStatisticsInput, predicate QueryStatisticOperationPredicate) (resp ListByServerCompleteResult, err error) {
	items := make([]QueryStatistic, 0)

	page, err := c.ListByServer(ctx, id, input)
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

	out := ListByServerCompleteResult{
		Items: items,
	}
	return out, nil
}
