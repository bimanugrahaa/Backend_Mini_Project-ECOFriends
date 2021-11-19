package data

import (
	"Backend_Mini_Project-ECOFriends/features/donation"

	"gorm.io/gorm"
)

//Contains field that already made before

type Donation struct {
	gorm.Model
	Title string `gorm:"column:donation_title"`
}

//DTO

func (d *Donation) toCore() donation.Core {
	return donation.Core{
		ID:         int(d.ID),
		Title:      d.Title,
		Created_at: d.CreatedAt,
	}
}

func toCoreList(resp []Donation) []donation.Core {
	d := []donation.Core{}
	for key := range resp {
		d = append(d, resp[key].toCore())
	}

	return d
}
