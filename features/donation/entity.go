package donation

import "time"

type Core struct {
	ID          int
	Title       string
	AuthorID    int
	Author      UserCore
	Description DescriptionCore
	Created_at  time.Time
}

type DescriptionCore struct {
	ID               int
	Description      string
	Target_Donation  int
	Current_Donation int
}

type UserCore struct {
	ID   int
	Name string
}

type Bussiness interface {
	CreateData(data Core) (resp Core, err error)
	GetAllData() (resp []Core)
	// GetUserById() (resp []UserCore, err error)
	//Another CRUD
}

//Initialize Port
type Data interface {
	InsertData(data Core) (resp Core, err error)
	SelectData() (resp []Core)
	// SelectUserById(id int) (resp UserCore)
}
