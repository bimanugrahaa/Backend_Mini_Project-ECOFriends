package user

import (
	"time"
)

type UserCore struct {
	ID         int
	Name       string
	Email      string
	Password   string
	Created_at time.Time
}

type Bussiness interface {
	CreateUser(data UserCore) (resp UserCore, err error)
	GetAllUser() (resp []UserCore)
	GetUserById(id int) (resp UserCore, err error)
	// UpdateUser(data UserCore) (resp UserCore, err error)
	// DeleteUser(id int) (err error)
}

type Data interface {
	InsertUser(data UserCore) (resp UserCore, err error)
	SelectAllUser() (resp []UserCore)
	SelectUserById(id int) (resp UserCore, err error)
	// EditUser(data UserCore) (resp UserCore, err error)
	// RemoveUser(id int) (err error)
}
