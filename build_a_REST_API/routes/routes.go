package routes

import (
	"fmt"
	"net/http"

	"example.com/web/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(server *gin.Engine) {

	server.GET("/:domain_nm/check", sayCheck)
	// server.GET("/:client_cd/:domain_nm/", sayCcDm)
	server.GET("/:domain_nm/", sayDomain)

	server.GET("/events", getEvents)
	server.GET("/event_:id", getEventById)
	server.GET("/users", getUsers)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/event", middlewares.Authenticate, createEvent)
	authenticated.PUT("/event_:id", updateEvent)
	authenticated.DELETE("/event_:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}

func sayDomain(context *gin.Context) {
	domain, _ := context.Params.Get("domain_nm")
	context.JSON(http.StatusOK, fmt.Sprintf("domain: %s", domain))
}

func sayCheck(context *gin.Context) {
	domain, _ := context.Params.Get("domain_nm")
	context.JSON(http.StatusOK, fmt.Sprintf("check: %s", domain))
}

// func sayCcDm(context *gin.Context) {
// 	domain, _ := context.Params.Get("domain_nm")
// 	cc, _ := context.Params.Get("client_cd")
// 	context.JSON(http.StatusOK, fmt.Sprintf("ccdm: %s %s", domain, cc))
// }
