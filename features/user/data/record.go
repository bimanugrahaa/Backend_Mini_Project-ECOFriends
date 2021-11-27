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

func toCore(u *User) user.UserCore {
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
	for _, value := range resp {
		// u[key] = toCore(&value)
		// u = append(u, resp[key].toCore())
		u = append(u, toCore(&value))
	}

	return u
}

// func toCoreUserId(resp User) user.UserCore {

// }
