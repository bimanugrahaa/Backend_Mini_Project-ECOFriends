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
	Token    string
}

func toCore(u *User) user.UserCore {
	return user.UserCore{
		ID:         int(u.ID),
		Name:       u.Name,
		Email:      u.Email,
		Password:   u.Password,
		Token:      u.Token,
		Created_at: u.CreatedAt,
	}
}

func toCoreList(resp []User) []user.UserCore {
	u := []user.UserCore{}
	for _, value := range resp {
		// u[key] = toCore(&value)
		// u = append(u, resp[key].toCore())
		u = append(u, toCore(&value))
	}

	return u
}

func fromCore(core user.UserCore) User {
	return User{
		Name:     core.Name,
		Email:    core.Email,
		Password: core.Password,
	}
}

// func toCoreUserId(resp User) user.UserCore {

// }
