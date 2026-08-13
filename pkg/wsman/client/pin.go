/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package client

import (
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"errors"
	"fmt"
)

// errCertificatePinning is returned when the peer's leaf certificate does not
// match the configured pinned SHA-256 fingerprint.
var errCertificatePinning = errors.New("certificate pinning failed")

// pinnedCertVerifier returns a tls.Config.VerifyPeerCertificate callback that
// compares the peer's leaf certificate against pinnedCert, a hex-encoded
// SHA-256 fingerprint of the DER-encoded certificate.
//
// Only the leaf certificate (rawCerts[0]) is compared; any other certificates
// in the presented chain are ignored.
func pinnedCertVerifier(pinnedCert string) func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
	return func(rawCerts [][]byte, _ [][]*x509.Certificate) error {
		if len(rawCerts) == 0 {
			return fmt.Errorf("%w: no certificate presented", errCertificatePinning)
		}

		cert, err := x509.ParseCertificate(rawCerts[0])
		if err != nil {
			return err
		}

		// Compare the leaf certificate with the pinned certificate.
		sha256Fingerprint := sha256.Sum256(cert.Raw)
		if hex.EncodeToString(sha256Fingerprint[:]) == pinnedCert {
			return nil // Success: the leaf certificate matches the pinned certificate
		}

		return errCertificatePinning
	}
}
