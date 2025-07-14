/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wsmantesting

import "fmt"

const (
	XMLHeader               = `<?xml version="1.0" encoding="utf-8"?>`
	Envelope                = `<Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd" xmlns="http://www.w3.org/2003/05/soap-envelope"><Header><a:Action>`
	EnumerationContext      = `AC070000-0000-0000-0000-000000000000`
	OperationTimeout        = `PT60S`
	Get                     = "http://schemas.xmlsoap.org/ws/2004/09/transfer/Get"
	Enumerate               = "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Enumerate"
	Pull                    = "http://schemas.xmlsoap.org/ws/2004/09/enumeration/Pull"
	Delete                  = "http://schemas.xmlsoap.org/ws/2004/09/transfer/Delete"
	Put                     = "http://schemas.xmlsoap.org/ws/2004/09/transfer/Put"
	Create                  = "http://schemas.xmlsoap.org/ws/2004/09/transfer/Create"
	EnumerateBody           = "<Enumerate xmlns=\"http://schemas.xmlsoap.org/ws/2004/09/enumeration\" />"
	SetBootConfigRole       = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootService/SetBootConfigRole"
	RequestStateChange      = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_BootService/RequestStateChange"
	ServerCertificateIssuer = `serverCertificateIssuer`
	ClientCertificate       = `clientCertificate`
	DigestRealm             = "Digest:Realm"
	AdminPassEncryptionType = 2
	AdminPassword           = `bebb3497d69b544c732651365cc3462d`
	MCNonce                 = `ZxxE0cFy590zDBIR39q6QU6iuII=`
	SigningAlgorithm        = 2
	DigitalSignature        = `T0NvoR7RUkOpVULIcNL0VhpEK5rO3j5/TBpN82q1YgPM5sRBxqymu7fKBgAGGN49oD8xsqW4X0SWxjuB3q/TLHjNJJNxoHHlXZnb77HTwfXHp59E/TM10UvOX96qEgKU5Mp+8/IE9LnYxC1ajQostSRA/X+HA5F6kRctLiCK+ViWUCk4sAtPzHhhHSTB/98KDWuacPepScSpref532hpD2/g43nD3Wg0SjmOMExPLMMnijWE9KDkxE00+Bos28DD3Yclj4BMhkoXDw6k4EcTWKbGhtF/9meXXmSPwRmXEaWe8COIDrQks1mpyLblYu8yHHnUjhssdcCQHtAOu7t0RA==`
	SetCertificates         = "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_IEEE8021xSettings/SetCertificates"
	AddNextCertInChain      = "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService/AddNextCertInChain"
	AdminSetup              = "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService/AdminSetup"
	UpgradeClientToAdmin    = "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService/UpgradeClientToAdmin"
	Setup                   = "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HostBasedSetupService/Setup"
	AddProxyAccessPoint     = "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_HTTPProxyService/AddProxyAccessPoint"
	SendOptInCode           = "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService/SendOptInCode"
	StartOptIn              = "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService/StartOptIn"
	CancelOptIn             = "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_OptInService/CancelOptIn"
	TrustedRootCert         = "MIIEOzCCAqOgAwIBAgIDAZiFMA0GCSqGSIb3DQEBDAUAMD0xFzAVBgNVBAMTDk1QU1Jvb3QtNjE0ZDg4MRAwDgYDVQQKEwd1bmtub3duMRAwDgYDVQQGEwd1bmtub3duMCAXDTIwMDgyNTE4MzMzN1oYDzIwNTEwODI1MTgzMzM3WjA9MRcwFQYDVQQDEw5NUFNSb290LTYxNGQ4ODEQMA4GA1UEChMHdW5rbm93bjEQMA4GA1UEBhMHdW5rbm93bjCCAaIwDQYJKoZIhvcNAQEBBQADggGPADCCAYoCggGBAOi1jx9L8DG6gBPxd9gmJ6vqQC/F/TBMTJvb3ZAuRbDxUKnxZk3PafyNM6fO8QTL4qZVhvyGEZaIzVePrdJj31aZ93mNY2TJee3/DLRsJUIZHGFufBvi8pgQL+JjE9JmFD5/S2yciHIEVpKmXo1CbGmZGsnb8NRjaQVwB94pI1mg8JFMxyKzU/cUoCBfI+wmeMgBVdOJPNpH2zjC/GxwEFNQaxGe9GHmYbwoeiDeMPo75E/o+Gw6kJm429cuhJBC3KqHevAJj9V2nSUvoO0oxKqzLVkUYcjHEGYjxIvP6a6uo7x9llwfshJsBZ3PE5hucNdWS3dY3GeCqOwcaAQQIj2jULpZ/KlgVAdBK/o5QjE+IIQXCVK9USvktGzz7I5oH98zy8jCFStbGM7PQCo+DEnHn/SANmVbcy3hjzrXC8zf5dvmKiUb2eKnpv+z3FHsi64sVwFqBArB2ipcTM/qv4nEM6uLW1t+7+NB0OyaBmLktJrpb6af7z/EW1QuPIfTcQIDAQABo0IwQDAMBgNVHRMEBTADAQH/MBEGCWCGSAGG+EIBAQQEAwIABzAdBgNVHQ4EFgQUYU2IeTFqWXI1rG+JqZq8eVDO/LMwDQYJKoZIhvcNAQEMBQADggGBANoKIsFOn8/Lrb98DjOP+LUeopoU9KQ70ndreNqchrkPmM61V9IdD9OZiLr/7OY/rLGZwNvkhQYRPUa842Mqjfpr4YcV6HC0j6Zg0lcpxQ5eGGBkLb/teBcboi3sZcJvbCFUW2DJjhy7uqYxzE4eqSsKx5fEjp/wa6oNzNrgWRXyxQlaOo42RjXnOXS7sB0jPrgO0FClL1Xzif06kFHzzyJCVUqzNEJv0ynLgkpzCVdUUfoMM1RcKc3xJes5C0zg64ugj2R9e4VwJfn9W3+rlYS1So1q1jL8w+3qOM7lXyvr8Bdgc5BMvrOvHxzdOnpZmUEJkbKty62e8fYKN+WP7BrpxnzFQSzczX5S0uN4rn0rLO4wxVf2rtnTqIhKKYTsPMRBVEjpbRT1smzPPdINKu5l/Rz/zZS0b5I4yKJrkTYNgoPC/QSq8A9uXZxxQvj6x1bWZJVWywmaqYolEp8NaVHd+JYnlTmr4XpMHm01TPi1laowtY3ZepnKm8I55Ly0JA=="
	CIMResourceURIBase      = "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	AMTResourceURIBase      = "http://intel.com/wbem/wscim/1/amt-schema/1/"
	IPSResourceURIBase      = "http://intel.com/wbem/wscim/1/ips-schema/1/"
	CurrentMessageEnumerate = "Enumerate"
	CurrentMessagePull      = "Pull"
	CurrentMessageGet       = "Get"
	CurrentMessagePut       = "Put"
	CurrentMessageCreate    = "Create"
	CurrentMessageDelete    = "Delete"
	CurrentMessageError     = "Error"
)

var PullBody = fmt.Sprintf(`<Pull xmlns="http://schemas.xmlsoap.org/ws/2004/09/enumeration"><EnumerationContext>%s</EnumerationContext><MaxElements>999</MaxElements><MaxCharacters>99999</MaxCharacters></Pull>`, EnumerationContext)

var ExpectedResponse = func(messageID int, resourceURIBase, method, action, extraHeader, body string) string {
	return fmt.Sprintf(`%s%s%s</a:Action><a:To>/wsman</a:To><w:ResourceURI>%s%s</w:ResourceURI><a:MessageID>%d</a:MessageID><a:ReplyTo><a:Address>http://schemas.xmlsoap.org/ws/2004/08/addressing/role/anonymous</a:Address></a:ReplyTo><w:OperationTimeout>%s</w:OperationTimeout>%s</Header><Body>%s</Body></Envelope>`, XMLHeader, Envelope, action, resourceURIBase, method, messageID, OperationTimeout, extraHeader, body)
}
