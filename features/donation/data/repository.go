package data

import (
	"Backend_Mini_Project-ECOFriends/features/donation"
	"fmt"

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

func (dr *mysqlDonationRepository) InsertDonation(data donation.Core) (resp donation.Core, err error) {

	record := fromCore(data)
	fmt.Println("record", record)
	if err := dr.Conn.Create(&record).Error; err != nil {

		return donation.Core{}, err
	}

	fmt.Println(dr.Conn.Create(&record))
	return donation.Core{}, nil
}

func (dr *mysqlDonationRepository) RemoveDonationsById(id int) (err error) {
	if err := dr.Conn.Delete(&Donation{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (dr *mysqlDonationRepository) EditDonation(data donation.Core) (resp donation.Core, err error) {
	record := fromCore(data)

	if err := dr.Conn.Model(&Donation{}).Where("id = ?", data.ID).Updates(&record).Error; err != nil {
		return donation.Core{}, err
	}

	if err := dr.Conn.Model(&DescriptionDonation{}).Where("id = ?", data.ID).Updates(&record.Description).Error; err != nil {
		return donation.Core{}, err
	}

	return donation.Core{}, nil
}

func (dr *mysqlDonationRepository) SelectAllDonations() (resp []donation.Core) {

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
