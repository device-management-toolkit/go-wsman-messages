/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package associatedpower

const (
	CIMAssociatedPowerManagementService string = "CIM_AssociatedPowerManagementService"
	ValueNotFound                       string = "Value not found in map"
)

const (
	AvailableRequestedPowerStatesOther AvailableRequestedPowerStates = iota + 1
	AvailableRequestedPowerStatesOn
	AvailableRequestedPowerStatesSleepLight
	AvailableRequestedPowerStatesSleepDeep
	AvailableRequestedPowerStatesPowerCycleSoft
	AvailableRequestedPowerStatesOffHard
	AvailableRequestedPowerStatesHibernate
	AvailableRequestedPowerStatesOffSoft
	AvailableRequestedPowerStatesPowerCycleHard
	AvailableRequestedPowerStatesMasterBusReset
	AvailableRequestedPowerStatesDiagnosticInterrupt
	AvailableRequestedPowerStatesPowerOffSoftGraceful
	AvailableRequestedPowerStatesPowerOffHardGraceful
	AvailableRequestedPowerStatesMasterBusResetGraceful
	AvailableRequestedPowerStatesPowerCycleSoftGraceful
	AvailableRequestedPowerStatesPowerCycleHardGraceful
)

// availableRequestedPowerStatesMap is a map of the AvailableRequestedPowerStates enumeration.
var availableRequestedPowerStatesMap = map[AvailableRequestedPowerStates]string{
	AvailableRequestedPowerStatesOther:                  "Other",
	AvailableRequestedPowerStatesOn:                     "On",
	AvailableRequestedPowerStatesSleepLight:             "SleepLight",
	AvailableRequestedPowerStatesSleepDeep:              "SleepDeep",
	AvailableRequestedPowerStatesPowerCycleSoft:         "PowerCycleSoft",
	AvailableRequestedPowerStatesOffHard:                "OffHard",
	AvailableRequestedPowerStatesHibernate:              "Hibernate",
	AvailableRequestedPowerStatesOffSoft:                "OffSoft",
	AvailableRequestedPowerStatesPowerCycleHard:         "PowerCycleHard",
	AvailableRequestedPowerStatesMasterBusReset:         "MasterBusReset",
	AvailableRequestedPowerStatesDiagnosticInterrupt:    "DiagnosticInterrupt",
	AvailableRequestedPowerStatesPowerOffSoftGraceful:   "PowerOffSoftGraceful",
	AvailableRequestedPowerStatesPowerOffHardGraceful:   "PowerOffHardGraceful",
	AvailableRequestedPowerStatesMasterBusResetGraceful: "MasterBusResetGraceful",
	AvailableRequestedPowerStatesPowerCycleSoftGraceful: "PowerCycleSoftGraceful",
	AvailableRequestedPowerStatesPowerCycleHardGraceful: "PowerCycleHardGraceful",
}

// String returns a human-readable string representation of the AvailableRequestedPowerStates enumeration.
func (e AvailableRequestedPowerStates) String() string {
	if s, ok := availableRequestedPowerStatesMap[e]; ok {
		return s
	}

	return ValueNotFound
}

const (
	PowerStateOther PowerState = iota + 1
	PowerStateOn
	PowerStateSleepLight
	PowerStateSleepDeep
	PowerStatePowerCycleSoft
	PowerStateOffHard
	PowerStateHibernate
	PowerStateOffSoft
	PowerStatePowerCycleHard
	PowerStateMasterBusReset
	PowerStateDiagnosticInterruptNMI
	PowerStatePowerOffSoftGraceful
	PowerStatePowerOffHardGraceful
	PowerStateMasterBusResetGraceful
	PowerStatePowerCycleSoftGraceful
	PowerStatePowerCycleHardGraceful
	PowerStateDiagnosticInterruptINIT
)

// powerStateMap is a map of the PowerState enumeration.
var powerStateMap = map[PowerState]string{
	PowerStateOther:                   "Other",
	PowerStateOn:                      "On",
	PowerStateSleepLight:              "SleepLight",
	PowerStateSleepDeep:               "SleepDeep",
	PowerStatePowerCycleSoft:          "PowerCycleSoft",
	PowerStateOffHard:                 "OffHard",
	PowerStateHibernate:               "Hibernate",
	PowerStateOffSoft:                 "OffSoft",
	PowerStatePowerCycleHard:          "PowerCycleHard",
	PowerStateMasterBusReset:          "MasterBusReset",
	PowerStateDiagnosticInterruptNMI:  "DiagnosticInterruptNMI",
	PowerStatePowerOffSoftGraceful:    "PowerOffSoftGraceful",
	PowerStatePowerOffHardGraceful:    "PowerOffHardGraceful",
	PowerStateMasterBusResetGraceful:  "MasterBusResetGraceful",
	PowerStatePowerCycleSoftGraceful:  "PowerCycleSoftGraceful",
	PowerStatePowerCycleHardGraceful:  "PowerCycleHardGraceful",
	PowerStateDiagnosticInterruptINIT: "DiagnosticInterruptINIT",
}

// String returns a human-readable string representation of the PowerState enumeration.
func (e PowerState) String() string {
	if s, ok := powerStateMap[e]; ok {
		return s
	}

	return ValueNotFound
}

