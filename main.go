package main

import (
	"jwt-test/routes"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// start http server
	router := httprouter.New()

	router.GET("/", routes.Index)

	router.POST("/user", routes.CreateUser)
	router.GET("/user/", routes.GetUsers)
	router.GET("/user/:id", routes.GetUser)
	router.PUT("/user/:id", routes.UpdateUser)
	router.DELETE("/user/:id", routes.DeleteUser)

	router.POST("/session", routes.Login)
	router.DELETE("/session", routes.Logout)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
