package aggregatedcost

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AggregatedCostClient struct {
	Client *resourcemanager.Client
}

func NewAggregatedCostClientWithBaseURI(api environments.Api) (*AggregatedCostClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "aggregatedcost", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AggregatedCostClient: %+v", err)
	}

	return &AggregatedCostClient{
		Client: client,
	}, nil
}
