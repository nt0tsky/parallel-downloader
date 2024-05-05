package downloader

import (
	"fmt"
	"github.com/nt0tsky/parallel-downloader/internal/downloader/internal"
	"github.com/nt0tsky/parallel-downloader/internal/storage"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func NewDownloader(logger *log.Logger, storage storage.Storage, threads int, url string) *Downloader {
	return &Downloader{
		logger,
		storage,
		threads,
		url,
	}
}

func (d *Downloader) Start() error {
	meta, err := d.getFileMetadata()
	if err != nil {
		d.logger.Println("Failed to get file metadata:", err)
		return err
	}

	fileName := internal.CreateFileName(d.url, meta.etag, meta.contentLength)
	file, err := d.storage.CreateFile(fileName, meta.contentLength)
	defer file.Close()
	if err != nil {
		d.logger.Printf("Failed to create file for download: %s", err)
		return err
	}

	err = d.start(file, meta)
	if err != nil {
		d.logger.Printf("Failed to start download: %s", err)
	}

	d.logger.Printf("Download finished")
	return nil
}

func (d *Downloader) start(file storage.FileHandle, meta *Metadata) error {
	wg := sync.WaitGroup{}
	wg.Add(d.threads)
	chunkSize := meta.contentLength / int64(d.threads)

	var startByte int64
	for range d.threads {
		endByte := startByte + chunkSize
		go func() {
			err := d.downloadChunk(file, startByte, endByte, &wg)
			if err != nil {
				d.logger.Printf("Failed to download chunk: %s", err)
			}
		}()
		startByte = endByte
	}
	wg.Wait()

	return nil
}

func (d *Downloader) downloadChunk(file storage.FileHandle, startByte int64, endByte int64, wg *sync.WaitGroup) error {
	defer wg.Done()
	bytes, err := d.fetchChunk(startByte, endByte)
	if err != nil {
		return err
	}

	err = d.storage.AddToFile(file, bytes, startByte)
	if err != nil {
		d.logger.Printf("Failed to write chunk: %s", err)
		return err
	}

	return nil
}

func (d *Downloader) fetchChunk(startByte int64, endByte int64) ([]byte, error) {
	d.logger.Printf("Downloading chunk bytes=%d-%d", startByte, endByte)
	req, err := http.NewRequest(http.MethodGet, d.url, nil)
	if err != nil {
		d.logger.Printf("Failed to create request: %s", err)
		return nil, err
	}

	req.Header.Add("Range", fmt.Sprintf("bytes=%d-%d", startByte, endByte-1))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		d.logger.Printf("Failed to download chunk: %s", err)
		return nil, err
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		d.logger.Printf("Failed to read chunk: %s", err)
		return nil, err
	}
	d.logger.Printf("Chunk downloaded bytes=%d-%d", startByte, endByte)

	return bytes, nil
}

func (d *Downloader) getFileMetadata() (*Metadata, error) {
	resp, err := http.Head(d.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	contentLength, err := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
	if err != nil {
		return nil, err
	}
	etag := resp.Header.Get("ETag")

	return &Metadata{
		contentLength,
		etag,
	}, nil
}
