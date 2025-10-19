package routes

import (
	"github.com/MichaelVenturi/go-practice-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// fetch events
	server.
		GET("/events", getEvents).
		GET("/events/:id", getEventById)

	authenticated := server.Group("/events")
	authenticated.Use(middlewares.Authenticate)
	// manipulate events (protected)
	authenticated.
		POST("/", createEvent).
		PUT("/:id", updateEvent).
		DELETE("/:id", deleteEvent)

	// registrations
	authenticated.
		POST("/:id/register", registerForEvent).
		DELETE("/:id/register", cancelRegistration)

	// user auth
	server.
		POST("/signup", signup).
		POST("/login", login)
}
