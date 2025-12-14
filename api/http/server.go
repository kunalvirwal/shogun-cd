package api

import "github.com/gorilla/mux"

func StartAPIServer() {
	r := mux.NewRouter()

	initRoutes(r)

}
