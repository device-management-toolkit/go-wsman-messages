/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package message

type WSManErrors string

const (
	HEADER                             WSManErrors = "missing header"
	BODY                               WSManErrors = "missing body"
	ACTION                             WSManErrors = "missing action"
	MESSAGE_ID                         WSManErrors = "missing messageId"
	RESOURCE_URI                       WSManErrors = "missing resourceUri"
	ENUMERATION_CONTEXT                WSManErrors = "missing enumerationContext"
	UNSUPPORTED_METHOD                 WSManErrors = "unsupported method"
	INPUT                              WSManErrors = "missing input"
	REQUESTED_STATE                    WSManErrors = "missing requestedState"
	SELECTOR                           WSManErrors = "missing selector"
	ROLE                               WSManErrors = "missing role"
	REQUESTED_POWER_STATE_CHANGE       WSManErrors = "missing powerState"
	ADMIN_PASS_ENCRYPTION_TYPE         WSManErrors = "missing adminPassEncryptionType"
	ADMIN_PASSWORD                     WSManErrors = "missing adminPassword"
	ETHERNET_PORT_OBJECT               WSManErrors = "missing ethernetPortObject"
	ENVIRONMENT_DETECTION_SETTING_DATA WSManErrors = "missing environmentDetectionSettingData"
	CERTIFICATE_BLOB                   WSManErrors = "missing certificateBlob"
	MP_SERVER                          WSManErrors = "missing mpServer"
	REMOTE_ACCESS_POLICY_RULE          WSManErrors = "missing remoteAccessPolicyRule"
	BOOT_SETTING_DATA                  WSManErrors = "missing bootSettingData"
	ADD_ALARM_DATA                     WSManErrors = "missing alarmClockOccurrence"
	IEEE8021X_SETTINGS                 WSManErrors = "missing ieee8021xSettings"
	OPT_IN_SERVICE_RESPONSE            WSManErrors = "missing OptInServiceResponse"
	OPT_IN_CODE                        WSManErrors = "missing OptInCode"
	KEY_PAIR                           WSManErrors = "missing KeyAlgorithm and/or KeyLength"
	DATA                               WSManErrors = "missing data"
	NONCE                              WSManErrors = "missing nonce"
	SIGNING_ALGORITHM                  WSManErrors = "missing signingAlgorithm"
	DIGITAL_SIGNATURE                  WSManErrors = "missing digitalSignature"
	IS_LEAF                            WSManErrors = "missing isLeaf"
	IS_ROOT                            WSManErrors = "missing isRoot"
	TLS_CREDENTIAL_CONTEXT             WSManErrors = "missing tlsCredentialContext"
	GENERAL_SETTINGS                   WSManErrors = "missing generalSettings"
	PASSWORD                           WSManErrors = "missing password"
	PKCS10_REQUEST                     WSManErrors = "missing PKCS10Request"
	USERNAME                           WSManErrors = "missing username"
	DIGEST_PASSWORD                    WSManErrors = "missing digestPassword"
	INSTANCE_ID                        WSManErrors = "missing InstanceID"
	MISSING_USER_ACL_ENTRY_INFORMATION WSManErrors = "Digest username and password or Kerberos SID is required"
	USERNAME_TOO_LONG                  WSManErrors = "Username is too long"
)

const (
	BaseActionsEnumerate = "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate"
	BaseActionsPull      = "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull"
	BaseActionsGet       = "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get"
	BaseActionsPut       = "http://schemas.xmlsoap.org/ws/2004/09/transfer/Put"
	BaseActionsCreate    = "http://schemas.xmlsoap.org/ws/2004/09/transfer/Create"
	BaseActionsDelete    = "http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete"
	DeleteBody           = "<Body></Body>"
	EnumerateBody        = "<Body><Enumerate xmlns=\"http://schemas.xmlsoap.org/ws/2004/09/enumeration\" /></Body>"
	GetBody              = "<Body></Body>"
	AMTSchema            = "http://intel.com/wbem/wscim/1/amt-schema/1/"
	CIMSchema            = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	IPSSchema            = "http://intel.com/wbem/wscim/1/ips-schema/1/"
	XMLBodySpace         = "http://www.w3.org/2003/05/soap-envelope"
	XMLPullResponseSpace = "http://schemas.xmlsoap.org/ws/2004/09/enumeration"
)
