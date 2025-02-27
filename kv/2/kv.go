package kv

import (
	"encoding/json"
	"errors"
	"fmt"
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
	err = json.NewDecoder(f).Decode(&s.data)
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

func (s store) All() map[string]string {
	return s.data
}

func (s store) Save() error {
	f, err := os.Create(s.path)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(s.data)
}

const Usage = `Usage: kv COMMAND [key] [value]

kv is a tool to manage a simple key-value store of strings. It
understands the following commands:

	kv all
		Lists all key-value pairs in the store file

	kv get KEY
		Prints the value associated with KEY, if one exists

	kv set KEY VALUE
		Sets KEY to be associated with VALUE, overwriting
		any previous associated value.

	The default store file is 'kv.store'. This file will be
	created automatically the first time a value is set using
	'kv set'.`

func Main() {
	if len(os.Args) < 2 {
		fmt.Println(Usage)
		return
	}
	command := os.Args[1]
	switch command {
	case "all":
		MainAll()
	case "get":
		MainGet()
	case "set":
		MainSet()
	default:
		fmt.Fprintln(os.Stderr, Usage)
		os.Exit(1)
	}
}

func MainAll() {
	s, err := OpenStore("kv.store")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for k, v := range s.All() {
		fmt.Printf("%s=%s\n", k, v)
	}
}

func MainGet() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, Usage)
		os.Exit(1)
	}
	s, err := OpenStore("kv.store")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	value, ok := s.Get(os.Args[2])
	if !ok {
		fmt.Fprintf(os.Stderr, "key %q not found\n", os.Args[2])
		os.Exit(1)
	}
	fmt.Println(value)
}

func MainSet() {
	if len(os.Args) < 4 {
		fmt.Fprintln(os.Stderr, Usage)
		os.Exit(1)
	}
	s, err := OpenStore("kv.store")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	s.Set(os.Args[2], os.Args[3])
	err = s.Save()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
