/*********************************************************************
 * Copyright (c) Intel Corporation 2024
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package client

import "testing"

func TestIsCompleteHTTPResponse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		data     string
		expected bool
	}{
		{
			name:     "incomplete headers",
			data:     "HTTP/1.1 200 OK\r\nContent-Length: 5\r\n",
			expected: false,
		},
		{
			name:     "content-length complete",
			data:     "HTTP/1.1 200 OK\r\nContent-Length: 5\r\n\r\nhello",
			expected: true,
		},
		{
			name:     "content-length incomplete",
			data:     "HTTP/1.1 200 OK\r\nContent-Length: 10\r\n\r\nhello",
			expected: false,
		},
		{
			name:     "content-length zero",
			data:     "HTTP/1.1 200 OK\r\nContent-Length: 0\r\n\r\n",
			expected: true,
		},
		{
			name:     "chunked complete",
			data:     "HTTP/1.1 200 OK\r\nTransfer-Encoding: chunked\r\n\r\n5\r\nhello\r\n0\r\n\r\n",
			expected: true,
		},
		{
			name:     "chunked incomplete",
			data:     "HTTP/1.1 200 OK\r\nTransfer-Encoding: chunked\r\n\r\n5\r\nhello\r\n",
			expected: false,
		},
		{
			name:     "204 no content",
			data:     "HTTP/1.1 204 No Content\r\n\r\n",
			expected: true,
		},
		{
			name:     "304 not modified",
			data:     "HTTP/1.1 304 Not Modified\r\n\r\n",
			expected: true,
		},
		{
			name:     "100 continue",
			data:     "HTTP/1.1 100 Continue\r\n\r\n",
			expected: true,
		},
		{
			name:     "no content-length or chunked",
			data:     "HTTP/1.1 200 OK\r\nServer: AMT\r\n\r\nsome data",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := isCompleteHTTPResponse([]byte(tt.data))
			if got != tt.expected {
				t.Errorf("isCompleteHTTPResponse() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestParseHTTPStatusCode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"200 OK", "http/1.1 200 ok\r\ncontent-length: 0", 200},
		{"204 No Content", "http/1.1 204 no content\r\n", 204},
		{"304 Not Modified", "http/1.1 304 not modified\r\n", 304},
		{"100 Continue", "http/1.1 100 continue\r\n", 100},
		{"not http", "garbage data", 0},
		{"empty", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := parseHTTPStatusCode(tt.input)
			if got != tt.expected {
				t.Errorf("parseHTTPStatusCode() = %v, want %v", got, tt.expected)
			}
		})
	}
}
