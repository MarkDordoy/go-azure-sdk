package trustedidproviders

import "fmt"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

const defaultApiVersion = "2016-11-01"

func userAgent() string {
	return fmt.Sprintf("hashicorp/go-azure-sdk/trustedidproviders/%s", defaultApiVersion)
}
