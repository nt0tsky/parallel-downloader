package downloader

import (
	"github.com/nt0tsky/parallel-downloader/internal/storage"
	"log"
)

type Metadata struct {
	contentLength int64
	etag          string
}

type Downloader struct {
	logger  *log.Logger
	storage storage.Storage
	threads int
	url     string
}
