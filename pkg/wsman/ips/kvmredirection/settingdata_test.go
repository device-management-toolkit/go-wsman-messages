package kvmredirection

import (
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestReturnValue_String(t *testing.T) {
	tests := []struct {
		state    ReturnValue
		expected string
	}{
		{ReturnValueSuccess, "Success"},
		{ReturnValueInternalError, "InternalError"},
		{ReturnValue(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		assert.Equal(t, test.expected, result)
	}
}

func TestPositiveIPS_KVMRedirectionSettingData(t *testing.T) {
	messageID := 0
	resourceURIBase := wsmantesting.IPSResourceURIBase
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceURIBase)

	client := wsmantesting.MockClient{
		PackageUnderTest: "ips/kvmredirection/settings",
	}
	elementUnderTest := NewKVMRedirectionSettingsWithClient(wsmanMessageCreator, &client)

	t.Run("ips_KVMRedirectionSettingData Tests", func(t *testing.T) {
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
				"should create a valid IPS_KVMRedirectionSettingData Get wsman message",
				"IPS_KVMRedirectionSettingData",
				wsmantesting.Get,
				"",
				"",
				func() (Response, error) {
					client.CurrentMessage = wsmantesting.CurrentMessageGet

					return elementUnderTest.Get()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					KVMRedirectionSettingsResponse: KVMRedirectionSettingsResponse{
						XMLName:            xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSKVMRedirectionSettingData), Local: IPSKVMRedirectionSettingData},
						BackToBackFbMode:   false,
						OptInPolicy:        true,
						OptInPolicyTimeout: 300,
						SessionTimeout:     3,
						ElementName:        "Intel(r) KVM Redirection Settings",
						InstanceID:         "Intel(r) KVM Redirection Settings",
						EnabledByMEBx:      true,
						Is5900PortEnabled:  false,
						RFBPassword:        "",
						DefaultScreen:      0,
					},
				},
			},
			// ENUMERATE
			{
				"should create a valid IPS_KVMRedirectionSettingData Enumerate wsman message",
				"IPS_KVMRedirectionSettingData",
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
						EnumerationContext: "9E000000-0000-0000-0000-000000000000",
					},
				},
			},
			// PULL
			{
				"should create a valid IPS_KVMRedirectionSettingData Pull wsman message",
				"IPS_KVMRedirectionSettingData",
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
						KVMRedirectionSettingsItems: []KVMRedirectionSettingsResponse{
							{
								XMLName:            xml.Name{Space: fmt.Sprintf("%s%s", message.IPSSchema, IPSKVMRedirectionSettingData), Local: IPSKVMRedirectionSettingData},
								BackToBackFbMode:   false,
								OptInPolicy:        true,
								OptInPolicyTimeout: 300,
								SessionTimeout:     3,
								ElementName:        "Intel(r) KVM Redirection Settings",
								InstanceID:         "Intel(r) KVM Redirection Settings",
								EnabledByMEBx:      true,
								Is5900PortEnabled:  false,
								RFBPassword:        "",
								DefaultScreen:      0,
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
