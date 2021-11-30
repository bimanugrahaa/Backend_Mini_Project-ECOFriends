package c_donation

import (
	"Backend_Mini_Project-ECOFriends/features/donation"
	presentation_request "Backend_Mini_Project-ECOFriends/features/donation/presentation/request"
	presentation_response "Backend_Mini_Project-ECOFriends/features/donation/presentation/response"
	"Backend_Mini_Project-ECOFriends/middleware"
	"fmt"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

type DonationHandler struct {
	donationBussiness donation.Bussiness
}

func NewDonationHandler(dbu donation.Bussiness) *DonationHandler {
	return &DonationHandler{
		donationBussiness: dbu,
	}
}

func (dh *DonationHandler) GetAllDonation(c echo.Context) error {
	result := dh.donationBussiness.GetAllDonations()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    presentation_response.FromCoreSlice(result),
	})
}

func (dh *DonationHandler) GetDonationsById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	result := dh.donationBussiness.GetDonationsById(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    presentation_response.FromCoreDetail(result),
	})
}

func (dh *DonationHandler) CreateDonation(c echo.Context) error {

	claim := middleware.ExtractTokenUserId(c)
	user_id := int(claim["user_id"].(float64))
	newDonation := presentation_request.Donation{}

	newDonation.AuthorID = user_id
	newDonation.Author.ID = user_id

	c.Bind(&newDonation)

	result, err := dh.donationBussiness.CreateDonation(presentation_request.ToCore(newDonation))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	fmt.Println("present resp", presentation_response.FromCoreDetail(result))
	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "success",
		"data":    presentation_response.FromCoreDetail(result),
	})
}

func (dh *DonationHandler) DeleteDonationsById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	claim := middleware.ExtractTokenUserId(c)
	user_id := int(claim["user_id"].(float64))
	fmt.Println(donation.Core{ID: id})
	err := dh.donationBussiness.DeleteDonationsById(user_id, donation.Core{ID: id})

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": http.StatusUnauthorized,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "delete donation by id success",
	})
}

func (dh *DonationHandler) UpdateDonation(c echo.Context) error {
	UpdateDonation := presentation_request.Donation{}

	c.Bind(&UpdateDonation)

	result, err := dh.donationBussiness.UpdateDonation(presentation_request.ToCore(UpdateDonation))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "success",
		"data":    presentation_response.FromCoreDetail(result),
	})
}

func (dh *DonationHandler) CreateComment(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	newComment := presentation_request.CommentDonation{}

	c.Bind(&newComment)

	result, err := dh.donationBussiness.CreateComment(id, presentation_request.ToCommentCore(id, newComment))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "success",
		"data":    presentation_response.FromCommentCore(result),
	})
}

func (dh *DonationHandler) UpdateComment(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	updateComment := presentation_request.CommentDonation{}

	c.Bind(&updateComment)

	result, err := dh.donationBussiness.UpdateComment(presentation_request.ToCommentCore(id, updateComment))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "success",
		"data":    presentation_response.FromCommentUpdateCore(result),
	})
}

func (dh *DonationHandler) DeleteComment(c echo.Context) error {
	comment := presentation_request.CommentDonation{}

	c.Bind(&comment)
	fmt.Println(comment)
	err := dh.donationBussiness.DeleteComment(comment.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "delete comment by id success",
	})
}
