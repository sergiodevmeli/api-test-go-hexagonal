package use_case

import (
	entities2 "api_test_hexagonal/pkg/video_games/domain/entities"
	"api_test_hexagonal/pkg/video_games/domain/repositories"
)

type VideoGameUseCase struct {
	Repository repositories.IVideoGamesRepository
}

func (repo *VideoGameUseCase) GetAllVideoGames() (entities2.IVideoGamesEntity, error) {
	return repo.Repository.GetAllVideoGames()
}

func (repo *VideoGameUseCase) CreateVideoGame(videoGame entities2.IVideoGameEntity) error {
	return repo.Repository.CreateVideoGame(videoGame)
}

func (repo *VideoGameUseCase) UpdateVideoGame(videoGames entities2.IVideoGameEntity, videoGameID string) error {
	return repo.Repository.UpdateVideoGame(videoGames, videoGameID)
}

func (repo *VideoGameUseCase) DeleteVideoGame(videoGameID string) error {
	return repo.Repository.DeleteVideoGames(videoGameID)
}