/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package message

import (
	"encoding/xml"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"
)

func NewWSManMessageCreator(resourceURIBase string) *WSManMessageCreator {
	return &WSManMessageCreator{
		MessageID:        0,
		XMLCommonPrefix:  `<?xml version="1.0" encoding="utf-8"?><Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope">`,
		XMLCommonEnd:     `</Envelope>`,
		AnonymousAddress: "http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous",
		DefaultTimeout:   "PT60S",
		ResourceURIBase:  resourceURIBase,
	}
}

func (w WSManMessageCreator) CreateXML(header, body string) string {
	return w.XMLCommonPrefix + header + body + w.XMLCommonEnd
}

func (w *WSManMessageCreator) CreateHeader(action, wsmanClass string, selectorSet []Selector, address, timeout string) string {
	header := "<Header>"
	header += fmt.Sprintf(`<a:Action>%s</a:Action><a:To>/wsman</a:To><w:ResourceURI>%s%s</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo>`, action, w.ResourceURIBase, wsmanClass, w.MessageID)

	w.MessageID++

	if address != "" {
		header += fmt.Sprintf(`<a:Address>%s</a:Address>`, address)
	} else {
		header += fmt.Sprintf(`<a:Address>%s</a:Address>`, w.AnonymousAddress)
	}

	header += "</a:ReplyTo>"

	if timeout != "" {
		header += fmt.Sprintf(`<w:OperationTimeout>%s</w:OperationTimeout>`, timeout)
	} else {
		header += fmt.Sprintf(`<w:OperationTimeout>%s</w:OperationTimeout>`, w.DefaultTimeout)
	}

	if selectorSet != nil {
		header += w.createSelector(selectorSet)
	}

	header += "</Header>"

	return header
}

func IsSlice(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Slice
}

func (w WSManMessageCreator) namespaceMe(subj interface{}, wsmanClass string) {
	ifaceValue := reflect.ValueOf(subj)
	// Check if the interface value is a pointer
	if ifaceValue.Kind() == reflect.Ptr {
		// Get the underlying value of the interface
		stype := ifaceValue.Elem()
		if stype.Kind() == reflect.Struct {
			field := stype.FieldByName("H")

			fieldPtr, ok := field.Addr().Interface().(*string)
			if !ok {
				logrus.Error("Failed to convert H to string")
			}

			*fieldPtr = fmt.Sprintf("%s%s", w.ResourceURIBase, wsmanClass)
		}
	}
}

func (w WSManMessageCreator) CreateBody(method, wsmanClass string, data interface{}) string {
	var str strings.Builder

	str.WriteString("<Body>")

	if data != nil {
		w.namespaceMe(data, wsmanClass)

		xmlString, err := xml.Marshal(data)
		if err != nil {
			log.Println(err)
		}

		str.WriteString(string(xmlString))
	} else {
		str.WriteString(fmt.Sprintf(`<h:%s xmlns:h="%s%s">`, method, w.ResourceURIBase, wsmanClass))
		str.WriteString(fmt.Sprintf(`</h:%s>`, method))
	}

	str.WriteString("</Body>")

	return str.String()
}

// createSelector creates a WSMAN string based on Selector Set information provided.
// It can be used in the header or body.
// selectorSet is the selector data being passed in. It could take many forms depending on the WSMAN call.
func (w *WSManMessageCreator) createSelector(selectorSet []Selector) string {
	var selectors strings.Builder

	if len(selectorSet) == 0 {
		return ""
	}

	selectors.WriteString("<w:SelectorSet>")

	for _, selector := range selectorSet {
		selectors.WriteString(fmt.Sprintf(`<w:Selector Name=%q>%s</w:Selector>`, selector.Name, selector.Value))
	}

	selectors.WriteString("</w:SelectorSet>")

	return selectors.String()
}

// createSelectorObjectForBody creates an object for the body using the given selector.
func (w *WSManMessageCreator) CreateSelectorObjectForBody(selector Selector) map[string]interface{} {
	obj := map[string]interface{}{
		"Selector": []map[string]interface{}{
			{
				"_": selector.Value,
				"$": map[string]string{
					"Name": selector.Name,
				},
			},
		},
	}

	return obj
}

func createCommonBodyPull(enumerationContext string, maxElements, maxCharacters int) string {
	if maxElements == 0 {
		maxElements = 999
	}

	if maxCharacters == 0 {
		maxCharacters = 99999
	}

	return fmt.Sprintf(`<Body><Pull xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration"><EnumerationContext>%s</EnumerationContext><MaxElements>%d</MaxElements><MaxCharacters>%d</MaxCharacters></Pull></Body>`, enumerationContext, maxElements, maxCharacters)
}

func (w WSManMessageCreator) createCommonBodyCreateOrPut(wsmanClass string, data interface{}) string {
	return w.CreateBody(wsmanClass, wsmanClass, data)
}

func createCommonBodyRequestStateChange(input string, requestedState int) string {
	return fmt.Sprintf(`<Body><h:RequestStateChange_INPUT xmlns:h=%q><h:RequestedState>%d</h:RequestedState></h:RequestStateChange_INPUT></Body>`, input, requestedState)
}
