package virtualmachines

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineRestrictMovementState string

const (
	VirtualMachineRestrictMovementStateDisabled VirtualMachineRestrictMovementState = "Disabled"
	VirtualMachineRestrictMovementStateEnabled  VirtualMachineRestrictMovementState = "Enabled"
)

func PossibleValuesForVirtualMachineRestrictMovementState() []string {
	return []string{
		string(VirtualMachineRestrictMovementStateDisabled),
		string(VirtualMachineRestrictMovementStateEnabled),
	}
}

func parseVirtualMachineRestrictMovementState(input string) (*VirtualMachineRestrictMovementState, error) {
	vals := map[string]VirtualMachineRestrictMovementState{
		"disabled": VirtualMachineRestrictMovementStateDisabled,
		"enabled":  VirtualMachineRestrictMovementStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualMachineRestrictMovementState(input)
	return &out, nil
}
