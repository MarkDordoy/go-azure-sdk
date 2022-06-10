package vmcollectionupdate

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMCollectionUpdate struct {
	OperationName *OperationName `json:"operationName,omitempty"`
	VmResourceId  *string        `json:"vmResourceId,omitempty"`
}
