/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package general

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

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
		XMLName           xml.Name `xml:"Body"`
		GetResponse       GeneralSettingsResponse
		EnumerateResponse common.EnumerateResponse
		PullResponse      PullResponse
		PutResponse       PutResponse
	}

	PullResponse struct {
		XMLName              xml.Name                  `xml:"PullResponse"`
		GeneralSettingsItems []GeneralSettingsResponse `xml:"Items>AMT_GeneralSettings"`
	}
	GeneralSettingsResponse struct {
		XMLName                       xml.Name                 `xml:"AMT_GeneralSettings"`
		ElementName                   string                   `xml:"ElementName,omitempty"`                   // The user-friendly name for this instance of SettingData. In addition, the user-friendly name can be used as an index property for a search or query. (Note: The name does not have to be unique within a namespace.)
		InstanceID                    string                   `xml:"InstanceID,omitempty"`                    // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class. This is a read-only property.
		NetworkInterfaceEnabled       bool                     `xml:"NetworkInterfaceEnabled,omitempty"`       // Indicates whether the network interface is enabled. This is a read-only property.
		DigestRealm                   string                   `xml:"DigestRealm,omitempty"`                   // The Intel® AMT device Digest Authentication Realm parameter as defined by RFC 2617. This is a read-only property.
		IdleWakeTimeout               int                      `xml:"IdleWakeTimeout,omitempty"`               // Defines the minimum time value, in minutes, that Intel® AMT will be powered after waking up from a sleep power state, or after the host enters sleep or off state.This timer value will be reloaded whenever Intel® AMT is servicing requests. Note: this setting may not be applicable under some power package definitions. The minimum value for this property is 1, maximum is 65535
		HostName                      string                   `xml:"HostName,omitempty"`                      // Intel® AMT host setting. In Intel AMT Release 6.0 and later releases, maximum length is 63 characters. Starting from Intel CSME 18.0, the hostname can contain Unicode characters, where each character is encoded as an html entity number, for example U+003C is represented by the ASCII string &#x3c; or &#60;. Maximum length of the string remains 63 bytes when encoded in UTF-8.
		DomainName                    string                   `xml:"DomainName,omitempty"`                    // Intel® AMT domain name setting. In Intel AMT Release 6.0 and later releases, maximum length is 191 characters.
		PingResponseEnabled           bool                     `xml:"PingResponseEnabled,omitempty"`           // Indicates whether Intel® AMT should respond to ping Echo Request messages. Additional Notes: 'PingResponseEnabled' is a required field for the Put command.
		WsmanOnlyMode                 bool                     `xml:"WsmanOnlyMode,omitempty"`                 // Indicates whether Intel® AMT should block network interfaces other than WS-Management. By default AMT enables both WS-Management and legacy interfaces. If set to true, only WS-Management will be enabled. Additional Notes: 'WsmanOnlyMode' is a required field for the Put command.
		PreferredAddressFamily        PreferredAddressFamily   `xml:"PreferredAddressFamily,omitempty"`        // Preferred Address Family (IPv4/IPv6). Preferred Address Family (IPv4/IPv6) used for controlling outbound traffic such as events and user initiated traffic. For such traffic, the preferred addressing family will be attempted first, but other considerations also apply, depending on the traffic and the destination.
		DHCPv6ConfigurationTimeout    int                      `xml:"DHCPv6ConfigurationTimeout,omitempty"`    // Defines the Maximum Duration (DHCPv6 MRD for the Solicit Message) in seconds during which the Intel® ME FW tries to locate a DHCPv6 server. 0 - means try forever. The default value for this property is 0.
		DDNSUpdateEnabled             bool                     `xml:"DDNSUpdateEnabled,omitempty"`             // Defines whether the Dynamic DNS Update Client in FW is enabled or not. (The default value for this property is disabled)
		DDNSUpdateByDHCPServerEnabled bool                     `xml:"DDNSUpdateByDHCPServerEnabled,omitempty"` // If the DDNS Update client in FW is disabled then this property will define whether DDNS Update should be requested from the DHCP Server for the shared IPv4 address and shared FQDN. (The default value for this property is enabled)
		SharedFQDN                    bool                     `xml:"SharedFQDN,omitempty"`                    // Defines Whether the FQDN (HostName.DomainName) is shared with the Host or dedicated to ME. (The default value for this property is shared - TRUE).
		HostOSFQDN                    string                   `xml:"HostOSFQDN,omitempty"`                    // Intel® AMT host OS FQDN. This value of host FQDN is needed for the case that FW is set with a dedicated FQDN - this allows the SW to correlate the FW name with the Host name.
		DDNSTTL                       int                      `xml:"DDNSTTL,omitempty"`                       // Defines the Time To Live value (cachable time) of RRs registered by the FW DDNSUpdateClient. Units are seconds. (The default value for this property is 15 minutes). Maximum value is 2147483647 (2^31-1) - according to RFC2181
		AMTNetworkEnabled             AMTNetwork               `xml:"AMTNetworkEnabled,omitempty"`             // When set to Disabled, the AMT OOB network interfaces (LAN and WLAN) are disabled including AMT user initiated applications, Environment Detection and RMCPPing. Since OOB networking is disabled, there will not be an option to enable it back remotely.
		RmcpPingResponseEnabled       bool                     `xml:"RmcpPingResponseEnabled,omitempty"`       // Indicates whether Intel® AMT should respond to RMCP ping Echo Request messages.
		DDNSPeriodicUpdateInterval    int                      `xml:"DDNSPeriodicUpdateInterval,omitempty"`    // Defines the interval at which the FW DDNS Update client will send periodic updates for all the RRs registered by FW. Should be set according to corporate DNS scavenging policy. Units are minutes. Can be : either 0, or 20 and over. A value of 0 disables periodic update. (The default value for this property is 24 hours - 1440 minutes).
		PresenceNotificationInterval  int                      `xml:"PresenceNotificationInterval,omitempty"`  // Defines the interval at which the FW will send periodic WS-management events notifications (for the subscribed clients) whenever network settings are changed. Units are minutes. A value of 0 disables periodic events. The default value for this property is 0 (notifications are disabled). The minimal allowed value is 15 minutes.
		PrivacyLevel                  PrivacyLevel             `xml:"PrivacyLevel,omitempty"`                  // Defines the Privacy and Security Level setting. This is a read-only property.
		PowerSource                   PowerSource              `xml:"PowerSource,omitempty"`                   // The system current power source. This is a read-only property.
		ThunderboltDockEnabled        ThunderboltDock          `xml:"ThunderboltDockEnabled,omitempty"`        // When set to Disabled, a management console cannot communicate with Intel AMT via a Thunderbolt dock. Available in Release 15.0 and later releases.
		OemID                         int                      `xml:"OemID,omitempty"`                         // The OEM's vendor ID as listed in the Peripheral Component Interconnect Special Interest Group (PCI-SIG) list of member companies. Available in Release 16.1 and later releases.
		DHCPSyncRequiresHostname      DHCPSyncRequiresHostname `xml:"DHCPSyncRequiresHostname,omitempty"`      // When set to Enabled, the Intel AMT device will require the client to provide a hostname when requesting an IP address from a DHCP server. This setting is only applicable when DHCP is enabled. Values: 0=Disabled, 1=Enabled. Default: Disabled.
	}

	PutResponse struct{}
)

