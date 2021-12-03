package request

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
}

type DonationDescription struct {
	Description      string `json:"desc"`
	Target_Donation  int    `json:"target_donation"`
	Current_Donation int    `json:"current_donation"`
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

func ToCore(req Donation) donation.Core {
	return donation.Core{
		ID:          req.ID,
		Title:       req.Title,
		AuthorID:    req.AuthorID,
		Description: ToDescriptionCore(req.Description),
	}
}

func ToDescriptionCore(req DonationDescription) donation.DescriptionCore {
	return donation.DescriptionCore{
		Description:      req.Description,
		Target_Donation:  req.Target_Donation,
		Current_Donation: req.Current_Donation,
	}
}

func ToCommentCore(id int, req CommentDonation) donation.CommentCore {
	return donation.CommentCore{
		ID:      req.ID,
		Comment: req.Comment,
		PostID:  id,
		UserID:  req.UserID,
		Status:  req.Status,
	}
}
