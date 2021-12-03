package factory

import (
	"Backend_Mini_Project-ECOFriends/config"
	_donation_bussiness "Backend_Mini_Project-ECOFriends/features/donation/bussiness"
	_donation_data "Backend_Mini_Project-ECOFriends/features/donation/data"
	_donation_presentation "Backend_Mini_Project-ECOFriends/features/donation/presentation"

	_user_bussiness "Backend_Mini_Project-ECOFriends/features/user/bussiness"
	_user_data "Backend_Mini_Project-ECOFriends/features/user/data"
	_user_presentation "Backend_Mini_Project-ECOFriends/features/user/presentation"
)

type Presenter struct {
	DonationPresentation *_donation_presentation.DonationHandler
	UserPresentation     *_user_presentation.UserHandler
}

func Init() Presenter {

	//User
	userData := _user_data.NewUserRepository(config.DB)
	userBussiness := _user_bussiness.NewUserBussiness(userData)
	userPresentation := _user_presentation.NewUserHandler(userBussiness)

	//Donation
	donationData := _donation_data.NewDonationRepository(config.DB)
	donationBussiness := _donation_bussiness.NewDonationBussiness(donationData, userBussiness)
	donationPresentation := _donation_presentation.NewDonationHandler(donationBussiness)

	return Presenter{
		DonationPresentation: donationPresentation,
		UserPresentation:     userPresentation,
	}
}
