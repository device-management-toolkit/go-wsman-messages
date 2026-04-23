/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package powermanagementcapabilities

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestPositiveCIMPowerManagementCapabilities(t *testing.T) {
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/powermanagementcapabilities",
	}
	elementUnderTest := NewPowerManagementCapabilitiesWithClient(wsmanMessageCreator, &client)

	t.Run("cim_PowerManagementCapabilities Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			responseFunc func() (Response, error)
		}{
			{
				"should create and parse a valid CIM_PowerManagementCapabilities Get call",
				CIMPowerManagementCapabilities,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
			},
			{
				"should create a valid CIM_PowerManagementCapabilities Enumerate call",
				CIMPowerManagementCapabilities,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
			},
			{
				"should create a valid CIM_PowerManagementCapabilities Pull call",
				CIMPowerManagementCapabilities,
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

func TestNegativeCIMPowerManagementCapabilities(t *testing.T) {
	resourceURIBase := wsmantesting.CIMResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/powermanagementcapabilities",
	}
	elementUnderTest := NewPowerManagementCapabilitiesWithClient(wsmanMessageCreator, &client)

	t.Run("cim_PowerManagementCapabilities Error Handling Tests", func(t *testing.T) {
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