// INPUTS
// Request Types.
type (
	GeneralSettingsRequest struct {
		XMLName                       xml.Name               `xml:"h:AMT_GeneralSettings"`
		H                             string                 `xml:"xmlns:h,attr"`
		ElementName                   string                 `xml:"h:ElementName,omitempty"`                   // The user-friendly name for this instance of SettingData. In addition, the user-friendly name can be used as an index property for a search or query. (Note: The name does not have to be unique within a namespace.)
		InstanceID                    string                 `xml:"h:InstanceID,omitempty"`                    // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class. This is a read-only property.
		IdleWakeTimeout               int                    `xml:"h:IdleWakeTimeout,omitempty"`               // Defines the minimum time value, in minutes, that Intel® AMT will be powered after waking up from a sleep power state, or after the host enters sleep or off state.This timer value will be reloaded whenever Intel® AMT is servicing requests. Note: this setting may not be applicable under some power package definitions. The minimum value for this property is 1, maximum is 65535
		HostName                      string                 `xml:"h:HostName,omitempty"`                      // Intel® AMT host setting. In Intel AMT Release 6.0 and later releases, maximum length is 63 characters. Starting from Intel CSME 18.0, the hostname can contain Unicode characters, where each character is encoded as an html entity number, for example U+003C is represented by the ASCII string &#x3c; or &#60;. Maximum length of the string remains 63 bytes when encoded in UTF-8.
		DomainName                    string                 `xml:"h:DomainName,omitempty"`                    // Intel® AMT domain name setting. In Intel AMT Release 6.0 and later releases, maximum length is 191 characters.
		PingResponseEnabled           bool                   `xml:"h:PingResponseEnabled,omitempty"`           // Indicates whether Intel® AMT should respond to ping Echo Request messages. Additional Notes: 'PingResponseEnabled' is a required field for the Put command.
		WsmanOnlyMode                 bool                   `xml:"h:WsmanOnlyMode,omitempty"`                 // Indicates whether Intel® AMT should block network interfaces other than WS-Management. By default AMT enables both WS-Management and legacy interfaces. If set to true, only WS-Management will be enabled. Additional Notes: 'WsmanOnlyMode' is a required field for the Put command.
		PreferredAddressFamily        PreferredAddressFamily `xml:"h:PreferredAddressFamily,omitempty"`        // Preferred Address Family (IPv4/IPv6). Preferred Address Family (IPv4/IPv6) used for controlling outbound traffic such as events and user initiated traffic. For such traffic, the preferred addressing family will be attempted first, but other considerations also apply, depending on the traffic and the destination.
		DHCPv6ConfigurationTimeout    int                    `xml:"h:DHCPv6ConfigurationTimeout,omitempty"`    // Defines the Maximum Duration (DHCPv6 MRD for the Solicit Message) in seconds during which the Intel® ME FW tries to locate a DHCPv6 server. 0 - means try forever. The default value for this property is 0.
		DDNSUpdateEnabled             bool                   `xml:"h:DDNSUpdateEnabled,omitempty"`             // Defines whether the Dynamic DNS Update Client in FW is enabled or not. (The default value for this property is disabled)
		DDNSUpdateByDHCPServerEnabled bool                   `xml:"h:DDNSUpdateByDHCPServerEnabled,omitempty"` // If the DDNS Update client in FW is disabled then this property will define whether DDNS Update should be requested from the DHCP Server for the shared IPv4 address and shared FQDN. (The default value for this property is enabled)
		SharedFQDN                    bool                   `xml:"h:SharedFQDN,omitempty"`                    // Defines Whether the FQDN (HostName.DomainName) is shared with the Host or dedicated to ME. (The default value for this property is shared - TRUE).
		HostOSFQDN                    string                 `xml:"h:HostOSFQDN,omitempty"`                    // Intel® AMT host OS FQDN. This value of host FQDN is needed for the case that FW is set with a dedicated FQDN - this allows the SW to correlate the FW name with the Host name.
		DDNSTTL                       int                    `xml:"h:DDNSTTL,omitempty"`                       // Defines the Time To Live value (cachable time) of RRs registered by the FW DDNSUpdateClient. Units are seconds. (The default value for this property is 15 minutes). Maximum value is 2147483647 (2^31-1) - according to RFC2181
		AMTNetworkEnabled             AMTNetwork             `xml:"h:AMTNetworkEnabled,omitempty"`             // When set to Disabled, the AMT OOB network interfaces (LAN and WLAN) are disabled including AMT user initiated applications, Environment Detection and RMCPPing. Since OOB networking is disabled, there will not be an option to enable it back remotely.
		RmcpPingResponseEnabled       bool                   `xml:"h:RmcpPingResponseEnabled,omitempty"`       // Indicates whether Intel® AMT should respond to RMCP ping Echo Request messages.
		DDNSPeriodicUpdateInterval    int                    `xml:"h:DDNSPeriodicUpdateInterval,omitempty"`    // Defines the interval at which the FW DDNS Update client will send periodic updates for all the RRs registered by FW. Should be set according to corporate DNS scavenging policy. Units are minutes. Can be : either 0, or 20 and over. A value of 0 disables periodic update. (The default value for this property is 24 hours - 1440 minutes).
		PresenceNotificationInterval  int                    `xml:"h:PresenceNotificationInterval,omitempty"`  // Defines the interval at which the FW will send periodic WS-management events notifications (for the subscribed clients) whenever network settings are changed. Units are minutes. A value of 0 disables periodic events. The default value for this property is 0 (notifications are disabled). The minimal allowed value is 15 minutes.
		ThunderboltDockEnabled        ThunderboltDock        `xml:"h:ThunderboltDockEnabled,omitempty"`        // When set to Disabled, a management console cannot communicate with Intel AMT via a Thunderbolt dock. Available in Release 15.0 and later releases.
		OemID                         int                    `xml:"h:OemID,omitempty"`                         // The OEM's vendor ID as listed in the Peripheral Component Interconnect Special Interest Group (PCI-SIG) list of member companies. Available in Release 16.1 and later releases.
		DHCPSyncRequiresHostname      int                    `xml:"h:DHCPSyncRequiresHostname,omitempty"`      // When set to Enabled, the Intel AMT device will require the client to provide a hostname when requesting an IP address from a DHCP server. This setting is only applicable when DHCP is enabled. Values: 0=Disabled, 1=Enabled. Default: Disabled.
	}
)

