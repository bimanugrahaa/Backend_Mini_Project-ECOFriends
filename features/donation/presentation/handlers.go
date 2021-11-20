package c_donation

import (
	"Backend_Mini_Project-ECOFriends/features/donation"
	presentation_response "Backend_Mini_Project-ECOFriends/features/donation/presentation/response"

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
	result := dh.donationBussiness.GetAllData("")
	return c.JSON(http.StatusOK, map[string]interface{}{
		// "claims":  middleware.ExtractClaim(c),
		"message": "Success",
		"data":    presentation_response.FromCoreSlice(result),
	})

}
