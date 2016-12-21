package controllers

import "bitbucket.org/kenorld/eject-core"

type App struct {
	*eject.Controller
}

func (c App) Index() eject.Result {
	return c.Render()
}
