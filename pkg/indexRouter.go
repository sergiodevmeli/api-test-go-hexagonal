package pkg

import (
	route "api_test_hexagonal/pkg/infrastructure/routes"
	"net/http"
)

func IndexRouter(mux *http.ServeMux) {
	route.RouteProducts(mux)
	route.RouteVideoGames(mux)
}