/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestJson(t *testing.T) {
	response := Response{
		Body: Body{
			BootCapabilitiesGetResponse: BootCapabilitiesResponse{},
		},
	}
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"BootSettingDataGetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"BIOSLastStatus\":null,\"BIOSPause\":false,\"BIOSSetup\":false,\"BootMediaIndex\":0,\"BootguardStatus\":0,\"ConfigurationDataReset\":false,\"ElementName\":\"\",\"EnforceSecureBoot\":false,\"FirmwareVerbosity\":0,\"ForcedProgressEvents\":false,\"IDERBootDevice\":0,\"InstanceID\":\"\",\"LockKeyboard\":false,\"LockPowerButton\":false,\"LockResetButton\":false,\"LockSleepButton\":false,\"OptionsCleared\":false,\"OwningEntity\":\"\",\"PlatformErase\":false,\"RPEEnabled\":false,\"RSEPassword\":\"\",\"ReflashBIOS\":false,\"SecureBootControlEnabled\":false,\"SecureErase\":false,\"UEFIHTTPSBootEnabled\":false,\"UEFIBootParametersArray\":null,\"UEFILocalPBABootEnabled\":false,\"UefiBootNumberOfParams\":0,\"UseIDER\":false,\"UseSOL\":false,\"UseSafeMode\":false,\"UserPasswordBypass\":false,\"WinREBootEnabled\":false},\"BootCapabilitiesGetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"InstanceID\":\"\",\"ElementName\":\"\",\"IDER\":false,\"SOL\":false,\"BIOSReflash\":false,\"BIOSSetup\":false,\"BIOSPause\":false,\"ForcePXEBoot\":false,\"ForceHardDriveBoot\":false,\"ForceHardDriveSafeModeBoot\":false,\"ForceDiagnosticBoot\":false,\"ForceCDorDVDBoot\":false,\"VerbosityScreenBlank\":false,\"PowerButtonLock\":false,\"ResetButtonLock\":false,\"KeyboardLock\":false,\"SleepButtonLock\":false,\"UserPasswordBypass\":false,\"ForcedProgressEvents\":false,\"VerbosityVerbose\":false,\"VerbosityQuiet\":false,\"ConfigurationDataReset\":false,\"BIOSSecureBoot\":false,\"SecureErase\":false,\"ForceWinREBoot\":false,\"ForceUEFILocalPBABoot\":false,\"ForceUEFIHTTPSBoot\":false,\"AMTSecureBootControl\":false,\"UEFIWiFiCoExistenceAndProfileShare\":false,\"PlatformErase\":0},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"BootSettingDataItems\":null,\"BootCapabilitiesItems\":null}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			BootCapabilitiesGetResponse: BootCapabilitiesResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\nbootsettingdatagetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    bioslaststatus: []\n    biospause: false\n    biossetup: false\n    bootmediaindex: 0\n    bootguardstatus: 0\n    configurationdatareset: false\n    elementname: \"\"\n    enforcesecureboot: false\n    firmwareverbosity: 0\n    forcedprogressevents: false\n    iderbootdevice: 0\n    instanceid: \"\"\n    lockkeyboard: false\n    lockpowerbutton: false\n    lockresetbutton: false\n    locksleepbutton: false\n    optionscleared: false\n    owningentity: \"\"\n    platformerase: false\n    rpeenabled: false\n    rsepassword: \"\"\n    reflashbios: false\n    securebootcontrolenabled: false\n    secureerase: false\n    uefihttpsbootenabled: false\n    uefibootparametersarray: []\n    uefilocalpbabootenabled: false\n    uefibootnumberofparams: 0\n    useider: false\n    usesol: false\n    usesafemode: false\n    userpasswordbypass: false\n    winrebootenabled: false\nbootcapabilitiesgetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    instanceid: \"\"\n    elementname: \"\"\n    ider: false\n    sol: false\n    biosreflash: false\n    biossetup: false\n    biospause: false\n    forcepxeboot: false\n    forceharddriveboot: false\n    forceharddrivesafemodeboot: false\n    forcediagnosticboot: false\n    forcecdordvdboot: false\n    verbosityscreenblank: false\n    powerbuttonlock: false\n    resetbuttonlock: false\n    keyboardlock: false\n    sleepbuttonlock: false\n    userpasswordbypass: false\n    forcedprogressevents: false\n    verbosityverbose: false\n    verbosityquiet: false\n    configurationdatareset: false\n    biossecureboot: false\n    secureerase: false\n    forcewinreboot: false\n    forceuefilocalpbaboot: false\n    forceuefihttpsboot: false\n    amtsecurebootcontrol: false\n    uefiwificoexistenceandprofileshare: false\n    platformerase: 0\nenumerateresponse:\n    enumerationcontext: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    bootsettingdataitems: []\n    bootcapabilitiesitems: []\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveAMT_BootCapabilities(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/boot/capabilities",
	}
	elementUnderTest := NewBootCapabilitiesWithClient(wsmanMessageCreator, &client)

	t.Run("amt_BootCapabilities Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// GETS
			{
				"should create a valid AMT_BootCapabilities Get wsman message",
				AMTBootCapabilities,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					BootCapabilitiesGetResponse: BootCapabilitiesResponse{
						XMLName:                    xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_BootCapabilities", Local: "AMT_BootCapabilities"},
						BIOSPause:                  false,
						BIOSReflash:                true,
						BIOSSecureBoot:             true,
						BIOSSetup:                  true,
						ConfigurationDataReset:     false,
						ElementName:                "Intel(r) AMT: Boot Capabilities",
						ForceCDorDVDBoot:           true,
						ForceDiagnosticBoot:        false,
						ForceHardDriveBoot:         true,
						ForceHardDriveSafeModeBoot: false,
						ForcePXEBoot:               true,
						ForcedProgressEvents:       true,
						IDER:                       true,
						InstanceID:                 "Intel(r) AMT:BootCapabilities 0",
						KeyboardLock:               true,
						PowerButtonLock:            false,
						ResetButtonLock:            false,
						SOL:                        true,
						SecureErase:                false,
						SleepButtonLock:            false,
						UserPasswordBypass:         true,
						VerbosityQuiet:             false,
						VerbosityScreenBlank:       false,
						VerbosityVerbose:           false,
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_BootCapabilities Enumerate wsman message",
				AMTBootCapabilities,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "E6070000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_BootCapabilities Pull wsman message",
				AMTBootCapabilities,
				wsmantesting.Pull,
				wsmantesting.PullBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessagePull

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						BootCapabilitiesItems: []BootCapabilitiesResponse{
							{
								XMLName:                    xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_BootCapabilities", Local: "AMT_BootCapabilities"},
								BIOSPause:                  false,
								BIOSReflash:                true,
								BIOSSecureBoot:             true,
								BIOSSetup:                  true,
								ConfigurationDataReset:     false,
								ElementName:                "Intel(r) AMT: Boot Capabilities",
								ForceCDorDVDBoot:           true,
								ForceDiagnosticBoot:        false,
								ForceHardDriveBoot:         true,
								ForceHardDriveSafeModeBoot: false,
								ForcePXEBoot:               true,
								ForcedProgressEvents:       true,
								IDER:                       true,
								InstanceID:                 "Intel(r) AMT:BootCapabilities 0",
								KeyboardLock:               true,
								PowerButtonLock:            false,
								ResetButtonLock:            false,
								SOL:                        true,
								SecureErase:                false,
								SleepButtonLock:            false,
								UserPasswordBypass:         true,
								VerbosityQuiet:             false,
								VerbosityScreenBlank:       false,
								VerbosityVerbose:           false,
							},
						},
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceURIBase, test.method, test.action, "", test.body)
				messageID++
				response, err := test.responseFunc()
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.Body)
			})
		}
	})
}

