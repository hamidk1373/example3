package users

import (
	"github.com/julienschmidt/httprouter"
)

// Routers defines all routes for users crud
func Routers(router *httprouter.Router) {
	router.GET("/users", readAll)
	router.GET("/users/:id", readOne)
	router.POST("/users", create)
	router.PUT("/users/:id", update)
	router.DELETE("/users/:id", delete)

	router.POST("/auth/register", register)
	router.POST("/auth/login", login)
}
