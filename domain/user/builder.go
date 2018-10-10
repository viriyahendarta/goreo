package user

import "github.com/viriyahendarta/goreo/model"

//BuildProfileFromUserModel build user profile from user model
func BuildProfileFromUserModel(mUser *model.User) Profile {
	return Profile{
		Email: mUser.Email,
		Name:  mUser.Name,
	}
}
