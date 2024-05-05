package internal

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateFileName(t *testing.T) {
	testCases := []struct {
		url            string
		etag           string
		contentLength  int64
		expectedOutput string
	}{
		{
			url:            "https://example.com/file",
			etag:           "abc123",
			contentLength:  1000,
			expectedOutput: "acff6b2e542c0b151548dffe055f652ddceff7a48618319c976b41abb72683b1",
		},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("%s with length %d", test.url, test.contentLength), func(t *testing.T) {
			fileName := CreateFileName(test.url, test.etag, test.contentLength)

			assert.Equal(t, test.expectedOutput, fileName)
		})
	}
}
