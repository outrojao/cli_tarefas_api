package routes

import (
	"net/http"
	"net/url"
	"testing"
)

func TestInitRoutes(t *testing.T) {
	InitRoutes()

	routes := []string{"/create", "/do/", "/remove/", "/list", "/health"}
	for _, route := range routes {
		_, pattern := http.DefaultServeMux.Handler(&http.Request{URL: &url.URL{Path: route}})
		if pattern == "" {
			t.Errorf("Route %s not registered", route)
		}
	}
}
