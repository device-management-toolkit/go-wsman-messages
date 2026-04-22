/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"encoding/xml"
	"fmt"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// Instantiates a new Boot Setting Data service.
func NewBootSettingDataWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) SettingData {
	return SettingData{
		base.NewService[Response](wsmanMessageCreator, AMTBootSettingData, client),
	}
}

// Put overrides the generic Put because AMT expects only a specific subset of
// BootSettingData fields on the wire; marshaling the full request struct would
// cause firmware to reject the request. Keep the hand-crafted body.
func (settingData SettingData) Put(bootSettingData BootSettingDataRequest) (response Response, err error) {
	header := settingData.Base.WSManMessageCreator.CreateHeader(message.BaseActionsPut, AMTBootSettingData, nil, "", "")
	body := fmt.Sprintf(
		`<Body><h:AMT_BootSettingData xmlns:h="%sAMT_BootSettingData"><h:BIOSPause>%t</h:BIOSPause><h:BIOSSetup>%t</h:BIOSSetup><h:BootMediaIndex>%d</h:BootMediaIndex><h:ConfigurationDataReset>%t</h:ConfigurationDataReset><h:ElementName>%s</h:ElementName><h:EnforceSecureBoot>%t</h:EnforceSecureBoot><h:FirmwareVerbosity>%d</h:FirmwareVerbosity><h:ForcedProgressEvents>%t</h:ForcedProgressEvents><h:IDERBootDevice>%d</h:IDERBootDevice><h:InstanceID>%s</h:InstanceID><h:LockKeyboard>%t</h:LockKeyboard><h:LockPowerButton>%t</h:LockPowerButton><h:LockResetButton>%t</h:LockResetButton><h:LockSleepButton>%t</h:LockSleepButton><h:OwningEntity>%s</h:OwningEntity><h:PlatformErase>%t</h:PlatformErase><h:RSEPassword>%s</h:RSEPassword><h:ReflashBIOS>%t</h:ReflashBIOS><h:SecureErase>%t</h:SecureErase><h:UefiBootParametersArray>%s</h:UefiBootParametersArray><h:UefiBootNumberOfParams>%d</h:UefiBootNumberOfParams><h:UseIDER>%t</h:UseIDER><h:UseSOL>%t</h:UseSOL><h:UseSafeMode>%t</h:UseSafeMode><h:UserPasswordBypass>%t</h:UserPasswordBypass></h:AMT_BootSettingData></Body>`,
		settingData.Base.WSManMessageCreator.ResourceURIBase,
		bootSettingData.BIOSPause,
		bootSettingData.BIOSSetup,
		bootSettingData.BootMediaIndex,
		bootSettingData.ConfigurationDataReset,
		bootSettingData.ElementName,
		bootSettingData.EnforceSecureBoot,
		bootSettingData.FirmwareVerbosity,
		bootSettingData.ForcedProgressEvents,
		bootSettingData.IDERBootDevice,
		bootSettingData.InstanceID,
		bootSettingData.LockKeyboard,
		bootSettingData.LockPowerButton,
		bootSettingData.LockResetButton,
		bootSettingData.LockSleepButton,
		bootSettingData.OwningEntity,
		bootSettingData.PlatformErase,
		bootSettingData.RSEPassword,
		bootSettingData.ReflashBIOS,
		bootSettingData.SecureErase,
		bootSettingData.UefiBootParametersArray,
		bootSettingData.UefiBootNumberOfParams,
		bootSettingData.UseIDER,
		bootSettingData.UseSOL,
		bootSettingData.UseSafeMode,
		bootSettingData.UserPasswordBypass)

	response = Response{
		Message: &client.Message{
			XMLInput: settingData.Base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	// send the message to AMT
	err = settingData.Base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, err
}
