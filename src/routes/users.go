package routes

import (
	"Gin/src/common"
	"Gin/src/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.POST(common.USER, controllers.CreateUser)
	r.GET(common.USERS, controllers.GetUsers)
	r.DELETE(common.USER_BY_ID, controllers.DeleteUser)
	r.PUT(common.USER_BY_ID, controllers.UpdateUser)
	r.POST(common.LOGIN, controllers.Login)
}
