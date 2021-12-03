package donation

import (
	"time"
)

type Core struct {
	ID          int
	Title       string
	AuthorID    int
	Author      UserCore
	Description DescriptionCore
	Comment     []CommentCore
	Created_at  time.Time
}

type DescriptionCore struct {
	ID                  int
	Description         string
	Target_Donation     int
	Current_Donation    int
	Percentage_Donation float64
}

type UserCore struct {
	ID   int
	Name string
}

type CommentCore struct {
	ID         int
	Comment    string
	PostID     int
	UserID     int
	Status     bool
	Updated_at time.Time
}

type Bussiness interface {
	CreateDonation(data Core) (resp Core, err error)
	GetAllDonations() (resp []Core)
	GetDonationsById(id int) (resp Core)
	GetDonationTrending() (resp []Core)
	GetDonationLatest() (resp []Core)
	DeleteDonationsById(data Core) (err error)
	UpdateDonation(data Core) (resp Core, err error)
	UpdateDonationValue(data DescriptionCore) (resp DescriptionCore, err error)

	CreateComment(id int, data CommentCore) (resp CommentCore, err error)
	GetCommentByPostId(id int) (resp []CommentCore, err error)
	UpdateComment(data CommentCore) (resp CommentCore, err error)
	DeleteComment(data CommentCore) (err error)
	//Another CRUD
}

//Initialize Port
type Data interface {
	InsertDonation(data Core) (resp Core, err error)
	SelectAllDonations() (resp []Core)
	SelectDonationsById(id int) (resp Core)
	SelectDonationsTrending() (resp []Core)
	SelectDonationsLatest() (resp []Core)
	RemoveDonationsById(data Core) (err error)
	EditDonation(data Core) (resp Core, err error)
	EditDonationValue(data DescriptionCore) (resp DescriptionCore, err error)

	InsertComment(id int, data CommentCore) (resp CommentCore, err error)
	SelectCommentByPostId(id int) (resp []CommentCore, err error)
	SelectCommentById(id int) (resp CommentCore, err error)
	EditComment(data CommentCore) (resp CommentCore, err error)
	RemoveComment(data CommentCore) (err error)
}
