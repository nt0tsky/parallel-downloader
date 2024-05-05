package downloader

import (
	"github.com/nt0tsky/parallel-downloader/internal/storage"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type MockFileHandle struct{}

func (m *MockFileHandle) Close() error {
	return nil
}

type MockStorage struct{}

func (s *MockStorage) CreateFile(fileName string, size int64) (storage.FileHandle, error) {
	return &MockFileHandle{}, nil
}

func (s *MockStorage) AddToFile(file storage.FileHandle, bytes []byte, startByte int64) error {
	return nil
}

func TestDownloader_Start(t *testing.T) {
	logger := log.New(os.Stdout, "test ", log.LstdFlags)
	storage := &MockStorage{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "100")
		w.Header().Set("ETag", "abc123")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("dummy content"))
	}))
	defer server.Close()

	downloader := NewDownloader(logger, storage, 2, server.URL)
	err := downloader.Start()

	assert.NoError(t, err)
}
