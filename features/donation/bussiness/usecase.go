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

func (du *donationUsecase) CreateData(data donation.Core) (resp donation.Core, err error) {
	// if err := du.validate.Struct(data); err != nil {
	// 	return donation.Core{}, err
	// }

	// resp, err = du.donationData.InsertData(data)
	// if err != nil {
	// 	return donation.Core{}, err
	// }

	// return resp, nil

	return donation.Core{}, nil
}

func (du *donationUsecase) GetAllData() (resp []donation.Core) {
	resp = du.donationData.SelectData()
	fmt.Println("'data'", du.userData.GetAllUser())

	for key, value := range resp {
		user, _ := du.userData.GetUserById(value.AuthorID)
		fmt.Println("'value'", value)
		resp[key].Author.ID = user.ID
		resp[key].Author.Name = user.Name

	}

	fmt.Println("'resp'", resp)
	return
}

// func (du *donationUsecase) GetUserById() (resp []donation.UserCore) {
// 	// user := du.userData.GetUserById()
// 	// user
// 	// resp = du.userData.GetAllUser()

// 	// for key := range user {
// 	// 	user = append(user, user[key])
// 	// }
// 	// return user
// }
