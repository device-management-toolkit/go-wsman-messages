/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package client

import (
	"crypto/md5"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
)

type AuthChallenge struct {
	Username   string
	Password   string
	Realm      string
	CSRFToken  string
	Domain     string
	Nonce      string
	Opaque     string
	Stale      string
	Algorithm  string
	Qop        string
	CNonce     string
	NonceCount int
}

func hashWithMD5(data string) string {
	md5Hash := md5.New()

	_, err := io.WriteString(md5Hash, data)
	if err != nil {
		log.Println("failed to write string to md5 hash")
	}

	return fmt.Sprintf("%x", md5Hash.Sum(nil))
}

func hashWithHash(secret, data string) string {
	return hashWithMD5(fmt.Sprintf("%s:%s", secret, data))
}

func (c *AuthChallenge) HashCredentials() string {
	return hashWithMD5(fmt.Sprintf("%s:%s:%s", c.Username, c.Realm, c.Password))
}

func (c *AuthChallenge) hashURI(method, uri string) string {
	return hashWithMD5(fmt.Sprintf("%s:%s", method, uri))
}

func (c *AuthChallenge) GetFormattedNonceData(nonceData string) string {
	nonceData = fmt.Sprintf("%s:%08x:%s:%s", nonceData, c.NonceCount, c.CNonce, c.Qop)

	return nonceData
}

func (c *AuthChallenge) ComputeDigestHash(method, uri, nonceData string) string {
	hashedCredentials := c.HashCredentials()
	hashedURI := c.hashURI(method, uri)
	response := hashWithHash(hashedCredentials, fmt.Sprintf("%s:%s", nonceData, hashedURI))

	return response
}

func (c *AuthChallenge) response(method, uri, cnonce string) (string, error) {
	c.NonceCount++

	if strings.Contains(c.Qop, "auth") || c.Qop == "" {
		nonceData := c.Nonce

		if strings.Contains(c.Qop, "auth") {
			if cnonce != "" {
				c.CNonce = cnonce
			} else {
				b := make([]byte, 8)
				if _, err := io.ReadFull(rand.Reader, b); err != nil {
					errRandRead := errors.New("failed to generate random bytes")

					return "", fmt.Errorf("%w: %w", errRandRead, err)
				}

				c.CNonce = fmt.Sprintf("%x", b)[:6]
			}

			c.Qop = "auth"
			nonceData = c.GetFormattedNonceData(nonceData)
		}

		return c.ComputeDigestHash(method, uri, nonceData), nil
	}

	errNotImplemented := errors.New("not implemented")

	return "", fmt.Errorf("%w", errNotImplemented)
}

func (c *AuthChallenge) authorize(method, uri string) (string, error) {
	if !strings.Contains(c.Qop, "auth") && c.Qop != "" {
		errQopNotImplemented := errors.New("qop not implemented")

		return "", fmt.Errorf("%w", errQopNotImplemented)
	}

	response, err := c.response(method, uri, "")
	if err != nil {
		return "", err
	}

	var sb strings.Builder

	sb.WriteString(`Digest username="`)
	sb.WriteString(c.Username)
	sb.WriteString(`",realm="`)
	sb.WriteString(c.Realm)
	sb.WriteString(`",nonce="`)
	sb.WriteString(c.Nonce)
	sb.WriteString(`",uri="`)
	sb.WriteString(uri)
	sb.WriteString(`",response="`)
	sb.WriteString(response)
	sb.WriteString(`"`)

	// if c.Algorithm != "" {
	// 	sb.WriteString(`, algorithm="`)
	// 	sb.WriteString(c.Algorithm)
	// 	sb.WriteString(`"`)
	// }

	if c.Opaque != "" {
		sb.WriteString(`,opaque="`)
		sb.WriteString(c.Opaque)
		sb.WriteString(`"`)
	}

	if c.Qop != "" {
		sb.WriteString(`,qop="`)
		sb.WriteString(c.Qop)
		sb.WriteString(`",nc="`)
		fmt.Fprintf(&sb, "%08x", c.NonceCount)
		sb.WriteString(`",cnonce="`)
		sb.WriteString(c.CNonce)
		sb.WriteString(`"`)
	}

	return sb.String(), nil
}

func (c *AuthChallenge) parseChallenge(input string) error {
	errBadChallenge := errors.New("bad challenge")

	const ws = " \n\r\t"

	const qs = "\""

	s := strings.Trim(input, ws)
	if !strings.HasPrefix(s, "Digest ") {
		return fmt.Errorf("%w, missing digest prefix: %s", errBadChallenge, input)
	}

	s = strings.Trim(s[7:], ws)
	sl := strings.Split(s, "\",")
	c.Algorithm = "MD5"

	var r []string

	for _, elem := range sl {
		r = strings.SplitN(elem, "=", 2)
		if len(r) != 2 {
			return fmt.Errorf("%w, malformed token: %s", errBadChallenge, elem)
		}

		key := strings.TrimSpace(r[0])

		value := strings.Trim(strings.TrimSpace(r[1]), qs)

		switch key {
		case "realm":
			c.Realm = value
		case "domain":
			c.Domain = value
		case "nonce":
			c.Nonce = value
		case "opaque":
			c.Opaque = value
		case "stale":
			c.Stale = value
		case "algorithm":
			c.Algorithm = value
		case "qop":
			c.Qop = value
		default:
			return fmt.Errorf("%w, unexpected token: %s", errBadChallenge, sl)
		}
	}

	return nil
}
