package response

import (
	"Backend_Mini_Project-ECOFriends/features/donation"
	"time"
)

type Donation struct {
	ID          int                 `json:"id"`
	Title       string              `json:"title"`
	Created_at  time.Time           `json:"created_at"`
	Description DescriptionDonation `json:"desc"`
}

type DescriptionDonation struct {
	ID               int    `json:"post_id"`
	Description      string `json:"desc"`
	Target_Donation  int    `json:"target_donation"`
	Current_Donation int    `json:"current_donation"`
}

// func FromDescriptionDonationCore(core donation.DescriptionCore) DescriptionDonation {

// 	return DescriptionDonation{
// 		ID:               core.ID,
// 		Description:      core.Description,
// 		Target_Donation:  core.Target_Donation,
// 		Current_Donation: core.Current_Donation,
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

func FromCore(core donation.Core) Donation {

	return Donation{
		ID:          core.ID,
		Title:       core.Title,
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
