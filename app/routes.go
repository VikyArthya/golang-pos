package app

import "github.com/VikyArthya/golang-pos/app/controllers"

func (server *Server) InitializeRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}