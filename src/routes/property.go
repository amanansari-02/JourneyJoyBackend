package routes

import (
	"JourneyJoyBackend/src/common"
	"JourneyJoyBackend/src/controllers"

	"github.com/gin-gonic/gin"
)

func PropertyRoutes(r *gin.Engine) {
	r.POST(common.PROPERTY, controllers.AddProperty)
}
