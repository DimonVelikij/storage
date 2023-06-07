package file

import "github.com/google/uuid"

type File struct {
	Id   uuid.UUID
	Name string
	Data []byte
}

func NewFile(filename string, blob []byte) (*File, error) {
	fileId, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	return &File{
		Id:   fileId,
		Name: filename,
		Data: blob,
	}, nil
}
