/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package managementpresence

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsmantesting"
)

func TestAMT_ManagementPresenceRemoteSAP(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	elementUnderTest := NewManagementPresenceRemoteSAP(wsmanMessageCreator)

	t.Run("amt_* Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			action       string
			body         string
			extraHeader  string
			responseFunc func() string
		}{
			//GETS
			{"should create a valid AMT_ManagementPresenceRemoteSAP Get wsman message", "AMT_ManagementPresenceRemoteSAP", wsmantesting.GET, "", "", elementUnderTest.Get},
			//ENUMERATES
			{"should create a valid AMT_ManagementPresenceRemoteSAP Enumerate wsman message", "AMT_ManagementPresenceRemoteSAP", wsmantesting.ENUMERATE, wsmantesting.ENUMERATE_BODY, "", elementUnderTest.Enumerate},
			//PULLS
			{"should create a valid AMT_ManagementPresenceRemoteSAP Pull wsman message", "AMT_ManagementPresenceRemoteSAP", wsmantesting.PULL, wsmantesting.PULL_BODY, "", func() string { return elementUnderTest.Pull(wsmantesting.EnumerationContext) }},
			//DELETE
			{"should create a valid AMT_ManagementPresenceRemoteSAP Delete wsman message", "AMT_ManagementPresenceRemoteSAP", wsmantesting.DELETE, "", "<w:SelectorSet><w:Selector Name=\"Name\">instanceID123</w:Selector></w:SelectorSet>", func() string { return elementUnderTest.Delete("instanceID123") }},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				correctResponse := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, test.extraHeader, test.body)
				messageID++
				response := test.responseFunc()
				if response != correctResponse {
					assert.Equal(t, correctResponse, response)
				}
			})
		}
	})
}
