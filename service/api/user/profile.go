package user

import (
	"net/http"
)

//GetUserProfile handles get user profile request
func (a *iUserAPI) GetUserProfile(r *http.Request) (interface{}, error) {
	return a.userBusiness.GetUserProfile(r.Context(), 1)
}