// Preferred Address Family (IPv4/IPv6). Preferred Address Family (IPv4/IPv6) used for controlling outbound traffic such as events and user initiated traffic. For such traffic, the preferred addressing family will be attempted first, but other considerations also apply, depending on the traffic and the destination.
//
// ValueMap={0, 1, 2..}
//
// Values={IPv4, IPv6, Reserved}.
type PreferredAddressFamily int

// When set to Disabled, the AMT OOB network interfaces (LAN and WLAN) are disabled including AMT user initiated applications, Environment Detection and RMCPPing. Since OOB networking is disabled, there will not be an option to enable it back remotely.
//
// ValueMap={0, 1, 2..}
//
// Values={Disabled, Enabled, Reserved}.
type AMTNetwork int

// When set to Disabled, a management console cannot communicate with Intel AMT via a Thunderbolt dock. Available in Release 15.0 and later releases.
//
// Values: 0=Disabled, 1=Enabled. Default: Enabled.
type ThunderboltDock int

// Defines the Privacy and Security Level setting
//
// Default: SOL enabled = true, IDER enabled = true, KVM enabled = true, Opt-in can be disabled = true, opt-in configurable remotely = true. From Intel ME 8: Also Client Control Mode allowed=true and RCFG enabled=true.
//
// Enhanced: SOL enabled = true, IDER enabled = true, KVM enabled = true, Opt-in can be disabled = false, opt-in configurable remotely = true. From Intel ME 8: Also Client Control Mode allowed=true and RCFG enabled=true.
//
// Extreme: SOL enabled = false, IDER enabled = false, KVM enabled = false, Opt-in can be disabled = false, opt-in configurable remotely = false. From Intel ME 8: Also Client Control Mode allowed=false and RCFG enabled = false.
type PrivacyLevel int

// The system current power source.
type PowerSource int

// When set to Enabled, the Intel AMT device will require the client to provide a hostname when requesting an IP address from a DHCP server. This setting is only applicable when DHCP is enabled. Values: 0=Disabled, 1=Enabled. Default: Disabled.
type DHCPSyncRequiresHostname int
