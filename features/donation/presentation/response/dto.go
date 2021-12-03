package response

import (
	"Backend_Mini_Project-ECOFriends/features/donation"
	"time"
)

type Donation struct {
	ID          int                 `json:"id"`
	Title       string              `json:"title"`
	Created_at  time.Time           `json:"created_at"`
	Description DonationDescription `json:"description"`
	AuthorID    int                 `json:"author_id"`
	Author      AuthorDonation      `json:"author"`
	Comment     []CommentDonation   `json:"comment"`
}

type DonationList struct {
	ID         int            `json:"id"`
	Title      string         `json:"title"`
	Created_at time.Time      `json:"created_at"`
	AuthorID   int            `json:"author_id"`
	Author     AuthorDonation `json:"author"`
}

type DonationDescription struct {
	ID               int    `json:"post_id"`
	Description      string `json:"desc"`
	Target_Donation  int    `json:"target_donation"`
	Current_Donation int    `json:"current_donation"`
}

type DonationAmount struct {
	ID                  int     `json:"post_id"`
	Current_Donation    int     `json:"current_donation"`
	Percentage_Donation float64 `json:"percentage_donation"`
}

type AuthorDonation struct {
	ID   int    `json:"author_id"`
	Name string `json:"name"`
}

type CommentDonation struct {
	ID      int    `json:"comment_id"`
	Comment string `json:"comment"`
	PostID  int    `json:"post_id"`
	UserID  int    `json:"user_id"`
	Status  bool   `json:"status"`
}

type CommentUpdateDonation struct {
	ID      int    `json:"comment_id"`
	Comment string `json:"comment"`
	PostID  int    `json:"post_id"`
	Status  bool   `json:"status"`
}

func FromDescriptionDonationCore(resp donation.DescriptionCore) DonationDescription {
	return DonationDescription{
		ID:               resp.ID,
		Description:      resp.Description,
		Target_Donation:  resp.Target_Donation,
		Current_Donation: resp.Current_Donation,
	}
}

func FromDonationAmount(resp donation.DescriptionCore) DonationAmount {
	return DonationAmount{
		ID:                  resp.ID,
		Current_Donation:    resp.Current_Donation,
		Percentage_Donation: resp.Percentage_Donation,
	}
}

func FromUserCore(uc donation.UserCore) AuthorDonation {
	return AuthorDonation{
		ID:   uc.ID,
		Name: uc.Name,
	}
}

func FromCommentCore(cc donation.CommentCore) CommentDonation {
	return CommentDonation{
		ID:      cc.ID,
		Comment: cc.Comment,
		PostID:  cc.PostID,
		UserID:  cc.UserID,
		Status:  cc.Status,
	}
}

func FromCommentUpdateCore(cc donation.CommentCore) CommentUpdateDonation {
	return CommentUpdateDonation{
		ID:      cc.ID,
		Comment: cc.Comment,
		PostID:  cc.PostID,
		Status:  cc.Status,
	}
}

func FromCore(core donation.Core) DonationList {

	return DonationList{
		ID:         core.ID,
		Title:      core.Title,
		AuthorID:   core.AuthorID,
		Author:     FromUserCore(core.Author),
		Created_at: core.Created_at,
	}
}

func FromCoreDetail(core donation.Core) Donation {
	return Donation{
		ID:          core.ID,
		Title:       core.Title,
		AuthorID:    core.AuthorID,
		Author:      FromUserCore(core.Author),
		Created_at:  core.Created_at,
		Description: FromDescriptionDonationCore(core.Description),
		Comment:     FromCommentSlice(core.Comment),
	}
}

func FromCommentSlice(cc []donation.CommentCore) []CommentDonation {
	var commentArray []CommentDonation
	for key := range cc {
		commentArray = append(commentArray, FromCommentCore(cc[key]))
	}

	return commentArray
}

func FromCoreSlice(core []donation.Core) []DonationList {
	var donationArray []DonationList
	for key := range core {
		donationArray = append(donationArray, FromCore(core[key]))
	}

	return donationArray
}
