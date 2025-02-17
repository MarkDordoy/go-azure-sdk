package projectresource

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProjectProvisioningState string

const (
	ProjectProvisioningStateDeleting  ProjectProvisioningState = "Deleting"
	ProjectProvisioningStateSucceeded ProjectProvisioningState = "Succeeded"
)

func PossibleValuesForProjectProvisioningState() []string {
	return []string{
		string(ProjectProvisioningStateDeleting),
		string(ProjectProvisioningStateSucceeded),
	}
}

func parseProjectProvisioningState(input string) (*ProjectProvisioningState, error) {
	vals := map[string]ProjectProvisioningState{
		"deleting":  ProjectProvisioningStateDeleting,
		"succeeded": ProjectProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProjectProvisioningState(input)
	return &out, nil
}

type ProjectSourcePlatform string

const (
	ProjectSourcePlatformMongoDb    ProjectSourcePlatform = "MongoDb"
	ProjectSourcePlatformMySQL      ProjectSourcePlatform = "MySQL"
	ProjectSourcePlatformPostgreSql ProjectSourcePlatform = "PostgreSql"
	ProjectSourcePlatformSQL        ProjectSourcePlatform = "SQL"
	ProjectSourcePlatformUnknown    ProjectSourcePlatform = "Unknown"
)

func PossibleValuesForProjectSourcePlatform() []string {
	return []string{
		string(ProjectSourcePlatformMongoDb),
		string(ProjectSourcePlatformMySQL),
		string(ProjectSourcePlatformPostgreSql),
		string(ProjectSourcePlatformSQL),
		string(ProjectSourcePlatformUnknown),
	}
}

func parseProjectSourcePlatform(input string) (*ProjectSourcePlatform, error) {
	vals := map[string]ProjectSourcePlatform{
		"mongodb":    ProjectSourcePlatformMongoDb,
		"mysql":      ProjectSourcePlatformMySQL,
		"postgresql": ProjectSourcePlatformPostgreSql,
		"sql":        ProjectSourcePlatformSQL,
		"unknown":    ProjectSourcePlatformUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProjectSourcePlatform(input)
	return &out, nil
}

type ProjectTargetPlatform string

const (
	ProjectTargetPlatformAzureDbForMySql      ProjectTargetPlatform = "AzureDbForMySql"
	ProjectTargetPlatformAzureDbForPostgreSql ProjectTargetPlatform = "AzureDbForPostgreSql"
	ProjectTargetPlatformMongoDb              ProjectTargetPlatform = "MongoDb"
	ProjectTargetPlatformSQLDB                ProjectTargetPlatform = "SQLDB"
	ProjectTargetPlatformSQLMI                ProjectTargetPlatform = "SQLMI"
	ProjectTargetPlatformUnknown              ProjectTargetPlatform = "Unknown"
)

func PossibleValuesForProjectTargetPlatform() []string {
	return []string{
		string(ProjectTargetPlatformAzureDbForMySql),
		string(ProjectTargetPlatformAzureDbForPostgreSql),
		string(ProjectTargetPlatformMongoDb),
		string(ProjectTargetPlatformSQLDB),
		string(ProjectTargetPlatformSQLMI),
		string(ProjectTargetPlatformUnknown),
	}
}

func parseProjectTargetPlatform(input string) (*ProjectTargetPlatform, error) {
	vals := map[string]ProjectTargetPlatform{
		"azuredbformysql":      ProjectTargetPlatformAzureDbForMySql,
		"azuredbforpostgresql": ProjectTargetPlatformAzureDbForPostgreSql,
		"mongodb":              ProjectTargetPlatformMongoDb,
		"sqldb":                ProjectTargetPlatformSQLDB,
		"sqlmi":                ProjectTargetPlatformSQLMI,
		"unknown":              ProjectTargetPlatformUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProjectTargetPlatform(input)
	return &out, nil
}
