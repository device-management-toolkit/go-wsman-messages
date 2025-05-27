/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package publicprivate

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestJson(t *testing.T) {
	response := Response{
		Body: Body{
			GetResponse: PublicPrivateKeyPair{},
		},
	}
	expectedResult := "{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"GetResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"ElementName\":\"\",\"InstanceID\":\"\",\"DERKey\":\"\"},\"EnumerateResponse\":{\"EnumerationContext\":\"\"},\"PullResponse\":{\"XMLName\":{\"Space\":\"\",\"Local\":\"\"},\"PublicPrivateKeyPairItems\":null},\"RefinedPullResponse\":{\"PublicPrivateKeyPairItems\":null}}"
	result := response.JSON()
	assert.Equal(t, expectedResult, result)
}

func TestYaml(t *testing.T) {
	response := Response{
		Body: Body{
			GetResponse: PublicPrivateKeyPair{},
		},
	}
	expectedResult := "xmlname:\n    space: \"\"\n    local: \"\"\ngetresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    elementname: \"\"\n    instanceid: \"\"\n    derkey: \"\"\nenumerateresponse:\n    enumerationcontext: \"\"\npullresponse:\n    xmlname:\n        space: \"\"\n        local: \"\"\n    publicprivatekeypairitems: []\nrefinedpullresponse:\n    publicprivatekeypairitems: []\n"
	result := response.YAML()
	assert.Equal(t, expectedResult, result)
}

func testSetup() (messageID int, resourceURIBase string, client *wsmantesting.MockClient, elementUnderTest KeyPair) {
	messageID = 0
	resourceURIBase = message.AMTSchema
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client = &wsmantesting.MockClient{
		PackageUnderTest: "amt/publicprivate",
	}
	elementUnderTest = NewPublicPrivateKeyPairWithClient(wsmanMessageCreator, client)

	return messageID, resourceURIBase, client, elementUnderTest
}

