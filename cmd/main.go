package main

import (
	"flag"
	"github.com/nt0tsky/parallel-downloader/internal/downloader"
	"github.com/nt0tsky/parallel-downloader/internal/preferences"
	"github.com/nt0tsky/parallel-downloader/internal/storage"
	"log"
	"os"
)

func main() {
	var threads int
	var url, destinationFolder string

	flag.IntVar(&threads, "threads", 3, "limit the number of downloading goroutines")
	flag.StringVar(&url, "url", "", "URL")
	flag.StringVar(&destinationFolder, "destinationFolder", "", "Destination folder")
	flag.Parse()

	logger := log.New(os.Stdout, "parallel-downloader", log.Ldate|log.Ltime)

	p, err := preferences.NewPreferences(threads, url, destinationFolder)
	if err != nil {
		logger.Fatalf("ERR: %v", err)
		return
	}
	logger.Printf("preferences: %#v", p)

	s := storage.NewMemoryStorage(logger, p.DestinationFolder)
	d := downloader.NewDownloader(logger, s, p.Threads, p.Url)

	err = d.Start()
	if err != nil {
		logger.Printf("ERR: %#v", err)
	}
}
