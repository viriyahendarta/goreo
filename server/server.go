package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/viriyahendarta/goreo/resource/service"
)

//Server holds contract for server implementation
type Server interface {
	Run() error
}

//httpServer holds data needed to run http server
type httpServer struct {
	serviceResource *service.Resource
	router          *mux.Router
	port            int
}

//httpHandlerFunc defines function for handle http request
type httpHandlerFunc func(r *http.Request) (interface{}, error)
