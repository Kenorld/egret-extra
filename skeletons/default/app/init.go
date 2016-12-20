package app

import "github.com/eject/eject"

func init() {
	// Filters is the default set of global filters.
	eject.Filters = []eject.Filter{
		eject.PanicFilter,             // Recover from panics and display an error page instead.
		eject.RouterFilter,            // Use the routing table to select the right Action
		eject.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		eject.ParamsFilter,            // Parse parameters into Controller.Params.
		eject.SessionFilter,           // Restore and write the session cookie.
		eject.FlashFilter,             // Restore and write the flash cookie.
		eject.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		eject.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		eject.InterceptorFilter,       // Run interceptors around the action.
		eject.CompressFilter,          // Compress the result.
		eject.ActionInvoker,           // Invoke the action.
	}

	// register startup functions with OnAppStart
	// ( order dependent )
	// eject.OnAppStart(InitDB)
	// eject.OnAppStart(FillCache)
}

// TODO turn this into eject.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *eject.Controller, fc []eject.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}
