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

func (ur *mysqlUserRepository) InsertUser(data user.UserCore) (resp user.UserCore, err error) {
	record := fromCore(data)

	if err := ur.Conn.Create(&record).Error; err != nil {
		return user.UserCore{}, err
	}
	return user.UserCore{}, nil
}

func (ur *mysqlUserRepository) SelectAllUser() (resp []user.UserCore) {
	record := []User{}
	if err := ur.Conn.Find(&record).Error; err != nil {
		return []user.UserCore{}
	}
	return toCoreList(record)
}

func (ur *mysqlUserRepository) SelectUserById(id int) (resp user.UserCore, err error) {
	record := User{}

	if err := ur.Conn.First(&record, id).Error; err != nil {
		return user.UserCore{}, err
	}

	return toCore(&record), err
}
