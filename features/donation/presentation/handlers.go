package c_donation

import (
	"Backend_Mini_Project-ECOFriends/features/donation"
	presentation_request "Backend_Mini_Project-ECOFriends/features/donation/presentation/request"
	presentation_response "Backend_Mini_Project-ECOFriends/features/donation/presentation/response"
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
		// "claims":  middleware.ExtractClaim(c),
		"message": "Success",
		"data":    presentation_response.FromCoreSlice(result),
	})
}

func (dh *DonationHandler) GetDonationsById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println(id)
	result := dh.donationBussiness.GetDonationsById(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    presentation_response.FromCoreDetail(result),
	})
}

func (dh *DonationHandler) CreateDonation(c echo.Context) error {

	newDonation := presentation_request.Donation{}

	c.Bind(&newDonation)

	result, err := dh.donationBussiness.CreateDonation(presentation_request.ToCore(newDonation))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "success",
		"data":    presentation_response.FromCoreDetail(result),
	})
}

func (dh *DonationHandler) DeleteDonationsById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println(id)
	err := dh.donationBussiness.DeleteDonationsById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
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
	UpdateComment := presentation_request.CommentDonation{}

	c.Bind(&UpdateComment)

	result, err := dh.donationBussiness.UpdateComment(presentation_request.ToCommentCore(id, UpdateComment))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "success",
		"data":    presentation_response.FromCommentCore(result),
	})
}
