package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/go-chi/chi"
)

func Test_application_routes(t *testing.T) {
	var registerd = []struct {
		route  string
		method string
	}{
		{"/", "GET"},
		{"/login", "POST"},
		{"/static/*", "GET"},
	}

	mux := app.routes()

	chiRoutes := mux.(chi.Routes)

	for _, route := range registerd {
		// check to see if the route exists
		if !routeExists(route.route, route.method, chiRoutes) {
			t.Errorf("route %s is not registerd", route.route)
		}
	}
}

// return true if this exists and false if it not
func routeExists(testRoute, testMethod string, chiRoutes chi.Routes) bool {
	found := false

	_ = chi.Walk(chiRoutes, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		if strings.EqualFold(method, testMethod) && strings.EqualFold(route, testRoute) {
			found = true
		}
		return nil
	})

	return found
}
