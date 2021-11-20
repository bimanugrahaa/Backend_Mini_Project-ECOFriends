package migrate

import (
	"Backend_Mini_Project-ECOFriends/config"
	m_donation "Backend_Mini_Project-ECOFriends/features/donation/data"
)

func AutoMigrate() {
	config.DB.AutoMigrate(
		&m_donation.Donation{},
		&m_donation.DescriptionDonation{})
}

//To migrate two and up
// func AutoMigrate() {
// 	config.DB.AutoMigrate(&m_donation.Donation{}, &m_user.User{})
// }

// var (
// 	mCreateTableUser createTableUser = createTableUser{config.DB.Migrator()}
// )

// func MigrateUp() {
// 	mCreateTableUser.up()
// }
