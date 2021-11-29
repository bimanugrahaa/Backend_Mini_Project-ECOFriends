package request

import (
	"Backend_Mini_Project-ECOFriends/features/user"
	"time"
)

type User struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Created_at time.Time `json:"created_at"`
}

func ToCore(req User) user.UserCore {
	return user.UserCore{
		ID:       req.ID,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}
