package restorables

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorableMongodbResourcesListResult struct {
	Value *[]DatabaseRestoreResource `json:"value,omitempty"`
}
