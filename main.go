package main

import (
	mainrouter "hamid/example3/mainRouter"
	"hamid/example3/migrate"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()
	mainrouter.Routers(router)

	err := godotenv.Load()
	if err != nil {
		panic("can not get env variables")
	}
	err = migrate.MakeMigrations()
	if err != nil {
		panic("can not migrate database")
	}

	log.Fatal(http.ListenAndServe(":11111", router))

}
