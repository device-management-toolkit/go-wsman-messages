/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package models

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
)

type SettingData struct {
	XMLName xml.Name `xml:"CIM_SettingData"`
	ManagedElement
	InstanceID string `xml:"InstanceID,omitempty"` // MaxLen=256
}

type ManagedElement struct {
	XMLName     xml.Name `xml:"h:ManagedElement"`
	Caption     string   `xml:"h:Caption,omitempty"`     // MaxLen=64
	Description string   `xml:"h:Description,omitempty"` // MaxLen=256
	ElementName string   `xml:"h:ElementName,omitempty"` // MaxLen=256
}
type Collection struct {
	XMLName xml.Name `xml:"CIM_Collection"`
	ManagedElement
}

type Role struct {
	XMLName xml.Name `xml:"CIM_Role"`
	Collection
	CreationClassName   string `xml:"CreationClassName,omitempty"`
	Name                string `xml:"Name,omitempty"`
	CommonName          string `xml:"CommonName,omitempty"`
	RoleCharacteristics []int  `xml:"RoleCharacteristics,omitempty"`
}

type LogicalPort struct {
	LogicalDevice
	Speed            int    `xml:"Speed,omitempty"`
	MaxSpeed         int    `xml:"MaxSpeed,omitempty"`
	RequestedSpeed   int    `xml:"RequestedSpeed,omitempty"`
	UsageRestriction int    `xml:"UsageRestriction,omitempty"`
	PortType         int    `xml:"PortType,omitempty"`
	OtherPortType    string `xml:"OtherPortType,omitempty"`
}

type NetworkPort struct {
	LogicalPort
	PortNumber                       int
	LinkTechnology                   int      `xml:"LinkTechnology,omitempty"`
	OtherLinkTechnology              string   `xml:"OtherLinkTechnology,omitempty"`
	PermanentAddress                 string   `xml:"PermanentAddress,omitempty"`
	NetworkAddresses                 []string `xml:"NetworkAddresses,omitempty"`
	FullDuplex                       bool     `xml:"FullDuplex,omitempty"`
	AutoSense                        bool     `xml:"AutoSense,omitempty"`
	SupportedMaximumTransmissionUnit int      `xml:"SupportedMaximumTransmissionUnit,omitempty"`
	ActiveMaximumTransmissionUnit    int      `xml:"OtherActiveMaximumTransmissionUnitPortType,omitempty"`
}

type EthernetPort struct {
	NetworkPort
}

type WiFiPort struct {
	NetworkPort
}

type Antecedent struct {
	XMLName             xml.Name `xml:"Antecedent,omitempty"`
	Address             string   `xml:"Address,omitempty"`
	ReferenceParameters ReferenceParameters
}

type Dependent struct {
	XMLName             xml.Name `xml:"Dependent,omitempty"`
	Address             string   `xml:"Address,omitempty"`
	ReferenceParameters ReferenceParameters
}

type SystemPackage struct {
	Antecedent   Antecedent
	Dependent    Dependent
	PlatformGUID string `xml:"PlatformGUID,omitempty"`
}

type PhysicalElement struct {
	ManagedSystemElement
	Tag                  string   `xml:"Tag,omitempty"`                  // MaxLen=256
	CreationClassName    string   `xml:"CreationClassName,omitempty"`    // MaxLen=256
	Manufacturer         string   `xml:"Manufacturer,omitempty"`         // MaxLen=256
	Model                string   `xml:"Model,omitempty"`                // MaxLen=256
	Sku                  string   `xml:"Sku,omitempty"`                  // MaxLen=64
	SerialNumber         string   `xml:"SerialNumber,omitempty"`         // MaxLen=256
	Version              string   `xml:"Version,omitempty"`              // MaxLen=64
	PartNumber           string   `xml:"PartNumber,omitempty"`           // MaxLen=256
	OtherIdentifyingInfo string   `xml:"OtherIdentifyingInfo,omitempty"` // MaxLen=256
	PoweredOn            bool     `xml:"PoweredOn,omitempty"`
	ManufactureDate      DateTime `xml:"ManufactureDate,omitempty"`
	VendorEquipmentType  string   `xml:"VendorEquipmentType,omitempty"` // MaxLen=256
	UserTracking         string   `xml:"UserTracking,omitempty"`        // MaxLen=256
	CanBeFRUed           bool     `xml:"CanBeFRUed,omitempty"`
}

