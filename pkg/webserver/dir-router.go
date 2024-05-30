package webserver

import (
	"fmt"
	"net/http"

	"github.com/panprogramadorgh/quickgoserver/pkg/fileutils"
)

func DirRouter(rootDirPath string, dir *[]fileutils.File) {
	for _, file := range *dir {
		// routePath se queda con la parte negativa de una mascara entre el root path y el file path
		routePath := fileutils.SwapFsPathToRoutePath(file.Path[len(rootDirPath):len(file.Path)])
		if file.IsDir {
			DirRouter(rootDirPath, file.DirFiles)
			break
		}
		Route(routePath, func(w http.ResponseWriter, r *http.Request) {
			html, err := fileutils.ReadFile(file.Path)
			if err != nil {
				fmt.Println(err)
				fmt.Fprint(w, "internal server error | 500")
			}
			fmt.Fprint(w, html)
		})
	}
}
