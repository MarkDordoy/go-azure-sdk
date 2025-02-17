package reservationtransactions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationTransactionsClient struct {
	Client *resourcemanager.Client
}

func NewReservationTransactionsClientWithBaseURI(api environments.Api) (*ReservationTransactionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "reservationtransactions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReservationTransactionsClient: %+v", err)
	}

	return &ReservationTransactionsClient{
		Client: client,
	}, nil
}
