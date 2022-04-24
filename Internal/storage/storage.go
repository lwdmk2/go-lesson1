package storage

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/lwdmk2/go-lesson1/Internal/file"
)

type Storage struct {
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, blob)
	if err != nil {
		return nil, err
	}
	s.files[newFile.Id] = newFile

	return newFile, nil

}

func (s *Storage) GetById(fileId uuid.UUID) (*file.File, error) {
	foundFile, ok := s.files[fileId]
	if !ok {
		return nil, fmt.Errorf("file %v not found", fileId)
	}

	return foundFile, nil
}
