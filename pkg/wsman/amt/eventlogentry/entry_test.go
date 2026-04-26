/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package eventlogentry

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestPositiveAMT_EventLogEntry(t *testing.T) {
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/eventlogentry",
	}
	elementUnderTest := NewServiceWithClient(wsmanMessageCreator, &client)

	t.Run("amt_EventLogEntry Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			responseFunc func() (Response, error)
		}{
			{
				"should create and parse a valid AMT_EventLogEntry Get call",
				AMTEventLogEntry,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
			},
			{
				"should create a valid AMT_EventLogEntry Enumerate call",
				AMTEventLogEntry,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
			},
			{
				"should create a valid AMT_EventLogEntry Pull call",
				AMTEventLogEntry,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessagePull

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				response, err := test.responseFunc()

				assert.NoError(t, err)
				assert.NotNil(t, response)
				assert.Contains(t, response.XMLInput, test.method)
			})
		}
	})
}

func TestNegativeAMT_EventLogEntry(t *testing.T) {
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/eventlogentry",
	}
	elementUnderTest := NewServiceWithClient(wsmanMessageCreator, &client)

	t.Run("amt_EventLogEntry Error Handling Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			responseFunc func() (Response, error)
		}{
			{
				"should handle error when Get fails",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Get()
				},
			},
			{
				"should handle error when Enumerate fails",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Enumerate()
				},
			},
			{
				"should handle error when Pull fails",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageError

					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				_, err := test.responseFunc()
				assert.Error(t, err)
			})
		}
	})
}
