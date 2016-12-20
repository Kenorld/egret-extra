package controllers

import "github.com/eject/eject"

type App struct {
	*eject.Controller
}

func (c App) Index() eject.Result {
	return c.Render()
}