type PhysicalComponent struct {
	PhysicalElement
	RemovalConditions int  `xml:"RemovalConditions,omitempty"`
	Removable         bool `xml:"Removable,omitempty"`
	Replaceable       bool `xml:"Replaceable,omitempty"`
	HotSwappable      bool `xml:"HotSwappable,omitempty"`
}
type Chip struct {
	PhysicalComponent
}

type PhysicalPackage struct {
	PhysicalElement
	PackageType int `xml:"PackageType,omitempty"`
}

/**
 * Enabled:1 | Disabled:2 | Enabled For Debug:3.
 */
type (
	Enabled       int
	PhysicalFrame struct {
		PhysicalPackage
		VendorCompatibilityStrings []string `xml:"VendorCompatibilityStrings,omitempty"`
		OtherPackageType           string   `xml:"OtherPackageType,omitempty"`
		Weight                     int      `xml:"Weight,omitempty"`
		Width                      int      `xml:"Width,omitempty"`
		Depth                      int      `xml:"Depth,omitempty"`
		Height                     int      `xml:"Height,omitempty"`
		RemovalConditions          int      `xml:"RemovalConditions,omitempty"`
		Removable                  bool     `xml:"Removable,omitempty"`
		Replaceable                bool     `xml:"Replaceable,omitempty"`
		HotSwappable               bool     `xml:"HotSwappable,omitempty"`
		CableManagementStrategy    string   `xml:"CableManagementStrategy,omitempty"`
		ServicePhilosophy          int      `xml:"ServicePhilosophy,omitempty"`
		ServiceDescriptions        []string `xml:"ServiceDescriptions,omitempty"`
		LockPresent                bool     `xml:"LockPresent,omitempty"`
		AudibleAlarm               bool     `xml:"AudibleAlarm,omitempty"`
		VisibleAlarm               bool     `xml:"VisibleAlarm,omitempty"`
		SecurityBreach             int      `xml:"SecurityBreach,omitempty"`
		BreachDescription          string   `xml:"BreachDescription,omitempty"`
		IsLocked                   bool     `xml:"IsLocked,omitempty"`
	}
)

type Chassis struct {
	PhysicalFrame
	ChassisPackageType int `xml:"ChassisPackageType,omitempty"`
}

type LogicalElement struct {
	ManagedSystemElement
}

type SoftwareElement struct {
	LogicalElement
	Version               string                `xml:"Version,omitempty"`
	SoftwareElementState  int                   `xml:"SoftwareElementState,omitempty"`
	SoftwareElementId     string                `xml:"SoftwareElementId,omitempty"`
	TargetOperatingSystem TargetOperatingSystem `xml:"TargetOperatingSystem,omitempty"`
	OtherTargetOs         string                `xml:"OtherTargetOs,omitempty"`
	Manufacturer          string                `xml:"Manufacturer,omitempty"`
	BuildNumber           string                `xml:"BuildNumber,omitempty"`
	SerialNumber          string                `xml:"SerialNumber,omitempty"`
	CodeSet               string                `xml:"CodeSet,omitempty"`
	IdentificationCode    string                `xml:"IdentificationCode,omitempty"`
	LanguageEdition       string                `xml:"LanguageEdition,omitempty"`
}
type BIOSElement struct {
	SoftwareElement
	PrimaryBIOS bool `xml:"PrimaryBIOS,omitempty"`
	ReleaseDate DateTime
}
type BootSettingData struct {
	SettingData
	OwningEntity string `xml:"OwningEntity,omitempty"` // MaxLen=256
}