func TestPositiveAMT_PublicPrivateKeyPair(t *testing.T) {
	messageID, resourceURIBase, client, elementUnderTest := testSetup()

	t.Run("amt_PublicPrivateKeyPair Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			extraHeader      string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// GETS
			{
				"should create a valid AMT_PublicPrivateKeyPair Get wsman message",
				AMTPublicPrivateKeyPair, wsmantesting.Get,
				"",
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT Key: Handle: 0</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get("Intel(r) AMT Key: Handle: 0")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: PublicPrivateKeyPair{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicPrivateKeyPair), Local: AMTPublicPrivateKeyPair},
						ElementName: "Intel(r) AMT Key",
						InstanceID:  "Intel(r) AMT Key: Handle: 0",
						DERKey:      "MIIBCgKCAQEA4y00wezZ1XwsSITMvqeYf61tgfVhlGbBVwq9Au0BaEgofPFCLuWMnKaTnMhUlJEGaeB2y6F8qjId0xMwLtNY6XWhmMoCP0R+ymgClT0treqtYp2zL1QPK1R04KTgF0KZh247oQpPGnB2nIe7PKCjPaY8BfOyBC6eNLeWUVIOA5TLL0gSTuk8y3iaadKo+LoWBaH/WDrIJ21Dzn6yU3zGueA8tphPH7yXaOJuNiijOUYZjVT7J0Ia8qMxUv1CrbfL2+N0lrcCG/E4f0QF1XgoCJnwIHdYaNhWzKVhfh2TTZIxJo8bXngckNOLzdYM35hUq98CxPiMSO8+G7J8RZaobQIDAQAB",
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_PublicPrivateKeyPair Enumerate wsman message",
				AMTPublicPrivateKeyPair,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "56080000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_PublicPrivateKeyPair Pull wsman message",
				AMTPublicPrivateKeyPair,
				wsmantesting.Pull,
				wsmantesting.PullBody,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessagePull

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: message.XMLPullResponseSpace, Local: "PullResponse"},
						PublicPrivateKeyPairItems: []PublicPrivateKeyPair{
							{
								XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicPrivateKeyPair), Local: AMTPublicPrivateKeyPair},
								ElementName: "Intel(r) AMT Key",
								InstanceID:  "Intel(r) AMT Key: Handle: 0",
								DERKey:      "MIIBCgKCAQEA4y00wezZ1XwsSITMvqeYf61tgfVhlGbBVwq9Au0BaEgofPFCLuWMnKaTnMhUlJEGaeB2y6F8qjId0xMwLtNY6XWhmMoCP0R+ymgClT0treqtYp2zL1QPK1R04KTgF0KZh247oQpPGnB2nIe7PKCjPaY8BfOyBC6eNLeWUVIOA5TLL0gSTuk8y3iaadKo+LoWBaH/WDrIJ21Dzn6yU3zGueA8tphPH7yXaOJuNiijOUYZjVT7J0Ia8qMxUv1CrbfL2+N0lrcCG/E4f0QF1XgoCJnwIHdYaNhWzKVhfh2TTZIxJo8bXngckNOLzdYM35hUq98CxPiMSO8+G7J8RZaobQIDAQAB",
							},
							{
								XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicPrivateKeyPair), Local: AMTPublicPrivateKeyPair},
								ElementName: "Intel(r) AMT Key",
								InstanceID:  "Intel(r) AMT Key: Handle: 1",
								DERKey:      "MIIBCgKCAQEAvMgYL2FyGuHOVvwYgjABqRlJ8j8LhMo2OCU1HU2WvDN3NoLmjAh2XmBS6ic5IjIc4VtjL7S8ImKP8+PSye9nxf+lv33AqcGsvQFcUuJ5gLTnYzrmqVk6XTcHf1qtvHEmVoykTV6bN7BQx0eTejTjhw3Ro6HZBMyStaTGIKjC9HLQySV6SnFGbrjdNZZoCYsaT8dVetn23npeses9f6dZT5K3IgpA13NcdJioS71uppjIcg8dXpcxA4QKgHLmmELPN9JLbywMvcCuU+xMDceWQlFld9ohmr8NiwgebLyVCh/Q+O+jkQT43snNolyTGLRWQFR4M6DT5fdgXivoFhzMcwIDAQAB",
							},
						},
					},
					RefinedPullResponse: RefinedPullResponse{
						PublicPrivateKeyPairItems: []RefinedPublicPrivateKeyPair{
							{
								ElementName:       "Intel(r) AMT Key",
								InstanceID:        "Intel(r) AMT Key: Handle: 0",
								DERKey:            "MIIBCgKCAQEA4y00wezZ1XwsSITMvqeYf61tgfVhlGbBVwq9Au0BaEgofPFCLuWMnKaTnMhUlJEGaeB2y6F8qjId0xMwLtNY6XWhmMoCP0R+ymgClT0treqtYp2zL1QPK1R04KTgF0KZh247oQpPGnB2nIe7PKCjPaY8BfOyBC6eNLeWUVIOA5TLL0gSTuk8y3iaadKo+LoWBaH/WDrIJ21Dzn6yU3zGueA8tphPH7yXaOJuNiijOUYZjVT7J0Ia8qMxUv1CrbfL2+N0lrcCG/E4f0QF1XgoCJnwIHdYaNhWzKVhfh2TTZIxJo8bXngckNOLzdYM35hUq98CxPiMSO8+G7J8RZaobQIDAQAB",
								CertificateHandle: "",
							},
							{
								ElementName:       "Intel(r) AMT Key",
								InstanceID:        "Intel(r) AMT Key: Handle: 1",
								DERKey:            "MIIBCgKCAQEAvMgYL2FyGuHOVvwYgjABqRlJ8j8LhMo2OCU1HU2WvDN3NoLmjAh2XmBS6ic5IjIc4VtjL7S8ImKP8+PSye9nxf+lv33AqcGsvQFcUuJ5gLTnYzrmqVk6XTcHf1qtvHEmVoykTV6bN7BQx0eTejTjhw3Ro6HZBMyStaTGIKjC9HLQySV6SnFGbrjdNZZoCYsaT8dVetn23npeses9f6dZT5K3IgpA13NcdJioS71uppjIcg8dXpcxA4QKgHLmmELPN9JLbywMvcCuU+xMDceWQlFld9ohmr8NiwgebLyVCh/Q+O+jkQT43snNolyTGLRWQFR4M6DT5fdgXivoFhzMcwIDAQAB",
								CertificateHandle: "",
							},
						},
					},
				},
			},
			// DELETE
			{
				"should create a valid AMT_PublicPrivateKeyPair Delete wsman message",
				AMTPublicPrivateKeyPair,
				wsmantesting.Delete,
				"",
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT Key: Handle: 0</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageDelete

					return elementUnderTest.Delete("Intel(r) AMT Key: Handle: 0")
				},
				Body{XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"}},
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

func TestNegativeAMT_PublicPrivateKeyPair(t *testing.T) {
	messageID, resourceURIBase, client, elementUnderTest := testSetup()

	t.Run("amt_PublicPrivateKeyPair Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			extraHeader      string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// GETS
			{
				"should create a valid AMT_PublicPrivateKeyPair Get wsman message",
				AMTPublicPrivateKeyPair, wsmantesting.Get,
				"",
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT Key: Handle: 0</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get("Intel(r) AMT Key: Handle: 0")
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					GetResponse: PublicPrivateKeyPair{
						XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicPrivateKeyPair), Local: AMTPublicPrivateKeyPair},
						ElementName: "Intel(r) AMT Key",
						InstanceID:  "Intel(r) AMT Key: Handle: 0",
						DERKey:      "MIIBCgKCAQEA4y00wezZ1XwsSITMvqeYf61tgfVhlGbBVwq9Au0BaEgofPFCLuWMnKaTnMhUlJEGaeB2y6F8qjId0xMwLtNY6XWhmMoCP0R+ymgClT0treqtYp2zL1QPK1R04KTgF0KZh247oQpPGnB2nIe7PKCjPaY8BfOyBC6eNLeWUVIOA5TLL0gSTuk8y3iaadKo+LoWBaH/WDrIJ21Dzn6yU3zGueA8tphPH7yXaOJuNiijOUYZjVT7J0Ia8qMxUv1CrbfL2+N0lrcCG/E4f0QF1XgoCJnwIHdYaNhWzKVhfh2TTZIxJo8bXngckNOLzdYM35hUq98CxPiMSO8+G7J8RZaobQIDAQAB",
					},
				},
			},
			// ENUMERATES
			{
				"should create a valid AMT_PublicPrivateKeyPair Enumerate wsman message",
				AMTPublicPrivateKeyPair,
				wsmantesting.Enumerate,
				wsmantesting.EnumerateBody,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "56080000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULLS
			{
				"should create a valid AMT_PublicPrivateKeyPair Pull wsman message",
				AMTPublicPrivateKeyPair,
				wsmantesting.Pull,
				wsmantesting.PullBody,
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: message.XMLPullResponseSpace, Local: "PullResponse"},
						PublicPrivateKeyPairItems: []PublicPrivateKeyPair{
							{
								XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicPrivateKeyPair), Local: AMTPublicPrivateKeyPair},
								ElementName: "Intel(r) AMT Key",
								InstanceID:  "Intel(r) AMT Key: Handle: 0",
								DERKey:      "MIIBCgKCAQEA4y00wezZ1XwsSITMvqeYf61tgfVhlGbBVwq9Au0BaEgofPFCLuWMnKaTnMhUlJEGaeB2y6F8qjId0xMwLtNY6XWhmMoCP0R+ymgClT0treqtYp2zL1QPK1R04KTgF0KZh247oQpPGnB2nIe7PKCjPaY8BfOyBC6eNLeWUVIOA5TLL0gSTuk8y3iaadKo+LoWBaH/WDrIJ21Dzn6yU3zGueA8tphPH7yXaOJuNiijOUYZjVT7J0Ia8qMxUv1CrbfL2+N0lrcCG/E4f0QF1XgoCJnwIHdYaNhWzKVhfh2TTZIxJo8bXngckNOLzdYM35hUq98CxPiMSO8+G7J8RZaobQIDAQAB",
							},
							{
								XMLName:     xml.Name{Space: fmt.Sprintf("%s%s", message.AMTSchema, AMTPublicPrivateKeyPair), Local: AMTPublicPrivateKeyPair},
								ElementName: "Intel(r) AMT Key",
								InstanceID:  "Intel(r) AMT Key: Handle: 1",
								DERKey:      "MIIBCgKCAQEAvMgYL2FyGuHOVvwYgjABqRlJ8j8LhMo2OCU1HU2WvDN3NoLmjAh2XmBS6ic5IjIc4VtjL7S8ImKP8+PSye9nxf+lv33AqcGsvQFcUuJ5gLTnYzrmqVk6XTcHf1qtvHEmVoykTV6bN7BQx0eTejTjhw3Ro6HZBMyStaTGIKjC9HLQySV6SnFGbrjdNZZoCYsaT8dVetn23npeses9f6dZT5K3IgpA13NcdJioS71uppjIcg8dXpcxA4QKgHLmmELPN9JLbywMvcCuU+xMDceWQlFld9ohmr8NiwgebLyVCh/Q+O+jkQT43snNolyTGLRWQFR4M6DT5fdgXivoFhzMcwIDAQAB",
							},
						},
					},
				},
			},
			// DELETE
			{
				"should create a valid AMT_PublicPrivateKeyPair Delete wsman message",
				AMTPublicPrivateKeyPair,
				wsmantesting.Delete,
				"",
				"<w:SelectorSet><w:Selector Name=\"InstanceID\">Intel(r) AMT Key: Handle: 0</w:Selector></w:SelectorSet>",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Delete("Intel(r) AMT Key: Handle: 0")
				},
				Body{XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"}},
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
