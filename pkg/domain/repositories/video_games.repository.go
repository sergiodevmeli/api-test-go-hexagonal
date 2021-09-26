package repositories

import (
	"api_test_hexagonal/pkg/domain/entities"
)

type IVideoGamesRepository interface {
	CreateVideoGame(videoGame entities.IVideoGameEntity) error
	GetAllVideoGames() (entities.IVideoGamesEntity, error)
	UpdateVideoGame(entities.IVideoGameEntity, string) error
	DeleteVideoGames(string) error
}
