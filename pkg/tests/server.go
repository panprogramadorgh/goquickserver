package tests

import (
	"fmt"

	configutils "github.com/panprogramadorgh/quickgoserver/pkg/confutils"
	"github.com/panprogramadorgh/quickgoserver/pkg/fileutils"
	"github.com/panprogramadorgh/quickgoserver/pkg/webserver"
)

func InitServer() {
	// lee la opcion site_dir del fichero de configuracion
	siteDir, err := configutils.ReadConf("site_dir", nil)
	if err != nil {
		panic(err)
	}
	// crea un mapa de lectura de directorios para la ruta de la carpeta indicada
	dir, err := fileutils.ReadDir(siteDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	// crea un enrutador web de directorios. se le debe pasar la ruta del directorio raiz junto al mapa de lectura
	webserver.DirRouter(siteDir, &dir)

	// define el puerto del servidor
	port := 3000
	// arranca el servidor
	webserver.Listen(fmt.Sprint(":", port), func() {
		fmt.Println("Server running on http://localhost:", port)
	})

}
