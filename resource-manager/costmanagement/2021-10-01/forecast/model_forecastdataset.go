package forecast

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ForecastDataset struct {
	Aggregation   *map[string]QueryAggregation `json:"aggregation,omitempty"`
	Configuration *QueryDatasetConfiguration   `json:"configuration,omitempty"`
	Filter        *QueryFilter                 `json:"filter,omitempty"`
	Granularity   *GranularityType             `json:"granularity,omitempty"`
}
