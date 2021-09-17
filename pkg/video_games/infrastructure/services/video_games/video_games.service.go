package video_games

import (
	"api_test_hexagonal/config/databases/mysql"
	"api_test_hexagonal/pkg/video_games/domain/entities"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type VideoGameService struct {}

func (v *VideoGameService) GetAllVideoGames() (entities.IVideoGamesEntity, error) {
	var videoGames entities.IVideoGamesEntity
	bd, err := mysql.Connect()
	if err != nil {
		fmt.Println("erro x aqui?1: ", err)
		return nil, err
	}

	rows, err := bd.Query("SELECT id, name, genre, year FROM video_games")
	if err != nil {
		fmt.Println("erro x aqui?: ", err)
		return nil, err
	}
	for rows.Next() {
		var videoGame entities.IVideoGameEntity
		err = rows.Scan(&videoGame.Id, &videoGame.Name, &videoGame.Genre, &videoGame.Year)
		if err != nil {
			return nil, err
		}
		videoGames = append(videoGames, &videoGame)
	}
	return videoGames, nil
}


func (v *VideoGameService) CreateVideoGame(videoGame entities.IVideoGameEntity) error {
	bd, err := mysql.Connect()
	if err != nil {
		return err
	}

	_, err = bd.Exec("INSERT INTO video_games (name, genre, year) VALUES (?, ?, ?)", videoGame.Name, videoGame.Genre, videoGame.Year)
	return err
}

func (v *VideoGameService) UpdateVideoGame(videoGame entities.IVideoGameEntity, videoGameID string) error {
	bd, err := mysql.Connect()
	if err != nil {
		return err
	}

	_, err = bd.Exec("UPDATE video_games SET name = ?, genre = ?, year = ? WHERE id = ?", videoGame.Name, videoGame.Genre, videoGame.Year, videoGameID)
	return err
}

func(v *VideoGameService) DeleteVideoGames(videoGameID string) error {
	bd, err := mysql.Connect()
	if err != nil {
		return err
	}

	_, err = bd.Exec("DELETE FROM video_games WHERE id = ?", videoGameID)
	return err
}