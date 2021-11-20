package bussiness

import (
	"Backend_Mini_Project-ECOFriends/features/donation"

	"github.com/go-playground/validator/v10"
)

type donationUsecase struct {
	donationData donation.Data
	validate     *validator.Validate
}

func NewDonationBussiness(donationData donation.Data) donation.Bussiness {
	return &donationUsecase{
		donationData: donationData,
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

func (du *donationUsecase) GetAllData(search string) (resp []donation.Core) {
	resp = du.donationData.SelectData(search)
	return
}
