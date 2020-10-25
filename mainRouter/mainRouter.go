package mainrouter

import (
	"hamid/example3/users"

	"github.com/julienschmidt/httprouter"
)

// Routers defines all routes of the app
func Routers(router *httprouter.Router) {
	users.Routers(router)
}
