package screensettings

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestPositiveIPS_ScreenSettingData(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.IPSResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)

	client := wsmantesting.MockClient{
		PackageUnderTest: "ips/screensetting",
	}
	elementUnderTest := NewScreenSettingDataWithClient(wsmanMessageCreator, &client)

	t.Run("IPS_ScreenSettingData Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			extraHeader      string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			// GET
			{
				"should create a valid IPS_ScreenSettingData Get wsman message",
				"IPS_ScreenSettingData",
				wsmantesting.Get,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					ScreenSettingDataResponse: ScreenSettingDataResponse{
						XMLName:        xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSScreenSettingData), Local: IPSScreenSettingData},
						ElementName:    "test",
						InstanceID:     "Intel(r) Screen Settings",
						PrimaryIndex:   1,
						SecondaryIndex: 2,
						TertiaryIndex:  3,
						QuadraryIndex:  4,
						IsActive:       []bool{true, true, false, false},
						UpperLeftX:     []int32{0, 1920, -1, -1},
						UpperLeftY:     []int32{0, 0, -1, -1},
						ResolutionX:    []uint32{1920, 1920, 0, 0},
						ResolutionY:    []uint32{1080, 1080, 0, 0},
					},
				},
			},
			// ENUMERATE
			{
				"should create a valid IPS_ScreenSettingData Enumerate wsman message",
				"IPS_ScreenSettingData",
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
						EnumerationContext: "A4000000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULL
			{
				"should create a valid IPS_ScreenSettingData Pull wsman message",
				"IPS_ScreenSettingData",
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
						ScreenSettingDataItems: []ScreenSettingDataResponse{
							{
								XMLName:        xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSScreenSettingData), Local: IPSScreenSettingData},
								ElementName:    "test",
								InstanceID:     "Intel(r) Screen Settings",
								PrimaryIndex:   1,
								SecondaryIndex: 2,
								TertiaryIndex:  3,
								QuadraryIndex:  4,
								IsActive:       []bool{true, true, false, false},
								UpperLeftX:     []int32{0, 1920, -1, -1},
								UpperLeftY:     []int32{0, 0, -1, -1},
								ResolutionX:    []uint32{1920, 1920, 0, 0},
								ResolutionY:    []uint32{1080, 1080, 0, 0},
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
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.Body)
			})
		}
	})
}
