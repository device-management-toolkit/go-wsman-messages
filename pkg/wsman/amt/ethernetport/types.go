/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package ethernetport

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type Settings struct {
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
		XMLName                   xml.Name `xml:"Body"`
		GetAndPutResponse         SettingsResponse
		EnumerateResponse         common.EnumerateResponse
		PullResponse              PullResponse
		SetLinkPreferenceResponse SetLinkPreferenceResponse
	}

	PullResponse struct {
		XMLName           xml.Name           `xml:"PullResponse"`
		EthernetPortItems []SettingsResponse `xml:"Items>AMT_EthernetPortSettings"`
	}

	SettingsResponse struct {
		XMLName                      xml.Name                `xml:"AMT_EthernetPortSettings"`
		ElementName                  string                  `xml:"ElementName,omitempty"`                  // The user-friendly name for this instance of SettingData. In addition, the user-friendly name can be used as an index property for a search or query. (Note: The name does not have to be unique within a namespace.)
		InstanceID                   string                  `xml:"InstanceID,omitempty"`                   // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class.
		VLANTag                      int                     `xml:"VLANTag,omitempty"`                      // Indicates whether VLAN is in use and what is the VLAN tag when used.
		SharedMAC                    bool                    `xml:"SharedMAC,omitempty"`                    // Indicates whether Intel® AMT shares it's MAC address with the host system.
		MACAddress                   string                  `xml:"MACAddress,omitempty"`                   // The MAC address used by Intel® AMT in a string format. For Example: 01-02-3f-b0-99-99. (This property can only be read and can't be changed.)
		LinkIsUp                     bool                    `xml:"LinkIsUp,omitempty"`                     // Indicates whether the network link is up
		LinkPolicy                   []LinkPolicy            `xml:"LinkPolicy,omitempty"`                   // Enumeration values for link policy restrictions for better power consumption. If Intel® AMT will not be able to determine the exact power state, the more restrictive closest configuration applies.
		LinkPreference               LinkPreference          `xml:"LinkPreference,omitempty"`               // Determines whether the link is preferred to be owned by ME or host
		LinkControl                  LinkControl             `xml:"LinkControl,omitempty"`                  // Determines whether the link is owned by ME or host.  Additional Notes: This property is read-only.
		SharedStaticIp               bool                    `xml:"SharedStaticIp,omitempty"`               // Indicates whether the static host IP is shared with ME.
		SharedDynamicIP              bool                    `xml:"SharedDynamicIP,omitempty"`              // Indicates whether the dynamic host IP is shared with ME. This property is read only.
		IpSyncEnabled                bool                    `xml:"IpSyncEnabled,omitempty"`                // Indicates whether the IP synchronization between host and ME is enabled.
		DHCPEnabled                  bool                    `xml:"DHCPEnabled,omitempty"`                  // Indicates whether DHCP is in use. Additional Notes: 'DHCPEnabled' is a required field for the Put command.
		IPAddress                    string                  `xml:"IPAddress,omitempty"`                    // String representation of IP address. Get operation - reports the acquired IP address (whether in static or DHCP mode). Put operation - sets the IP address (in static mode only).
		SubnetMask                   string                  `xml:"SubnetMask,omitempty"`                   // Subnet mask in a string format.For example: 255.255.0.0
		DefaultGateway               string                  `xml:"DefaultGateway,omitempty"`               // Default Gateway in a string format. For example: 10.12.232.1
		PrimaryDNS                   string                  `xml:"PrimaryDNS,omitempty"`                   // Primary DNS in a string format. For example: 10.12.232.1
		SecondaryDNS                 string                  `xml:"SecondaryDNS,omitempty"`                 // Secondary DNS in a string format. For example: 10.12.232.1
		ConsoleTcpMaxRetransmissions int                     `xml:"ConsoleTcpMaxRetransmissions,omitempty"` // Indicates the number of retransmissions host TCP SW tries ifno ack is accepted
		WLANLinkProtectionLevel      WLANLinkProtectionLevel `xml:"WLANLinkProtectionLevel,omitempty"`      // Defines the level of the link protection feature activation. Read only property.
		PhysicalConnectionType       PhysicalConnectionType  `xml:"PhysicalConnectionType,omitempty"`       // Indicates the physical connection type of this network interface. Note: Applicable in Intel AMT 15.0 and later.
		PhysicalNicMedium            PhysicalNicMedium       `xml:"PhysicalNicMedium,omitempty"`            // Indicates which medium is currently used by Intel® AMT to communicate with the NIC. Note: Applicable in Intel AMT 15.0 and later.
	}
)

