/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package cim

import (
	"reflect"
	"testing"

	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/bios"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/boot"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/card"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/chassis"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/chip"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/computer"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/concrete"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/credential"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/kvm"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/mediaaccess"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/physical"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/power"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/processor"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/service"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/software"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/system"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/cim/wifi"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/ieee8021x"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestNewMessages(t *testing.T) {
	mock := &wsmantesting.MockClient{}
	m := NewMessages(mock)

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

	if reflect.DeepEqual(m.Card, card.Package{}) {
		t.Error("Card is not initialized")
	}

	if reflect.DeepEqual(m.Chassis, chassis.Package{}) {
		t.Error("Chassis is not initialized")
	}

	if reflect.DeepEqual(m.Chip, chip.Package{}) {
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

	if reflect.DeepEqual(m.IEEE8021xSettings, ieee8021x.IEEE8021xSettingsRequest{}) {
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

	if reflect.DeepEqual(m.Processor, processor.Package{}) {
		t.Error("Processor is not initialized")
	}

	if reflect.DeepEqual(m.ServiceAvailableToElement, service.AvailableToElement{}) {
		t.Error("ServiceAvailableToElement is not initialized")
	}

	if reflect.DeepEqual(m.SoftwareIdentity, software.Identity{}) {
		t.Error("SoftwareIdentity is not initialized")
	}

	if reflect.DeepEqual(m.SystemPackaging, system.Package{}) {
		t.Error("SystemPackaging is not initialized")
	}

	if reflect.DeepEqual(m.WiFiEndpointSettings, wifi.EndpointSettings{}) {
		t.Error("WiFiEndpointSettings is not initialized")
	}

	if reflect.DeepEqual(m.WiFiPort, wifi.Port{}) {
		t.Error("WiFiPort is not initialized")
	}
}
