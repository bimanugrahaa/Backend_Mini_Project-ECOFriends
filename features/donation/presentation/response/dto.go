package response

import (
	"Backend_Mini_Project-ECOFriends/features/donation"
	"fmt"
	"time"
)

type Donation struct {
	ID          int                 `json:"id"`
	Title       string              `json:"title"`
	AuthorID    int                 `json:"author_id"`
	Created_at  time.Time           `json:"created_at"`
	Description DescriptionDonation `json:"desc"`
	Author      AuthorDonation      `json:"author"`
}

//Add author from user from donation.bussiness
type DescriptionDonation struct {
	ID               int    `json:"post_id"`
	Description      string `json:"desc"`
	Target_Donation  int    `json:"target_donation"`
	Current_Donation int    `json:"current_donation"`
}

type AuthorDonation struct {
	ID   int    `json:"author_id"`
	Name string `json:"name"`
}

// func FromDescriptionDonationCore(core donation.DescriptionCore) DescriptionDonation {

// 	return DescriptionDonation{
// 		ID:               core.ID,
// 		Description:      core.Description,
// 		Target_Donation:  core.Target_Donation,
// 		Current_Donation: core.Current_Donation,
// 	}
// }

// func FromUsers(resp donation.Bussiness) AuthorDonation {
// 	AuthorDonation = resp.GetAllData()
// 	return AuthorDonation{
// 		ID: resp.GetAllData(),
// 	}
// }

func FromDescriptionDonationCoreList(resp donation.DescriptionCore) DescriptionDonation {
	// fmt.Println(resp)

	return DescriptionDonation{
		ID:               resp.ID,
		Description:      resp.Description,
		Target_Donation:  resp.Target_Donation,
		Current_Donation: resp.Current_Donation,
	}
}

func FromUserCore(uc donation.UserCore) AuthorDonation {
	fmt.Println("uc", uc)
	return AuthorDonation{
		ID:   uc.ID,
		Name: uc.Name,
	}
}

func FromCore(core donation.Core) Donation {

	return Donation{
		ID:          core.ID,
		Title:       core.Title,
		AuthorID:    core.AuthorID,
		Author:      FromUserCore(core.Author),
		Created_at:  core.Created_at,
		Description: FromDescriptionDonationCoreList(core.Description),
	}
}

func FromCoreSlice(core []donation.Core) []Donation {
	var donationArray []Donation
	for key := range core {
		donationArray = append(donationArray, FromCore(core[key]))
	}

	return donationArray
}
