package preferences

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPreferences(t *testing.T) {
	tests := []struct {
		name              string
		threads           int
		url               string
		destinationFolder string
		expectedError     error
	}{
		{
			name:              "ValidArguments",
			threads:           5,
			url:               "http://example.com",
			destinationFolder: "/path/to/folder",
			expectedError:     nil,
		},
		{
			name:              "MissingURL",
			threads:           5,
			url:               "",
			destinationFolder: "/path/to/folder",
			expectedError:     errors.New("invalid url"),
		},
		{
			name:              "MissingDestinationFolder",
			threads:           5,
			url:               "http://example.com",
			destinationFolder: "",
			expectedError:     errors.New("invalid DestinationFolder"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pref, err := NewPreferences(tt.threads, tt.url, tt.destinationFolder)

			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
				assert.Nil(t, pref)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, pref)
				assert.Equal(t, tt.threads, pref.Threads)
				assert.Equal(t, tt.url, pref.Url)
				assert.Equal(t, tt.destinationFolder, pref.DestinationFolder)
			}
		})
	}
}
