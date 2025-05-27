package secio

import (
	"encoding/xml"
	"fmt"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/ips/methods"
)

// NewSecIOService creates a new instance of SecIOService.
func NewSecIOServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Service {
	return Service{
		base: message.NewBaseWithClient(wsmanMessageCreator, IPSSecIOService, client),
	}
}

// Get retrieves the representation of the instance.
func (settings Service) Get() (response Response, err error) {
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
func (settings Service) Enumerate() (response Response, err error) {
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
func (settings Service) Pull(enumerationContext string) (response Response, err error) {
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

// Put updates the SecIO settings.
func (settings Service) Put(secioSettings SecIOServiceRequest) (response Response, err error) {
	secioSettings.H = fmt.Sprintf("%s%s", message.IPSSchema, IPSSecIOService)
	response = Response{
		Message: &client.Message{
			XMLInput: settings.base.Put(secioSettings, false, nil),
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

// RequestStateChange changes the operational state of SecIO.
func (settings Service) RequestStateChange(requestedState uint16) (response Response, err error) {
	header := settings.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(IPSSecIOService, RequestStateChange), IPSSecIOService, nil, "", "")
	body := settings.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(RequestStateChange), IPSSecIOService,
		struct {
			RequestedState uint16 `xml:"h:RequestedState"`
		}{
			RequestedState: requestedState,
		})

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
