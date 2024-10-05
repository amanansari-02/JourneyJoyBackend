package controllers

import (
	"JourneyJoyBackend/src/common"
	"JourneyJoyBackend/src/config"
	"JourneyJoyBackend/src/email"
	"JourneyJoyBackend/src/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddBooking(c *gin.Context) {
	var booking models.Booking

	if err := c.BindJSON(&booking); err != nil {
		common.ErrorJsonResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	startDate, err := time.Parse(time.RFC3339, booking.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid StartDate format"})
		return
	}

	endDate, err := time.Parse(time.RFC3339, booking.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid EndDate format"})
		return
	}

	booking.StartDate = startDate.Format("2006-01-02")
	booking.EndDate = endDate.Format("2006-01-02")

	if err := config.DB.Create(&booking).Error; err != nil {
		common.ErrorJsonResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := config.DB.Preload("User").Preload("Property").Find(&booking, booking.Id).Error; err != nil {
		common.ErrorJsonResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := email.SendBookingConfirmationMessage(booking); err != nil {
		common.ErrorJsonResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	common.JsonResponse(c, http.StatusOK, common.CREATE_BOOKING, booking)
}

func GetBooking(c *gin.Context) {
	var bookings []models.Booking

	if err := config.DB.Preload("User").Preload("Property").Find(&bookings).Error; err != nil {
		common.ErrorJsonResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	common.JsonResponse(c, http.StatusOK, common.GET_ALL_BOOKING, bookings)
}

func FormatDate(date time.Time) string {
	return date.Format("2006-01-02")
}

func GetBookingByPropertyId(c *gin.Context) {
	var booking []models.Booking
	id := c.Param("propertyId")

	if err := config.DB.Where("property_id = ?", id).Find(&booking).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			common.ErrorJsonResponse(c, http.StatusNotFound, "No booking found for the specified property.")
		}

		common.ErrorJsonResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	common.JsonResponse(c, http.StatusOK, "All booking data", booking)
}

func GetBookingByUserId(c *gin.Context) {
	var bookings []models.Booking
	id := c.Param("userId")

	if err := config.DB.Where("user_id = ?", id).Preload("Property").Find(&bookings).Error; err != nil {
		common.ErrorJsonResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(bookings) == 0 {
		common.ErrorJsonResponse(c, http.StatusNotFound, "No booking found for the specified user.")
		return
	}

	common.JsonResponse(c, http.StatusOK, "All booking data by user id", bookings)
}

func GetAllBookingData(c *gin.Context) {
	var bookings []models.Booking

	if err := config.DB.Preload("Property").Find(&bookings).Error; err != nil {
		common.ErrorJsonResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	common.JsonResponse(c, http.StatusOK, "All booking data", bookings)
}

func Dashboard(c *gin.Context) {
	var bookings []models.Booking
	var property []models.Property
	var users []models.User
	var total_users int64
	var total_properties int64
	var total_booking int64

	// Count all users
	res := config.DB.Find(&users).Count(&total_users)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to count users",
		})
		return
	}

	// Count all properties
	pro_res := config.DB.Find(&property).Count(&total_properties)
	if pro_res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to count users",
		})
		return
	}

	// Count all booking
	booking_res := config.DB.Find(&bookings).Count(&total_booking)
	if booking_res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to count users",
		})
		return
	}

	data := gin.H{
		"status":           200,
		"message":          "Dashboard data",
		"total_users":      total_users,
		"total_properties": total_properties,
		"total_booking":    total_booking,
	}
	c.JSON(http.StatusOK, data)
}
