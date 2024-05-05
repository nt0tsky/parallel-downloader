package preferences

import (
	"errors"
)

type Preferences struct {
	Threads           int
	Url               string
	DestinationFolder string
}

func NewPreferences(threads int, url string, destinationFolder string) (*Preferences, error) {
	if len(url) == 0 {
		return nil, errors.New("invalid url")
	}

	if len(destinationFolder) == 0 {
		return nil, errors.New("invalid DestinationFolder")
	}

	return &Preferences{
		threads,
		url,
		destinationFolder,
	}, nil
}
