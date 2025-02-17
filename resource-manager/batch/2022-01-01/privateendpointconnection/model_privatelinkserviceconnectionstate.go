package privateendpointconnection

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateLinkServiceConnectionState struct {
	ActionRequired *string                            `json:"actionRequired,omitempty"`
	Description    *string                            `json:"description,omitempty"`
	Status         PrivateLinkServiceConnectionStatus `json:"status"`
}
