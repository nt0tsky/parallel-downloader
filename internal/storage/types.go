package storage

type FileHandle interface {
	Close() error
}

type Storage interface {
	CreateFile(fileName string, fileSize int64) (FileHandle, error)
	AddToFile(file FileHandle, bytes []byte, startByte int64) error
}
