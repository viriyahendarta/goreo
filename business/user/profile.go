package user

import (
	"context"
	"errors"

	"github.com/viriyahendarta/goreo/infra/rror"
)

//GetUserProfile find user profile by userID
func (b *userBusiness) GetUserProfile(ctx context.Context, userID int64) (interface{}, error) {
	return nil, rror.New(rror.CodeGeneralQuery, "error to query get user profile", errors.New("wkwkkw"))
	// mUser := b.UserDB.Find(userID)
	// return userdomain.BuildUserProfile(mUser), nil
}
