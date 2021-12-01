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
	if err := dr.Conn.Create(&record).Error; err != nil {

		return donation.Core{}, err
	}

	return toCoreDetail(&record), nil
}

func (dr *mysqlDonationRepository) RemoveDonationsById(data donation.Core) (err error) {

	if err := dr.Conn.Delete(&DescriptionDonation{}, data.ID).Error; err != nil {
		return err
	}
	if err := dr.Conn.Delete(&Donation{}, data.ID).Error; err != nil {
		return err
	}
	if err := dr.Conn.Where("post_id", data.ID).Delete(&CommentDonation{}).Error; err != nil {
		return err
	}
	return nil
}

func (dr *mysqlDonationRepository) EditDonation(data donation.Core) (resp donation.Core, err error) {
	record := fromCore(data)
	record.Description.ID = data.ID
	fmt.Println(record)

	if err := dr.Conn.Model(&Donation{}).Where("id = ?", data.ID).Updates(&record).Error; err != nil {
		return donation.Core{}, err
	}

	if err := dr.Conn.Model(&DescriptionDonation{}).Where("id = ?", data.ID).Updates(&record.Description).Error; err != nil {
		return donation.Core{}, err
	}

	return toCoreDetail(&record), nil
}

func (dr *mysqlDonationRepository) EditDonationValue(data donation.DescriptionCore) (resp donation.DescriptionCore, err error) {
	record := fromDescriptionCore(data)
	record.ID = data.ID

	if err := dr.Conn.Model(&DescriptionDonation{}).Where("id = ?", data.ID).Updates(&record).Error; err != nil {
		fmt.Println("err", err)
	}

	return toDescriptionCore(&record), nil
}

func (dr *mysqlDonationRepository) SelectAllDonations() (resp []donation.Core) {

	var record []Donation

	if err := dr.Conn.Model(&Donation{}).Find(&record).Error; err != nil {
		return []donation.Core{}
	}

	return toCoreList(record)
}

func (dr *mysqlDonationRepository) SelectDonationsTrending() (resp []donation.Core) {

	var record []Donation

	if err := dr.Conn.Model(&Donation{}).Joins("Description").Order("description.donation_percentage desc").Limit(10).Find(&record).Error; err != nil {

		return []donation.Core{}
	}

	return toCoreList(record)
}

func (dr *mysqlDonationRepository) SelectDonationsLatest() (resp []donation.Core) {

	var record []Donation

	if err := dr.Conn.Model(&Donation{}).Order("created_at desc").Limit(10).Find(&record).Error; err != nil {

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

func (dr *mysqlDonationRepository) InsertComment(id int, data donation.CommentCore) (resp donation.CommentCore, err error) {

	record := fromCommentCore(id, data)
	if err := dr.Conn.Model(&CommentDonation{}).Create(&record).Error; err != nil {

		return donation.CommentCore{}, err
	}

	return toCommentCore(&record), nil
}

func (dr *mysqlDonationRepository) SelectCommentByPostId(id int) (resp []donation.CommentCore, err error) {

	var record []CommentDonation

	if err := dr.Conn.Model(&CommentDonation{}).Where("post_id = ?", id).Find(&record).Error; err != nil {
		return []donation.CommentCore{}, err
	}

	return toCommentList(record), nil
}

func (dr *mysqlDonationRepository) SelectCommentById(id int) (resp donation.CommentCore, err error) {
	var record CommentDonation

	if err := dr.Conn.Model(&CommentDonation{}).First(&record, id).Error; err != nil {
		return donation.CommentCore{}, err
	}

	return toCommentCore(&record), nil
}

func (dr *mysqlDonationRepository) EditComment(data donation.CommentCore) (resp donation.CommentCore, err error) {
	record := fromCommentCore(data.PostID, data)
	record.ID = data.ID

	if err := dr.Conn.Model(&CommentDonation{}).Where("id = ?", data.ID).Updates(&record).Error; err != nil {
		return donation.CommentCore{}, err
	}

	return toCommentCore(&record), nil
}

func (dr *mysqlDonationRepository) RemoveComment(data donation.CommentCore) (err error) {
	if err := dr.Conn.Delete(&CommentDonation{}, data.ID).Error; err != nil {
		return err
	}

	return nil
}
