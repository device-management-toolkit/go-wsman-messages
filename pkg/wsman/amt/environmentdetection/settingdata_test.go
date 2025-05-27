/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package environmentdetection

import (
	"encoding/xml"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestJson(t *testing.T) {
	response := Response{
		Body: Body{
			GetAndPutResponse: EnvironmentDetectionSettingDataResponse{},
		},
	}
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"GetAndPutResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ElementName\":\"\",\"InstanceID\":\"\",\"DetectionAlgorithm\":0,\"DetectionStrings\":null,\"DetectionIPv6LocalPrefixes\":null},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"EnvironmentDetectionSettingDataItems\":null}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			GetAndPutResponse: EnvironmentDetectionSettingDataResponse{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\ngetandputresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    elementname: \"\"\n    instanceid: \"\"\n    detectionalgorithm: 0\n    detectionstrings: []\n    detectionipv6localprefixes: []\nenumerateresponse:\n    enumerationcontext: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    environmentdetectionsettingdataitems: []\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func TestPositiveAMT_EnvironmentDetectionSettingData(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/environmentdetection",
	}
	elementUnderTest := NewEnvironmentDetectionSettingDataWithClient(wsmanMessageCreator, &client)

	t.Run("amt_EnvironmentDetectionSettingData Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			extraHeader      string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// GETS
			{
				"should create a valid AMT_EnvironmentDetectionSettingData Get wsman message",
				AMTEnvironmentDetectionSettingData,
				wsmantesting.Get,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetAndPutResponse: EnvironmentDetectionSettingDataResponse{
						XMLName:            xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EnvironmentDetectionSettingData", Local: "AMT_EnvironmentDetectionSettingData"},
						DetectionStrings:   []string{"b332bb28-ef3a-43b0-b998-342285ac1e26.com", "test.com"},
						DetectionAlgorithm: 0,
						ElementName:        "Intel(r) AMT Environment Detection Settings",
						InstanceID:         "Intel(r) AMT Environment Detection Settings",
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_EnvironmentDetectionSettingData Enumerate wsman message",
				AMTEnvironmentDetectionSettingData,
				wsmantesting.Enumerate,
				"",
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate
					if elementUnderTest.base.WSManMessageCreator == nil {
						logrus.Println("Error")
					}

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "61000000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_EnvironmentDetectionSettingData Pull wsman message",
				AMTEnvironmentDetectionSettingData,
				wsmantesting.Pull,
				"",
				wsmantesting.PullBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessagePull

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						EnvironmentDetectionSettingDataItems: []EnvironmentDetectionSettingDataResponse{
							{
								XMLName:            xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EnvironmentDetectionSettingData", Local: "AMT_EnvironmentDetectionSettingData"},
								DetectionAlgorithm: 0,
								DetectionStrings:   []string{"00d032fb-4341-42a5-a353-aaf83ff9d410.com"},
								ElementName:        "Intel(r) AMT Environment Detection Settings",
								InstanceID:         "Intel(r) AMT Environment Detection Settings",
							},
						},
					},
				},
			},
			// PUT
			{
				"should create a valid AMT_EnvironmentDetectionSettingData Put wsman message",
				AMTEnvironmentDetectionSettingData,
				wsmantesting.Put,
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT Environment Detection Settings</w:Selector></w:SelectorSet>",
				"<h:AMT_EnvironmentDetectionSettingData xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EnvironmentDetectionSettingData\"><h:ElementName>Intel(r) AMT Environment Detection Settings</h:ElementName><h:InstanceID>Intel(r) AMT Environment Detection Settings</h:InstanceID><h:DetectionAlgorithm>0</h:DetectionAlgorithm><h:DetectionStrings>2b14eacc-7f20-4a11-99bc-fdc6a162160b.com</h:DetectionStrings></h:AMT_EnvironmentDetectionSettingData>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessagePut
					edsd := EnvironmentDetectionSettingDataRequest{
						H:                  "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EnvironmentDetectionSettingData",
						ElementName:        "Intel(r) AMT Environment Detection Settings",
						InstanceID:         "Intel(r) AMT Environment Detection Settings",
						DetectionAlgorithm: 0,
						DetectionStrings:   []string{"2b14eacc-7f20-4a11-99bc-fdc6a162160b.com"},
					}

					return elementUnderTest.Put(edsd)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetAndPutResponse: EnvironmentDetectionSettingDataResponse{
						XMLName:            xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EnvironmentDetectionSettingData", Local: "AMT_EnvironmentDetectionSettingData"},
						DetectionStrings:   []string{"2b14eacc-7f20-4a11-99bc-fdc6a162160b.com"},
						DetectionAlgorithm: 0,
						ElementName:        "Intel(r) AMT Environment Detection Settings",
						InstanceID:         "Intel(r) AMT Environment Detection Settings",
					},
				},
			},
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceURIBase, test.method, test.action, test.extraHeader, test.body)
				messageID++
				response, err := test.responseFunc()
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.Body)
			})
		}
	})
}

