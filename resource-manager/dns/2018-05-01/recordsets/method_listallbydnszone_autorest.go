package recordsets

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

type ListAllByDnsZoneOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]RecordSet

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListAllByDnsZoneOperationResponse, error)
}

type ListAllByDnsZoneCompleteResult struct {
	Items []RecordSet
}

func (r ListAllByDnsZoneOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListAllByDnsZoneOperationResponse) LoadMore(ctx context.Context) (resp ListAllByDnsZoneOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListAllByDnsZoneOperationOptions struct {
	Recordsetnamesuffix *string
	Top                 *int64
}

func DefaultListAllByDnsZoneOperationOptions() ListAllByDnsZoneOperationOptions {
	return ListAllByDnsZoneOperationOptions{}
}

func (o ListAllByDnsZoneOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListAllByDnsZoneOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Recordsetnamesuffix != nil {
		out["$recordsetnamesuffix"] = *o.Recordsetnamesuffix
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// ListAllByDnsZone ...
func (c RecordSetsClient) ListAllByDnsZone(ctx context.Context, id DnsZoneId, options ListAllByDnsZoneOperationOptions) (resp ListAllByDnsZoneOperationResponse, err error) {
	req, err := c.preparerForListAllByDnsZone(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recordsets.RecordSetsClient", "ListAllByDnsZone", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "recordsets.RecordSetsClient", "ListAllByDnsZone", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListAllByDnsZone(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recordsets.RecordSetsClient", "ListAllByDnsZone", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// ListAllByDnsZoneComplete retrieves all of the results into a single object
func (c RecordSetsClient) ListAllByDnsZoneComplete(ctx context.Context, id DnsZoneId, options ListAllByDnsZoneOperationOptions) (ListAllByDnsZoneCompleteResult, error) {
	return c.ListAllByDnsZoneCompleteMatchingPredicate(ctx, id, options, RecordSetOperationPredicate{})
}

// ListAllByDnsZoneCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c RecordSetsClient) ListAllByDnsZoneCompleteMatchingPredicate(ctx context.Context, id DnsZoneId, options ListAllByDnsZoneOperationOptions, predicate RecordSetOperationPredicate) (resp ListAllByDnsZoneCompleteResult, err error) {
	items := make([]RecordSet, 0)

	page, err := c.ListAllByDnsZone(ctx, id, options)
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

	out := ListAllByDnsZoneCompleteResult{
		Items: items,
	}
	return out, nil
}

// preparerForListAllByDnsZone prepares the ListAllByDnsZone request.
func (c RecordSetsClient) preparerForListAllByDnsZone(ctx context.Context, id DnsZoneId, options ListAllByDnsZoneOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/all", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListAllByDnsZoneWithNextLink prepares the ListAllByDnsZone request with the given nextLink token.
func (c RecordSetsClient) preparerForListAllByDnsZoneWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListAllByDnsZone handles the response to the ListAllByDnsZone request. The method always
// closes the http.Response Body.
func (c RecordSetsClient) responderForListAllByDnsZone(resp *http.Response) (result ListAllByDnsZoneOperationResponse, err error) {
	type page struct {
		Values   []RecordSet `json:"value"`
		NextLink *string     `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListAllByDnsZoneOperationResponse, err error) {
			req, err := c.preparerForListAllByDnsZoneWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "recordsets.RecordSetsClient", "ListAllByDnsZone", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "recordsets.RecordSetsClient", "ListAllByDnsZone", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListAllByDnsZone(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "recordsets.RecordSetsClient", "ListAllByDnsZone", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}
