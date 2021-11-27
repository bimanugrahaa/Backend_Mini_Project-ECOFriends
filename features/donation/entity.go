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
	CreateDonation(data Core) (resp Core, err error)
	// CreateDescriptionDonation(data DescriptionCore) (resp DescriptionCore, err error)
	GetAllDonations() (resp []Core)
	GetDonationsById(id int) (resp Core)
	//Another CRUD
}

//Initialize Port
type Data interface {
	InsertDonation(data Core) (resp Core, err error)
	// InsertDescriptionDonation(data DescriptionCore) (resp DescriptionCore, err error)
	SelectAllDonations() (resp []Core)
	SelectDonationsById(id int) (resp Core)
}
