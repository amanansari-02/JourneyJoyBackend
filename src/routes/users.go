package routes

import (
	"JourneyJoyBackend/src/common"
	"JourneyJoyBackend/src/controllers"
	"JourneyJoyBackend/src/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.POST(common.USER, controllers.CreateUser)
	r.DELETE(common.USER_BY_ID, controllers.DeleteUser)
	r.PUT(common.USER_BY_ID, controllers.UpdateUser)
	r.POST(common.LOGIN, controllers.Login)
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	protected.GET(common.USERS, controllers.GetUsers)
	r.GET(common.USER_BY_ID, controllers.GetUserById)

	// contactUs
	r.POST(common.CONTACT_US, controllers.AddContactUs)
}
