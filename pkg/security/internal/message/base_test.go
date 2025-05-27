/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package message

import (
	"crypto/tls"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type MockClient struct {
	Err error
}

const (
	TestData    = "test-data"
	TestAction  = "test-action"
	TestContext = "test-context"
)

func (c *MockClient) Post(msg string) ([]byte, error) {
	var response []byte

	return response, c.Err
}
func (c *MockClient) Send(data []byte) error                          { return nil }
func (c *MockClient) Receive() ([]byte, error)                        { return nil, nil }
func (c *MockClient) CloseConnection() error                          { return nil }
func (c *MockClient) Connect() error                                  { return nil }
func (c *MockClient) IsAuthenticated() bool                           { return true }
func (c *MockClient) GetServerCertificate() (*tls.Certificate, error) { return nil, nil }
func TestBaseWithClient(t *testing.T) {
	mockWsmanMessageCreator := NewWSManMessageCreator("test-uri")
	mockClient := MockClient{}
	MessageID := 0
	base := NewBaseWithClient(mockWsmanMessageCreator, "TestClass", &mockClient)

	t.Run("Execute with no error", func(t *testing.T) {
		mockClient.Err = nil
		message := client.Message{
			XMLInput: "TestMessage",
		}
		err := base.Execute(&message)
		assert.NoError(t, err)
	})
	t.Run("Execute returns error", func(t *testing.T) {
		mockClient.Err = errors.New("test error")
		message := client.Message{
			XMLInput: "TestMessage",
		}
		err := base.Execute(&message)
		assert.Error(t, err)
	})
	t.Run("Enumerate", func(t *testing.T) {
		expected := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><Enumerate xmlns=\"http://schemas.xmlsoap.org/ws/2004/09/enumeration\" /></Body></Envelope>", MessageID)
		MessageID++
		actual := base.Enumerate()
		assert.Equal(t, expected, actual)
	})

	t.Run("Get", func(t *testing.T) {
		selector := &Selector{Name: "Key", Value: "Value"}
		expected := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Get</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name=\"Key\">Value</w:Selector></w:SelectorSet></Header><Body></Body></Envelope>", MessageID)
		MessageID++
		actual := base.Get(selector)
		assert.Equal(t, expected, actual)
	})

	t.Run("Pull", func(t *testing.T) {
		enumerationContext := TestContext
		expected := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><Pull xmlns=\"http://schemas.xmlsoap.org/ws/2004/09/enumeration\"><EnumerationContext>test-context</EnumerationContext><MaxElements>999</MaxElements><MaxCharacters>99999</MaxCharacters></Pull></Body></Envelope>", MessageID)
		MessageID++
		actual := base.Pull(enumerationContext)
		assert.Equal(t, expected, actual)
	})

	t.Run("Delete", func(t *testing.T) {
		expected := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name=\"Name\">Value</w:Selector></w:SelectorSet></Header><Body></Body></Envelope>", MessageID)
		MessageID++
		actual := base.Delete(Selector{Name: "Name", Value: "Value"})
		assert.Equal(t, expected, actual)
	})

	t.Run("Put", func(t *testing.T) {
		data := TestData
		customSelector := &Selector{Name: "Key", Value: "Value"}
		expected := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Put</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name=\"Key\">Value</w:Selector></w:SelectorSet></Header><Body><string>test-data</string></Body></Envelope>", MessageID)
		MessageID++
		actual := base.Put(data, true, customSelector)
		assert.Equal(t, expected, actual)

		expectedDefaultSelector := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Put</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name=\"InstanceID\">test-data</w:Selector></w:SelectorSet></Header><Body><string>test-data</string></Body></Envelope>", MessageID)
		MessageID++
		actualDefaultSelector := base.Put(data, true, nil)
		assert.Equal(t, expectedDefaultSelector, actualDefaultSelector)

		expectedNoSelector := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Put</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><string>test-data</string></Body></Envelope>", MessageID)
		MessageID++
		actualNoSelector := base.Put(data, false, customSelector)
		assert.Equal(t, expectedNoSelector, actualNoSelector)
	})

	t.Run("Create", func(t *testing.T) {
		data := TestData
		selector := &Selector{Name: "Key", Value: "Value"}
		expected := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Create</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name=\"Key\">Value</w:Selector></w:SelectorSet></Header><Body><string>test-data</string></Body></Envelope>", MessageID)
		MessageID++
		actual := base.Create(data, selector)
		assert.Equal(t, expected, actual)
	})

	t.Run("RequestStateChange", func(t *testing.T) {
		actionName := TestAction
		requestedState := 2
		expected := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>test-action</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><h:RequestStateChange_INPUT xmlns:h=\"test-uriTestClass\"><h:RequestedState>2</h:RequestedState></h:RequestStateChange_INPUT></Body></Envelope>", MessageID)
		MessageID++
		actual := base.RequestStateChange(actionName, requestedState)
		assert.Equal(t, expected, actual)
	})
}