// SharedCredential represents a shared credential for a device.
type SharedCredential struct {
	InstanceID string `xml:"InstanceID,omitempty"`
	RemoteID   string `xml:"RemoteID,omitempty"`
	Secret     string `xml:"Secret,omitempty"`
	Algorithm  string `xml:"Algorithm,omitempty"`
	Protocol   string `xml:"Protocol,omitempty"`
}

// IEEE8021xSettings represents the IEEE 802.1x settings for a network interface.
type IEEE8021xSettings struct {
	XMLName                         xml.Name                        `xml:"h:IEEE8021xSettingsInput,omitempty"`
	H                               string                          `xml:"xmlns:q,attr"`
	ElementName                     string                          `xml:"q:ElementName,omitempty"`
	InstanceID                      string                          `xml:"q:InstanceID,omitempty"`
	AuthenticationProtocol          AuthenticationProtocol          `xml:"q:AuthenticationProtocol"`
	RoamingIdentity                 string                          `xml:"q:RoamingIdentity,omitempty"`
	ServerCertificateName           string                          `xml:"q:ServerCertificateName,omitempty"`
	ServerCertificateNameComparison ServerCertificateNameComparison `xml:"q:ServerCertificateNameComparison,omitempty"`
	Username                        string                          `xml:"q:Username,omitempty"`
	Password                        string                          `xml:"q:Password,omitempty"`
	Domain                          string                          `xml:"q:Domain,omitempty"`
	ProtectedAccessCredential       string                          `xml:"q:ProtectedAccessCredential,omitempty"`
	PACPassword                     string                          `xml:"q:PACPassword,omitempty"`
	PSK                             string                          `xml:"q:PSK,omitempty"`
}

type AuthenticationProtocol int

// BootSourceSetting represents the boot source settings for a device.
type BootSourceSetting struct {
	SettingData
	StructuredBootString string               `xml:"StructuredBootString,omitempty"`
	BIOSBootString       string               `xml:"BIOSBootString,omitempty"`
	BootString           string               `xml:"BootString,omitempty"`
	FailThroughSupported FailThroughSupported `xml:"FailThroughSupported,omitempty"`
}

type Service struct {
	EnabledLogicalElement
	SystemCreationClassName string `xml:"SystemCreationClassName,omitempty"`
	SystemName              string `xml:"SystemName,omitempty"`
	CreationClassName       string `xml:"CreationClassName,omitempty"`
	PrimaryOwnerName        string `xml:"PrimaryOwnerName,omitempty"`
	PrimaryOwnerContact     string `xml:"PrimaryOwnerContact,omitempty"`
	StartMode               string `xml:"StartMode,omitempty"`
	Started                 bool   `xml:"Started,omitempty"`
}

type SecurityService struct {
	Service
}

type AuthenticationService struct {
	SecurityService
}
type NetworkPortConfigurationService struct {
	Service
}
type CredentialManagementService struct {
	AuthenticationService
	// InstanceID is an optional property that may be used to opaquely and uniquely identify an instance of this class within the scope of the instantiating Namespace . . .
	InstanceID string // MaxLen=256
}

