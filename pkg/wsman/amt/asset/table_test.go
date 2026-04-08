/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package asset

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestPositiveAMT_AssetTable(t *testing.T) {
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/asset",
	}
	elementUnderTest := NewTableWithClient(wsmanMessageCreator, &client)

	t.Run("amt_AssetTable Tests", func(t *testing.T) {
		tests := []struct {
			name         string
			method       string
			responseFunc func() (Response, error)
		}{
			// GET
			{
				"should create and parse a valid AMT_AssetTable Get call",
				AMTAssetTable,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
			},
			// ENUMERATE
			{
				"should create a valid AMT_AssetTable Enumerate call",
				AMTAssetTable,
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageEnumerate

					return elementUnderTest.Enumerate()
				},
			},
			// PULL
			{
				"should create a valid AMT_AssetTable Pull call",
				AMTAssetTable,
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

				// Verify XML input contains expected method
				assert.Contains(t, response.XMLInput, test.method)
			})
		}
	})
}

func TestNegativeAMT_AssetTable(t *testing.T) {
	resourceURIBase := wsmantesting.AMTResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/asset",
	}
	elementUnderTest := NewTableWithClient(wsmanMessageCreator, &client)

	t.Run("amt_AssetTable Error Handling Tests", func(t *testing.T) {
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
				// Error is expected from empty error response
				assert.Error(t, err)
			})
		}
	})
}

func TestDecodeAssetTable(t *testing.T) {
	tests := []struct {
		name     string
		response Response
		expected int
	}{
		{
			name: "empty response",
			response: Response{
				Body: Body{
					PullResponse:         PullResponse{AssetTableItems: []AssetTable{}},
					DecodedTableResponse: []AssetTableEntry{},
				},
			},
			expected: 0,
		},
		{
			name: "single asset table entry",
			response: Response{
				Body: Body{
					PullResponse: PullResponse{AssetTableItems: []AssetTable{
						{
							AssetTableIndex:   1,
							Name:              "Asset1",
							CreationClassName: "AMT_AssetTable",
							SystemName:        "AMT",
							TableID:           1,
							TableData:         "testdata",
							TableSize:         8,
							Checksum:          "abc123",
						},
					}},
					DecodedTableResponse: []AssetTableEntry{},
				},
			},
			expected: 1,
		},
		{
			name: "multiple asset table entries",
			response: Response{
				Body: Body{
					PullResponse: PullResponse{AssetTableItems: []AssetTable{
						{
							AssetTableIndex: 1,
							TableID:         1,
							TableData:       "data1",
						},
						{
							AssetTableIndex: 2,
							TableID:         2,
							TableData:       "data2",
						},
					}},
					DecodedTableResponse: []AssetTableEntry{},
				},
			},
			expected: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := DecodeAssetTable(test.response)
			assert.Equal(t, test.expected, len(result))

			// Verify data correctness
			for i, entry := range result {
				expectedIndex := i + 1
				assert.Equal(t, expectedIndex, entry.Index)
			}
		})
	}
}
