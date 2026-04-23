/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package cim implements CIM classes to support communicating with Intel® AMT Devices
package cim

import (
	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/bios"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/biosfeature"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/boot"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/card"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/chassis"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/chip"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/computer"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/concrete"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/credential"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/ethernetport"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/ieee8021x"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/kvm"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/mediaaccess"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/physical"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/power"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/powermanagementcapabilities"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/processor"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/redirectionservice"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/service"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/software"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/system"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/wifi"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/wifiendpoint"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/wifiendpointcapabilities"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/wifiportcapabilities"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

type Messages struct {
	wsmanMessageCreator         *message.WSManMessageCreator
	BIOSElement                 bios.Element
	BIOSFeature                 biosfeature.Feature
	BootConfigSetting           boot.ConfigSetting
	BootService                 boot.Service
	BootSourceSetting           boot.SourceSetting
	Card                        card.Package
	Chassis                     chassis.Package
	Chip                        chip.Package
	ComputerSystemPackage       computer.SystemPackage
	ConcreteDependency          concrete.Dependency
	CredentialContext           credential.Context
	EthernetPort                ethernetport.Port
	IEEE8021xSettings           ieee8021x.Settings
	KVMRedirectionSAP           kvm.RedirectionSAP
	MediaAccessDevice           mediaaccess.Device
	PhysicalMemory              physical.Memory
	PhysicalPackage             physical.Package
	PowerManagementCapabilities powermanagementcapabilities.Capabilities
	PowerManagementService      power.ManagementService
	Processor                   processor.Package
	RedirectionService          redirectionservice.Service
	ServiceAvailableToElement   service.AvailableToElement
	SoftwareIdentity            software.Identity
	SystemPackaging             system.Package
	WiFiEndpoint                wifiendpoint.Endpoint
	WiFiEndpointCapabilities    wifiendpointcapabilities.EndpointCapabilities
	WiFiEndpointSettings        wifi.EndpointSettings
	WiFiPort                    wifi.Port
	WiFiPortCapabilities        wifiportcapabilities.PortCapabilities
}

func NewMessages(client client.WSMan) Messages {
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	m := Messages{
		wsmanMessageCreator: wsmanMessageCreator,
	}
	m.BIOSElement = bios.NewBIOSElementWithClient(wsmanMessageCreator, client)
	m.BIOSFeature = biosfeature.NewBIOSFeatureWithClient(wsmanMessageCreator, client)
	m.BootConfigSetting = boot.NewBootConfigSettingWithClient(wsmanMessageCreator, client)
	m.BootService = boot.NewBootServiceWithClient(wsmanMessageCreator, client)
	m.BootSourceSetting = boot.NewBootSourceSettingWithClient(wsmanMessageCreator, client)
	m.Card = card.NewCardWithClient(wsmanMessageCreator, client)
	m.Chassis = chassis.NewChassisWithClient(wsmanMessageCreator, client)
	m.Chip = chip.NewChipWithClient(wsmanMessageCreator, client)
	m.ComputerSystemPackage = computer.NewComputerSystemPackageWithClient(wsmanMessageCreator, client)
	m.ConcreteDependency = concrete.NewDependencyWithClient(wsmanMessageCreator, client)
	m.CredentialContext = credential.NewContextWithClient(wsmanMessageCreator, client)
	m.EthernetPort = ethernetport.NewEthernetPortWithClient(wsmanMessageCreator, client)
	m.IEEE8021xSettings = ieee8021x.NewIEEE8021xSettingsWithClient(wsmanMessageCreator, client)
	m.KVMRedirectionSAP = kvm.NewKVMRedirectionSAPWithClient(wsmanMessageCreator, client)
	m.MediaAccessDevice = mediaaccess.NewMediaAccessDeviceWithClient(wsmanMessageCreator, client)
	m.PhysicalMemory = physical.NewPhysicalMemoryWithClient(wsmanMessageCreator, client)
	m.PhysicalPackage = physical.NewPhysicalPackageWithClient(wsmanMessageCreator, client)
	m.PowerManagementCapabilities = powermanagementcapabilities.NewPowerManagementCapabilitiesWithClient(wsmanMessageCreator, client)
	m.PowerManagementService = power.NewPowerManagementServiceWithClient(wsmanMessageCreator, client)
	m.Processor = processor.NewProcessorWithClient(wsmanMessageCreator, client)
	m.RedirectionService = redirectionservice.NewRedirectionServiceWithClient(wsmanMessageCreator, client)
	m.ServiceAvailableToElement = service.NewServiceAvailableToElementWithClient(wsmanMessageCreator, client)
	m.SoftwareIdentity = software.NewSoftwareIdentityWithClient(wsmanMessageCreator, client)
	m.SystemPackaging = system.NewSystemPackageWithClient(wsmanMessageCreator, client)
	m.WiFiEndpoint = wifiendpoint.NewWiFiEndpointWithClient(wsmanMessageCreator, client)
	m.WiFiEndpointCapabilities = wifiendpointcapabilities.NewWiFiEndpointCapabilitiesWithClient(wsmanMessageCreator, client)
	m.WiFiEndpointSettings = wifi.NewWiFiEndpointSettingsWithClient(wsmanMessageCreator, client)
	m.WiFiPort = wifi.NewWiFiPortWithClient(wsmanMessageCreator, client)
	m.WiFiPortCapabilities = wifiportcapabilities.NewWiFiPortCapabilitiesWithClient(wsmanMessageCreator, client)

	return m
}
