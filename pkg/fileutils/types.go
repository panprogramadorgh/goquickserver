package fileutils

type File struct {
	Name     string
	IsDir    bool
	Path     string
	DirFiles *[]File
}
