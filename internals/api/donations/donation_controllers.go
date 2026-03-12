package donations

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"

	"github.com/P47H4N/shafar_foundation_api/internals/helpers"
	"github.com/gin-gonic/gin"
)

func InitDonationController(ds *DonationServices) *DonationControllers {
	return &DonationControllers{
		service: ds,
	}
}

func (dc *DonationControllers) GetDonations(c *gin.Context) {
	donations, err := dc.service.GetDonations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.APIResponse{
			Status: "failed",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, helpers.APIResponse{
		Status: "success",
		Message: fmt.Sprintf("%d donations found.", len(donations)),
		Data: donations,
	})
}

func (dc *DonationControllers) CreateDonation(c *gin.Context) {
	var donationBody DonationBody
	if err := c.ShouldBindBodyWithJSON(&donationBody); err != nil {
		c.JSON(http.StatusBadRequest, helpers.APIResponse{
			Status: "failed",
			Message: "Invalid Data.",
		})
		return
	}
	if err := dc.service.CreateDonation(&donationBody); err != nil {
		c.JSON(http.StatusInternalServerError, helpers.APIResponse{
			Status: "failed",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, helpers.APIResponse{
		Status: "success",
		Message: "Donation Created.",
		Data: donationBody,
	})
}

func (dc *DonationControllers) UpdateDonation(c *gin.Context) {
	status := c.Param("status")
	validStatus := []string{"approve", "reject"}
	if !slices.Contains(validStatus, status) {
		c.JSON(http.StatusBadRequest, helpers.APIResponse{
			Status: "failed",
			Message: "Invalid Status.",
		})
		return
	}
	paramId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.APIResponse{
			Status: "failed",
			Message: "Invalid donation id.",
		})
		return
	}
	if err := dc.service.UpdateDonation(status, uint(paramId)); err != nil {
		c.JSON(http.StatusInternalServerError, helpers.APIResponse{
			Status: "failed",
			Message: err.Error(),
		})
		return
	}
	message := "Donation Rejected."
	if status == "approve" {
		message = "Donation Approved."
	}
	c.JSON(http.StatusOK, helpers.APIResponse{
		Status: "success",
		Message: message,
	})
}

func (dc *DonationControllers) DeleteDonation(c *gin.Context) {
	paramId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.APIResponse{
			Status: "failed",
			Message: "Invalid donation id.",
		})
		return
	}
	if err := dc.service.DeleteDonation(uint(paramId)); err != nil {
		c.JSON(http.StatusInternalServerError, helpers.APIResponse{
			Status: "failed",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, helpers.APIResponse{
		Status: "success",
		Message: "Donation Deleted Successfully.",
	})
}
