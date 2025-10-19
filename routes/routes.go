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

	// manipulate events (protected)
	authenticated := server.Group("/events")
	authenticated.Use(middlewares.Authenticate)
	authenticated.
		POST("/", createEvent).
		PUT("/:id", updateEvent).
		DELETE("/:id", deleteEvent)

	// user auth
	server.
		POST("/signup", signup).
		POST("/login", login)
}
