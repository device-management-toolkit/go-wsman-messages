//go:build integration

/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package remoteaccess_test

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/device-management-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/amt/remoteaccess"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestCaptureFirmwareXMLRemoteAccessCapabilities(t *testing.T) {
	params, err := parametersFromEnv()
	if err != nil {
		t.Skipf("capture skipped: %v", err)
	}

	wsmanClient := client.NewWsman(params)
	wsmanMessageCreator := message.NewWSManMessageCreator(wsmantesting.AMTResourceURIBase)
	capabilities := remoteaccess.NewCapabilitiesWithClient(wsmanMessageCreator, wsmanClient)
	defer wsmanClient.CloseConnection()

	enumerateResponse, err := capabilities.Enumerate()
	if err != nil {
		t.Fatalf("enumerate error: %v", err)
	}

	pullResponse, err := capabilities.Pull(enumerateResponse.Body.EnumerateResponse.EnumerationContext)
	if err != nil {
		t.Fatalf("pull error: %v", err)
	}

	getResponse, err := capabilities.Get()
	if err != nil {
		t.Fatalf("get error: %v", err)
	}

	_, thisFile, _, _ := runtime.Caller(0)
	defaultOutputDir := filepath.Join(
		filepath.Dir(thisFile),
		"..", "..", "..", "wsmantesting", "responses", "amt", "remoteaccess", "capabilities",
	)
	outputDir := envOrDefault("OUTPUT_DIR", defaultOutputDir)

	if err := writeFWXML(outputDir, getResponse.XMLOutput, enumerateResponse.XMLOutput, pullResponse.XMLOutput); err != nil {
		t.Fatalf("write error: %v", err)
	}

	t.Logf("saved get.xml, enumerate.xml, and pull.xml to %s", outputDir)
}

func parametersFromEnv() (client.Parameters, error) {
	target := os.Getenv("AMT_TARGET")
	username := os.Getenv("AMT_USERNAME")
	password := os.Getenv("AMT_PASSWORD")

	if target == "" || username == "" || password == "" {
		return client.Parameters{}, fmt.Errorf("set AMT_TARGET, AMT_USERNAME, and AMT_PASSWORD")
	}

	return client.Parameters{
		Target:            target,
		Username:          username,
		Password:          password,
		UseDigest:         true,
		UseTLS:            true,
		SelfSignedAllowed: true,
		LogAMTMessages:    true,
	}, nil
}

func writeFWXML(outputDir, getXML, enumerateXML, pullXML string) error {
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return err
	}

	files := map[string]string{
		"get.xml":       getXML,
		"enumerate.xml": enumerateXML,
		"pull.xml":      pullXML,
	}

	for name, contents := range files {
		path := filepath.Join(outputDir, name)
		if err := os.WriteFile(path, []byte(contents), 0o600); err != nil {
			return err
		}
	}

	return nil
}

func envOrDefault(name, fallback string) string {
	value := os.Getenv(name)
	if value == "" {
		return fallback
	}

	return value
}
