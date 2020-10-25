package courses

import (
	"github.com/julienschmidt/httprouter"
)

// Routers defines all routes for courses crud
func Routers(router *httprouter.Router) {
	router.GET("/courses", readAll)
	router.GET("/courses/:id", readOne)
	router.POST("/courses", create)
	router.PUT("/courses/:id", update)
	router.DELETE("/courses/:id", delete)
}
