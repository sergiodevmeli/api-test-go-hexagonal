package routes

import (
	"api_test_hexagonal/pkg/infrastructure/handlers"
	"net/http"
)

func RouteProducts(mux *http.ServeMux) {
	handler := handlers.Product()

	mux.HandleFunc("/v1/product/get-all", handler.GetAll)
	mux.HandleFunc("/v1/product/create", handler.Create)
	mux.HandleFunc("/v1/product/update", handler.Update)
}