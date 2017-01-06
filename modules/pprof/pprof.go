package pprof

import (
	"net/http/pprof"
	"strings"

	"github.com/kataras/eject"
)

// New returns the pprof (profile, debug usage) Handler/ middleware
// NOTE: Route MUST have the last named parameter wildcard named '*action'
// Usage:
// eject.Get("debug/pprof/*action", pprof.New())
func New() eject.HandlerFunc {
	indexHandler := eject.ToHandler(pprof.Index)
	cmdlineHandler := eject.ToHandler(pprof.Cmdline)
	profileHandler := eject.ToHandler(pprof.Profile)
	symbolHandler := eject.ToHandler(pprof.Symbol)
	goroutineHandler := eject.ToHandler(pprof.Handler("goroutine"))
	heapHandler := eject.ToHandler(pprof.Handler("heap"))
	threadcreateHandler := eject.ToHandler(pprof.Handler("threadcreate"))
	debugBlockHandler := eject.ToHandler(pprof.Handler("block"))

	return eject.HandlerFunc(func(ctx *eject.Context) {
		ctx.SetContentType("text/html; charset=" + ctx.Framework().Config.Charset)

		action := ctx.Param("action")
		if len(action) > 1 {
			if strings.Contains(action, "cmdline") {
				cmdlineHandler.Serve((ctx))
			} else if strings.Contains(action, "profile") {
				profileHandler.Serve(ctx)
			} else if strings.Contains(action, "symbol") {
				symbolHandler.Serve(ctx)
			} else if strings.Contains(action, "goroutine") {
				goroutineHandler.Serve(ctx)
			} else if strings.Contains(action, "heap") {
				heapHandler.Serve(ctx)
			} else if strings.Contains(action, "threadcreate") {
				threadcreateHandler.Serve(ctx)
			} else if strings.Contains(action, "debug/block") {
				debugBlockHandler.Serve(ctx)
			}
		} else {
			indexHandler.Serve(ctx)
		}
	})
}
