package data

import (
	"Backend_Mini_Project-ECOFriends/features/donation"

	"gorm.io/gorm"
)

//Contains field that already made before

type Donation struct {
	gorm.Model
	Title                 string `gorm:"column:donation_title"`
	DescriptionDonationID int
	AuthorID              int
	Description           DescriptionDonation `gorm:"foreignKey:id"`
	// Comment               []CommentDonation   `gorm:"foreignKey:CommentDonation.PostID"`
}

type DescriptionDonation struct {
	gorm.Model
	ID               int
	Description      string `gorm:"column:donation_desc"`
	Target_Donation  int    `gorm:"column:donation_target"`
	Current_Donation int    `gorm:"column:donation_current"`
}

type CommentDonation struct {
	gorm.Model
	ID      int
	Comment string `gorm:"column:comment"`
	PostID  int
	UserID  int
	Status  bool `gorm:"default:false"`
}

//DTO

func toDescriptionCore(dd *DescriptionDonation) donation.DescriptionCore {
	return donation.DescriptionCore{
		ID:               dd.ID,
		Description:      dd.Description,
		Target_Donation:  dd.Target_Donation,
		Current_Donation: dd.Current_Donation,
	}
}

func toCommentCore(cd *CommentDonation) donation.CommentCore {
	return donation.CommentCore{
		ID:      cd.ID,
		Comment: cd.Comment,
		PostID:  cd.PostID,
		UserID:  cd.UserID,
		Status:  cd.Status,
	}
}

func toCore(d *Donation) donation.Core {
	return donation.Core{
		ID:         int(d.ID),
		Title:      d.Title,
		AuthorID:   d.AuthorID,
		Created_at: d.CreatedAt,
	}
}

func toCommentList(resp []CommentDonation) []donation.CommentCore {
	cc := []donation.CommentCore{}

	for _, value := range resp {
		cc = append(cc, toCommentCore(&value))
	}

	return cc
}

func toCoreDetail(d *Donation) donation.Core {
	// cc := []donation.CommentCore{}

	// fmt.Println(d.Comment)
	// for _, value := range d.Comment {
	// 	cc = append(cc, toCommentCore(&value))
	// }
	// fmt.Println(cc)

	return donation.Core{
		ID:          int(d.ID),
		Title:       d.Title,
		AuthorID:    d.AuthorID,
		Created_at:  d.CreatedAt,
		Description: toDescriptionCore(&d.Description),
		// Comment:     cc,
	}
}

func toCoreList(resp []Donation) []donation.Core {
	d := []donation.Core{}

	// fmt.Println("d", resp)
	for _, value := range resp {
		d = append(d, toCore(&value))
	}

	return d
}

func fromDescriptionCore(dc donation.DescriptionCore) DescriptionDonation {
	return DescriptionDonation{
		Description:      dc.Description,
		Target_Donation:  dc.Target_Donation,
		Current_Donation: dc.Current_Donation,
	}
}

func fromCommentCore(id int, cc donation.CommentCore) CommentDonation {
	return CommentDonation{
		Comment: cc.Comment,
		PostID:  id,
		UserID:  cc.UserID,
		Status:  cc.Status,
	}
}

func fromCore(core donation.Core) Donation {
	return Donation{
		Title:       core.Title,
		AuthorID:    core.AuthorID,
		Description: fromDescriptionCore(core.Description),
	}
}
