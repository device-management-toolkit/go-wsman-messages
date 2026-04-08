/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package asset

import (
	"encoding/xml"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

// OUTPUTS
// Response Types.
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName                   xml.Name `xml:"Body"`
		EnumerateResponse         common.EnumerateResponse
		GetResponse               AssetTable        `xml:"AMT_AssetTable"`
		AssetTableServiceResponse AssetTableService `xml:"AMT_AssetTableService"`
		PullResponse              PullResponse      `xml:"PullResponse"`
		DecodedTableResponse      []AssetTableEntry
		GetAssetTableDataResponse GetAssetTableData_OUTPUT `xml:"GetAssetTableData_OUTPUT"`
		GetAssetTableSizeResponse GetAssetTableSize_OUTPUT `xml:"GetAssetTableSize_OUTPUT"`
	}
	PullResponse struct {
		XMLName            xml.Name            `xml:"PullResponse"`
		AssetTableItems    []AssetTable        `xml:"Items>AMT_AssetTable"`
		AssetTableServices []AssetTableService `xml:"Items>AMT_AssetTableService"`
		EnumerationContext string              `xml:"EnumerationContext"`
	}

	// AssetTable represents a single asset table entry in AMT.
	AssetTable struct {
		XMLName                 xml.Name `xml:"AMT_AssetTable"`
		AssetTableIndex         int      `xml:"AssetTableIndex,omitempty"` // Unique identifier for the asset table entry
		ElementName             string   `xml:"ElementName,omitempty"`     // A user-friendly name for the object
		InstanceID              string   `xml:"InstanceID,omitempty"`
		Name                    string   `xml:"Name,omitempty"`                    // The Name property uniquely identifies the Service
		CreationClassName       string   `xml:"CreationClassName,omitempty"`       // CreationClassName indicates the name of the class or the subclass
		SystemName              string   `xml:"SystemName,omitempty"`              // The Name of the scoping System
		SystemCreationClassName string   `xml:"SystemCreationClassName,omitempty"` // The CreationClassName of the scoping System
		TableID                 int      `xml:"TableID,omitempty"`                 // Table identifier
		TableType               int      `xml:"TableType,omitempty"`
		TableTypeInfo           string   `xml:"TableTypeInfo,omitempty"`
		TableData               string   `xml:"TableData,omitempty"` // The actual table data in XML format
		TableSize               int      `xml:"TableSize,omitempty"` // Size of the table data
		Checksum                string   `xml:"Checksum,omitempty"`  // Checksum of table data for integrity verification
	}

	// AssetTableEntry represents a decoded entry from the asset table.
	AssetTableEntry struct {
		Index         int    `json:"index"`
		InstanceID    string `json:"instanceID"`
		Name          string `json:"name"`
		CreationClass string `json:"creationClass"`
		SystemName    string `json:"systemName"`
		TableID       int    `json:"tableID"`
		TableType     int    `json:"tableType"`
		TableTypeInfo string `json:"tableTypeInfo"`
		TableData     string `json:"tableData"`
		TableSize     int    `json:"tableSize"`
		Checksum      string `json:"checksum"`
	}

	// AssetTableService represents the asset table service in AMT.
	AssetTableService struct {
		XMLName                 xml.Name `xml:"AMT_AssetTableService"`
		CreationClassName       string   `xml:"CreationClassName"`         // CreationClassName indicates the name of the class or the subclass
		ElementName             string   `xml:"ElementName"`               // A user-friendly name for the object
		EnabledState            int      `xml:"EnabledState,omitempty"`    // EnabledState is an integer enumeration
		Name                    string   `xml:"Name"`                      // The Name property uniquely identifies the Service
		RequestedState          int      `xml:"RequestedState,omitempty"`  // RequestedState is an integer enumeration
		SystemCreationClassName string   `xml:"SystemCreationClassName"`   // The CreationClassName of the scoping System
		SystemName              string   `xml:"SystemName"`                // The Name of the scoping System
		AssetTableCount         int      `xml:"AssetTableCount,omitempty"` // Total number of asset tables
		TableTypes              []string `xml:"TableTypes"`
	}

	GetAssetTableData_OUTPUT struct {
		XMLName     xml.Name `xml:"GetAssetTableData_OUTPUT"`
		ReturnValue int      `xml:"ReturnValue"`
		TableData   string   `xml:"TableData"`
	}
	GetAssetTableSize_OUTPUT struct {
		XMLName     xml.Name `xml:"GetAssetTableSize_OUTPUT"`
		ReturnValue int      `xml:"ReturnValue"`
		TableSize   int      `xml:"TableSize"`
	}

	// ReturnValue is an integer enumeration that indicates the success or failure of an operation.
	ReturnValue int
)

// INPUTS
// Request Types.
type (
	GetAssetTableData_INPUT struct {
		XMLName xml.Name `xml:"h:GetAssetTableData_INPUT"`
		H       string   `xml:"xmlns:h,attr"`
		TableID int      `xml:"h:TableID"` // Specifies the asset table to retrieve data from
	}

	GetAssetTableSize_INPUT struct {
		XMLName xml.Name `xml:"h:GetAssetTableSize_INPUT"`
		H       string   `xml:"xmlns:h,attr"`
		TableID int      `xml:"h:TableID"` // Specifies the asset table to get size for
	}
)
