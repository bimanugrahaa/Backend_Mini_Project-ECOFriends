package donation

import "time"

type Core struct {
	ID         int
	Title      string
	Created_at time.Time
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
