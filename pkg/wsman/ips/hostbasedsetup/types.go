/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package hostbasedsetup

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

// OUTPUT
// Response Types.
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName                     xml.Name              `xml:"Body"`
		GetResponse                 HostBasedSetupService `xml:"IPS_HostBasedSetupService"`
		EnumerateResponse           common.EnumerateResponse
		PullResponse                PullResponse
		Setup_OUTPUT                Setup_OUTPUT                `xml:"Setup_OUTPUT"`
		AdminSetup_OUTPUT           AdminSetup_OUTPUT           `xml:"AdminSetup_OUTPUT"`
		AddNextCertInChain_OUTPUT   AddNextCertInChain_OUTPUT   `xml:"AddNextCertInChain_OUTPUT"`
		UpgradeClientToAdmin_OUTPUT UpgradeClientToAdmin_OUTPUT `xml:"UpgradeClientToAdmin_OUTPUT"`
	}
	PullResponse struct {
		XMLName                    xml.Name                `xml:"PullResponse"`
		HostBasedSetupServiceItems []HostBasedSetupService `xml:"Items>IPS_HostBasedSetupService"`
	}
	HostBasedSetupService struct {
		XMLName                 xml.Name `xml:"IPS_HostBasedSetupService"`
		ElementName             string
		SystemCreationClassName string
		SystemName              string
		CreationClassName       string
		Name                    string
		CurrentControlMode      CurrentControlMode
		AllowedControlModes     []AllowedControlModes
		ConfigurationNonce      string
		CertChainStatus         CertChainStatus
	}
	AddNextCertInChain_OUTPUT struct {
		XMLName     xml.Name `xml:"AddNextCertInChain_OUTPUT"`
		ReturnValue SetupReturnValue
	}

	AdminSetup_OUTPUT struct {
		XMLName     xml.Name `xml:"AdminSetup_OUTPUT"`
		ReturnValue SetupReturnValue
	}

	Setup_OUTPUT struct {
		XMLName     xml.Name `xml:"Setup_OUTPUT"`
		ReturnValue SetupReturnValue
	}
	UpgradeClientToAdmin_OUTPUT struct {
		XMLName     xml.Name `xml:"UpgradeClientToAdmin_OUTPUT"`
		ReturnValue SetupReturnValue
	}
)

// INPUT
// Request Types.
type (
	AddNextCertInChainInput struct {
		XMLName           xml.Name `xml:"h:AddNextCertInChain_INPUT"`
		H                 string   `xml:"xmlns:h,attr"`
		NextCertificate   string   `xml:"h:NextCertificate"`
		IsLeafCertificate bool     `xml:"h:IsLeafCertificate"`
		IsRootCertificate bool     `xml:"h:IsRootCertificate"`
	}
	AdminSetupInput struct {
		XMLName                    xml.Name `xml:"h:AdminSetup_INPUT"`
		H                          string   `xml:"xmlns:h,attr"`
		NetAdminPassEncryptionType int      `xml:"h:NetAdminPassEncryptionType"`
		NetworkAdminPassword       string   `xml:"h:NetworkAdminPassword"`
		McNonce                    string   `xml:"h:McNonce"`
		SigningAlgorithm           int      `xml:"h:SigningAlgorithm"`
		DigitalSignature           string   `xml:"h:DigitalSignature"`
	}
	SetupInput struct {
		XMLName                    xml.Name `xml:"h:Setup_INPUT"`
		H                          string   `xml:"xmlns:h,attr"`
		NetAdminPassEncryptionType int      `xml:"h:NetAdminPassEncryptionType"`
		NetworkAdminPassword       string   `xml:"h:NetworkAdminPassword"`
	}
	UpgradeClientToAdminInput struct {
		XMLName          xml.Name `xml:"h:UpgradeClientToAdmin_INPUT"`
		H                string   `xml:"xmlns:h,attr"`
		McNonce          string   `xml:"h:McNonce"`
		SigningAlgorithm int      `xml:"h:SigningAlgorithm"`
		DigitalSignature string   `xml:"h:DigitalSignature"`
	}
)

// AdminPassEncryptionType is the encryption type for the network admin password.
type AdminPassEncryptionType int

// SigningAlgorithm is the algorithm used to sign the setup operation.
type SigningAlgorithm int

// CurrentControlMode is an enumeration value that indicates the current control mode of the Intel(r) AMT subsystem after provisioning.
type CurrentControlMode int

// CertChainStatus is an enumeration value that indicates the status of "AddNextCertInChain" progress.
type CertChainStatus int

// AllowedControlModes is an array of values that indicates the allowed control modes for the Intel(r) AMT subsystem.
type AllowedControlModes int

// SetupReturnValue is an enumeration value that indicates the status of the operation.
type SetupReturnValue int

// AddNextCertInChainReturnValue is an enumeration value that indicates the status of the operation.
type AddNextCertInChainReturnValue int

// AdminSetupReturnValue is an enumeration value that indicates the status of the operation.
type AdminSetupReturnValue int
