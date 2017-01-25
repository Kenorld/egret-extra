package www

import "bitbucket.org/kenorld/eject-core"

type home struct {
}

func (h *home) Index(ctx *eject.Context) {
	data := &struct {
		Title string
	}{
		"Hello World",
	}
	ctx.RenderTemplate("index.html", data, nil)
}

var Home = &home{}
