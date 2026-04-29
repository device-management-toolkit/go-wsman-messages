/*********************************************************************
 * Copyright (c) Intel Corporation 2026
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package asset facilitates communication with Intel® AMT devices to retrieve asset table information.
package asset

import (
	"encoding/xml"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/base"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type Table struct {
	base.WSManService[Response]
}

// NewTableWithClient instantiates a new Asset Table service.
func NewTableWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Table {
	return Table{
		base.NewService[Response](wsmanMessageCreator, AMTAssetTable, client),
	}
}

// GetByAssetTableIndex retrieves a specific asset table instance by its unique index.
func (t Table) GetByAssetTableIndex(index int) (response Response, err error) {
	selector := &message.Selector{
		Name:  "AssetTableIndex",
		Value: strconv.Itoa(index),
	}

	msg := &client.Message{XMLInput: t.Base.Get(selector)}
	response.Message = msg

	if err = t.Base.Execute(msg); err != nil {
		return response, err
	}

	if err = xml.Unmarshal([]byte(msg.XMLOutput), &response); err != nil {
		return response, err
	}

	response.Message = msg

	return response, nil
}

// GetByInstanceIDAndTableType retrieves a specific asset table instance using the
// selector combination observed from firmware enumeration output.
func (t Table) GetByInstanceIDAndTableType(instanceID string, tableType int) (response Response, err error) {
	selectors := []message.Selector{
		{
			Name:  "InstanceID",
			Value: instanceID,
		},
		{
			Name:  "TableType",
			Value: strconv.Itoa(tableType),
		},
	}

	header := t.Base.WSManMessageCreator.CreateHeader(message.BaseActionsGet, t.Base.ClassName, selectors, "", "")
	msg := &client.Message{XMLInput: t.Base.WSManMessageCreator.CreateXML(header, message.GetBody)}
	response.Message = msg

	if err = t.Base.Execute(msg); err != nil {
		return response, err
	}

	if err = xml.Unmarshal([]byte(msg.XMLOutput), &response); err != nil {
		return response, err
	}

	response.Message = msg

	return response, nil
}

// DecodeAssetTable decodes the asset table response into a structured format.
func DecodeAssetTable(response Response) []AssetTableEntry {
	var entries []AssetTableEntry

	if len(response.Body.PullResponse.AssetTableItems) == 0 && len(response.Body.DecodedTableResponse) == 0 {
		if response.Body.GetResponse.AssetTableIndex != 0 ||
			response.Body.GetResponse.InstanceID != "" ||
			response.Body.GetResponse.TableData != "" ||
			response.Body.GetResponse.TableType != 0 {
			return []AssetTableEntry{
				{
					Index:         response.Body.GetResponse.AssetTableIndex,
					InstanceID:    response.Body.GetResponse.InstanceID,
					Name:          response.Body.GetResponse.Name,
					CreationClass: response.Body.GetResponse.CreationClassName,
					SystemName:    response.Body.GetResponse.SystemName,
					TableID:       response.Body.GetResponse.TableID,
					TableType:     response.Body.GetResponse.TableType,
					TableTypeInfo: response.Body.GetResponse.TableTypeInfo,
					TableData:     response.Body.GetResponse.TableData,
					TableSize:     response.Body.GetResponse.TableSize,
					Checksum:      response.Body.GetResponse.Checksum,
				},
			}
		}

		logrus.Debug("No asset table items found in response")

		return entries
	}

	// If already decoded, return as-is
	if len(response.Body.DecodedTableResponse) > 0 {
		return response.Body.DecodedTableResponse
	}

	// Decode from AssetTable items
	for _, item := range response.Body.PullResponse.AssetTableItems {
		entry := AssetTableEntry{
			Index:         item.AssetTableIndex,
			InstanceID:    item.InstanceID,
			Name:          item.Name,
			CreationClass: item.CreationClassName,
			SystemName:    item.SystemName,
			TableID:       item.TableID,
			TableType:     item.TableType,
			TableTypeInfo: item.TableTypeInfo,
			TableData:     item.TableData,
			TableSize:     item.TableSize,
			Checksum:      item.Checksum,
		}
		entries = append(entries, entry)
	}

	return entries
}
