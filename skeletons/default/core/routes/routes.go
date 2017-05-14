package routes

import (
	"bitbucket.org/kenorld/eject-core"
)

func Register() {
	router := eject.NewRouter()
	//	router.Before("*", eject.SharedHandlers...)
	// register index using a 'Handler'
	//router.Path("/").Get(routes.Index)
	eject.Static(router.Path("/favicon.ico"), []string{"/public/favicon.ico"}, nil)

	//router.Path("/").Get(Home.Index)
	root := router.Path("/").Get(Home.Index)
	user := root.Path("users").Get(User.List)

	user.Path("create").Get(User.Create)
	user.Path("update").Post(User.Update)
	user.Path("trash").Post(User.Trash)
	user.Path("<id>").Get(User.Show)
	
	eject.Static(router.Path("/<*path>"), []string{"/public"}, eject.StaticOptions{"listing": true})
}
