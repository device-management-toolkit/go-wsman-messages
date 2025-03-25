package screensetting

import (
	"encoding/xml"
	"fmt"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/ips/methods"
)

// NewScreenSettingData creates a new instance of ScreenSettingData.
func NewScreenSettingDataWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Data {
	return Data{
		base: message.NewBaseWithClient(wsmanMessageCreator, IPSScreenSettingData, client),
	}
}

// Get retrieves the representation of the instance.
func (settings Data) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: settings.base.Get(nil),
		},
	}

	err = settings.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// Enumerate returns an enumeration context which is used in a subsequent Pull call.
func (settings Data) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: settings.base.Enumerate(),
		},
	}

	err = settings.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// Pull returns the instances of this class using an enumeration context from Enumerate.
func (settings Data) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: settings.base.Pull(enumerationContext),
		},
	}

	err = settings.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// Put updates the screen settings.
func (settings Data) Put(screenSettings ScreenSettingDataRequest) (response Response, err error) {
	screenSettings.H = fmt.Sprintf("%s%s", message.IPSSchema, IPSScreenSettingData)
	response = Response{
		Message: &client.Message{
			XMLInput: settings.base.Put(screenSettings, false, nil),
		},
	}

	err = settings.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// ResetToDefault resets the screen settings to default.
func (settings Data) ResetToDefault() (response Response, err error) {
	header := settings.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(IPSScreenSettingData, ResetToDefault), IPSScreenSettingData, nil, "", "")
	body := settings.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(ResetToDefault), IPSScreenSettingData, nil)

	response = Response{
		Message: &client.Message{
			XMLInput: settings.base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	err = settings.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