type EnabledLogicalElement struct {
	LogicalElement
	EnabledState          EnabledState   `xml:"EnabledState,omitempty"`
	OtherEnabledState     string         `xml:"OtherEnabledState,omitempty"`
	RequestedState        RequestedState `xml:"RequestedState,omitempty"`
	EnabledDefault        EnabledDefault `xml:"EnabledDefault,omitempty"`
	TimeOfLastStateChange DateTime
}
type LogicalDevice struct {
	EnabledLogicalElement       EnabledLogicalElement
	SystemCreationClassName     string                              `xml:"SystemCreationClassName,omitempty"`
	SystemName                  string                              `xml:"SystemName,omitempty"`
	CreationClassName           string                              `xml:"CreationClassName,omitempty"`
	DeviceId                    string                              `xml:"DeviceId,omitempty"`
	PowerManagementSupported    bool                                `xml:"PowerManagementSupported,omitempty"`
	PowerManagementCapabilities []PowerManagementCapabilitiesValues `xml:"PowerManagementCapabilities,omitempty"`
	Availability                AvailabilityValues                  `xml:"Availability,omitempty"`
	StatusInfo                  StatusInfoValues                    `xml:"StatusInfo,omitempty"`
	LastErrorCode               int                                 `xml:"LastErrorCode,omitempty"`
	ErrorDescription            string                              `xml:"ErrorDescription,omitempty"`
	ErrorCleared                bool                                `xml:"ErrorCleared,omitempty"`
	OtherIdentifyingInfo        []string                            `xml:"OtherIdentifyingInfo,omitempty"`
	PowerOnHours                int                                 `xml:"PowerOnHours,omitempty"`
	TotalPowerOnHours           int                                 `xml:"TotalPowerOnHours,omitempty"`
	IdentifyingDescriptions     []string                            `xml:"IdentifyingDescriptions,omitempty"`
	AdditionalAvailability      []AdditionalAvailabilityValues      `xml:"AdditionalAvailability,omitempty"`
	MaxQuiesceTime              int                                 `xml:"MaxQuiesceTime,omitempty"`
}

type Job struct {
	LogicalElement
	InstanceId          string              `xml:"InstanceId,omitempty"`
	CommunicationStatus CommunicationStatus `xml:"CommunicationStatus,omitempty"`
	DetailedStatus      DetailedStatus      `xml:"DetailedStatus,omitempty"`
	OperatingStatus     OperatingStatus     `xml:"OperatingStatus,omitempty"`
	PrimaryStatus       PrimaryStatus       `xml:"PrimaryStatus,omitempty"`
	JobStatus           string              `xml:"JobStatus,omitempty"`
	TimeSubmitted       DateTime            `xml:"TimeSubmitted,omitempty"`
	ScheduledStartTime  DateTime            `xml:"ScheduledStartTime,omitempty"`
	StartTime           DateTime            `xml:"StartTime,omitempty"`
	ElapsedTime         DateTime            `xml:"ElapsedTime,omitempty"`
	JobRunTimes         int                 `xml:"JobRunTimes,omitempty"`
	RunMonth            RunMonth            `xml:"RunMonth,omitempty"`
	RunDay              int                 `xml:"RunDay,omitempty"`
	RunDayOfWeek        RunDayOfWeek        `xml:"RunDayOfWeek,omitempty"`
	RunStartInterval    DateTime            `xml:"RunStartInterval,omitempty"`
	LocalOrUtcTime      LocalOrUtcTime      `xml:"LocalOrUtcTime,omitempty"`
	Notify              string              `xml:"Notify,omitempty"`
	Owner               string              `xml:"Owner,omitempty"`
	Priority            int                 `xml:"Priority,omitempty"`
	PercentComplete     int                 `xml:"PercentComplete,omitempty"`
	DeleteOnCompletion  bool                `xml:"DeleteOnCompletion,omitempty"`
	ErrorCode           int                 `xml:"ErrorCode,omitempty"`
	ErrorDescription    string              `xml:"ErrorDescription,omitempty"`
	RecoveryAction      RecoveryAction      `xml:"RecoveryAction,omitempty"`
	OtherRecoveryAction string              `xml:"OtherRecoveryAction,omitempty"`
}
type ConcreteJob struct {
	Job
	UntilTime             DateTime         `xml:"UntilTime,omitempty"`
	JobState              ConcreteJobState `xml:"JobState,omitempty"`
	TimeOfLastStateChange DateTime         `xml:"TimeOfLastStateChange,omitempty"`
	TimeBeforeRemoval     DateTime         `xml:"TimeBeforeRemoval,omitempty"`
}