// INPUTS
// Request Types.
type SettingsRequest struct {
	XMLName                      xml.Name                     `xml:"h:AMT_EthernetPortSettings"`
	H                            string                       `xml:"xmlns:h,attr"`
	ElementName                  string                       `xml:"h:ElementName,omitempty"`                  // The user-friendly name for this instance of SettingData. In addition, the user-friendly name can be used as an index property for a search or query. (Note: The name does not have to be unique within a namespace.)
	InstanceID                   string                       `xml:"h:InstanceID,omitempty"`                   // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class.
	VLANTag                      int                          `xml:"h:VLANTag,omitempty"`                      // Indicates whether VLAN is in use and what is the VLAN tag when used.
	SharedMAC                    bool                         `xml:"h:SharedMAC"`                              // Indicates whether Intel® AMT shares it's MAC address with the host system.
	LinkIsUp                     bool                         `xml:"h:LinkIsUp"`                               // Indicates whether the network link is up
	LinkPolicy                   []LinkPolicy                 `xml:"h:LinkPolicy,omitempty"`                   // Enumeration values for link policy restrictions for better power consumption. If Intel® AMT will not be able to determine the exact power state, the more restrictive closest configuration applies.
	LinkPreference               LinkPreference               `xml:"h:LinkPreference,omitempty"`               // Determines whether the link is preferred to be owned by ME or host
	SharedStaticIp               bool                         `xml:"h:SharedStaticIp"`                         // Indicates whether the static host IP is shared with ME.
	IpSyncEnabled                bool                         `xml:"h:IpSyncEnabled"`                          // Indicates whether the IP synchronization between host and ME is enabled.
	DHCPEnabled                  bool                         `xml:"h:DHCPEnabled"`                            // Indicates whether DHCP is in use. Additional Notes: 'DHCPEnabled' is a required field for the Put command.
	IPAddress                    string                       `xml:"h:IPAddress,omitempty"`                    // String representation of IP address. Get operation - reports the acquired IP address (whether in static or DHCP mode). Put operation - sets the IP address (in static mode only).
	SubnetMask                   string                       `xml:"h:SubnetMask,omitempty"`                   // Subnet mask in a string format.For example: 255.255.0.0
	DefaultGateway               string                       `xml:"h:DefaultGateway,omitempty"`               // Default Gateway in a string format. For example: 10.12.232.1
	PrimaryDNS                   string                       `xml:"h:PrimaryDNS,omitempty"`                   // Primary DNS in a string format. For example: 10.12.232.1
	SecondaryDNS                 string                       `xml:"h:SecondaryDNS,omitempty"`                 // Secondary DNS in a string format. For example: 10.12.232.1
	ConsoleTcpMaxRetransmissions ConsoleTCPMaxRetransmissions `xml:"h:ConsoleTcpMaxRetransmissions,omitempty"` // Indicates the number of retransmissions host TCP SW tries if no ack is accepted
	PhysicalConnectionType       PhysicalConnectionType       `xml:"h:PhysicalConnectionType,omitempty"`       // Indicates the physical connection type of this network interface. Note: Applicable in Intel AMT 15.0 and later.
	PhysicalNicMedium            PhysicalNicMedium            `xml:"h:PhysicalNicMedium,omitempty"`            // Indicates which medium is currently used by Intel® AMT to communicate with the NIC. Note: Applicable in Intel AMT 15.0 and later.
}

// Enumeration values for link policy restrictions for better power consumption. If Intel® AMT will not be able to determine the exact power state, the more restrictive closest configuration applies.
//
// ValueMap={1, 14, 16, 224}
//
// Values={available on S0 AC, available on Sx AC, available on S0 DC, available on Sx DC}.
type LinkPolicy int

// Indicates the number of retransmissions host TCP SW tries if no ack is accepted
//
// MinValue=5
//
// MaxValue=7.
type ConsoleTCPMaxRetransmissions int

// Determines whether the link is preferred to be owned by ME or host
//
// ValueMap={1, 2, 3..}
//
// Values={ME, HOST, Reserved}.
type LinkPreference int

// Determines whether the link is owned by ME or host.  Additional Notes: This property is read-only.
//
// ValueMap={1, 2, 3..}
//
// Values={ME, HOST, Reserved}.
type LinkControl int

// Defines the level of the link protection feature activation. Read only property.
//
// ValueMap={0, 1, 2, 3, 4..}
//
// Values={OVERRIDE, NONE, PASSIVE, HIGH, RESERVED}.
type WLANLinkProtectionLevel int

// Indicates the physical connection type of this network interface. Note: Applicable in Intel AMT 15.0 and later.
//
// ValueMap={"0", "1", "2", "3", "4.."}
//
// Values={"Integrated LAN NIC", "Discrete LAN NIC", "LAN via a Thunderbolt dock", "Wireless LAN", "Reserved"}.
type PhysicalConnectionType int

// Indicates which medium is currently used by Intel® AMT to communicate with the NIC. Note: Applicable in Intel AMT 15.0 and later.
//
// ValueMap={"0", "1", "2.."}
//
// Values={"SMBUS", "PCIe", "Reserved"}.
type PhysicalNicMedium int

// SetLinkPreferenceResponse is the response from SetLinkPreference method call.
type SetLinkPreferenceResponse struct {
	XMLName     xml.Name `xml:"SetLinkPreference_OUTPUT"`
	ReturnValue int      `xml:"ReturnValue"`
}

// SetLinkPreferenceRequest is the input for SetLinkPreference method call.
type SetLinkPreferenceRequest struct {
	XMLName        xml.Name `xml:"h:SetLinkPreference_INPUT"`
	H              string   `xml:"xmlns:h,attr"`
	LinkPreference int      `xml:"h:LinkPreference"`
	Timeout        int      `xml:"h:Timeout"`
}
