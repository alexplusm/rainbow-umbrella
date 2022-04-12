package utils

import (
	"fmt"
	"os"
)

func MakeDirs(dirsPath []string) error {
	for _, dirPath := range dirsPath {
		if err := os.MkdirAll(dirPath, 0777); err != nil && !os.IsExist(err) {
			return fmt.Errorf("[MakeDirs][1]: %w", err)
		}
	}
	return nil
}
