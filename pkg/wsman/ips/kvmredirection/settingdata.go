package kvmredirection

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/methods"
)

type SettingData struct {
	base.WSManService[Response]
}

// NewKVMRedirectionSettingDataWithClient returns a new instance of the KVMRedirectionSettings struct.
func NewKVMRedirectionSettingDataWithClient(creator *message.WSManMessageCreator, wsclient client.WSMan) SettingData {
	return SettingData{
		base.NewService[Response](creator, IPSKVMRedirectionSettingData, wsclient),
	}
}

func (settings *SettingData) TerminateSession() (response Response, err error) {
	// TerminateSession stops an active KVM session.
	header := settings.Base.WSManMessageCreator.CreateHeader(methods.GenerateAction(IPSKVMRedirectionSettingData, TerminateSession), IPSKVMRedirectionSettingData, nil, "", "")
	body := settings.Base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(TerminateSession), IPSKVMRedirectionSettingData, nil)

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
