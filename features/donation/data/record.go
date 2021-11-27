package data

import (
	"Backend_Mini_Project-ECOFriends/features/donation"
	"fmt"

	"gorm.io/gorm"
)

//Contains field that already made before

type Donation struct {
	gorm.Model
	Title                 string `gorm:"column:donation_title"`
	DescriptionDonationID int
	AuthorID              int
	// Author                Users
	Description DescriptionDonation `gorm:"foreignKey:id"`
}

type DescriptionDonation struct {
	gorm.Model
	ID               int
	Description      string `gorm:"column:donation_desc"`
	Target_Donation  int    `gorm:"column:donation_target"`
	Current_Donation int    `gorm:"column:donation_current"`
}

// type Users struct {
// 	ID   int
// 	Name string
// }

//DTO
func (dd *DescriptionDonation) toDescriptionCore() donation.DescriptionCore {
	return donation.DescriptionCore{
		ID:               int(dd.ID),
		Description:      dd.Description,
		Target_Donation:  dd.Target_Donation,
		Current_Donation: dd.Current_Donation,
	}
}

// func toUsersCore(resp Users) donation.UserCore {
// 	return donation.UserCore{
// 		ID:   resp.ID,
// 		Name: resp.Name,
// 	}
// }

func toDescriptionCoreList(resp DescriptionDonation) donation.DescriptionCore {
	// descriptionDonation := donation.DescriptionCore{
	// 	ID:               resp.ID,
	// 	Description:      resp.Description,
	// 	Target_Donation:  resp.Target_Donation,
	// 	Current_Donation: resp.Current_Donation,
	// }
	// fmt.Println(resp)
	// fmt.Println(donation.DescriptionCore{})
	return donation.DescriptionCore{
		ID:               resp.ID,
		Description:      resp.Description,
		Target_Donation:  resp.Target_Donation,
		Current_Donation: resp.Current_Donation,
	}

}

func (d *Donation) toCore() donation.Core {

	fmt.Println(d.Description)
	return donation.Core{
		ID:       int(d.ID),
		Title:    d.Title,
		AuthorID: d.AuthorID,
		// Author:      toUsersCore(d.Author),
		Created_at:  d.CreatedAt,
		Description: toDescriptionCoreList(d.Description),
	}
}

func toCoreList(resp []Donation) []donation.Core {
	d := []donation.Core{}

	for key := range resp {
		d = append(d, resp[key].toCore())
	}
	fmt.Println(d)
	return d
}
