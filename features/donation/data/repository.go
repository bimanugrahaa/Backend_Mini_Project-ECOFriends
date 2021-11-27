package data

import (
	"Backend_Mini_Project-ECOFriends/features/donation"

	"gorm.io/gorm"
)

type mysqlDonationRepository struct {
	Conn *gorm.DB
}

func NewDonationRepository(conn *gorm.DB) donation.Data {
	return &mysqlDonationRepository{
		Conn: conn,
	}
}

func (dr *mysqlDonationRepository) InsertData(data donation.Core) (resp donation.Core, err error) {
	return donation.Core{}, nil
}

func (dr *mysqlDonationRepository) SelectAllDonations() (resp []donation.Core) {
	// record := []Donation{}
	// if err := dr.Conn.Find(&record).Error; err != nil {
	// 	return []donation.Core{}
	// }
	var record []Donation
	if err := dr.Conn.Find(&record).Error; err != nil {
		return []donation.Core{}
	}
	return toCoreList(record)
}

func (dr *mysqlDonationRepository) SelectDonationsById(id int) (resp donation.Core) {
	var record Donation
	if err := dr.Conn.Preload("Description").First(&record, id).Error; err != nil {
		return donation.Core{}
	}
	return toCoreDetail(&record)
}

// func (dr *mysqlDonationRepository) GetUserById(id int) {

// }
