package models

import (
	"database/sql"
	"log"
)

type Usermoviemodels struct {
	DB *sql.DB
}

type Usermovie struct {
	UsermovieId   int
	UsermovieName string
}

func (m *Usermoviemodels) SelectById(usermovieId int) (*Usermovie, error) {
	query := "SELECT * FROM usermovie WHERE usermovieId = ?" //? la de nhap du lieu
	reslut := m.DB.QueryRow(query, usermovieId)
	log.Println(usermovieId)
	usermovie := &Usermovie{}
	err := reslut.Scan(&usermovie.UsermovieId, &usermovie.UsermovieName)
	if err != nil {
		return nil, err
	}
	return usermovie, nil
}

func (m *Usermoviemodels) SelectByName(usermovieName string) {

}

func (m *Usermoviemodels) Insert(usermovieName string) {

}

func (m *Usermoviemodels) Update(usermovieName string) {

}
