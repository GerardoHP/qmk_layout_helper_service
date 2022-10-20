package utils

import (
	"bufio"
	"fmt"
	"os"

	"github.com/GerardoHP/qmk_layout_helper_service/src/models"
)

type MappingType int

const (
	Basic MappingType = iota
	Quantum
)

var conf = map[MappingType]string{
	Basic:   "../config/keycodes_basic.csv",
	Quantum: "../config/keycodes_quantum.csv",
}

// Key mapper from file read
type FileKeyMapper struct {
	keyMaps map[string]models.KeyMap
	configs []MappingType
}

var _ KeyMapper = (*FileKeyMapper)(nil)

// Gets a new instance of the file key mapper
func NewFileKeyMapper(c []MappingType) *FileKeyMapper {
	if len(c) == 0 {
		c = append(c, Basic)
	}

	return &FileKeyMapper{configs: c, keyMaps: map[string]models.KeyMap{}}
}

// Read all the files
func (f *FileKeyMapper) Read() error {
	for _, v := range f.configs {
		err := readFromFile(v, f.keyMaps)
		if err != nil {
			return err
		}
	}

	return nil
}

// Gets a key value
func (f FileKeyMapper) GetKey(k string) (*models.KeyMap, error) {
	if v, found := f.keyMaps[k]; len(f.keyMaps) == 0 || !found {
		return nil, fmt.Errorf("key maps not found")
	} else {
		return &v, nil
	}
}

// Gets all the mapped keys
func (f FileKeyMapper) GetAllKeys() map[string]models.KeyMap {
	return f.keyMaps
}

// Read from file
// TODO: Move this files to YAML files which are more human readable
func readFromFile(t MappingType, m map[string]models.KeyMap) error {
	f, found := conf[t]
	if !found {
		return fmt.Errorf("file not found")
	}

	file, err := os.Open(f)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		km := models.NewKeyMapFromFileLine(scanner.Text())
		m[km.Key] = *km
	}

	return nil
}
