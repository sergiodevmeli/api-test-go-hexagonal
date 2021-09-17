package utils

import (
	"net/http"
)

func VerifyMethodHttp(writer http.ResponseWriter, methodRequest, methodHttp string) bool {
	if methodRequest != methodHttp {
		response := NewResponse(Error, "Metodo no permitido", nil)
		ResponseJSON(writer, http.StatusBadRequest, response)
		return true
	}
	return false
}