package routes

import (
	"net/http"
	"net/http/pprof"

	"bitbucket.org/kenorld/eject-core"
)

type Pprof struct {
	*eject.Context
}

// The PprofHandler type makes it easy to call the net/http/pprof handler methods
// since they all have the same method signature
type PprofHandler func(http.ResponseWriter, *http.Request)

func (r PprofHandler) Apply(req *eject.Request, resp *eject.Response) {
	r(resp.Out, req.Request)
}

func (c Pprof) Profile() eject.Result {
	return PprofHandler(pprof.Profile)
}

func (c Pprof) Symbol() eject.Result {
	return PprofHandler(pprof.Symbol)
}

func (c Pprof) Cmdline() eject.Result {
	return PprofHandler(pprof.Cmdline)
}

func (c Pprof) Trace() eject.Result {
	return PprofHandler(pprof.Trace)
}

func (c Pprof) Index() eject.Result {
	return PprofHandler(pprof.Index)
}
