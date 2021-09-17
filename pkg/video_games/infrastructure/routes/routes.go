package routes

import (
	"api_test_hexagonal/pkg/video_games/infrastructure/handlers"
	"net/http"
)

func RouteVideoGames(mux *http.ServeMux) {
	handler := handlers.VideoGame()

	mux.HandleFunc("/v1/videogame/get-all", handler.GetAll)
	mux.HandleFunc("/v1/videogame/create", handler.Create)
	//mux.HandleFunc("/v1/product/update", handler.Update)
}