package kvmredirection

import (
	"encoding/xml"
	"fmt"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/ips/methods"
)

// NewKVMRedirectionSettings returns a new instance of the KVMRedirectionSettings struct.
func NewKVMRedirectionSettingsWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) SettingsData {
	return SettingsData{
		base: message.NewBaseWithClient(wsmanMessageCreator, IPSKVMRedirectionSettingData, client),
	}
}

// Get retrieves the representation of the instance.
func (settings SettingsData) Get() (response Response, err error) {
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
func (settings SettingsData) Enumerate() (response Response, err error) {
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
func (settings SettingsData) Pull(enumerationContext string) (response Response, err error) {
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

// Put updates the KVM settings.
func (settings SettingsData) Put(kvmSettings KVMRedirectionSettingsRequest) (response Response, err error) {
	kvmSettings.H = fmt.Sprintf("%s%s", message.IPSSchema, IPSKVMRedirectionSettingData)
	response = Response{
		Message: &client.Message{
			XMLInput: settings.base.Put(kvmSettings, false, nil),
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

// TerminateSession stops an active KVM session.
func (settings SettingsData) TerminateSession() (response Response, err error) {
	header := settings.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(IPSKVMRedirectionSettingData, TerminateSession), IPSKVMRedirectionSettingData, nil, "", "")
	body := settings.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(TerminateSession), IPSKVMRedirectionSettingData, nil)

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
