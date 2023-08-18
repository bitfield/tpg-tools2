package kv

import (
	"encoding/gob"
	"errors"
	"io/fs"
	"os"
)

type store struct {
	path string
	data map[string]string
}

func OpenStore(path string) (*store, error) {
	s := &store{
		path: path,
		data: map[string]string{},
	}
	f, err := os.Open(path)
	if errors.Is(err, fs.ErrNotExist) {
		return s, nil
	}
	if err != nil {
		return nil, err
	}
	defer f.Close()
	err = gob.NewDecoder(f).Decode(&s.data)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *store) Set(k, v string) {
	s.data[k] = v
}

func (s store) Get(k string) (string, bool) {
	v, ok := s.data[k]
	return v, ok
}

func (s store) Save() error {
	f, err := os.Create(s.path)
	if err != nil {
		return err
	}
	defer f.Close()
	return gob.NewEncoder(f).Encode(s.data)
}
