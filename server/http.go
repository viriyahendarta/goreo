package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/viriyahendarta/goreo/config"
	cont "github.com/viriyahendarta/goreo/infra/context"
	serviceresource "github.com/viriyahendarta/goreo/resource/service"
	"github.com/viriyahendarta/goreo/service/api/user"
)

//InitHTTPServer initializes http server and return http server implementation
func InitHTTPServer(serviceResource *serviceresource.Resource, port int) Server {
	return &httpServer{
		serviceResource: serviceResource,
		router:          serviceResource.Router,
		port:            port,
	}
}

//Run register request handlers and start the server
func (s *httpServer) Run() error {
	s.registerAPI()

	address := fmt.Sprint("0.0.0.0:", s.port)
	log.Printf("Starting [%s] server at %s\n", config.GetEnv(), address)
	return gracehttp.Serve(&http.Server{
		Addr:         address,
		Handler:      s.router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	})
}

//registerAPI register handlers for apis
func (s *httpServer) registerAPI() {
	userAPI := user.GetAPI(s.serviceResource)

	s.router.HandleFunc("/user/{user_id}/profile", s.handle(userAPI.GetUserProfile)).Methods(http.MethodGet)
}

//handle is a helper function to wrap http handleFunc
func (s *httpServer) handle(handler httpHandlerFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), cont.StartTime, time.Now())
		response, httpCode, err := handler(r)
		s.serviceResource.RenderJSON(ctx, w, response, httpCode, err)
	}
}
