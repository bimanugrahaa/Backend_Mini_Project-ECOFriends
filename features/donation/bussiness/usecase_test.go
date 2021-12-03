package bussiness

import (
	"Backend_Mini_Project-ECOFriends/features/donation"
	"Backend_Mini_Project-ECOFriends/features/donation/mocks"
	"Backend_Mini_Project-ECOFriends/features/user"
	t_users "Backend_Mini_Project-ECOFriends/features/user/mocks"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	donationValue       donation.Core
	donationUseCase     donation.Bussiness
	donationDescription donation.DescriptionCore
	donationComment     donation.CommentCore
	donationAuthor      donation.UserCore
	donationData        mocks.Data
	donationsValue      []donation.Core
	donationComments    []donation.CommentCore

	userData t_users.Bussiness
)

func TestMain(m *testing.M) {
	donationUseCase = NewDonationBussiness(&donationData, &userData)

	donationValue = donation.Core{
		ID:       1,
		Title:    "Reboisasi hutan",
		AuthorID: 1,
		Author: donation.UserCore{
			ID:   1,
			Name: "Barry Keoghan",
		},
		Description: donation.DescriptionCore{
			ID:                  1,
			Description:         "Reboisasi hutan sebagai paru-paru dunia",
			Target_Donation:     100000000,
			Current_Donation:    50000000,
			Percentage_Donation: 50.00,
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
		Created_at: time.Now(),
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

	donationDescription = donation.DescriptionCore{
		ID:                  1,
		Description:         "Reboisasi hutan sebagai paru-paru dunia",
		Target_Donation:     100000000,
		Current_Donation:    50000000,
		Percentage_Donation: 50.00,
	}

	donationComments = []donation.CommentCore{
		{
			ID:         1,
			Comment:    "Mari dukung kegiatan reboisasi",
			PostID:     1,
			UserID:     1,
			Status:     false,
			Updated_at: time.Now(),
		},
	}

	os.Exit(m.Run())
}

func TestGetAllDonations(t *testing.T) {
	t.Run("valid - get all donations", func(t *testing.T) {

		data := []int{1}
		userData.On("GetUserById", mock.AnythingOfType("int")).Return(user.UserCore{}, nil).Times(len(data))

		donationData.On("SelectAllDonations").Return(donationsValue).Once()

		resp := donationUseCase.GetAllDonations()

		assert.NotEqual(t, len(resp), 0)
	})
}

func TestGetDonationTrending(t *testing.T) {
	t.Run("valid - get donations trending", func(t *testing.T) {

		data := []int{1}
		userData.On("GetUserById", mock.AnythingOfType("int")).Return(user.UserCore{}, nil).Times(len(data))

		donationData.On("SelectDonationsTrending").Return(donationsValue).Once()

		resp := donationUseCase.GetDonationTrending()

		assert.NotEqual(t, len(resp), 0)
	})
}

func TestGetDonationLatest(t *testing.T) {
	t.Run("valid - get donations latest", func(t *testing.T) {

		data := []int{1}
		userData.On("GetUserById", mock.AnythingOfType("int")).Return(user.UserCore{}, nil).Times(len(data))

		donationData.On("SelectDonationsLatest").Return(donationsValue).Once()

		resp := donationUseCase.GetDonationLatest()

		assert.NotEqual(t, len(resp), 0)
	})
}

func TestGetDonationsById(t *testing.T) {
	t.Run("valid - get donations by id", func(t *testing.T) {

		data := []int{1}
		userData.On("GetUserById", mock.AnythingOfType("int")).Return(user.UserCore{}, nil).Times(len(data))
		donationData.On("SelectCommentByPostId", mock.AnythingOfType("int")).Return([]donation.CommentCore{}, nil).Once()
		donationData.On("SelectDonationsById", mock.AnythingOfType("int")).Return(donationValue, nil).Once()

		resp := donationUseCase.GetDonationsById(donationValue.ID)

		assert.Equal(t, resp.ID, donationValue.ID)
	})
}
func TestCreateDonation(t *testing.T) {
	t.Run("valid - create donation", func(t *testing.T) {
		data := []int{1}
		userData.On("GetUserById", mock.AnythingOfType("int")).Return(user.UserCore{}, nil).Times(len(data))

		donationData.On("InsertDonation", mock.AnythingOfType("donation.Core")).Return(donationValue, nil)

		resp, err := donationUseCase.CreateDonation(donationValue)

		assert.Nil(t, err)
		assert.Equal(t, resp.ID, donationValue.ID)
	})

	t.Run("invalid - error create donation", func(t *testing.T) {
		data := []int{1}
		userData.On("GetUserById", mock.AnythingOfType("int")).Return(user.UserCore{}, errors.New("error create donation")).Times(len(data))

		donationData.On("InsertDonation", mock.AnythingOfType("donation.Core")).Return(donationValue, errors.New("error create donation")).Once()

		// _, err := donationUseCase.CreateDonation(donation.Core{
		// 	Title: "Reboisasi",
		// })

		// assert.Nil(t, err)
		// assert.Equal(t, err.Error(), "error create donation")
	})
}

func TestDeleteDonationsById(t *testing.T) {
	t.Run("valid - delete donation by id", func(t *testing.T) {
		donationData.On("SelectDonationsById", mock.AnythingOfType("int")).Return(donationValue, nil).Once()
		donationData.On("RemoveDonationsById", mock.AnythingOfType("donation.Core")).Return(nil).Once()

		err := donationUseCase.DeleteDonationsById(donationValue)
		assert.Nil(t, err)
	})

	t.Run("invalid - error delete donation by id", func(t *testing.T) {
		donationData.On("SelectDonationsById", mock.AnythingOfType("int")).Return(donationValue, nil).Once()
		donationData.On("RemoveDonationsById", mock.AnythingOfType("donation.Core")).Return(errors.New("error delete donation")).Once()

		err := donationUseCase.DeleteDonationsById(donationValue)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "error delete donation")
	})

	t.Run("invalid - error delete donation unauthorized", func(t *testing.T) {
		donationData.On("SelectDonationsById", mock.AnythingOfType("int")).Return(donationValue, nil).Once()

		donationData.On("RemoveDonationsById", mock.AnythingOfType("donation.Core")).Return(errors.New("unauthorized")).Once()

		err := donationUseCase.DeleteDonationsById(donation.Core{
			ID: 5,
		})
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "unauthorized")
	})
}

