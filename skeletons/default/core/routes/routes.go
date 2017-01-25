package routes

import (
	"bitbucket.org/kenorld/eject-core"
	"bitbucket.org/kenorld/xlh-server/core/routes/www"
)

func Register(router *eject.Router) {
	//	router.Before("*", eject.SharedHandlers...)
	// register index using a 'Handler'
	//router.Path("/").Get(routes.Index)
	eject.Static(router.Path("/favicon.ico"), []string{"/public/favicon.ico"}, nil)
	eject.Static(router.Path("/public/<*path>"), []string{"/public"}, eject.StaticOptions{"listing": true})

	//router.Path("/").Get(www.Home.Index)
	root := router.Path("/").Get(www.Home.Index)
	user := root.Path("users").Get(www.User.List)

	user.Path("create").Get(www.User.Create)
	user.Path("update").Post(www.User.Update)
	user.Path("trash").Post(www.User.Trash)
	user.Path("<id>").Get(www.User.Show)
}
