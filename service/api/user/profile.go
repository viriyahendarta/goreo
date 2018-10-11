package user

import (
	"net/http"
)

//GetUserProfile handles get user profile request
func (a *iUserAPI) GetUserProfile(r *http.Request) (interface{}, int, error) {
	result, err := a.userBusiness.GetUserProfile(r.Context(), 1)
	if err != nil {
		return nil, http.StatusGatewayTimeout, err
	}

	return result, 200, nil
}
