package fileutils

import (
	"fmt"
	"os"
)

func ReadDir(path string) ([]File, error) {
	dir, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return []File{}, err
	}
	defer dir.Close()

	var dirFiles []File = []File{}
	files, err := (*dir).Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return []File{}, nil
	}
	for _, file := range files {
		var currentDirFile File
		filePath := fmt.Sprint(path, "\\", file.Name())
		if file.IsDir() {
			innerFiles, err := ReadDir(filePath)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
			currentDirFile = File{Name: file.Name(), Path: filePath, IsDir: file.IsDir(), DirFiles: &innerFiles}
		} else {
			currentDirFile = File{Name: file.Name(), Path: filePath, IsDir: file.IsDir()}
		}

		dirFiles = append(dirFiles, currentDirFile)
	}
	return dirFiles, nil
}
