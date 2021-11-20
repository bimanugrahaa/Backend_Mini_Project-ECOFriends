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
}

type Data interface {
	InsertData(data UserCore) (resp UserCore, err error)
	SelectData() (resp []UserCore)
}
