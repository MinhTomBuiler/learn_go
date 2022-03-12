package models

import (
	"database/sql"
	"log"
)

type Usermodels struct {
	DB *sql.DB
}

type User struct {
	UserId   int
	UserName string
}

func (m *Usermodels) SelectById(userId int) (*User, error) {
	query := "SELECT * FROM user WHERE userId = ?" //? la de nhap du lieu
	reslut := m.DB.QueryRow(query, userId)
	log.Println(userId)
	user := &User{}
	err := reslut.Scan(&user.UserId, &user.UserName)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m *Usermodels) SelectByName(userName string) {

}

func (m *Usermodels) Insert(userName string) {

}

func (m *Usermodels) Update(userName string) {

}
