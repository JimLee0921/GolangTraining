package counter

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Store interface {
	Load() (int, error)
	Save(v int) error
}

type FileStore struct {
	Path string
}

func (s FileStore) Load() (int, error) {
	b, err := os.ReadFile(s.Path)
	if err != nil {
		return 0, err
	}
	txt := strings.TrimSpace(string(b))
	if txt == "" {
		return 0, nil
	}
	v, err := strconv.Atoi(txt)
	if err != nil {
		return 0, fmt.Errorf("invalid data in store: %s", err.Error())
	}
	return v, nil
}

func (s FileStore) Save(v int) error {
	return os.WriteFile(s.Path, []byte(strconv.Itoa(v)+"\n"), 0o600)
}
