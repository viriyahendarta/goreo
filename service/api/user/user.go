package user

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/viriyahendarta/goreo/business/user"
	"github.com/viriyahendarta/goreo/resource/service"
)

//API holds all user services
type API interface {
	GetUserProfile(r *http.Request) (interface{}, error)
}

//iUserAPI holds data needed for user api
type iUserAPI struct {
	userBusiness user.Business
	Router       *mux.Router
}

var userAPI API
var once sync.Once

//GetAPI initializes and return user api implementation
func GetAPI(serviceResource *service.Resource) API {
	once.Do(func() {
		userAPI = &iUserAPI{
			userBusiness: user.GetUserBusiness(serviceResource.BusinessResource.UserDatabase),
			Router:       serviceResource.Router,
		}
	})
	return userAPI
}