func TestNegativeAMT_EnvironmentDetectionSettingData(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/environmentdetection",
	}
	elementUnderTest := NewEnvironmentDetectionSettingDataWithClient(wsmanMessageCreator, &client)

	t.Run("amt_EnvironmentDetectionSettingData Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			extraHeader      string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// GETS
			{
				"should create a valid AMT_EnvironmentDetectionSettingData Get wsman message",
				AMTEnvironmentDetectionSettingData,
				wsmantesting.Get,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetAndPutResponse: EnvironmentDetectionSettingDataResponse{
						XMLName:            xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EnvironmentDetectionSettingData", Local: "AMT_EnvironmentDetectionSettingData"},
						DetectionStrings:   []string{"b332bb28-ef3a-43b0-b998-342285ac1e26.com", "test.com"},
						DetectionAlgorithm: 0,
						ElementName:        "Intel(r) AMT Environment Detection Settings",
						InstanceID:         "Intel(r) AMT Environment Detection Settings",
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_EnvironmentDetectionSettingData Enumerate wsman message",
				AMTEnvironmentDetectionSettingData,
				wsmantesting.Enumerate,
				"",
				wsmantesting.EnumerateBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError
					if elementUnderTest.base.WSManMessageCreator == nil {
						logrus.Print("Error")
					}

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "61000000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_EnvironmentDetectionSettingData Pull wsman message",
				AMTEnvironmentDetectionSettingData,
				wsmantesting.Pull,
				"",
				wsmantesting.PullBody,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						EnvironmentDetectionSettingDataItems: []EnvironmentDetectionSettingDataResponse{
							{
								XMLName:            xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EnvironmentDetectionSettingData", Local: "AMT_EnvironmentDetectionSettingData"},
								DetectionAlgorithm: 0,
								DetectionStrings:   []string{"00d032fb-4341-42a5-a353-aaf83ff9d410.com"},
								ElementName:        "Intel(r) AMT Environment Detection Settings",
								InstanceID:         "Intel(r) AMT Environment Detection Settings",
							},
						},
					},
				},
			},
			// PUT
			{
				"should create a valid AMT_EnvironmentDetectionSettingData Put wsman message",
				AMTEnvironmentDetectionSettingData,
				wsmantesting.Put,
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT Environment Detection Settings</w:Selector></w:SelectorSet>",
				"<h:AMT_EnvironmentDetectionSettingData xmlns:h=\"http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EnvironmentDetectionSettingData\"><h:ElementName>Intel(r) AMT Environment Detection Settings</h:ElementName><h:InstanceID>Intel(r) AMT Environment Detection Settings</h:InstanceID><h:DetectionAlgorithm>0</h:DetectionAlgorithm><h:DetectionStrings>2b14eacc-7f20-4a11-99bc-fdc6a162160b.com</h:DetectionStrings></h:AMT_EnvironmentDetectionSettingData>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError
					edsd := EnvironmentDetectionSettingDataRequest{
						H:                  "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EnvironmentDetectionSettingData",
						ElementName:        "Intel(r) AMT Environment Detection Settings",
						InstanceID:         "Intel(r) AMT Environment Detection Settings",
						DetectionAlgorithm: 0,
						DetectionStrings:   []string{"2b14eacc-7f20-4a11-99bc-fdc6a162160b.com"},
					}

					return elementUnderTest.Put(edsd)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetAndPutResponse: EnvironmentDetectionSettingDataResponse{
						XMLName:            xml.Name{Space: "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_EnvironmentDetectionSettingData", Local: "AMT_EnvironmentDetectionSettingData"},
						DetectionStrings:   []string{"2b14eacc-7f20-4a11-99bc-fdc6a162160b.com"},
						DetectionAlgorithm: 0,
						ElementName:        "Intel(r) AMT Environment Detection Settings",
						InstanceID:         "Intel(r) AMT Environment Detection Settings",
					},
				},
			},
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceURIBase, test.method, test.action, test.extraHeader, test.body)
				messageID++
				response, err := test.responseFunc()
				assert.Error(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.Body)
			})
		}
	})
}
