package integrationruntimes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationRuntimeStatus struct {
	DataFactoryName *string                  `json:"dataFactoryName,omitempty"`
	State           *IntegrationRuntimeState `json:"state,omitempty"`
	Type            IntegrationRuntimeType   `json:"type"`
}
