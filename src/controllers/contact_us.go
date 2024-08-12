package controllers

import (
	"JourneyJoyBackend/src/common"
	"JourneyJoyBackend/src/config"
	"JourneyJoyBackend/src/email"
	"JourneyJoyBackend/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddContactUs(c *gin.Context) {
	var contactUs models.ContactUs
	if err := c.BindJSON(&contactUs); err != nil {
		common.ErrorJsonResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := config.DB.Save(&contactUs).Error; err != nil {
		common.ErrorJsonResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := email.SendContactUsEmail(contactUs); err != nil {
		common.ErrorJsonResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	common.JsonResponse(c, http.StatusOK, common.CONTACT_US_SUCCESS_MSG, contactUs)
}
