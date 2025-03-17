/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package message

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

type Base struct {
	WSManMessageCreator *WSManMessageCreator
	className           string
	client              client.WSMan
}

type Header struct {
	XMLName     xml.Name `xml:"Header"`
	To          string   `xml:"To"`
	RelatesTo   int      `xml:"RelatesTo"`
	Action      Action   `xml:"Action"`
	MessageID   string   `xml:"MessageID"`
	ResourceURI string   `xml:"ResourceURI"`
}
type Action struct {
	XMLName        xml.Name `xml:"Action"`
	MustUnderstand string   `xml:"mustUnderstand,attr"`
	Value          string   `xml:",chardata"`
}

type SelectorSet struct {
	XMLName  xml.Name `xml:"SelectorSet"`
	Selector Selector
}

type Selector struct {
	XMLName xml.Name `xml:"Selector,omitempty"`
	Name    string   `xml:"Name,attr"`
	Value   string   `xml:",chardata"`
}
type Selector_OUTPUT struct {
	XMLName xml.Name `xml:"Selector,omitempty"`
	Name    string   `xml:"Name,attr"`
	Value   string   `xml:",chardata"`
}

type ReturnValue struct {
	XMLName        xml.Name `xml:"RequestStateChange_OUTPUT,omitempty"`
	ReturnValue    int      `xml:"ReturnValue,omitempty"`
	ReturnValueStr string   `xml:"ReturnValueStr,omitempty"`
}
type WSManMessageCreator struct {
	MessageID        int
	XMLCommonPrefix  string
	XMLCommonEnd     string
	AnonymousAddress string
	DefaultTimeout   string
	ResourceURIBase  string
}
