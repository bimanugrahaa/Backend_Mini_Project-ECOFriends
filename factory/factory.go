package factory

import (
	"Backend_Mini_Project-ECOFriends/config"
	_donation_bussiness "Backend_Mini_Project-ECOFriends/features/donation/bussiness"
	_donation_data "Backend_Mini_Project-ECOFriends/features/donation/data"
	_donation_presentation "Backend_Mini_Project-ECOFriends/features/donation/presentation"
)

type Presenter struct {
	DonationPresentation *_donation_presentation.DonationHandler
}

func Init() Presenter {

	donationData := _donation_data.NewDonationRepository(config.DB)
	donationBussiness := _donation_bussiness.NewDonationBussiness(donationData)

	return Presenter{
		DonationPresentation: _donation_presentation.NewDonationHandler(donationBussiness),
	}
}
