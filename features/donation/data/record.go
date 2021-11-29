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
}

type DescriptionDonation struct {
	gorm.Model
	ID               int
	Description      string `gorm:"column:donation_desc"`
	Target_Donation  int    `gorm:"column:donation_target"`
	Current_Donation int    `gorm:"column:donation_current"`
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

func toCore(d *Donation) donation.Core {
	return donation.Core{
		ID:         int(d.ID),
		Title:      d.Title,
		AuthorID:   d.AuthorID,
		Created_at: d.CreatedAt,
	}
}

func toCoreDetail(d *Donation) donation.Core {
	return donation.Core{
		ID:          int(d.ID),
		Title:       d.Title,
		AuthorID:    d.AuthorID,
		Created_at:  d.CreatedAt,
		Description: toDescriptionCore(&d.Description),
	}
}

func toCoreList(resp []Donation) []donation.Core {
	d := []donation.Core{}

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

func fromCore(core donation.Core) Donation {
	return Donation{
		Title:       core.Title,
		AuthorID:    core.AuthorID,
		Description: fromDescriptionCore(core.Description),
	}
}
