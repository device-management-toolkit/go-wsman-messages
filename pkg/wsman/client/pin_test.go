/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package client

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/hex"
	"errors"
	"math/big"
	"testing"
)

// newSelfSignedDER creates a throwaway self-signed certificate and returns its
// DER bytes plus its hex-encoded SHA-256 fingerprint (the pin format).
func newSelfSignedDER(t *testing.T, commonName string) (der []byte, fingerprint string) {
	t.Helper()

	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("generate key: %v", err)
	}

	tmpl := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: commonName},
	}

	der, err = x509.CreateCertificate(rand.Reader, tmpl, tmpl, &key.PublicKey, key)
	if err != nil {
		t.Fatalf("create certificate: %v", err)
	}

	sum := sha256.Sum256(der)

	return der, hex.EncodeToString(sum[:])
}

func TestPinnedCertVerifier(t *testing.T) {
	pinnedDER, pin := newSelfSignedDER(t, "pinned")
	otherDER, _ := newSelfSignedDER(t, "other")

	tests := []struct {
		name     string
		pin      string
		rawCerts [][]byte
		wantErr  bool
	}{
		{
			name:     "leaf matches pin",
			pin:      pin,
			rawCerts: [][]byte{pinnedDER},
			wantErr:  false,
		},
		{
			// Only the leaf is compared: a matching certificate that appears
			// later in the presented chain does not satisfy the pin.
			name:     "pinned cert only as a non-leaf chain entry does not match",
			pin:      pin,
			rawCerts: [][]byte{otherDER, pinnedDER},
			wantErr:  true,
		},
		{
			name:     "leaf does not match pin",
			pin:      pin,
			rawCerts: [][]byte{otherDER},
			wantErr:  true,
		},
		{
			name:     "no certificate presented",
			pin:      pin,
			rawCerts: nil,
			wantErr:  true,
		},
		{
			name:     "unparseable leaf certificate",
			pin:      pin,
			rawCerts: [][]byte{[]byte("not a certificate")},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := pinnedCertVerifier(tt.pin)(tt.rawCerts, nil)
			if (err != nil) != tt.wantErr {
				t.Fatalf("pinnedCertVerifier() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPinnedCertVerifier_ErrorIsSentinel(t *testing.T) {
	_, pin := newSelfSignedDER(t, "pinned")
	otherDER, _ := newSelfSignedDER(t, "other")

	err := pinnedCertVerifier(pin)([][]byte{otherDER}, nil)
	if !errors.Is(err, errCertificatePinning) {
		t.Errorf("expected errCertificatePinning, got %v", err)
	}
}
