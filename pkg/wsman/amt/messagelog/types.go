/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package messagelog

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
		XMLName                       xml.Name `xml:"Body"`
		GetResponse                   MessageLogResponse
		EnumerateResponse             common.EnumerateResponse
		PullResponse                  PullResponse
		GetRecordsResponse            GetRecordsResponse
		PositionToFirstRecordResponse PositionToFirstRecordResponse
	}

	PullResponse struct {
		XMLName         xml.Name             `xml:"PullResponse"`
		MessageLogItems []MessageLogResponse `xml:"Items>AMT_MessageLog"`
	}

	MessageLogResponse struct {
		XMLName                xml.Name            `xml:"AMT_MessageLog"`
		Capabilities           []Capabilities      `xml:"Capabilities"`           // An array of integers indicating the Log capabilities. Information such as "Write Record Supported" (value= 2) or "Variable Length Records Supported" (8) is specified in this property.
		CharacterSet           CharacterSet        `xml:"CharacterSet"`           // An enumeration describing the character set used to record data in the individual Log entries. For example, the Log records may contain ASCII data (value=2), or be raw octet strings (value=10).
		CreationClassName      string              `xml:"CreationClassName"`      // CreationClassName indicates the name of the class or the subclass used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
		CurrentNumberOfRecords int                 `xml:"CurrentNumberOfRecords"` // Current number of records in the Log.
		ElementName            string              `xml:"ElementName"`            // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information. Note that the Name property of ManagedSystemElement is also defined as a user-friendly name. But, it is often subclassed to be a Key. It is not reasonable that the same property can convey both identity and a user-friendly name, without inconsistencies. Where Name exists and is not a Key (such as for instances of LogicalDevice), the same information can be present in both the Name and ElementName properties. Note that if there is an associated instance of CIM_EnabledLogicalElementCapabilities, restrictions on this properties may exist as defined in ElementNameMask and MaxElementNameLen properties defined in that class.
		EnabledDefault         EnabledDefault      `xml:"EnabledDefault"`         // An enumerated value indicating an administrator's default or startup configuration for the Enabled State of an element. By default, the element is "Enabled" (value=2).
		EnabledState           EnabledState        `xml:"EnabledState"`           // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		HealthState            HealthState         `xml:"HealthState"`            // Indicates the current health of the element.
		IsFrozen               bool                `xml:"IsFrozen"`               // Boolean indicating that the Log is currently frozen and modifications are not allowed.
		LastChange             LastChange          `xml:"LastChange"`             // An enumeration describing the last change to the MessageLog.
		LogState               LogState            `xml:"LogState"`               // LogState is an integer enumeration that indicates the current state of a log represented by CIM_Log subclasses.
		MaxLogSize             int                 `xml:"MaxLogSize"`             // The maximum size, in bytes, to which the Log can grow. If there is no maximum, then MaxLogSize should be set to 0.
		MaxNumberOfRecords     int                 `xml:"MaxNumberOfRecords"`     // Maximum number of records that can be captured in the Log. If undefined, a value of zero should be specified.
		MaxRecordSize          int                 `xml:"MaxRecordSize"`          // Maximum size, in bytes, to which an individual Log entry (record) can grow - if the Capabilities array includes a value of 7 ("Variable Length Records Supported"). If the Capabilities array does not include a 7, then the Log only supports fixed length entries. The size of these entries is described by this property.
		Name                   string              `xml:"Name"`                   // The inherited Name serves as part of the key (a unique identifier) for the MessageLog instance.
		OperationalStatus      []OperationalStatus `xml:"OperationalStatus"`      // Indicates the current statuses of the element.
		OverwritePolicy        OverwritePolicy     `xml:"OverwritePolicy"`        // An enumeration describing the behavior of the Log, when it becomes full or near full.
		PercentageNearFull     int                 `xml:"PercentageNearFull"`     // If the OverwritePolicy is based on clearing records when the Log is near full (value=3), this property defines the record capacity (in percentage) that is considered to be 'near full'.
		RequestedState         RequestedState      `xml:"RequestedState"`         // RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
		SizeOfHeader           int                 `xml:"SizeOfHeader"`           // The size of the Log header, in bytes, if one is present. If there is no Log header, then this property should be set to 0. Headers may include general information about the Log such as the current number of records, time of last update, or a pointer to the location of the first Log entry. Note that this property is NOT the size of the header for an individual Log entry. The latter is described by the property, SizeOfRecordHeader.
		SizeOfRecordHeader     int                 `xml:"SizeOfRecordHeader"`     // The size of the header for the Log's individual entries, in bytes, if record headers are defined. If there are no record headers, then this property should be set to 0. Record headers may include information such as the type of the Log entry, the date/time that the entry was last updated, or a pointer to the start of optional data. Note that this property defines the header size for individual records in the Log, while the SizeOfHeader property describes the Log's overall header, typically located at the start of the MessageLog.
		Status                 string              `xml:"Status"`                 // A string indicating the current status of the object. This property is deprecated in lieu of OperationalStatus, which includes the same semantics in its enumeration.
	}

	GetRecordsResponse struct {
		XMLName             xml.Name              `xml:"GetRecords_OUTPUT"`
		IterationIdentifier int                   `xml:"IterationIdentifier"` // An identifier for the iterator.
		NoMoreRecords       bool                  `xml:"NoMoreRecords"`       // Indicates that there are no more records to read
		RecordArray         []string              `xml:"RecordArray"`         // Array of records encoded as Base64
		RawEventData        []RawEventData        `xml:"RawEventData"`        // Slice of raw event data
		RefinedEventData    []RefinedEventData    `xml:"RefinedEventData"`    // Slice of refined event data
		ReturnValue         GetRecordsReturnValue `xml:"ReturnValue"`         // ValueMap={0, 1, 2, 3} Values={Completed with No Error, Not Supported, Invalid record pointed, No record exists in log}
	}

	PositionToFirstRecordResponse struct {
		XMLName             xml.Name                         `xml:"PositionToFirstRecord_OUTPUT"`
		IterationIdentifier int                              `xml:"IterationIdentifier"` // An identifier for the iterator.
		ReturnValue         PositionToFirstRecordReturnValue `xml:"ReturnValue"`         // ValueMap={0, 1, 2} Values={Completed with No Error, Not Supported, No record exists}
	}

	RawEventData struct {
		TimeStamp       uint32
		DeviceAddress   uint8
		EventSensorType uint8
		EventType       uint8
		EventOffset     uint8
		EventSourceType uint8
		EventSeverity   uint8
		SensorNumber    uint8
		Entity          uint8
		EntityInstance  uint8
		EventData       []uint8
	}

	RefinedEventData struct {
		TimeStamp       time.Time
		DeviceAddress   uint8
		Description     string
		Entity          string
		EntityInstance  uint8
		EventData       []uint8
		EventSensorType uint8
		EventType       uint8
		EventOffset     uint8
		EventSourceType uint8
		EventSeverity   string
		SensorNumber    uint8
	}

	// Capabilities is an array of integers indicating the Log capabilities.
	Capabilities int

	// CharacterSet is an enumeration describing the character set used to record data in the individual Log entries.
	CharacterSet int

	// EnabledDefault is an enumerated value indicating an administrator's default or startup configuration for the Enabled State of an element.
	EnabledDefault int

	// EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
	EnabledState int

	// HealthState indicates the current health of the element.
	HealthState int

	// LastChange is an enumeration describing the last change to the MessageLog.
	LastChange int

	// LogState is an integer enumeration that indicates the current state of a log represented by CIM_Log subclasses.
	LogState int

	// OperationalStatus indicates the current statuses of the element.
	OperationalStatus int

	// OverwritePolicy is an enumeration describing the behavior of the Log, when it becomes full or near full.
	OverwritePolicy int

	// RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
	RequestedState int

	// GetRecordsReturnValue is an integer indicating the return value of the GetRecords operation.
	GetRecordsReturnValue int

	// PositionToFirstRecordReturnValue is an integer indicating the return value of the PositionToFirstRecord operation.
	PositionToFirstRecordReturnValue int
)

// INPUTS.
type GetRecords_INPUT struct {
	XMLName             xml.Name `xml:"h:GetRecords_INPUT"`
	H                   string   `xml:"xmlns:h,attr"`
	IterationIdentifier int      `xml:"h:IterationIdentifier"` // An identifier for the iterator.
	MaxReadRecords      int      `xml:"h:MaxReadRecords"`      // Maximum number of records to read
}

const (
	// Intel AMT can return 400 records in a single GetRecords call, but we limit it to 390.
	MaxAMTRecords = 390
	// DefaultRecords is the default number of records to retrieve when no valid count is specified.
	DefaultRecords = 120
)
