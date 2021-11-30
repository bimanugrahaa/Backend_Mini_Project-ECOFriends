package bussiness

import (
	"Backend_Mini_Project-ECOFriends/features/donation"
	"Backend_Mini_Project-ECOFriends/features/user"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type donationUsecase struct {
	donationData donation.Data
	userData     user.Bussiness
	validate     *validator.Validate
}

func NewDonationBussiness(donationData donation.Data, userData user.Bussiness) donation.Bussiness {
	return &donationUsecase{
		donationData: donationData,
		userData:     userData,
		validate:     validator.New(),
	}
}

func (du *donationUsecase) CreateDonation(data donation.Core) (resp donation.Core, err error) {
	// if err := du.validate.Struct(data); err != nil {
	// 	return donation.Core{}, err
	// }

	resp, err = du.donationData.InsertDonation(data)
	user, _ := du.userData.GetUserById(resp.AuthorID)
	resp.Author.ID = user.ID
	resp.Author.Name = user.Name

	if err != nil {
		return donation.Core{}, err
	}

	return resp, nil
}

func (du *donationUsecase) DeleteDonationsById(id int, data donation.Core) (err error) {

	// resp := du.GetDonationsById(id)
	resp := du.donationData.SelectDonationsById(data.ID)
	fmt.Println("resp", resp)
	err = du.donationData.RemoveDonationsById(id, resp)

	if err != nil {
		return err
	}

	return nil
}

func (du *donationUsecase) UpdateDonation(data donation.Core) (resp donation.Core, err error) {
	resp, err = du.donationData.EditDonation(data)
	user, _ := du.userData.GetUserById(resp.AuthorID)
	resp.Author.ID = user.ID
	resp.Author.Name = user.Name

	if err != nil {
		return donation.Core{}, err
	}

	return resp, nil
}

func (du *donationUsecase) GetAllDonations() (resp []donation.Core) {
	resp = du.donationData.SelectAllDonations()

	for key, value := range resp {
		user, _ := du.userData.GetUserById(value.AuthorID)
		resp[key].Author.ID = user.ID
		resp[key].Author.Name = user.Name
	}

	return
}

func (du *donationUsecase) GetDonationsById(id int) (resp donation.Core) {
	resp = du.donationData.SelectDonationsById(id)

	user, _ := du.userData.GetUserById(resp.AuthorID)
	comment, _ := du.donationData.SelectCommentByPostId(id)
	resp.Author.ID = user.ID
	resp.Author.Name = user.Name
	resp.Comment = comment

	return
}

func (du *donationUsecase) CreateComment(id int, data donation.CommentCore) (resp donation.CommentCore, err error) {
	// if err := du.validate.Struct(data); err != nil {
	// 	return donation.Core{}, err
	// }

	resp, err = du.donationData.InsertComment(id, data)
	if err != nil {
		return donation.CommentCore{}, err
	}

	return resp, nil
}

func (du *donationUsecase) GetCommentByPostId(id int) (resp []donation.CommentCore, err error) {
	resp, err = du.donationData.SelectCommentByPostId(id)
	return
}

func (du *donationUsecase) UpdateComment(data donation.CommentCore) (resp donation.CommentCore, err error) {
	resp, err = du.donationData.EditComment(data)

	if err != nil {
		return donation.CommentCore{}, err
	}

	return resp, nil
}

func (du *donationUsecase) DeleteComment(id int) (err error) {
	err = du.donationData.RemoveComment(id)

	if err != nil {
		return err
	}

	return nil
}
