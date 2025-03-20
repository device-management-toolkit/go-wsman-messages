package secio

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
	"github.com/stretchr/testify/assert"
)

func TestPositiveIPS_SecIOService(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.IPSResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)

	client := wsmantesting.MockClient{
		PackageUnderTest: "ips/secio/service",
	}
	elementUnderTest := NewSecIOServiceWithClient(wsmanMessageCreator, &client)

	t.Run("IPS_SecIOService Tests", func(t *testing.T) {
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
				"should create a valid IPS_SecIOService Get wsman message",
				"IPS_SecIOService",
				wsmantesting.Get,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					SecIOServiceResponse: SecIOServiceResponse{
						XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSSecIOService), Local: IPSSecIOService},
						CreationClassName:       "IPS_SecIOService",
						SystemCreationClassName: "CIM_ComputerSystem",
						SystemName:              "ManagedSystem",
						Name:                    "SecIO",
						Started:                 true,
						Status:                  "OK",
						EnabledState:            2,
						Language:                1,
						RequestedLanguage:       2,
						Zoom:                    65535,
						DefaultScreen:           1,
					},
				},
			},
			// ENUMERATE
			{
				"should create a valid IPS_SecIOService Enumerate wsman message",
				"IPS_SecIOService",
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
						EnumerationContext: "A2000000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULL
			{
				"should create a valid IPS_SecIOService Pull wsman message",
				"IPS_SecIOService",
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
						SecIOServiceItems: []SecIOServiceResponse{
							{
								XMLName:                 xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSSecIOService), Local: IPSSecIOService},
								CreationClassName:       "IPS_SecIOService",
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "ManagedSystem",
								Name:                    "SecIO",
								Started:                 true,
								Status:                  "OK",
								EnabledState:            2,
								Language:                1,
								RequestedLanguage:       2,
								Zoom:                    65535,
								DefaultScreen:           1,
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
