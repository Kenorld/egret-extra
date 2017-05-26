package www

import "github.com/kenorld/egret-core"

type home struct {
}

func (h *home) Index(ctx *egret.Context) {
	data := &struct {
		Title string
	}{
		"Hello World",
	}
	ctx.RenderTemplate("index.html", data, nil)
}

var Home = &home{}
