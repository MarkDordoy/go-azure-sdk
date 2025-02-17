package operations

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationResourceType string

const (
	ConfigurationResourceTypeMicrosoftPointAppConfigurationConfigurationStores ConfigurationResourceType = "Microsoft.AppConfiguration/configurationStores"
)

func PossibleValuesForConfigurationResourceType() []string {
	return []string{
		string(ConfigurationResourceTypeMicrosoftPointAppConfigurationConfigurationStores),
	}
}

func parseConfigurationResourceType(input string) (*ConfigurationResourceType, error) {
	vals := map[string]ConfigurationResourceType{
		"microsoft.appconfiguration/configurationstores": ConfigurationResourceTypeMicrosoftPointAppConfigurationConfigurationStores,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConfigurationResourceType(input)
	return &out, nil
}
