package bussiness

import (
	"Backend_Mini_Project-ECOFriends/features/donation"
	"Backend_Mini_Project-ECOFriends/features/donation/mocks"
	t_users "Backend_Mini_Project-ECOFriends/features/user/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// var (
// 	donationBussiness   donation.Bussiness
// 	donationValue        donation.Core
// 	donationDescription donation.DescriptionCore
// 	donationComment     donation.CommentCore
// 	donationData  mocks.donationData
// 	donations           []donation.Core
// 	// donationData
// 	// donationValue donation.Core

// )

var (
	donationValue       donation.Core
	donationUseCase     donation.Bussiness
	donationDescription donation.DescriptionCore
	donationComment     donation.CommentCore
	donationData        mocks.Data
	donationsValue      []donation.Core

	userData t_users.Bussiness
)

func TestMain(m *testing.M) {
	donationUseCase = NewDonationBussiness(&donationData, &userData)

	donationValue = donation.Core{
		ID:         1,
		Title:      "Reboisasi hutan",
		AuthorID:   1,
		Created_at: time.Now(),
		Author: donation.UserCore{
			ID:   1,
			Name: "Barry Keoghan",
		},
		Comment: []donation.CommentCore{
			{
				ID:         1,
				Comment:    "Mari dukung kegiatan reboisasi",
				PostID:     1,
				UserID:     1,
				Status:     false,
				Updated_at: time.Now(),
			},
		},
	}

	donationsValue = []donation.Core{
		{
			ID:         1,
			Title:      "Reboisasi hutan",
			AuthorID:   1,
			Created_at: time.Now(),
			Author: donation.UserCore{
				ID:   1,
				Name: "Barry Keoghan",
			},
		},
	}

	os.Exit(m.Run())
}

func TestGetAllDonations(t *testing.T) {
	t.Run("valid - get all donations", func(t *testing.T) {
		donationData.On("SelectAllDonations").Return(donationsValue).Once()
		resp := donationUseCase.GetAllDonations()

		assert.NotEqual(t, len(resp), 0)
	})
}

// type Core struct {
// 	ID          int
// 	Title       string
// 	AuthorID    int
// 	Author      UserCore
// 	Description DescriptionCore
// 	Comment     []CommentCore
// 	Created_at  time.Time
// }

// type DescriptionCore struct {
// 	ID                  int
// 	Description         string
// 	Target_Donation     int
// 	Current_Donation    int
// 	Percentage_Donation float64
// }

// type UserCore struct {
// 	ID   int
// 	Name string
// }

// type CommentCore struct {
// 	ID         int
// 	Comment    string
// 	PostID     int
// 	UserID     int
// 	Status     bool
// 	Updated_at time.Time
// }
