package data

import (
	"Backend_Mini_Project-ECOFriends/features/user"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) user.Data {
	return &mysqlUserRepository{
		Conn: conn,
	}
}

func (ur *mysqlUserRepository) InsertData(data user.UserCore) (resp user.UserCore, err error) {
	return user.UserCore{}, nil
}

func (ur *mysqlUserRepository) SelectData() (resp []user.UserCore) {
	record := []User{}
	if err := ur.Conn.Find(&record).Error; err != nil {
		return []user.UserCore{}
	}
	return toCoreList(record)
}
