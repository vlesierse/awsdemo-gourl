package api

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"RedirectUrl",
		"GET",
		"/{slug}",
		HandleRedirectToUrl,
	},
	Route{
		"CreateUrl",
		"POST",
		"/",
		HandleCreateUrl,
	},
}
