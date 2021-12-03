package migrate

import (
	"Backend_Mini_Project-ECOFriends/config"
	m_donation "Backend_Mini_Project-ECOFriends/features/donation/data"
	m_user "Backend_Mini_Project-ECOFriends/features/user/data"
)

func AutoMigrate() {
	config.DB.AutoMigrate(
		&m_donation.Donation{},
		&m_donation.DescriptionDonation{},
		&m_donation.CommentDonation{},
		&m_user.User{})
}
