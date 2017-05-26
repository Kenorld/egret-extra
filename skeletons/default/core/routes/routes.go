package routes

import (
	"github.com/kenorld/egret-core"
)

func Register() {
	router := egret.NewRouter()
	//	router.Before("*", egret.SharedHandlers...)
	// register index using a 'Handler'
	//router.Path("/").Get(routes.Index)
	egret.Static(router.Path("/favicon.ico"), []string{"/public/favicon.ico"}, nil)

	//router.Path("/").Get(Home.Index)
	root := router.Path("/").Get(Home.Index)
	user := root.Path("users").Get(User.List)

	user.Path("create").Get(User.Create)
	user.Path("update").Post(User.Update)
	user.Path("trash").Post(User.Trash)
	user.Path("<id>").Get(User.Show)
	
	egret.Static(router.Path("/<*path>"), []string{"/public"}, egret.StaticOptions{"listing": true})
}
