package donation

import "time"

type Core struct {
	ID          int
	Title       string
	Description DescriptionCore
	Created_at  time.Time
}

type DescriptionCore struct {
	ID               int
	Description      string
	Target_Donation  int
	Current_Donation int
}

type Bussiness interface {
	CreateData(data Core) (resp Core, err error)
	GetAllData(search string) (resp []Core)
	//Another CRUD
}

//Initialize Port
type Data interface {
	InsertData(data Core) (resp Core, err error)
	SelectData(title string) (resp []Core)
}
