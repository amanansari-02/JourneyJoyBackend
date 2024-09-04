package routes

import (
	"JourneyJoyBackend/src/common"
	"JourneyJoyBackend/src/controllers"
	"JourneyJoyBackend/src/middleware"

	"github.com/gin-gonic/gin"
)

func PropertyRoutes(r *gin.Engine) {
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	protected.POST(common.PROPERTY, controllers.AddProperty)
	protected.GET(common.PROPERTY, controllers.GetProperties)
	protected.GET(common.PROPERTY_BY_ID, controllers.GetPropertyById)
	protected.GET(common.LATEST_PROPERTY_URL, controllers.GetLatestProperty)
	protected.GET(common.SEARCH_BY_NAME_TYPE, controllers.SearchByNameAndPrice)
	protected.PUT(common.PROPERTY_BY_ID, controllers.UpdateProperty)
	protected.DELETE(common.PROPERTY_BY_ID, controllers.DeleteProperty)
}
