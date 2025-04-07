package utils

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

// Exists determine whether the file exists
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// CreateNestedDirectory create nested directory
func CreateNestedDirectory(path string) error {
	err := os.MkdirAll(path, 0700)
	if err != nil {
		log.Errorf("can't create folder, %s", err)
	}
	return err
}

// CreateNestedFile create nested file
func CreateNestedFile(path string) (*os.File, error) {
	basePath := filepath.Dir(path)
	if err := CreateNestedDirectory(basePath); err != nil {
		return nil, err
	}
	return os.Create(path)
}
