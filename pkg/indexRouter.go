package pkg

import (
	product "api_test_hexagonal/pkg/products/infrastructure/routes"
	videoGame "api_test_hexagonal/pkg/video_games/infrastructure/routes"
	"net/http"
)

func IndexRouter(mux *http.ServeMux) {
	product.RouteProducts(mux)
	videoGame.RouteVideoGames(mux)
}