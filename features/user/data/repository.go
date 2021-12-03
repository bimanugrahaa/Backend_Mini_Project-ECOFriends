package data

import (
	"Backend_Mini_Project-ECOFriends/features/user"
	"Backend_Mini_Project-ECOFriends/middleware"
	"fmt"

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
	return toCore(&record), nil
}

func (ur *mysqlUserRepository) EditUser(data user.UserCore) (resp user.UserCore, err error) {
	record := fromCore(data)

	if err := ur.Conn.Model(&User{}).Where("id = ?", data.ID).Updates(&record).Error; err != nil {
		return user.UserCore{}, err
	}

	return toCore(&record), nil

}

func (ur *mysqlUserRepository) RemoveUser(id int) (err error) {
	if err := ur.Conn.Delete(&User{}, id).Error; err != nil {
		return err
	}

	return nil
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

func (ur *mysqlUserRepository) SelectUserEmail(data user.UserCore) (resp user.UserCore, err error) {
	record := fromCore(data)

	if err := ur.Conn.Model(&User{}).Where("email = ?", data.Email).First(&record).Error; err != nil {
		return user.UserCore{}, err
	}

	return toCore(&record), err
}

func (ur *mysqlUserRepository) Login(data user.UserCore) (resp user.UserCore, err error) {

	record := fromCore(data)

	if err := ur.Conn.Model(&User{}).Where("email = ? AND password = ?", data.Email, data.Password).First(&record).Error; err != nil {
		return user.UserCore{}, err
	}

	record.Token, _ = middleware.CreateToken(int(record.ID))

	if err != nil {
		return user.UserCore{}, err
	}

	if err := ur.Conn.Model(&User{}).Where("id = ?", data.ID).Updates(&record).Error; err != nil {
		return user.UserCore{}, err
	}

	fmt.Println(record.Token)
	return toCore(&record), err
}
