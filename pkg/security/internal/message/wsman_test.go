/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package message

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateXML(t *testing.T) {
	messageID := 0
	enumerationContext := "A4070000-0000-0000-0000-000000000000"
	wsmanMessageCreator := NewWSManMessageCreator("http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/")

	t.Run("creates an enumerate wsman string when provided a header and body to CreateXML", func(t *testing.T) {
		header := wsmanMessageCreator.CreateHeader(BaseActionsEnumerate, "CIM_ServiceAvailableToElement", nil, "", "")
		response := wsmanMessageCreator.CreateXML(header, EnumerateBody)
		correctResponse := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ServiceAvailableToElement</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><Enumerate xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration" /></Body></Envelope>`, messageID)
		messageID++

		assert.Equal(t, correctResponse, response)
	})

	t.Run("creates a pull wsman string when provided a header and body to createXML", func(t *testing.T) {
		header := wsmanMessageCreator.CreateHeader(BaseActionsPull, "CIM_ServiceAvailableToElement", nil, "", "")
		PullBody := fmt.Sprintf(`<Body><Pull xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration"><EnumerationContext>%s</EnumerationContext><MaxElements>999</MaxElements><MaxCharacters>99999</MaxCharacters></Pull></Body>`, enumerationContext)
		response := wsmanMessageCreator.CreateXML(header, PullBody)
		correctResponse := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ServiceAvailableToElement</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><Pull xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration"><EnumerationContext>%s</EnumerationContext><MaxElements>999</MaxElements><MaxCharacters>99999</MaxCharacters></Pull></Body></Envelope>`, messageID, enumerationContext)
		messageID++

		assert.Equal(t, correctResponse, response)
	})
}

func TestCreateHeader(t *testing.T) {
	messageID := 0
	selector := Selector{Name: "InstanceID", Value: "Intel(r) AMT Device 0"}
	wsmanMessageCreator := NewWSManMessageCreator("http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/")

	t.Run("creates a correct header with action, resourceUri, and messageId provided for createHeader", func(t *testing.T) {
		correctHeader := fmt.Sprintf(`<Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ServiceAvailableToElement</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header>`, messageID)
		header := wsmanMessageCreator.CreateHeader(BaseActionsEnumerate, "CIM_ServiceAvailableToElement", nil, "", "")
		messageID++

		assert.Equal(t, correctHeader, header)
	})

	t.Run("applies custom address correctly in createHeader", func(t *testing.T) {
		correctHeader := fmt.Sprintf(`<Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ServiceAvailableToElement</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>customAddress</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header>`, messageID)
		header := wsmanMessageCreator.CreateHeader(BaseActionsEnumerate, "CIM_ServiceAvailableToElement", nil, "customAddress", "")
		messageID++

		assert.Equal(t, correctHeader, header)
	})

	t.Run("applies custom timeout correctly in createHeader", func(t *testing.T) {
		correctHeader := fmt.Sprintf(`<Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ServiceAvailableToElement</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT30S</w:OperationTimeout></Header>`, messageID)
		header := wsmanMessageCreator.CreateHeader(BaseActionsEnumerate, "CIM_ServiceAvailableToElement", nil, "", "PT30S")
		messageID++

		assert.Equal(t, correctHeader, header)
	})

	t.Run("applies custom selector correctly in createHeader", func(t *testing.T) {
		correctHeader := fmt.Sprintf(`<Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate</a:Action><a:To>/wsman</a:To><w:ResourceURI>http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_ServiceAvailableToElement</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT30S</w:OperationTimeout><w:SelectorSet><w:Selector Name="InstanceID">Intel(r) AMT Device 0</w:Selector></w:SelectorSet></Header>`, messageID)
		header := wsmanMessageCreator.CreateHeader(BaseActionsEnumerate, "CIM_ServiceAvailableToElement", &selector, "", "PT30S")
		messageID++

		assert.Equal(t, correctHeader, header)
	})
}

type TestStruct struct {
	XMLName   xml.Name `xml:"h:testMethod"`
	H         string   `xml:"xmlns:h,attr"`
	TestXmlns string   `xml:"h:testXmlns"`
}

func TestCreateBody(t *testing.T) {
	wsmanMessageCreator := NewWSManMessageCreator("http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/")

	t.Run("should convert obj to XML with object test values", func(t *testing.T) {
		testData := TestStruct{
			TestXmlns: "test",
		}
		result := wsmanMessageCreator.CreateBody("testMethod", "testUri", &testData)
		expectedResult := `<Body><h:testMethod xmlns:h="http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/testUri"><h:testXmlns>test</h:testXmlns></h:testMethod></Body>`
		assert.Equal(t, expectedResult, result)
	})

	t.Run("should create body without data", func(t *testing.T) {
		result := wsmanMessageCreator.CreateBody("testMethod", "testUri", nil)
		expectedResult := "<Body><h:testMethod xmlns:h=\"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/testUri\"></h:testMethod></Body>"
		assert.Equal(t, expectedResult, result)
	})
}

func TestCreateSelectorObjectForBody(t *testing.T) {
	wsmanMessageCreator := NewWSManMessageCreator("http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/")

	t.Run("should return map string interface", func(t *testing.T) {
		selector := Selector{
			Name:  "InstanceID",
			Value: "Value",
		}
		result := wsmanMessageCreator.CreateSelectorObjectForBody(selector)
		expectedResult := map[string]interface{}{"Selector": []map[string]interface{}{{"$": map[string]string{"Name": "InstanceID"}, "_": "Value"}}}
		assert.Equal(t, expectedResult, result)
	})
}

func TestIsSlice(t *testing.T) {
	t.Run("should return true if input is slice", func(t *testing.T) {
		slice := []string{"123", "456"}
		result := IsSlice(slice)
		assert.True(t, result)
	})

	t.Run("should return false if input is not slice", func(t *testing.T) {
		notSlice := 0
		result := IsSlice(notSlice)
		assert.False(t, result)
	})
}