const (
	RequestedPowerStateUnknown RequestedPowerState = iota
	RequestedPowerStateOther
	RequestedPowerStateOn
	RequestedPowerStateSleepLight
	RequestedPowerStateSleepDeep
	RequestedPowerStatePowerCycleSoft
	RequestedPowerStateOffHard
	RequestedPowerStateHibernate
	RequestedPowerStateOffSoft
	RequestedPowerStatePowerCycleHard
	RequestedPowerStateMasterBusReset
	RequestedPowerStateDiagnosticInterruptNMI
	RequestedPowerStateNotApplicable
	RequestedPowerStatePowerOffSoftGraceful
	RequestedPowerStatePowerOffHardGraceful
	RequestedPowerStateMasterBusResetGraceful
	RequestedPowerStatePowerCycleSoftGraceful
	RequestedPowerStatePowerCycleHardGraceful
	RequestedPowerStateDiagnosticInterruptINIT
)

// requestedPowerStateMap is a map of the RequestedPowerState enumeration.
var requestedPowerStateMap = map[RequestedPowerState]string{
	RequestedPowerStateUnknown:                 "Unknown",
	RequestedPowerStateOther:                   "Other",
	RequestedPowerStateOn:                      "On",
	RequestedPowerStateSleepLight:              "SleepLight",
	RequestedPowerStateSleepDeep:               "SleepDeep",
	RequestedPowerStatePowerCycleSoft:          "PowerCycleSoft",
	RequestedPowerStateOffHard:                 "OffHard",
	RequestedPowerStateHibernate:               "Hibernate",
	RequestedPowerStateOffSoft:                 "OffSoft",
	RequestedPowerStatePowerCycleHard:          "PowerCycleHard",
	RequestedPowerStateMasterBusReset:          "MasterBusReset",
	RequestedPowerStateDiagnosticInterruptNMI:  "DiagnosticInterruptNMI",
	RequestedPowerStateNotApplicable:           "NotApplicable",
	RequestedPowerStatePowerOffSoftGraceful:    "PowerOffSoftGraceful",
	RequestedPowerStatePowerOffHardGraceful:    "PowerOffHardGraceful",
	RequestedPowerStateMasterBusResetGraceful:  "MasterBusResetGraceful",
	RequestedPowerStatePowerCycleSoftGraceful:  "PowerCycleSoftGraceful",
	RequestedPowerStatePowerCycleHardGraceful:  "PowerCycleHardGraceful",
	RequestedPowerStateDiagnosticInterruptINIT: "DiagnosticInterruptINIT",
}

// String returns a human-readable string representation of the RequestedPowerState enumeration.
func (e RequestedPowerState) String() string {
	if s, ok := requestedPowerStateMap[e]; ok {
		return s
	}

	return ValueNotFound
}

const (
	TransitioningToPowerStateOther TransitioningToPowerState = iota + 1
	TransitioningToPowerStateOn
	TransitioningToPowerStateSleepLight
	TransitioningToPowerStateSleepDeep
	TransitioningToPowerStatePowerCycleSoft
	TransitioningToPowerStateOffHard
	TransitioningToPowerStateHibernate
	TransitioningToPowerStateOffSoft
	TransitioningToPowerStatePowerCycleHard
	TransitioningToPowerStateMasterBusReset
	TransitioningToPowerStateDiagnosticInterruptNMI
	TransitioningToPowerStatePowerOffSoftGraceful
	TransitioningToPowerStatePowerOffHardGraceful
	TransitioningToPowerStateMasterBusResetGraceful
	TransitioningToPowerStatePowerCycleSoftGraceful
	TransitioningToPowerStatePowerCycleHardGraceful
	TransitioningToPowerStateDiagnosticInterruptINIT
	TransitioningToPowerStateNotApplicable
	TransitioningToPowerStateNoChange
)

// transitioningToPowerStateMap is a map of the TransitioningToPowerState enumeration.
var transitioningToPowerStateMap = map[TransitioningToPowerState]string{
	TransitioningToPowerStateOther:                   "Other",
	TransitioningToPowerStateOn:                      "On",
	TransitioningToPowerStateSleepLight:              "SleepLight",
	TransitioningToPowerStateSleepDeep:               "SleepDeep",
	TransitioningToPowerStatePowerCycleSoft:          "PowerCycleSoft",
	TransitioningToPowerStateOffHard:                 "OffHard",
	TransitioningToPowerStateHibernate:               "Hibernate",
	TransitioningToPowerStateOffSoft:                 "OffSoft",
	TransitioningToPowerStatePowerCycleHard:          "PowerCycleHard",
	TransitioningToPowerStateMasterBusReset:          "MasterBusReset",
	TransitioningToPowerStateDiagnosticInterruptNMI:  "DiagnosticInterruptNMI",
	TransitioningToPowerStatePowerOffSoftGraceful:    "PowerOffSoftGraceful",
	TransitioningToPowerStatePowerOffHardGraceful:    "PowerOffHardGraceful",
	TransitioningToPowerStateMasterBusResetGraceful:  "MasterBusResetGraceful",
	TransitioningToPowerStatePowerCycleSoftGraceful:  "PowerCycleSoftGraceful",
	TransitioningToPowerStatePowerCycleHardGraceful:  "PowerCycleHardGraceful",
	TransitioningToPowerStateDiagnosticInterruptINIT: "DiagnosticInterruptINIT",
	TransitioningToPowerStateNotApplicable:           "NotApplicable",
	TransitioningToPowerStateNoChange:                "NoChange",
}

// String returns a human-readable string representation of the TransitioningToPowerState enumeration.
func (e TransitioningToPowerState) String() string {
	if s, ok := transitioningToPowerStateMap[e]; ok {
		return s
	}

	return ValueNotFound
}
