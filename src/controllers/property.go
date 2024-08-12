package controllers

import "github.com/gin-gonic/gin"

// import (
// 	"Gin/src/config"
// 	"Gin/src/models"
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

func AddProperty(c *gin.Context) {

}

// 	PropertyName := c.PostForm("propertyName")
// 	PropertyType := c.PostForm("propertyType")
// 	PriceStr := c.PostForm("price")
// 	Description := c.PostForm("description")
// 	Location := c.PostForm("location")
// 	RoomsStr := c.PostForm("rooms")

// 	Price, err := strconv.ParseFloat(PriceStr, 64)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "error comes to convert Price string to float"})
// 		return
// 	}

// 	Rooms, err := strconv.ParseInt(RoomsStr, 10, 64)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "error comes to convert Rooms string to int"})
// 		return
// 	}

// 	form, err := c.MultipartForm()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	files := form.File["propertyImages"]
// 	ImagePaths := []string{}
// 	for _, file := range files {
// 		currentTime := time.Now().Format("20060102150405")
// 		fileName := fmt.Sprintf("%s_%s", currentTime, filepath.Base(file.Filename))
// 		savePath := fmt.Sprintf("uploads/property_image/%s", fileName)
// 		err := c.SaveUploadedFile(file, savePath)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file"})
// 			return
// 		}

// 		ImagePaths = append(ImagePaths, savePath)
// 	}

// 	property := models.Property{
// 		PropertyName:   PropertyName,
// 		PropertyType:   PropertyType,
// 		PropertyImages: ImagePaths,
// 		Price:          Price,
// 		Location:       Location,
// 		Rooms:          Rooms,
// 		Description:    Description,
// 	}

// 	if err := config.DB.Create(&property); err.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save property", "msg": err.Error.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusCreated, "message": "Property added successfully", "data": property})
// }

// func GetProperties(c *gin.Context) {
// 	var Property []models.Property
// 	res := config.DB.Find(&Property)
// 	if res.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "All propety data", "data": Property})
// }

// func SearchByNameAndPrice(c *gin.Context) {
// 	var Property []models.Property

// 	Location := c.Query("location")
// 	PriceStr := c.Query("price")

// 	query := config.DB.Table("property")

// 	if Location != "" {
// 		query = query.Where("LOWER(location) LIKE ?", "%"+strings.ToLower(Location)+"%")
// 	}

// 	if PriceStr != "" {
// 		price, err := strconv.ParseFloat(PriceStr, 64)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid price format"})
// 			return
// 		}
// 		query = query.Where("price <= ?", price)
// 	}

// 	if err := query.Find(&Property).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving properties"})
// 		return
// 	}

// 	if len(Property) == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Property not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Property by name and price", "data": Property})
// }

// func DeleteProperty(c *gin.Context) {
// 	var Property models.Property
// 	id := c.Param("id")

// 	if err := config.DB.First(&Property, id).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Property not found"})
// 		} else {
// 			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": err.Error()})
// 		}
// 		return
// 	}

// 	for _, imgName := range Property.PropertyImages {
// 		imgPath := filepath.Join(os.Getenv("DELETE_FILE_PATH"), imgName)
// 		if err := os.Remove(imgPath); err != nil {
// 			// c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusBadRequest, "error": "Error deleting image", "image": imgName, "msg": err.Error()})
// 			// return
// 			break
// 		}
// 	}

// 	if err := config.DB.Delete(&Property).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": "Error deleting property", "message": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Property deleted successfully"})

// }

// func UpdateProperty(c *gin.Context) {
// 	var property models.Property
// 	id := c.Param("id")

// 	if err := config.DB.Find(&property, id).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "error": "property not found"})
// 		} else {
// 			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": err.Error(), "msg": "internal server err"})
// 		}
// 		return
// 	}

// 	PropertyName := c.PostForm("propertyName")
// 	PropertyType := c.PostForm("propertyType")
// 	PriceStr := c.PostForm("price")
// 	Description := c.PostForm("description")
// 	Location := c.PostForm("location")
// 	RoomsStr := c.PostForm("rooms")

// 	Price, err := strconv.ParseFloat(PriceStr, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	Rooms, err := strconv.ParseInt(RoomsStr, 10, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := c.Request.ParseMultipartForm(32 << 20); err != nil { // Limit size to 32 MB
// 		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": "Unable to parse form"})
// 		return
// 	}

// 	for _, oldImg := range property.PropertyImages {
// 		fullImgPath := filepath.Join(os.Getenv("DELETE_FILE_PATH"), oldImg)
// 		if err := os.Remove(fullImgPath); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": "Error deleting old image", "image": oldImg, "msg": err.Error()})
// 			return
// 		}
// 	}

// 	form := c.Request.MultipartForm
// 	files := form.File["propertyImages"]
// 	var newImagePaths []string
// 	for _, file := range files {
// 		currentTime := time.Now().Format("20060102150405")
// 		fileName := fmt.Sprintf("%s_%s", currentTime, filepath.Base(file.Filename))
// 		savePath := fmt.Sprintf("uploads/property_image/%s", fileName)
// 		if err := c.SaveUploadedFile(file, savePath); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file", "msg": err.Error()})
// 			return
// 		}
// 		newImagePaths = append(newImagePaths, fmt.Sprintf("uploads/property_image/%s", fileName))
// 	}

// 	property.PropertyImages = newImagePaths
// 	property.PropertyName = PropertyName
// 	property.PropertyType = PropertyType
// 	property.Price = Price
// 	property.Location = Location
// 	property.Rooms = Rooms
// 	property.Description = Description

// 	if err := config.DB.Save(&property).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusBadRequest, "error": "Property cannot update"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Property updated successfully", "data": property})
// }

// func GetPropertyById(c *gin.Context) {
// 	var property models.Property
// 	id := c.Param("id")

// 	if err := config.DB.Find(&property, id).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "error": "property not found"})
// 		} else {
// 			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "error": "Internal Server Error"})
// 		}
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "All property data", "data": property})
// }
