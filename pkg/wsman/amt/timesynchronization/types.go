/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package timesynchronization

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

// Response Types.
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName                          xml.Name `xml:"Body"`
		GetResponse                      TimeSynchronizationServiceResponse
		EnumerateResponse                common.EnumerateResponse
		PullResponse                     PullResponse
		GetLowAccuracyTimeSynchResponse  GetLowAccuracyTimeSynchResponse
		SetHighAccuracyTimeSynchResponse SetHighAccuracyTimeSynchResponse
	}

	PullResponse struct {
		XMLName                         xml.Name                             `xml:"PullResponse"`
		TimeSynchronizationServiceItems []TimeSynchronizationServiceResponse `xml:"Items>AMT_TimeSynchronizationService"`
	}

	TimeSynchronizationServiceResponse struct {
		XMLName                 xml.Name             `xml:"AMT_TimeSynchronizationService"`
		Name                    string               `xml:"Name,omitempty"`                    // The Name property uniquely identifies the Service and provides an indication of the functionality that is managed. This functionality is described in more detail in the Description property of the object.
		CreationClassName       string               `xml:"CreationClassName,omitempty"`       // CreationClassName indicates the name of the class or the subclass that is used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
		SystemName              string               `xml:"SystemName,omitempty"`              // The Name of the scoping System.
		SystemCreationClassName string               `xml:"SystemCreationClassName,omitempty"` // The CreationClassName of the scoping System.
		ElementName             string               `xml:"ElementName,omitempty"`             // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		EnabledState            EnabledState         `xml:"EnabledState,omitempty"`            // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		RequestedState          RequestedState       `xml:"RequestedState,omitempty"`          // RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
		LocalTimeSyncEnabled    LocalTimeSyncEnabled `xml:"LocalTimeSyncEnabled,omitempty"`    // Determines if user with LOCAL_SYSTEM_REALM permission can set the time.
		TimeSource              TimeSource           `xml:"TimeSource,omitempty"`              // Determines if RTC was set to UTC by any configuration SW.
	}

	// ValueMap={0, 1}
	//
	// Values={PT_STATUS_SUCCESS, PT_STATUS_INTERNAL_ERROR}.
	GetLowAccuracyTimeSynchResponse struct {
		XMLName     xml.Name    `xml:"GetLowAccuracyTimeSynch_OUTPUT"`
		Ta0         int64       `xml:"Ta0"`
		ReturnValue ReturnValue `xml:"ReturnValue"`
	}

	// ValueMap={0, 1, 36, 38}
	//
	// Values={PT_STATUS_SUCCESS, PT_STATUS_INTERNAL_ERROR, PT_STATUS_INVALID_PARAMETER, PT_STATUS_FLASH_WRITE_LIMIT_EXCEEDED}.
	SetHighAccuracyTimeSynchResponse struct {
		XMLName     xml.Name    `xml:"SetHighAccuracyTimeSynch_OUTPUT"`
		ReturnValue ReturnValue `xml:"ReturnValue"`
	}

	// EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
	EnabledState int
	// LocalTimeSyncEnabled is an integer enumeration that determines if user with LOCAL_SYSTEM_REALM permission can set the time.
	LocalTimeSyncEnabled int
	// RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
	RequestedState int
	// ReturnValue is an integer enumeration that indicates the success or failure of the operation.
	ReturnValue int
	// TimeSource is an integer enumeration that determines if RTC was set to UTC by any configuration SW.
	TimeSource int
)

// Request Types.
type (
	// Ta0: The time value received from invoking GetLowAccuracyTimeSynch().
	//
	// Tm1: The remote client timestamp after getting a response from GetLowAccuracyTimeSynch().
	//
	// Tm2: The remote client timestamp obtained immediately prior to invoking this method.
	SetHighAccuracyTimeSynch_INPUT struct {
		XMLName xml.Name `xml:"h:SetHighAccuracyTimeSynch_INPUT"`
		H       string   `xml:"xmlns:h,attr"`
		Ta0     int64    `xml:"h:Ta0"`
		Tm1     int64    `xml:"h:Tm1"`
		Tm2     int64    `xml:"h:Tm2"`
	}
)
