package secio

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

// OUTPUT.
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName              xml.Name `xml:"Body"`
		PullResponse         PullResponse
		EnumerateResponse    common.EnumerateResponse
		SecIOServiceResponse SecIOServiceResponse
	}

	SecIOServiceResponse struct {
		XMLName                 xml.Name `xml:"IPS_SecIOService"`
		CreationClassName       string   `xml:"CreationClassName"`
		DefaultScreen           uint8    `xml:"DefaultScreen"`
		EnabledState            uint16   `xml:"EnabledState"`
		Name                    string   `xml:"Name"`
		RequestedLanguage       uint16   `xml:"RequestedLanguage"`
		Started                 bool     `xml:"Started"`
		Status                  string   `xml:"Status"`
		SystemCreationClassName string   `xml:"SystemCreationClassName"`
		SystemName              string   `xml:"SystemName"`
		Language                uint16   `xml:"language"`
		Zoom                    uint16   `xml:"zoom"`
	}
	PullResponse struct {
		XMLName           xml.Name               `xml:"PullResponse"`
		SecIOServiceItems []SecIOServiceResponse `xml:"Items>IPS_SecIOService"`
	}
)

// INPUT.
type (
	SecIOServiceRequest struct {
		XMLName           xml.Name `xml:"h:IPS_SecIOService,omitempty"`
		H                 string   `xml:"xmlns:h,attr"`
		ElementName       string   `xml:"h:ElementName,omitempty"`
		InstanceID        string   `xml:"h:InstanceID,omitempty"`
		Language          uint16   `xml:"h:language,omitempty"`
		RequestedLanguage uint16   `xml:"h:RequestedLanguage,omitempty"`
		Zoom              uint16   `xml:"h:zoom,omitempty"`
		DefaultScreen     uint8    `xml:"h:DefaultScreen,omitempty"`
	}
)