type ConcreteJobState int

type CommunicationStatus int

type DetailedStatus int

type OperatingStatus int

type PrimaryStatus int

type RunMonth int

type RunDayOfWeek int

type RecoveryAction int

type LocalOrUtcTime int

type PowerManagementCapabilitiesValues int

type AvailabilityValues int

type StatusInfoValues int

type AdditionalAvailabilityValues int

type Credential struct {
	ManagedElement
	Issued  DateTime `xml:"Issued,omitempty"`  // The date and time when the credential was issued. Default is current time
	Expires DateTime `xml:"Expires,omitempty"` // The date and time when the credential expires (and is not appropriate for use for authentication/authorization). Default is '99991231235959.999999+999'
}

type CredentialContext struct {
	ElementInContext        Credential
	ElementProvidingContext ManagedElement
}

type ManagedSystemElement struct {
	ManagedElement
	InstallDate        DateTime          `xml:"InstallDate,omitempty"`
	Name               string            `xml:"Name,omitempty"`
	OperationalStatus  OperationalStatus `xml:"OperationalStatus,omitempty"`
	StatusDescriptions []string          `xml:"Items>StatusDescriptions,omitempty"`
	Status             string            `xml:"Status,omitempty"`
	HealthState        HealthState       `xml:"HealthState,omitempty"`
}

type DateTime struct {
	DateTime string `xml:"DateTime,omitempty"`
}

type ServiceAvailableToElement struct {
	ServiceProvided ServiceProvider
	UserOfService   UserOfService
}

type AssociationReference struct {
	Address             string `xml:"Address,omitempty"`
	ReferenceParameters ReferenceParametersNoNamespace
}

type ReferenceParametersNoNamespace struct {
	XMLName     xml.Name            `xml:"ReferenceParameters,omitempty"`
	ResourceURI string              `xml:"ResourceURI,omitempty"`
	SelectorSet SelectorNoNamespace `xml:"SelectorSet,omitempty"`
}
type SelectorNoNamespace struct {
	XMLName   xml.Name           `xml:"SelectorSet,omitempty"`
	Selectors []SelectorResponse `xml:"Selector,omitempty"`
}

type SelectorResponse struct {
	XMLName xml.Name `xml:"Selector,omitempty"`
	Name    string   `xml:"Name,attr"`
	Text    string   `xml:",chardata"`
}

type ServiceProvider struct {
	Address             string `xml:"Address,omitempty"`
	ReferenceParameters ReferenceParameters
}

type UserOfService struct {
	Address             string `xml:"Address,omitempty"`
	ReferenceParameters ReferenceParameters
}

type ReferenceParameters struct {
	XMLName     xml.Name    `xml:"ReferenceParameters"`
	H           string      `xml:"xmlns:c,attr"`
	ResourceURI string      `xml:"c:ResourceURI,omitempty"`
	SelectorSet SelectorSet `xml:"c:SelectorSet,omitempty"`
}

type ReferenceParameters_OUTPUT struct {
	ResourceURI string             `xml:"ResourceURI,omitempty"`
	SelectorSet SelectorSet_OUTPUT `xml:"SelectorSet,omitempty"`
}
type SelectorSet_OUTPUT struct {
	XMLName  xml.Name `xml:"SelectorSet,omitempty"`
	Selector []message.Selector_OUTPUT
}

type SelectorSet struct {
	H        string   `xml:"xmlns:c,attr"`
	XMLName  xml.Name `xml:"c:SelectorSet,omitempty"`
	Selector []Selector
}

type Selector struct {
	H       string   `xml:"xmlns:c,attr"`
	XMLName xml.Name `xml:"c:Selector,omitempty"`
	Name    string   `xml:"Name,attr"`
	Value   string   `xml:",chardata"`
}

