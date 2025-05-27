package screensetting

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

// Data structure represents IPS_ScreenSettingData.
type (
	Data struct {
		base message.Base
	}
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
		XMLName                   xml.Name `xml:"Body"`
		PullResponse              PullResponse
		EnumerateResponse         common.EnumerateResponse
		ScreenSettingDataResponse ScreenSettingDataResponse
	}

	ScreenSettingDataResponse struct {
		XMLName        xml.Name `xml:"IPS_ScreenSettingData"`
		ElementName    string   `xml:"ElementName"`
		InstanceID     string   `xml:"InstanceID"`
		PrimaryIndex   uint8    `xml:"PrimaryIndex"`
		SecondaryIndex uint8    `xml:"SecondaryIndex"`
		TertiaryIndex  uint8    `xml:"TertiaryIndex"`
		QuadraryIndex  uint8    `xml:"QuadraryIndex"`
		IsActive       []bool   `xml:"IsActive"`
		UpperLeftX     []int32  `xml:"UpperLeftX"`
		UpperLeftY     []int32  `xml:"UpperLeftY"`
		ResolutionX    []uint32 `xml:"ResolutionX"`
		ResolutionY    []uint32 `xml:"ResolutionY"`
	}

	PullResponse struct {
		XMLName                xml.Name                    `xml:"PullResponse"`
		ScreenSettingDataItems []ScreenSettingDataResponse `xml:"Items>IPS_ScreenSettingData"`
	}
)

// INPUT.
type (
	ScreenSettingDataRequest struct {
		XMLName        xml.Name `xml:"h:IPS_ScreenSettingData,omitempty"`
		H              string   `xml:"xmlns:h,attr"`
		ElementName    string   `xml:"h:ElementName,omitempty"`
		InstanceID     string   `xml:"h:InstanceID,omitempty"`
		PrimaryIndex   uint8    `xml:"h:PrimaryIndex,omitempty"`
		SecondaryIndex uint8    `xml:"h:SecondaryIndex,omitempty"`
		TertiaryIndex  uint8    `xml:"h:TertiaryIndex,omitempty"`
		QuadraryIndex  uint8    `xml:"h:QuadraryIndex,omitempty"`
		IsActive       []bool   `xml:"h:IsActive,omitempty"`
		UpperLeftX     []int32  `xml:"h:UpperLeftX,omitempty"`
		UpperLeftY     []int32  `xml:"h:UpperLeftY,omitempty"`
		ResolutionX    []uint32 `xml:"h:ResolutionX,omitempty"`
		ResolutionY    []uint32 `xml:"h:ResolutionY,omitempty"`
	}
)
