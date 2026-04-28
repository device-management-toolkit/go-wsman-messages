/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package provisioningrecordlog

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName           xml.Name              `xml:"Body"`
		GetResponse       ProvisioningRecordLog `xml:"IPS_ProvisioningRecordLog"`
		PullResponse      PullResponse
		EnumerateResponse common.EnumerateResponse
	}

	ProvisioningRecordLog struct {
		XMLName                xml.Name `xml:"IPS_ProvisioningRecordLog"`
		CurrentNumberOfRecords int      `xml:"CurrentNumberOfRecords"`
		ElementName            string   `xml:"ElementName,omitempty"`
		EnabledState           int      `xml:"EnabledState"`
		HealthState            int      `xml:"HealthState"`
		InstanceID             string   `xml:"InstanceID,omitempty"`
		LogState               int      `xml:"LogState"`
		MaxNumberOfRecords     int      `xml:"MaxNumberOfRecords"`
		Name                   string   `xml:"Name,omitempty"`
		OverwritePolicy        int      `xml:"OverwritePolicy"`
		RequestedState         int      `xml:"RequestedState"`
	}

	PullResponse struct {
		XMLName                    xml.Name                `xml:"PullResponse"`
		ProvisioningRecordLogItems []ProvisioningRecordLog `xml:"Items>IPS_ProvisioningRecordLog"`
	}
)