func TestNegativeAMT_BootCapabilities(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/boot/capabilities",
	}
	elementUnderTest := NewBootCapabilitiesWithClient(wsmanMessageCreator, &client)

	t.Run("amt_BootCapabilities Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// GETS
			{
				"should create a valid AMT_BootCapabilities Get wsman message",
				AMTBootCapabilities,
				wsmantesting.Get,
				"<error></error>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					BootCapabilitiesGetResponse: BootCapabilitiesResponse{
						XMLName:                            xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_BootCapabilities", Local: "AMT_BootCapabilities"},
						InstanceID:                         "t",
						ElementName:                        "t",
						IDER:                               false,
						SOL:                                false,
						BIOSReflash:                        false,
						BIOSSetup:                          false,
						BIOSPause:                          false,
						ForcePXEBoot:                       false,
						ForceHardDriveBoot:                 false,
						ForceHardDriveSafeModeBoot:         false,
						ForceDiagnosticBoot:                false,
						ForceCDorDVDBoot:                   false,
						VerbosityScreenBlank:               false,
						PowerButtonLock:                    false,
						ResetButtonLock:                    false,
						KeyboardLock:                       false,
						SleepButtonLock:                    false,
						UserPasswordBypass:                 false,
						ForcedProgressEvents:               false,
						VerbosityVerbose:                   false,
						VerbosityQuiet:                     false,
						ConfigurationDataReset:             false,
						BIOSSecureBoot:                     false,
						SecureErase:                        false,
						ForceWinREBoot:                     false,
						ForceUEFILocalPBABoot:              false,
						ForceUEFIHTTPSBoot:                 false,
						AMTSecureBootControl:               false,
						UEFIWiFiCoExistenceAndProfileShare: false,
						PlatformErase:                      0,
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_BootCapabilities Enumerate wsman message",
				AMTBootCapabilities,
				wsmantesting.Enumerate,
				"<error></error>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "error",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_BootCapabilities Pull wsman message",
				AMTBootCapabilities,
				wsmantesting.Pull,
				"<error></error>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						BootCapabilitiesItems: []BootCapabilitiesResponse{
							{
								XMLName:                            xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_BootCapabilities", Local: "AMT_BootCapabilities"},
								InstanceID:                         "t",
								ElementName:                        "t",
								IDER:                               false,
								SOL:                                false,
								BIOSReflash:                        false,
								BIOSSetup:                          false,
								BIOSPause:                          false,
								ForcePXEBoot:                       false,
								ForceHardDriveBoot:                 false,
								ForceHardDriveSafeModeBoot:         false,
								ForceDiagnosticBoot:                false,
								ForceCDorDVDBoot:                   false,
								VerbosityScreenBlank:               false,
								PowerButtonLock:                    false,
								ResetButtonLock:                    false,
								KeyboardLock:                       false,
								SleepButtonLock:                    false,
								UserPasswordBypass:                 false,
								ForcedProgressEvents:               false,
								VerbosityVerbose:                   false,
								VerbosityQuiet:                     false,
								ConfigurationDataReset:             false,
								BIOSSecureBoot:                     false,
								SecureErase:                        false,
								ForceWinREBoot:                     false,
								ForceUEFILocalPBABoot:              false,
								ForceUEFIHTTPSBoot:                 false,
								AMTSecureBootControl:               false,
								UEFIWiFiCoExistenceAndProfileShare: false,
								PlatformErase:                      0,
							},
						},
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceURIBase, test.method, test.action, "", test.body)
				messageID++
				response, err := test.responseFunc()
				assert.Error(t, err)
				assert.NotEqual(t, expectedXMLInput, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.Body)
			})
		}
	})
}
