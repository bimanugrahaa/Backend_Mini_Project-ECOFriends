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
	if err := dr.Conn.Delete(&DescriptionDonation{}, id).Error; err != nil {
		return err
	}
	if err := dr.Conn.Delete(&Donation{}, id).Error; err != nil {
		return err
	}
	if err := dr.Conn.Where("post_id", id).Delete(&CommentDonation{}).Error; err != nil {
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

	if err := dr.Conn.Model(&Donation{}).Find(&record).Error; err != nil {
		return []donation.Core{}
	}

	return toCoreList(record)
}

func (dr *mysqlDonationRepository) SelectDonationsById(id int) (resp donation.Core) {
	var record Donation

	if err := dr.Conn.First(&record, id).Error; err != nil {
		return donation.Core{}
	}

	return toCoreDetail(&record)
}

func (dr *mysqlDonationRepository) InsertComment(id int, data donation.CommentCore) (resp donation.CommentCore, err error) {

	record := fromCommentCore(id, data)
	fmt.Println("record", record)
	if err := dr.Conn.Model(&CommentDonation{}).Create(&record).Error; err != nil {

		return donation.CommentCore{}, err
	}

	fmt.Println(dr.Conn.Create(&record))
	return donation.CommentCore{}, nil
}

func (dr *mysqlDonationRepository) SelectCommentByPostId(id int) (resp []donation.CommentCore, err error) {

	var record []CommentDonation

	// fmt.Println(&resp)
	if err := dr.Conn.Model(&CommentDonation{}).Where("post_id = ?", id).Find(&record).Error; err != nil {
		return []donation.CommentCore{}, err
	}
	fmt.Println(record)

	return toCommentList(record), nil
}

func (dr *mysqlDonationRepository) EditComment(data donation.CommentCore) (resp donation.CommentCore, err error) {
	record := fromCommentCore(data.PostID, data)

	if err := dr.Conn.Model(&CommentDonation{}).Where("id = ?", data.ID).Updates(&record).Error; err != nil {
		return donation.CommentCore{}, err
	}

	return donation.CommentCore{}, nil
}

func (dr *mysqlDonationRepository) RemoveComment(id int) (err error) {
	if err := dr.Conn.Delete(&CommentDonation{}, id).Error; err != nil {
		return err
	}

	return nil
}
