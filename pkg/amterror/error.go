/*********************************************************************
 * Copyright (c) Intel Corporation 2024
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
package amterror

import (
	"encoding/xml"
	"fmt"
)

func (e *AMTError) Error() string {
	return fmt.Sprintf("Error [SubCode: %s] Message: %s, Detail: %s", e.SubCode, e.Message, e.Detail)
}

func NewAMTError(subCode, message, detail string) *AMTError {
	return &AMTError{
		SubCode: subCode,
		Message: message,
		Detail:  detail,
	}
}

func DecodeAMTErrorString(s string) error {
	checkForErrorResponse := ErrorResponse{}

	err := xml.Unmarshal([]byte(s), &checkForErrorResponse)
	if err != nil {
		return err
	}

	return NewAMTError(checkForErrorResponse.Body.Fault.Code.SubCode.Value, checkForErrorResponse.Body.Fault.Reason.Text, checkForErrorResponse.Body.Fault.Detail)
}

// AMT WSMAN Error Response Types.
type (
	AMTError struct {
		SubCode string
		Message string
		Detail  string
	}

	Header struct {
		XMLName     xml.Name `xml:"Header"`
		To          string   `xml:"To"`
		RelatesTo   int      `xml:"RelatesTo"`
		Action      Action   `xml:"Action"`
		MessageID   string   `xml:"MessageID"`
		ResourceURI string   `xml:"ResourceURI"`
	}

	Action struct {
		XMLName        xml.Name `xml:"Action"`
		MustUnderstand string   `xml:"mustUnderstand,attr"`
		Value          string   `xml:",chardata"`
	}

	ErrorResponse struct {
		XMLName xml.Name  `xml:"Envelope"`
		Header  Header    `xml:"Header"`
		Body    ErrorBody `xml:"Body"`
	}

	ErrorBody struct {
		XMLName xml.Name `xml:"Body"`
		Fault   Fault    `xml:"Fault"`
	}

	Fault struct {
		XMLName xml.Name `xml:"Fault"`
		Code    Code     `xml:"Code"`
		Reason  Reason   `xml:"Reason"`
		Detail  string   `xml:"Detail"`
	}

	Code struct {
		XMLName xml.Name `xml:"Code"`
		Value   string   `xml:"Value"`
		SubCode SubCode  `xml:"Subcode"`
	}

	SubCode struct {
		XMLName xml.Name `xml:"Subcode"`
		Value   string   `xml:"Value"`
	}

	Reason struct {
		XMLName xml.Name `xml:"Reason"`
		Text    string   `xml:"Text"`
	}
)
