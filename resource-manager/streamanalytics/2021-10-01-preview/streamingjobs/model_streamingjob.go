package streamingjobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StreamingJob struct {
	Id         *string                 `json:"id,omitempty"`
	Identity   *Identity               `json:"identity,omitempty"`
	Location   *string                 `json:"location,omitempty"`
	Name       *string                 `json:"name,omitempty"`
	Properties *StreamingJobProperties `json:"properties,omitempty"`
	Sku        *Sku                    `json:"sku,omitempty"`
	Tags       *map[string]string      `json:"tags,omitempty"`
	Type       *string                 `json:"type,omitempty"`
}