func TestBase(t *testing.T) {
	mockWsmanMessageCreator := NewWSManMessageCreator("test-uri")
	base := NewBase(mockWsmanMessageCreator, "TestClass")
	MessageID := 0

	t.Run("Enumerate", func(t *testing.T) {
		expected := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><Enumerate xmlns=\"http://schemas.xmlsoap.org/ws/2004/09/enumeration\" /></Body></Envelope>", MessageID)
		MessageID++
		actual := base.Enumerate()
		assert.Equal(t, expected, actual)
	})

	t.Run("Get", func(t *testing.T) {
		selector := &Selector{Name: "Key", Value: "Value"}
		expected := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Get</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name=\"Key\">Value</w:Selector></w:SelectorSet></Header><Body></Body></Envelope>", MessageID)
		MessageID++
		actual := base.Get(selector)
		assert.Equal(t, expected, actual)
	})

	t.Run("Pull", func(t *testing.T) {
		enumerationContext := TestContext
		expected := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><Pull xmlns=\"http://schemas.xmlsoap.org/ws/2004/09/enumeration\"><EnumerationContext>test-context</EnumerationContext><MaxElements>999</MaxElements><MaxCharacters>99999</MaxCharacters></Pull></Body></Envelope>", MessageID)
		MessageID++
		actual := base.Pull(enumerationContext)
		assert.Equal(t, expected, actual)
	})

	t.Run("Delete", func(t *testing.T) {
		expected := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name=\"Name\">Value</w:Selector></w:SelectorSet></Header><Body></Body></Envelope>", MessageID)
		MessageID++
		actual := base.Delete(Selector{Name: "Name", Value: "Value"})
		assert.Equal(t, expected, actual)
	})

	t.Run("Put", func(t *testing.T) {
		data := TestData
		customSelector := &Selector{Name: "Key", Value: "Value"}
		expected := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Put</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name=\"Key\">Value</w:Selector></w:SelectorSet></Header><Body><string>test-data</string></Body></Envelope>", MessageID)
		MessageID++
		actual := base.Put(data, true, customSelector)
		assert.Equal(t, expected, actual)

		expectedDefaultSelector := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Put</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name=\"InstanceID\">test-data</w:Selector></w:SelectorSet></Header><Body><string>test-data</string></Body></Envelope>", MessageID)
		MessageID++
		actualDefaultSelector := base.Put(data, true, nil)
		assert.Equal(t, expectedDefaultSelector, actualDefaultSelector)

		expectedNoSelector := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Put</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><string>test-data</string></Body></Envelope>", MessageID)
		MessageID++
		actualNoSelector := base.Put(data, false, customSelector)
		assert.Equal(t, expectedNoSelector, actualNoSelector)
	})

	t.Run("Create", func(t *testing.T) {
		data := TestData
		selector := &Selector{Name: "Key", Value: "Value"}
		expected := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>http://schemas.xmlsoap.org/ws/2004/09/transfer/Create</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout><w:SelectorSet><w:Selector Name=\"Key\">Value</w:Selector></w:SelectorSet></Header><Body><string>test-data</string></Body></Envelope>", MessageID)
		MessageID++
		actual := base.Create(data, selector)
		assert.Equal(t, expected, actual)
	})

	t.Run("RequestStateChange", func(t *testing.T) {
		actionName := TestAction
		requestedState := 2
		expected := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"utf-8\"?><Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:a=\"http://schemas.xmlsoap.org/ws/2004/08/addressing\" xmlns:w=\"http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd\" xmlns=\"http://www.w3.org/2003/05/soap-envelope\"><Header><a:Action>test-action</a:Action><a:To>/wsman</a:To><w:ResourceURI>test-uriTestClass</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>PT60S</w:OperationTimeout></Header><Body><h:RequestStateChange_INPUT xmlns:h=\"test-uriTestClass\"><h:RequestedState>2</h:RequestedState></h:RequestStateChange_INPUT></Body></Envelope>", MessageID)
		MessageID++
		actual := base.RequestStateChange(actionName, requestedState)
		assert.Equal(t, expected, actual)
	})
}
