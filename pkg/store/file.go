package store

import (
	"encoding/json"
	"os"
)


type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}
 

type Type string 

const ( 
	FileType Type = "file"
)

type FileStore struct {
	FileName string
}

func New(store Type, filename string) Store {
	switch store {
	case FileType :
		return &FileStore{FileName: filename}
	}
	return nil
}

func (store FileStore) Read(data interface{}) error {
	file, err := os.ReadFile(store.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}

func (store FileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "\t", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(store.FileName, fileData, 0644) 
}



