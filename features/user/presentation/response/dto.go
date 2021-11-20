package response

import (
	"Backend_Mini_Project-ECOFriends/features/user"
	"time"
)

type User struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"passord"`
	Created_at time.Time `json:"created_at"`
}

func FromCore(core user.UserCore) User {
	return User{
		ID:         core.ID,
		Name:       core.Name,
		Email:      core.Email,
		Password:   core.Password,
		Created_at: core.Created_at,
	}
}

func FromCoreSlice(core []user.UserCore) []User {
	var userArray []User
	for key := range core {
		userArray = append(userArray, FromCore(core[key]))
	}

	return userArray
}
