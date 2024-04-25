package posters

import "database/sql"

type PosterRepo struct {
	Database *sql.DB
}
type Poster struct {
	Id              int
	MovieId         int
	MainPoster      string
	SecondaryPoster sql.NullString
	ThirdPoster     sql.NullString
	FourthPoster    sql.NullString
	FifthPoster     sql.NullString
}
type PosterJson struct {
	Id              int    `json:"id"`
	MovieId         int    `json:"movieId"`
	MainPoster      string `json:"mainPoster" binding:"required"`
	SecondaryPoster string `json:"secondaryPoster"`
	ThirdPoster     string `json:"thirdPoster"`
	FourthPoster    string `json:"fourthPoster"`
	FifthPoster     string `json:"fifthPoster"`
}
