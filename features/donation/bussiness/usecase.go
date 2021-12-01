package bussiness

import (
	"Backend_Mini_Project-ECOFriends/features/donation"
	"Backend_Mini_Project-ECOFriends/features/user"
	"errors"
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

func (du *donationUsecase) DeleteDonationsById(data donation.Core) (err error) {

	resp := du.donationData.SelectDonationsById(data.ID)
	if data.AuthorID != resp.AuthorID {
		err = errors.New("unauthorized")
		return err
	}
	err = du.donationData.RemoveDonationsById(resp)

	if err != nil {
		return err
	}

	return nil
}

func (du *donationUsecase) UpdateDonation(data donation.Core) (resp donation.Core, err error) {
	donationAuthor := du.donationData.SelectDonationsById(data.ID)
	if data.AuthorID != donationAuthor.AuthorID {
		fmt.Println("resp", resp)

		fmt.Println(data)
		err = errors.New("unauthorized")
		return donation.Core{}, err
	}
	resp, err = du.donationData.EditDonation(data)

	user, _ := du.userData.GetUserById(donationAuthor.AuthorID)
	resp.Author.ID = user.ID
	resp.Author.Name = user.Name

	if err != nil {
		return donation.Core{}, err
	}

	return resp, nil
}

func (du *donationUsecase) UpdateDonationValue(data donation.DescriptionCore) (resp donation.DescriptionCore, err error) {
	dn := du.donationData.SelectDonationsById(data.ID)

	//Update current donation and percentage
	dn.Description.Current_Donation = dn.Description.Current_Donation + data.Current_Donation
	data.Current_Donation = dn.Description.Current_Donation
	data.Percentage_Donation = (float64(dn.Description.Current_Donation) / float64(dn.Description.Target_Donation)) * 100.00
	resp, err = du.donationData.EditDonationValue(data)

	//Assign data to resp
	resp.ID = data.ID
	resp.Current_Donation = data.Current_Donation
	resp.Percentage_Donation = data.Percentage_Donation

	if err != nil {
		return donation.DescriptionCore{}, err
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

func (du *donationUsecase) GetDonationTrending() (resp []donation.Core) {
	resp = du.donationData.SelectDonationsTrending()

	for key, value := range resp {
		user, _ := du.userData.GetUserById(value.AuthorID)
		resp[key].Author.ID = user.ID
		resp[key].Author.Name = user.Name
	}

	return
}

func (du *donationUsecase) GetDonationLatest() (resp []donation.Core) {
	resp = du.donationData.SelectDonationsLatest()

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
	donationComment, _ := du.donationData.SelectCommentById(data.ID)

	if data.UserID != donationComment.UserID {
		err = errors.New("unauthorized")
		return donation.CommentCore{}, err
	}

	resp, err = du.donationData.EditComment(data)

	if err != nil {
		return donation.CommentCore{}, err
	}

	return resp, nil
}

func (du *donationUsecase) DeleteComment(data donation.CommentCore) (err error) {
	donationComment, _ := du.donationData.SelectCommentById(data.ID)

	if data.UserID != donationComment.UserID {
		err = errors.New("unauthorized")
		return err
	}

	err = du.donationData.RemoveComment(donationComment)

	if err != nil {
		return err
	}

	return nil
}
