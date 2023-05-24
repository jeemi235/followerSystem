package main

import (
	"fmt"

	routes "media/api/Routes"
	"media/middlewares"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

func main() {
	//Using chi we have created new router
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	//For assigning value of user id
	r.Use(middlewares.RequestId)
	//For assigning value of database to the context
	r.Use(middlewares.DbContext)

	//This function will call our Routes function
	routes.Routes(r)

	fmt.Println("Server started at :6000")
	//Started the sever on localhost port:6000
	http.ListenAndServe(":6000", r)

}
