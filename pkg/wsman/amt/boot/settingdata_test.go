/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

var boot_settings = BootSettingDataRequest{
	BIOSPause:              false,
	BIOSSetup:              false,
	BootMediaIndex:         0,
	ConfigurationDataReset: false,
	ElementName:            "Intel(r) AMT Boot Configuration Settings",
	EnforceSecureBoot:      false,
	FirmwareVerbosity:      0,
	ForcedProgressEvents:   false,
	IDERBootDevice:         0,
	InstanceID:             "Intel(r) AMT:BootSettingData 0",
	LockKeyboard:           false,
	LockPowerButton:        false,
	LockResetButton:        false,
	LockSleepButton:        false,
	OwningEntity:           "Intel(r) AMT",
	PlatformErase:          false,
	RSEPassword:            "",
	ReflashBIOS:            false,
	SecureErase:            false,
	UefiBootNumberOfParams: 0,
	UseIDER:                false,
	UseSOL:                 false,
	UseSafeMode:            false,
	UserPasswordBypass:     false,
}

func TestPositiveAMT_BootSettingData(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)

	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/boot/settingdata",
	}

	elementUnderTest := NewBootSettingDataWithClient(wsmanMessageCreator, &client)

	t.Run("amt_BootSettingData Tests", func(t *testing.T) {
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
				"should create a valid AMT_BootSettingData Get wsman message",
				AMTBootSettingData,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					BootSettingDataGetResponse: BootSettingDataResponse{
						XMLName:                  xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_BootSettingData", Local: "AMT_BootSettingData"},
						BIOSLastStatus:           []uint16{2, 0},
						BIOSSetup:                false,
						BIOSPause:                false,
						BootMediaIndex:           0,
						BootguardStatus:          127,
						ConfigurationDataReset:   false,
						ElementName:              "Intel(r) AMT Boot Configuration Settings",
						EnforceSecureBoot:        false,
						FirmwareVerbosity:        0,
						ForcedProgressEvents:     false,
						IDERBootDevice:           0,
						InstanceID:               "Intel(r) AMT:BootSettingData 0",
						LockKeyboard:             false,
						LockPowerButton:          false,
						LockResetButton:          false,
						LockSleepButton:          false,
						OptionsCleared:           true,
						OwningEntity:             "Intel(r) AMT",
						PlatformErase:            false,
						RPEEnabled:               true,
						RSEPassword:              "",
						ReflashBIOS:              false,
						SecureBootControlEnabled: true,
						SecureErase:              false,
						UEFIHTTPSBootEnabled:     true,
						UEFILocalPBABootEnabled:  true,
						UefiBootNumberOfParams:   0,
						UseIDER:                  false,
						UseSOL:                   false,
						UseSafeMode:              false,
						UserPasswordBypass:       false,
						WinREBootEnabled:         true,
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_BootSettingData Enumerate wsman message",
				AMTBootSettingData,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "DD070000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_BootSettingData Pull wsman message",
				AMTBootSettingData,
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
						BootSettingDataItems: []BootSettingDataResponse{
							{
								XMLName:                  xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_BootSettingData", Local: "AMT_BootSettingData"},
								BIOSLastStatus:           []uint16{2, 0},
								BIOSSetup:                false,
								BIOSPause:                false,
								BootMediaIndex:           0,
								BootguardStatus:          127,
								ConfigurationDataReset:   false,
								ElementName:              "Intel(r) AMT Boot Configuration Settings",
								EnforceSecureBoot:        false,
								FirmwareVerbosity:        0,
								ForcedProgressEvents:     false,
								IDERBootDevice:           0,
								InstanceID:               "Intel(r) AMT:BootSettingData 0",
								LockKeyboard:             false,
								LockPowerButton:          false,
								LockResetButton:          false,
								LockSleepButton:          false,
								OptionsCleared:           true,
								OwningEntity:             "Intel(r) AMT",
								PlatformErase:            false,
								RPEEnabled:               true,
								RSEPassword:              "",
								ReflashBIOS:              false,
								SecureBootControlEnabled: true,
								SecureErase:              false,
								UEFIHTTPSBootEnabled:     true,
								UEFILocalPBABootEnabled:  true,
								UefiBootNumberOfParams:   0,
								UseIDER:                  false,
								UseSOL:                   false,
								UseSafeMode:              false,
								UserPasswordBypass:       false,
								WinREBootEnabled:         true,
							},
						},
					},
				},
			},
			// PUT
			{
				"should create a valid AMT_BootSettingData Put wsman message",
				AMTBootSettingData,
				wsmantesting.Put,
				fmt.Sprintf(
					`<h:AMT_BootSettingData xmlns:h="%sAMT_BootSettingData"><h:BIOSPause>false</h:BIOSPause><h:BIOSSetup>false</h:BIOSSetup><h:BootMediaIndex>0</h:BootMediaIndex><h:ConfigurationDataReset>false</h:ConfigurationDataReset><h:ElementName>Intel(r) AMT Boot Configuration Settings</h:ElementName><h:EnforceSecureBoot>false</h:EnforceSecureBoot><h:FirmwareVerbosity>0</h:FirmwareVerbosity><h:ForcedProgressEvents>false</h:ForcedProgressEvents><h:IDERBootDevice>0</h:IDERBootDevice><h:InstanceID>Intel(r) AMT:BootSettingData 0</h:InstanceID><h:LockKeyboard>false</h:LockKeyboard><h:LockPowerButton>false</h:LockPowerButton><h:LockResetButton>false</h:LockResetButton><h:LockSleepButton>false</h:LockSleepButton><h:OwningEntity>Intel(r) AMT</h:OwningEntity><h:PlatformErase>false</h:PlatformErase><h:RSEPassword></h:RSEPassword><h:ReflashBIOS>false</h:ReflashBIOS><h:SecureErase>false</h:SecureErase><h:UefiBootParametersArray></h:UefiBootParametersArray><h:UefiBootNumberOfParams>0</h:UefiBootNumberOfParams><h:UseIDER>false</h:UseIDER><h:UseSOL>false</h:UseSOL><h:UseSafeMode>false</h:UseSafeMode><h:UserPasswordBypass>false</h:UserPasswordBypass></h:AMT_BootSettingData>`,
					"http://intel.com/wbem/wscim/1/amt-schema/1/"),
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessagePut

					return elementUnderTest.Put(boot_settings)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					BootSettingDataGetResponse: BootSettingDataResponse{
						XMLName:                  xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_BootSettingData", Local: "AMT_BootSettingData"},
						BIOSLastStatus:           []uint16{2, 0},
						BIOSSetup:                false,
						BIOSPause:                false,
						BootMediaIndex:           0,
						BootguardStatus:          127,
						ConfigurationDataReset:   false,
						ElementName:              "Intel(r) AMT Boot Configuration Settings",
						EnforceSecureBoot:        false,
						FirmwareVerbosity:        0,
						ForcedProgressEvents:     false,
						IDERBootDevice:           0,
						InstanceID:               "Intel(r) AMT:BootSettingData 0",
						LockKeyboard:             false,
						LockPowerButton:          false,
						LockResetButton:          false,
						LockSleepButton:          false,
						OptionsCleared:           true,
						OwningEntity:             "Intel(r) AMT",
						PlatformErase:            false,
						RPEEnabled:               false,
						RSEPassword:              "",
						ReflashBIOS:              false,
						SecureBootControlEnabled: false,
						SecureErase:              false,
						UEFIHTTPSBootEnabled:     false,
						UEFILocalPBABootEnabled:  false,
						UefiBootNumberOfParams:   0,
						UseIDER:                  false,
						UseSOL:                   false,
						UseSafeMode:              false,
						UserPasswordBypass:       false,
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

func TestNegativeAMT_BootSettingData(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/boot/settingdata",
	}
	elementUnderTest := NewBootSettingDataWithClient(wsmanMessageCreator, &client)

	t.Run("amt_BootSettingData Tests", func(t *testing.T) {
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
				"should create a valid AMT_BootSettingData Get wsman message",
				AMTBootSettingData,
				wsmantesting.Get,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					BootSettingDataGetResponse: BootSettingDataResponse{
						XMLName:                  xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_BootSettingData", Local: "AMT_BootSettingData"},
						InstanceID:               "t",
						ElementName:              "t",
						OwningEntity:             "t",
						UseSOL:                   false,
						UseSafeMode:              false,
						ReflashBIOS:              false,
						BIOSSetup:                false,
						BIOSPause:                false,
						LockPowerButton:          false,
						LockResetButton:          false,
						LockKeyboard:             false,
						LockSleepButton:          false,
						UserPasswordBypass:       false,
						ForcedProgressEvents:     false,
						FirmwareVerbosity:        0,
						ConfigurationDataReset:   false,
						IDERBootDevice:           1,
						UseIDER:                  false,
						EnforceSecureBoot:        false,
						BootMediaIndex:           0,
						SecureErase:              false,
						RSEPassword:              "",
						OptionsCleared:           true,
						WinREBootEnabled:         false,
						UEFILocalPBABootEnabled:  false,
						UEFIHTTPSBootEnabled:     false,
						SecureBootControlEnabled: false,
						BootguardStatus:          127,
						BIOSLastStatus:           []uint16{0, 0},
						UEFIBootParametersArray:  []uint8{0},
						UefiBootNumberOfParams:   0,
						RPEEnabled:               false,
						PlatformErase:            false,
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_BootSettingData Enumerate wsman message",
				AMTBootSettingData,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "5C000000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_BootSettingData Pull wsman message",
				AMTBootSettingData,
				wsmantesting.Pull,
				wsmantesting.PullBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						BootSettingDataItems: []BootSettingDataResponse{
							{
								XMLName:                  xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_BootSettingData", Local: "AMT_BootSettingData"},
								InstanceID:               "t",
								ElementName:              "t",
								OwningEntity:             "t",
								UseSOL:                   false,
								UseSafeMode:              false,
								ReflashBIOS:              false,
								BIOSSetup:                false,
								BIOSPause:                false,
								LockPowerButton:          false,
								LockResetButton:          false,
								LockKeyboard:             false,
								LockSleepButton:          false,
								UserPasswordBypass:       false,
								ForcedProgressEvents:     false,
								FirmwareVerbosity:        0,
								ConfigurationDataReset:   false,
								IDERBootDevice:           1,
								UseIDER:                  false,
								EnforceSecureBoot:        false,
								BootMediaIndex:           0,
								SecureErase:              false,
								RSEPassword:              "",
								OptionsCleared:           false,
								WinREBootEnabled:         false,
								UEFILocalPBABootEnabled:  false,
								UEFIHTTPSBootEnabled:     false,
								SecureBootControlEnabled: false,
								BootguardStatus:          127,
								BIOSLastStatus:           []uint16{0, 0},
								UEFIBootParametersArray:  []uint8{0},
								UefiBootNumberOfParams:   0,
								RPEEnabled:               false,
								PlatformErase:            false,
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
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.Body)
			})
		}
	})
}
