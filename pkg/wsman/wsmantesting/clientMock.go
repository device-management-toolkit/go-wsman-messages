/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wsmantesting

import (
	"crypto/tls"
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// MockClient is a mock implementation of the wsman.Client interface for testing.
type MockClient struct {
	CurrentMessage   string
	PackageUnderTest string
}

func (c *MockClient) IsAuthenticated() bool { return true }

func (c *MockClient) Post(msg string) ([]byte, error) {
	if strings.EqualFold(c.CurrentMessage, "error") {
		return []byte(""), nil
	}
	// read an xml file from disk:
	xmlFile, err := os.Open("../../wsmantesting/responses/" + c.PackageUnderTest + "/" + strings.ToLower(c.CurrentMessage) + ".xml")
	if err != nil {
		logrus.Print("Error opening file:", err)

		return nil, err
	}
	defer xmlFile.Close()
	// read file into string
	xmlData, err := io.ReadAll(xmlFile)
	if err != nil {
		logrus.Print("Error reading file:", err)

		return nil, err
	}
	// strip carriage returns and new line characters
	xmlData = []byte(strings.ReplaceAll(string(xmlData), "\r\n", ""))

	// Simulate a successful response for testing.
	return xmlData, nil
}
func (c *MockClient) Send(data []byte) error                          { return nil }
func (c *MockClient) Receive() ([]byte, error)                        { return nil, nil }
func (c *MockClient) CloseConnection() error                          { return nil }
func (c *MockClient) Connect() error                                  { return nil }
func (c *MockClient) GetServerCertificate() (*tls.Certificate, error) { return nil, nil }
