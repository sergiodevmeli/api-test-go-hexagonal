package handlers

import (
	"api_test_hexagonal/pkg/application/use-cases"
	"api_test_hexagonal/pkg/domain/entities"
	"api_test_hexagonal/pkg/infrastructure/services"
	"api_test_hexagonal/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type VideoGameHandler struct {
	service *services.VideoGameService
}

func VideoGame() VideoGameHandler {
	return VideoGameHandler{}
}

func (videoGame *VideoGameHandler) GetAll(writer http.ResponseWriter, request *http.Request){
	if utils.VerifyMethodHttp(writer, request.Method, http.MethodGet) { return }

	videoGameUseCase := use_cases.VideoGameUseCase{Repository: videoGame.service}

	videoGameList, err := videoGameUseCase.GetAllVideoGames()
	if err != nil {
		response := utils.NewResponse(utils.Error, "Error en el get", nil)
		utils.ResponseJSON(writer, http.StatusBadRequest, response)
		return
	}

	response := utils.NewResponse(utils.Message, "Lista de video games", videoGameList)
	utils.ResponseJSON(writer, http.StatusOK, response)
}

func (videoGame *VideoGameHandler) Create(writer http.ResponseWriter, request *http.Request) {
	if utils.VerifyMethodHttp(writer, request.Method, http.MethodPost) { return }

	videoGameUseCase := use_cases.VideoGameUseCase{Repository: videoGame.service}
	var newVideoGame entities.IVideoGameEntity

	err := json.NewDecoder(request.Body).Decode(&newVideoGame)
	if err != nil {
		response := utils.NewResponse(utils.Error, "JSON mal formado", nil)
		utils.ResponseJSON(writer, http.StatusBadRequest, response)
		return
	}

	errVideoGameCase := videoGameUseCase.CreateVideoGame(newVideoGame)
	fmt.Println("errVideoGameCase: ", errVideoGameCase)
	if errVideoGameCase != nil {
		response := utils.NewResponse(utils.Error, "Error en crear el video game", nil)
		utils.ResponseJSON(writer, http.StatusBadRequest, response)
		return
	}

	response := utils.NewResponse(utils.Message, "Video Game creado con exito", nil)
	utils.ResponseJSON(writer, http.StatusCreated, response)
}

func (videoGame *VideoGameHandler) Update(writer http.ResponseWriter, request *http.Response){

}

func (videoGame *VideoGameHandler) Delete(writer http.ResponseWriter, request *http.Response){

}