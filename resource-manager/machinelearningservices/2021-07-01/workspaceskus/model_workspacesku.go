package workspaceskus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceSku struct {
	Capabilities *[]SKUCapability           `json:"capabilities,omitempty"`
	LocationInfo *[]ResourceSkuLocationInfo `json:"locationInfo,omitempty"`
	Locations    *[]string                  `json:"locations,omitempty"`
	Name         *string                    `json:"name,omitempty"`
	ResourceType *string                    `json:"resourceType,omitempty"`
	Restrictions *[]Restriction             `json:"restrictions,omitempty"`
	Tier         *string                    `json:"tier,omitempty"`
}
