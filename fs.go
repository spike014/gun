package main

import (
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type strcutNil struct{}

var excludePath = map[string]struct{}{
	".git": strcutNil{},
}

var filetypes = []string{".go", ".toml", ".json", ".ini"}
var filepaths *[]string

func checkIfIsNotGoProject(path string) bool {
	return pathNotExists(filepath.Join(path, "main.go"))
}

// getExcutePath 获取程序所在目录
func getExcutePath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		log.Fatalln("Get program path ERROR:", err)
		return "", err
	}
	return filepath.Dir(ex), nil
}

func pathNotExists(path string) bool {
	_, err := os.Stat(path)
	return errors.Is(err, fs.ErrNotExist)
}

func updateFilespath() {
	var err error
	filepaths, err = getAllFilespath(basePath)
	if err != nil {
		panic(err)
	}
}

// getAllFilespath get all .go files under the root directory
func getAllFilespath(path string) (*[]string, error) {
	paths := make([]string, 0)
	err := filepath.WalkDir(path,
		func(path string, d fs.DirEntry, err error) error {
			if _, ok := excludePath[d.Name()]; ok ||
				strings.Contains(path, ".git") || d.IsDir() {
				return nil
			}

			for _, v := range filetypes {
				if strings.HasSuffix(d.Name(), v) {
					paths = append(paths, path)
				}
			}
			return err
		})
	return &paths, err
}
