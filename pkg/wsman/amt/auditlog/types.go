/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package auditlog

import (
	"encoding/xml"
	"time"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type Service struct {
	base message.Base
}

// INPUTS
// Request Types.
type ReadRecordsInput struct {
	XMLName    xml.Name `xml:"h:ReadRecords_INPUT"`
	H          string   `xml:"xmlns:h,attr"`
	StartIndex int      `xml:"h:StartIndex" json:"StartIndex"`
}

// OUTPUTS
// Response Types.
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName                xml.Name `xml:"Body"`
		EnumerateResponse      common.EnumerateResponse
		GetResponse            AuditLog
		PullResponse           PullResponse
		ReadRecordsResponse    ReadRecords_OUTPUT
		DecodedRecordsResponse []AuditLogRecord
	}
	PullResponse struct {
		XMLName       xml.Name   `xml:"PullResponse"`
		AuditLogItems []AuditLog `xml:"Items>AMT_AuditLog"`
	}

	AuditLog struct {
		XMLName                xml.Name        `xml:"AMT_AuditLog"`
		OverwritePolicy        OverwritePolicy `xml:"OverwritePolicy,omitempty"`        // OverwritePolicy is an integer enumeration that indicates whether the log, represented by the CIM_Log subclasses, can overwrite its entries.Unknown (0) indicates the log's overwrite policy is unknown
		CurrentNumberOfRecords int             `xml:"CurrentNumberOfRecords,omitempty"` // Current number of records in the Log
		MaxNumberOfRecords     int             `xml:"MaxNumberOfRecords,omitempty"`     // Maximum number of records that can be captured in the Log
		ElementName            string          `xml:"ElementName,omitempty"`            // A user-friendly name for the object
		EnabledState           int             `xml:"EnabledState,omitempty"`           // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element
		RequestedState         int             `xml:"RequestedState,omitempty"`         // RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested
		PercentageFree         int             `xml:"PercentageFree,omitempty"`         // Indicates the percentage of free space in the storage dedicated to the audit log
		Name                   string          `xml:"Name,omitempty"`                   // The Name property uniquely identifies the Service and provides an indication of the functionality that is managed
		TimeOfLastRecord       Datetime        `xml:"TimeOfLastRecord"`                 // Time stamp of the most recent entry in the log if such an entry exists
		AuditState             int             `xml:"AuditState,omitempty"`             // State of log
		MaxAllowedAuditors     int             `xml:"MaxAllowedAuditors,omitempty"`     // Maximum number of auditors allowed
		StoragePolicy          StoragePolicy   `xml:"StoragePolicy,omitempty"`          // AuditLog storage policy
		MinDaysToKeep          int             `xml:"MinDaysToKeep,omitempty"`          // Minimum number of days to keep records in the AuditLog
	}

	Datetime struct {
		Datetime string `xml:"Datetime,omitempty"`
	}

	ReadRecords_OUTPUT struct {
		XMLName          xml.Name `xml:"ReadRecords_OUTPUT,omitempty"`
		TotalRecordCount int      `xml:"TotalRecordCount,omitempty"` // The total number of records in the log.
		RecordsReturned  int      `xml:"RecordsReturned,omitempty"`  // The number of records returned + content of 10 records from the start index.
		EventRecords     []string `xml:"EventRecords,omitempty"`     // Notice: the values of this array are actually base64 encoded values. A list of event records.
		ReturnValue      int      `xml:"ReturnValue,omitempty"`      // ValueMap={0, 1, 2, 35} Values={PT_STATUS_SUCCESS, PT_STATUS_INTERNAL_ERROR, PT_STATUS_NOT_READY, PT_STATUS_INVALID_INDEX}
	}

	AuditLogRecord struct {
		AuditAppID     int       `json:"AuditAppId" binding:"required" example:"0"`
		EventID        int       `json:"EventId" binding:"required" example:"0"`
		InitiatorType  uint8     `json:"InitiatorType" binding:"required" example:"0"`
		AuditApp       string    `json:"AuditApp" binding:"required" example:"Security Admin"`
		Event          string    `json:"Event" binding:"required" example:"Provisioning Started"`
		Initiator      string    `json:"Initiator" binding:"required" example:"Local"`
		Time           time.Time `json:"Time" binding:"required" example:"2023-04-19T20:38:20.000Z"`
		MCLocationType uint8     `json:"MCLocationType" binding:"required" example:"0"`
		NetAddress     string    `json:"NetAddress" binding:"required" example:"127.0.0.1"`
		Ex             string    `json:"Ex" binding:"required" example:""`
		ExStr          string    `json:"ExStr" binding:"required" example:"Remote WSAMN"`
	}

	// OverwritePolicy is an integer enumeration that indicates whether the log, represented by the CIM_Log subclasses, can overwrite its entries.
	OverwritePolicy int

	// StoragePolicy is an integer enumeration that indicates the storage policy of the log.
	StoragePolicy int

	// EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
	EnabledState int

	// RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
	RequestedState int

	ProvisioningParameters struct {
		ProvisioningMethod        uint8
		HashType                  uint8
		TrustedRootCertHash       []byte
		NumberOfCertificates      uint8
		CertSerialNumbers         []string
		AdditionalCaSerialNumbers uint8
		ProvServFQDNLength        uint8
		ProvServFQDN              string
	}

	ACLEntry struct {
		ParameterModified uint8
		AccessType        uint8
		EntryState        uint8
		InitiatorType     uint8
		UsernameLength    uint8
		SID               uint32
		Username          string
		DomainLength      uint8
		Domain            string
	}

	RemoteControlEvent struct {
		SpecialCommand                  uint8
		SpecialCommandParameterHighByte uint8
		SpecialCommandParameterLowByte  uint8
		BootOptionsMaskByte1            uint8
		BootOptionsMaskByte2            uint8
		OEMParameterByte1               uint8
		OEMParameterByte2               uint8
	}

	FWVersion struct {
		Major  uint16
		Minor  uint16
		Hotfix uint16
		Build  uint16
	}

	FWUpdateFailure struct {
		Type   uint8
		Reason uint8
	}

	NetworkAdministrationEvent struct {
		InterfaceHandle    uint32
		DHCPEnabled        uint8
		IPV4Address        uint32
		SubnetMask         uint32
		Gateway            uint32
		PrimaryDNS         uint32
		SecondaryDNS       uint32
		HostNameLength     uint8
		HostName           string
		DomainNameLength   uint8
		DomainName         string
		VLANTag            uint16
		LinkPolicy         uint32
		IPV6Enabled        uint8
		InterfaceIDGenType uint8
		InterfaceID        []uint8
		IPV6Address        []uint8
		IPV6Gateway        []uint8
		IPV6PrimaryDNS     []uint8
		IPV6SecondaryDNS   []uint8
	}

	StorageAdministrationEvent struct {
		MaxPartnerStorage                uint32
		MaxNonPartnerTotalAllocationSize uint32
	}

	EventManagerEvent struct {
		PolicyID              uint8
		SubscriptionAlertType uint8
		IPAddrType            uint8
		AlertTargetIPAddress  []uint8
		Freeze                uint8
	}

	SystemDefenseManagerEvent struct {
		FilterHandle       uint32
		PolicyHandle       uint32
		HardwareInterface  uint32
		InterfaceHandle    uint32
		BlockAll           uint8
		BlockOffensivePort uint8
	}

	AgentPresenceManagerEvent struct {
		AgentID            []uint8
		AgentHeartBeatTime uint16
		AgentStartupTime   uint16
	}

	WirelessConfigurationEvent struct {
		SSID                   []uint8
		ProfilePriority        uint8
		ProfileNameLength      uint8
		ProfileName            []uint8
		ProfileSync            uint32
		Timeout                uint32
		LinkPreference         uint32
		ProfileSharingWithUEFI uint8
	}

	UserOptInEvent struct {
		PreviousOptInPolicy uint8
		CurrentOptInPolicy  uint8
		OperationStatus     uint8
	}
)
