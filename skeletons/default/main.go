package main

import (
	"flag"

	"bitbucket.org/kenorld/eject-core"
	"bitbucket.org/kenorld/xlh-server/core/routes"
)

var (
	runMode    *string = flag.String("runMode", "dev", "Run mode.")
	port       *int    = flag.Int("port", 4411, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "bitbucket.org/kenorld/xlh-server", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")
)

func main() {
	flag.Parse()
	eject.Init(*runMode, *importPath, *srcPath)
	// DB Main
	//DbMain()
	routes.Register(eject.MainRouter)

	// start the server
	eject.Serve(*port)
}

// func DbMain() {
// 	// Database Main Conexion
// 	Db := db.MgoDb{}
// 	Db.Init()
// 	// index keys
// 	keys := []string{"email"}
// 	Db.Index("auth", keys)
// }
