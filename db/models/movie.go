package models

import (
	"database/sql"
	"log"
)

type Moviemodels struct {
	DB *sql.DB
}

type Movie struct {
	MovieId   int
	MovieName string
}

func (m *Moviemodels) SelectById(movieId int) (*Movie, error) {
	query := "SELECT * FROM movie WHERE movieId = ?" //? la de nhap du lieu
	reslut := m.DB.QueryRow(query, movieId)
	log.Println(movieId)
	movie := &Movie{}
	err := reslut.Scan(&movie.MovieId, &movie.MovieName)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (m *Moviemodels) SelectByName(movieName string) {

}

func (m *Moviemodels) Insert(movieName string) {

}

func (m *Moviemodels) Update(movieName string) {

}
