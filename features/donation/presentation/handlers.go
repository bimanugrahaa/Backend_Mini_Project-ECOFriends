package c_donation

import (
	"Backend_Mini_Project-ECOFriends/features/donation"
	"Backend_Mini_Project-ECOFriends/features/donation/presentation/request"
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
	// var ids int
	// idx := c.Param(strconv.Itoa(ids))
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println(id)
	result := dh.donationBussiness.GetDonationsById(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    presentation_response.FromCoreDetail(result),
	})
}

func (dh *DonationHandler) CreateDonation(c echo.Context) error {
	// var newDonation presentation_request.Donation

	newDonation := request.Donation{}

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

// func (dh *DonationHandler) CreateDescriptionDonation(c echo.Context) error {
// 	newDonationDescription := request.DonationDescription{}

// 	c.Bind(newDonationDescription)

// 	resp, err := dh.donationBussiness.CreateDescriptionDonation(presentation_request.ToDescriptionCore(newDonationDescription))
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err)
// 	}

// 	return c.JSON(http.StatusAccepted, map[string]interface{}{
// 		"message": "Donation successfully added.",
// 		"data":    presentation_response.FromDescriptionDonationCore(resp),
// 	})

// }
