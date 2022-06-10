package flux

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

type ConfigurationsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]FluxConfiguration

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ConfigurationsListOperationResponse, error)
}

type ConfigurationsListCompleteResult struct {
	Items []FluxConfiguration
}

func (r ConfigurationsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ConfigurationsListOperationResponse) LoadMore(ctx context.Context) (resp ConfigurationsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// ConfigurationsList ...
func (c FluxClient) ConfigurationsList(ctx context.Context, id ProviderId) (resp ConfigurationsListOperationResponse, err error) {
	req, err := c.preparerForConfigurationsList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "flux.FluxClient", "ConfigurationsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "flux.FluxClient", "ConfigurationsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForConfigurationsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "flux.FluxClient", "ConfigurationsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// ConfigurationsListComplete retrieves all of the results into a single object
func (c FluxClient) ConfigurationsListComplete(ctx context.Context, id ProviderId) (ConfigurationsListCompleteResult, error) {
	return c.ConfigurationsListCompleteMatchingPredicate(ctx, id, FluxConfigurationOperationPredicate{})
}

// ConfigurationsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c FluxClient) ConfigurationsListCompleteMatchingPredicate(ctx context.Context, id ProviderId, predicate FluxConfigurationOperationPredicate) (resp ConfigurationsListCompleteResult, err error) {
	items := make([]FluxConfiguration, 0)

	page, err := c.ConfigurationsList(ctx, id)
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

	out := ConfigurationsListCompleteResult{
		Items: items,
	}
	return out, nil
}

// preparerForConfigurationsList prepares the ConfigurationsList request.
func (c FluxClient) preparerForConfigurationsList(ctx context.Context, id ProviderId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.KubernetesConfiguration/fluxConfigurations", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForConfigurationsListWithNextLink prepares the ConfigurationsList request with the given nextLink token.
func (c FluxClient) preparerForConfigurationsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForConfigurationsList handles the response to the ConfigurationsList request. The method always
// closes the http.Response Body.
func (c FluxClient) responderForConfigurationsList(resp *http.Response) (result ConfigurationsListOperationResponse, err error) {
	type page struct {
		Values   []FluxConfiguration `json:"value"`
		NextLink *string             `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ConfigurationsListOperationResponse, err error) {
			req, err := c.preparerForConfigurationsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "flux.FluxClient", "ConfigurationsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "flux.FluxClient", "ConfigurationsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForConfigurationsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "flux.FluxClient", "ConfigurationsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}
