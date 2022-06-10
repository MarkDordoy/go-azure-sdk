package experiments

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExperimentsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewExperimentsClientWithBaseURI(endpoint string) ExperimentsClient {
	return ExperimentsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
