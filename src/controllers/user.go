package controllers

import (
	"JourneyJoyBackend/src/common"
	"JourneyJoyBackend/src/config"
	"JourneyJoyBackend/src/models"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context) {
	// For FormData
	name := c.PostForm("Name")
	email := c.PostForm("Email")
	password := c.PostForm("Password")
	phoneNo := c.PostForm("PhoneNo")
	if name == "" || email == "" || password == "" {
		common.ErrorJsonResponse(c, http.StatusBadRequest, common.USER_ERR_MSG)
		return
	}

	var existsingUser models.User
	if common.FindJsonResponse(c, "email", email, &existsingUser, http.StatusBadRequest, common.EMAIL_ERR_MSG) {
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		common.ErrorJsonResponse(c, http.StatusInternalServerError, common.FAILED_HASH_MSG)
	}

	var profilePhotoPath string
	file, err := c.FormFile("ProfilePhoto")
	if err == nil {
		// Take Image from user
		ext := filepath.Ext(file.Filename)
		if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
			common.ErrorJsonResponse(c, http.StatusBadRequest, common.FILE_TYPE_MSG)
			return
		}

		// Saved Image on our server
		currentTime := time.Now().Format("20060102150405")
		fileName := fmt.Sprintf("%s_%s%s", currentTime, file.Filename, ext)
		filePath := "uploads/profile_photos/" + fileName
		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			common.ErrorJsonResponse(c, http.StatusBadRequest, common.FAILED_SAVED_FILE_MSG)
			return
		}

		profilePhotoPath = filePath
	}

	user := models.User{
		Name:         name,
		Email:        email,
		Password:     string(hashedPass),
		PhoneNo:      phoneNo,
		ProfilePhoto: profilePhotoPath,
	}

	// Create User
	result := config.DB.Create(&user)
	if result.Error != nil {
		common.ErrorJsonResponse(c, http.StatusInternalServerError, common.CREATE_USER_ERR_MSG)
	}

	common.JsonResponse(c, http.StatusCreated, common.USER_CREATE_SUCCESS_MSG, user)
}

func GetUsers(c *gin.Context) {
	var users []models.User
	allUsers := config.DB.Find(&users)
	if allUsers.Error != nil {
		common.ErrorJsonResponse(c, http.StatusInternalServerError, common.FAILED_FETCH_USER)
		return
	}

	common.JsonResponse(c, http.StatusOK, common.ALL_USER_SUCCESS_MSG, users)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := config.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.ErrorJsonResponse(c, http.StatusNotFound, common.USER_NOT_FOUND)
		} else {
			common.ErrorJsonResponse(c, http.StatusInternalServerError, common.USER_NOT_DELETE)
		}
		return
	}
	config.DB.Delete(&user, id)
	common.JsonResponse(c, http.StatusOK, common.USER_DEL_SUCCESS_MSG, nil)

}

func UpdateUser(c *gin.Context) {
	var req models.User
	id := c.Param("id")
	if res := c.BindJSON(&req); res != nil {
		common.ErrorJsonResponse(c, http.StatusBadRequest, res.Error())
		return
	}

	var user models.User
	if res := config.DB.Find(&user, id).Error; res != nil {
		common.ErrorJsonResponse(c, http.StatusNotFound, res.Error())
		return
	}
	user.Name = req.Name
	user.Email = req.Email
	user.PhoneNo = req.PhoneNo
	config.DB.Save(&user)
	common.JsonResponse(c, http.StatusOK, common.USER_UPD_SUCCESS_MSG, user)
}

func Login(c *gin.Context) {
	var req common.LoginStruct
	if err := c.BindJSON(&req); err != nil {
		common.ErrorJsonResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.ErrorJsonResponse(c, http.StatusNotFound, "User Not Found")
		}
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		common.ErrorJsonResponse(c, http.StatusBadRequest, common.INCORRECT_PASS_ERR_MSG)
		return
	}

	common.JsonResponse(c, http.StatusOK, common.LOGIN_SUCCESS_MSG, user)
}

func GetUserById(c *gin.Context) {
	// var user models.User
	// id := c.Param("id")
	// if common.FindJsonResponse(c, "id", id, &user, http.StatusBadRequest, common.USER_NOT_FOUND) {
	// 	return
	// }
	// common.JsonResponse(c, http.StatusOK, "Single User", user)
}
