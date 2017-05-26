package www

import "github.com/kenorld/eject-core"

// About the handler func, we don't use struct here
// this is the simplest method to declare a route
func About(ctx *eject.Context) {
	// MustRender, same as Render but sends status 500 internal server error if rendering failed
	x := struct {
		Title string
	}{
		"titile",
	}
	ctx.RenderTemplate("about.html", x, nil)
}