func TestUpdateDonation(t *testing.T) {
	t.Run("valid - update donation", func(t *testing.T) {

		donationData.On("SelectDonationsById", mock.AnythingOfType("int")).Return(donationValue).Once()
		donationData.On("EditDonation", mock.AnythingOfType("donation.Core")).Return(donationValue, nil).Once()

		userData.On("GetUserById", mock.AnythingOfType("int")).Return(user.UserCore{}, nil).Once()

		resp, err := donationUseCase.UpdateDonation(donation.Core{
			AuthorID: 1,
			Title:    "Reboisasi",
		})

		assert.Nil(t, err)
		assert.Equal(t, resp.Title, donationValue.Title)
	})

	t.Run("invalid - error update donation unauthorized", func(t *testing.T) {
		donationData.On("SelectDonationsById", mock.AnythingOfType("int")).Return(donationValue).Once()
		data := []int{1}
		userData.On("GetUserById", mock.AnythingOfType("int")).Return(user.UserCore{}, errors.New("unauthorized")).Times(len(data))

		donationData.On("EditDonation", mock.AnythingOfType("donation.Core")).Return(donationValue, nil).Once()
		_, err := donationUseCase.UpdateDonation(donation.Core{
			Title: "Reboisasi",
		})

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "unauthorized")

	})

	// t.Run("invalid - error update donation", func(t *testing.T) {
	// 	userData.On("GetUserById", mock.AnythingOfType("int")).Return(user.UserCore{}, nil).Once()
	// 	donationData.On("GetUserById", mock.AnythingOfType("int")).Return(user.UserCore{}, nil).Once()

	// 	donationData.On("SelectDonationsById", mock.AnythingOfType("int")).Return(donationValue).Once()
	// 	donationData.On("EditDonation", mock.AnythingOfType("donation.Core")).Return(donationValue, errors.New("error update donation")).Once()

	// 	resp, err := donationUseCase.UpdateDonation(donation.Core{
	// 		Title:    "Reboisasi",
	// 		AuthorID: 1,
	// 	})

	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, err.Error(), "error update donation")
	// 	assert.NotEqual(t, resp.Title, donationValue.Title)
	// })
}

func TestUpdateDonationValue(t *testing.T) {
	t.Run("valid - update donation value", func(t *testing.T) {
		donationData.On("SelectDonationsById", mock.AnythingOfType("int")).Return(donationValue).Once()
		donationData.On("EditDonationValue", mock.AnythingOfType("donation.DescriptionCore")).Return(donationDescription, nil).Once()

		resp, err := donationUseCase.UpdateDonationValue(donation.DescriptionCore{
			Current_Donation: 100000,
		})

		assert.Nil(t, err)
		assert.Equal(t, resp.Target_Donation, donationDescription.Target_Donation)
	})

	t.Run("invalid - error update donation value", func(t *testing.T) {
		donationData.On("SelectDonationsById", mock.AnythingOfType("int")).Return(donationValue).Once()
		donationData.On("EditDonationValue", mock.AnythingOfType("donation.DescriptionCore")).Return(donationDescription, errors.New("error update current donation")).Once()

		resp, err := donationUseCase.UpdateDonationValue(donation.DescriptionCore{
			Current_Donation: 100000,
		})

		assert.NotNil(t, err)
		assert.NotEqual(t, resp.Target_Donation, donationDescription.Target_Donation)
	})
}

