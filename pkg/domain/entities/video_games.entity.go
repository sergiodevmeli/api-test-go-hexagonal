package entities

type IVideoGameEntity struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Genre string `json:"genre"`
	Year int `json:"year"`
}

type IVideoGamesEntity []*IVideoGameEntity