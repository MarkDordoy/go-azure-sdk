package fluxconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GitRepositoryDefinition struct {
	HttpsCACert           *string                  `json:"httpsCACert,omitempty"`
	HttpsUser             *string                  `json:"httpsUser,omitempty"`
	LocalAuthRef          *string                  `json:"localAuthRef,omitempty"`
	RepositoryRef         *RepositoryRefDefinition `json:"repositoryRef,omitempty"`
	SshKnownHosts         *string                  `json:"sshKnownHosts,omitempty"`
	SyncIntervalInSeconds *int64                   `json:"syncIntervalInSeconds,omitempty"`
	TimeoutInSeconds      *int64                   `json:"timeoutInSeconds,omitempty"`
	Url                   *string                  `json:"url,omitempty"`
}
