package domains

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainsClient struct {
	Client *resourcemanager.Client
}

func NewDomainsClientWithBaseURI(api environments.Api) (*DomainsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "domains", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DomainsClient: %+v", err)
	}

	return &DomainsClient{
		Client: client,
	}, nil
}
