package storage

import (
	"io"
	"log"
	"os"
)

type MemoryStorage struct {
	logger            *log.Logger
	destinationFolder string
}

func NewMemoryStorage(logger *log.Logger, destinationFolder string) Storage {
	return &MemoryStorage{logger, destinationFolder}
}

func (s *MemoryStorage) CreateFile(fileName string, fileSize int64) (FileHandle, error) {
	fullPath := s.destinationFolder + "/" + fileName
	file, err := os.Create(fullPath)
	if err != nil {
		return nil, err
	}

	s.logger.Printf("Created file %s", fileName)
	file.Seek(fileSize-1, io.SeekStart)
	file.Write([]byte{0})
	return file, nil
}

func (s *MemoryStorage) AddToFile(file FileHandle, bytes []byte, startByte int64) error {
	fileOS := file.(*os.File)

	_, err := fileOS.WriteAt(bytes, startByte)
	if err != nil {
		s.logger.Printf("Failed to write chunk: %s", err)
		return err
	}

	return nil
}