func TestCreateComment(t *testing.T) {
	t.Run("valid - create comment", func(t *testing.T) {
		donationData.On("InsertComment", mock.AnythingOfType("int"), mock.AnythingOfType("donation.CommentCore")).Return(donationComment, nil).Once()

		resp, err := donationUseCase.CreateComment(donationComment.PostID, donation.CommentCore{
			Comment: "Bagus",
		})

		assert.Nil(t, err)
		assert.Equal(t, resp.Comment, donationComment.Comment)

	})

	t.Run("invalid - error create comment", func(t *testing.T) {
		donationData.On("InsertComment", mock.AnythingOfType("int"), mock.AnythingOfType("donation.CommentCore")).Return(donationComment, errors.New("error create comment")).Once()

		_, err := donationUseCase.CreateComment(donationComment.PostID, donation.CommentCore{
			Comment: "Bagus",
		})

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "error create comment")

	})
}

func TestGetCommentById(t *testing.T) {
	t.Run("valid - get comment by id", func(t *testing.T) {
		donationData.On("SelectCommentByPostId", mock.AnythingOfType("int")).Return(donationComments, nil).Once()

		resp, err := donationUseCase.GetCommentByPostId(donationComment.PostID)

		assert.Nil(t, err)
		assert.NotEqual(t, len(resp), 0)
	})
}

func TestUpdateComment(t *testing.T) {
	t.Run("valid - update comment", func(t *testing.T) {
		donationData.On("SelectCommentById", mock.AnythingOfType("int")).Return(donationComment, nil).Once()
		donationData.On("EditComment", mock.AnythingOfType("donation.CommentCore")).Return(donationComment, nil).Once()

		resp, err := donationUseCase.UpdateComment(donation.CommentCore{
			Comment: "Amazing",
		})

		assert.Nil(t, err)
		assert.Equal(t, resp.Comment, donationComment.Comment)
	})

	t.Run("invalid - unauthorized update comment", func(t *testing.T) {
		donationData.On("SelectCommentById", mock.AnythingOfType("int")).Return(donationComment, nil).Once()
		donationData.On("EditComment", mock.AnythingOfType("donation.CommentCore")).Return(donationComment, errors.New("unauthorized")).Once()

		_, err := donationUseCase.UpdateComment(donation.CommentCore{
			UserID:  5,
			Comment: "Amazing",
		})

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "unauthorized")
	})

	t.Run("invalid - error update comment", func(t *testing.T) {
		donationData.On("SelectCommentById", mock.AnythingOfType("int")).Return(donationComment, nil).Once()
		donationData.On("EditComment", mock.AnythingOfType("donation.CommentCore")).Return(donationComment, errors.New("error update comment")).Once()

		_, err := donationUseCase.UpdateComment(donation.CommentCore{
			Comment: "Amazing",
		})

		assert.NotNil(t, err)
	})
}

func TestDeleteComment(t *testing.T) {
	t.Run("valid - delete comment", func(t *testing.T) {
		donationData.On("SelectCommentById", mock.AnythingOfType("int")).Return(donationComment, nil).Once()
		donationData.On("RemoveComment", mock.AnythingOfType("donation.CommentCore")).Return(nil)

		err := donationUseCase.DeleteComment(donationComment)

		assert.Nil(t, err)
	})

	t.Run("invalid - unauthorized delete comment", func(t *testing.T) {
		donationData.On("SelectCommentById", mock.AnythingOfType("int")).Return(donationComment, nil).Once()
		donationData.On("RemoveComment", mock.AnythingOfType("donation.CommentCore")).Return(errors.New("unauthorized"))

		err := donationUseCase.DeleteComment(donation.CommentCore{
			UserID: 5,
		})

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "unauthorized")
	})

	// t.Run("invalid - error delete comment", func(t *testing.T) {
	// 	donationData.On("SelectCommentById", mock.AnythingOfType("int")).Return(donationComment, nil).Once()
	// 	donationData.On("RemoveComment", mock.AnythingOfType("donation.CommentCore")).Return(errors.New("error delete comment"))

	// 	err := donationUseCase.DeleteComment(donationComment)

	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, err.Error(), "error delete comment")
	// })

}
