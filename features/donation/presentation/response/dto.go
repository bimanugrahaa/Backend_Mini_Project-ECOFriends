package response

import (
	"Backend_Mini_Project-ECOFriends/features/donation"
	"time"
)

type Donation struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Created_at time.Time `json:"created_at"`
}

func FromCore(core donation.Core) Donation {
	return Donation{
		ID:         core.ID,
		Title:      core.Title,
		Created_at: core.Created_at,
	}
}

func FromCoreSlice(core []donation.Core) []Donation {
	var donationArray []Donation
	for key := range core {
		donationArray = append(donationArray, FromCore(core[key]))
	}

	return donationArray
}
