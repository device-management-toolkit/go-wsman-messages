/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package models

// Shared DMTF power-state enumerations used by CIM_AssociatedPowerManagementService
// and any other wrapper that needs to decode power-state integers on the wire.
// These values follow the canonical DMTF CIM schema (2.49.0).

type (
	AvailableRequestedPowerStates int
	PowerState                    int
	RequestedPowerState           int
	TransitioningToPowerState     int
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

func (e TransitioningToPowerState) String() string {
	if s, ok := transitioningToPowerStateMap[e]; ok {
		return s
	}

	return ValueNotFound
}
