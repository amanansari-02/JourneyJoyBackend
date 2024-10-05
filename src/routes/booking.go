package routes

import (
	"JourneyJoyBackend/src/common"
	"JourneyJoyBackend/src/controllers"
	"JourneyJoyBackend/src/middleware"

	"github.com/gin-gonic/gin"
)

func BookingRoutes(r *gin.Engine) {
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	protected.GET(common.BOOKING_BY_PROPERTY_ID, controllers.GetBookingByPropertyId)
	protected.GET(common.BOOKING_BY_USER_ID, controllers.GetBookingByUserId)
	protected.GET(common.ALL_BOOKING, controllers.GetAllBookingData)
	protected.POST(common.BOOKING, controllers.AddBooking)
	protected.GET(common.DASHBOARD_API, controllers.Dashboard)
}
