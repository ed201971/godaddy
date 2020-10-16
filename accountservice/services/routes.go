package services

import "net/http"

type Route struct {
	Name	string
	Method	string
	Pattern	string
	HandlerFunc	http.HandlerFunc
}

type Routes[]Route

var routes = Routes {

	Route {
		"GetAccount",
		"GET",
		"/accounts/{accountId}",
		GetAccount,
	},

	Route{
        "HealthCheck",
        "GET",
        "/healthz",
        HealthCheck,
	},

	Route {
		"Testability",
		"GET",
		"/testability/healthy/{state}",
		SetHealthyState,
	},

}
