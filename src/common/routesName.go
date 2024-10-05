package common

const (
	USER       = "/user"
	USERS      = "/users"
	USER_BY_ID = "/user/:id"
	LOGIN      = "/login"

	// Contact Us
	CONTACT_US = "/contact_us"

	// Property
	PROPERTY            = "/property"
	PROPERTY_BY_ID      = "/property/:id"
	LATEST_PROPERTY_URL = "/property/latest"
	SEARCH_BY_NAME_TYPE = "/property/SeacrhByNameAndType"

	// Booking
	BOOKING_BY_PROPERTY_ID = "/bookingByProperty/:propertyId"
	BOOKING_BY_USER_ID     = "/bookingByUser/:userId"
	BOOKING                = "/booking"
	ALL_BOOKING            = "/booking/all"
	DASHBOARD_API          = "/dashboard"
)
