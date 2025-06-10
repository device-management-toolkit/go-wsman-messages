package screensetting

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/methods"
)

type Data struct {
	base.WSManService[Response]
}

// NewScreenSettingData creates a new instance of ScreenSettingData.
func NewScreenSettingDataWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Data {
	return Data{
		base.NewService[Response](wsmanMessageCreator, IPSScreenSettingData, client),
	}
}

// ResetToDefault resets the screen settings to default.
func (settings Data) ResetToDefault() (response Response, err error) {
	header := settings.Base.WSManMessageCreator.CreateHeader(methods.GenerateAction(IPSScreenSettingData, ResetToDefault), IPSScreenSettingData, nil, "", "")
	body := settings.Base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(ResetToDefault), IPSScreenSettingData, nil)

	response = Response{
		Message: &client.Message{
			XMLInput: settings.Base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	err = settings.Base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