type AssociatedPowerManagementService struct {
	ServiceAvailableToElement
	CIMAssociatedPowerManagementService CIMAssociatedPowerManagementServiceItem
}

type CIMAssociatedPowerManagementServiceItem struct {
	ServiceAvailableToElement
	AvailableRequestedPowerStates []string `xml:"AvailableRequestedPowerStates,omitempty"`
	PowerState                    string   `xml:"PowerState,omitempty"`
}

type SoftwareIdentity struct {
	LogicalElement
	CIMSoftwareIdentity []CIMSoftwareIdentityItem
}

type CIMSoftwareIdentityItem struct {
	LogicalElement
	InstanceID    string `xml:"InstanceID,omitempty"`
	VersionString string `xml:"VersionString,omitempty"`
	IsEntity      bool   `xml:"IsEntity,omitempty"`
}

type Log struct {
	EnabledLogicalElement
	MaxNumberOfRecords     int             `xml:"MaxNumberOfRecords,omitempty"`
	CurrentNumberOfRecords int             `xml:"CurrentNumberOfRecords,omitempty"`
	OverwritePolicy        OverwritePolicy `xml:"OverwritePolicy,omitempty"`
	LogState               LogState        `xml:"LogState,omitempty"`
}

type MessageLog struct {
	Log
	CreationClassName        string               `xml:"CreationClassName,omitempty"`
	Capabilities             []CapabilitiesValues `xml:"Capabilities,omitempty"`
	CapabilitiesDescriptions []string             `xml:"CapabilitiesDescriptions,omitempty"`
	MaxLogSize               int                  `xml:"MaxLogSize,omitempty"`
	SizeOfHeader             int                  `xml:"SizeOfHeader,omitempty"`
	HeaderFormat             string               `xml:"HeaderFormat,omitempty"`
	MaxRecordSize            int                  `xml:"MaxRecordSize,omitempty"`
	SizeOfRecordHeader       int                  `xml:"SizeOfRecordHeader,omitempty"`
	RecordHeaderFormat       string               `xml:"RecordHeaderFormat,omitempty"`
	OtherPolicyDescription   string               `xml:"OtherPolicyDescription,omitempty"`
	TimeWhenOutdated         DateTime             `xml:"TimeWhenOutdated,omitempty"`
	PercentageNearFull       int                  `xml:"PercentageNearFull,omitempty"`
	LastChange               LastChange           `xml:"LastChange,omitempty"`
	TimeOfLastChange         DateTime             `xml:"TimeOfLastChange,omitempty"`
	RecordLastChanged        int                  `xml:"RecordLastChanged,omitempty"`
	IsFrozen                 bool                 `xml:"IsFrozen,omitempty"`
	CharacterSet             CharacterSet         `xml:"CharacterSet,omitempty"`
}

type ServiceAccessPoint struct {
	EnabledLogicalElement   EnabledLogicalElement
	SystemCreationClassName string `xml:"SystemCreationClassName,omitempty"`
	SystemName              string `xml:"SystemName,omitempty"`
	CreationClassName       string `xml:"CreationClassName,omitempty"`
}

type (
	OverwritePolicy             int
	LogState                    int
	CapabilitiesValues          int
	LastChange                  int
	CharacterSet                int
	BootConfigSettingInstanceID string
	FailThroughSupported        int
	RemovalConditions           int
	PackageType                 int
	ServicePhilosophy           int
	SecurityBreach              int
	ChassisPackageType          int
	SoftwareElementState        int
	TargetOperatingSystem       int
)

// ServerCertificateNameComparison represents the ServerCertificateNameComparison type for IEEE8021xProfile.
type ServerCertificateNameComparison int

// PowerState represents the PowerState type in the PowerManagementService namespace.

type (
	UpgradeMethod     int
	CPUStatus         int
	OperationalStatus int
	HealthState       int
	EnabledState      int
	RequestedState    int
	EnabledDefault    int
)
