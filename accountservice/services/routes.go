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
        "/healthy123",
        HealthCheck,
	},

	Route {
		"Testability",
		"GET",
		"/testability/healthxyz/{state}",
		SetHealthyState,
	},

}
