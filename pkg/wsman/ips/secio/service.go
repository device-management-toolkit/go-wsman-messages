package secio

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/methods"
)

type Service struct {
	base.WSManService[Response]
}

func NewSecIOServiceWithClient(creator *message.WSManMessageCreator, wsclient client.WSMan) Service {
	return Service{
		base.NewService[Response](creator, IPSSecIOService, wsclient),
	}
}

// RequestStateChange changes the operational state of SecIO.
func (settings Service) RequestStateChange(requestedState uint16) (response Response, err error) {
	header := settings.Base.WSManMessageCreator.CreateHeader(methods.GenerateAction(IPSSecIOService, RequestStateChange), IPSSecIOService, nil, "", "")
	body := settings.Base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(RequestStateChange), IPSSecIOService,
		struct {
			RequestedState uint16 `xml:"h:RequestedState"`
		}{
			RequestedState: requestedState,
		})

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
