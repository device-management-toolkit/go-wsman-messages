/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package cim

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/concrete"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/credential"
	"reflect"
	"testing"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/bios"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/boot"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/computer"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/kvm"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/mediaaccess"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/physical"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/power"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/service"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/software"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/system"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/wifi"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/ips/ieee8021x"
)

func TestNewMessages(t *testing.T) {
	m := NewMessages()

	if m.wsmanMessageCreator == nil {
		t.Error("wsmanMessageCreator is not initialized")
	}
	if reflect.DeepEqual(m.BIOSElement, bios.Element{}) {
		t.Error("BIOSElement is not initialized")
	}
	if reflect.DeepEqual(m.BootConfigSetting, boot.ConfigSetting{}) {
		t.Error("BootConfigSetting is not initialized")
	}
	if reflect.DeepEqual(m.BootService, boot.Service{}) {
		t.Error("BootService is not initialized")
	}
	if reflect.DeepEqual(m.BootSourceSetting, boot.SourceSetting{}) {
		t.Error("BootSourceSetting is not initialized")
	}
	if reflect.DeepEqual(m.Card, physical.Card{}) {
		t.Error("Card is not initialized")
	}
	if reflect.DeepEqual(m.Chassis, physical.Chassis{}) {
		t.Error("Chassis is not initialized")
	}
	if reflect.DeepEqual(m.Chip, physical.Chip{}) {
		t.Error("Chip is not initialized")
	}
	if reflect.DeepEqual(m.ComputerSystemPackage, computer.SystemPackage{}) {
		t.Error("ComputerSystemPackage is not initialized")
	}
	if reflect.DeepEqual(m.ConcreteDependency, concrete.Dependency{}) {
		t.Error("Dependency is not initialized")
	}
	if reflect.DeepEqual(m.CredentialContext, credential.Context{}) {
		t.Error("Context is not initialized")
	}
	if reflect.DeepEqual(m.IEEE8021xSettings, ieee8021x.IEEE8021xSettings{}) {
		t.Error("IEEE8021xSettings is not initialized")
	}
	if reflect.DeepEqual(m.KVMRedirectionSAP, kvm.RedirectionSAP{}) {
		t.Error("KVMRedirectionSAP is not initialized")
	}
	if reflect.DeepEqual(m.MediaAccessDevice, mediaaccess.Device{}) {
		t.Error("MediaAccessDevice is not initialized")
	}
	if reflect.DeepEqual(m.PhysicalMemory, physical.Memory{}) {
		t.Error("PhysicalMemory is not initialized")
	}
	if reflect.DeepEqual(m.PhysicalPackage, physical.Package{}) {
		t.Error("PhysicalPackage is not initialized")
	}
	if reflect.DeepEqual(m.PowerManagementService, power.ManagementService{}) {
		t.Error("PowerManagementService is not initialized")
	}
	if reflect.DeepEqual(m.Processor, physical.Processor{}) {
		t.Error("Processor is not initialized")
	}
	if reflect.DeepEqual(m.ServiceAvailableToElement, service.AvailableToElement{}) {
		t.Error("ServiceAvailableToElement is not initialized")
	}
	if reflect.DeepEqual(m.SoftwareIdentity, software.Identity{}) {
		t.Error("SoftwareIdentity is not initialized")
	}
	if reflect.DeepEqual(m.SystemPackaging, system.Packaging{}) {
		t.Error("SystemPackaging is not initialized")
	}
	if reflect.DeepEqual(m.WiFiEndpointSettings, wifi.EndpointSettings{}) {
		t.Error("WiFiEndpointSettings is not initialized")
	}
	if reflect.DeepEqual(m.WiFiPort, wifi.Port{}) {
		t.Error("WiFiPort is not initialized")
	}
}
