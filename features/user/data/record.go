package data

import (
	"Backend_Mini_Project-ECOFriends/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

func (u *User) toCore() user.UserCore {
	return user.UserCore{
		ID:         int(u.ID),
		Name:       u.Name,
		Email:      u.Email,
		Password:   u.Password,
		Created_at: u.CreatedAt,
	}
}

func toCoreList(resp []User) []user.UserCore {
	u := []user.UserCore{}
	for key := range resp {
		u = append(u, resp[key].toCore())
	}

	return u
}
