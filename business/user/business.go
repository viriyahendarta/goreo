package user

import (
	"context"
	"errors"
	"sync"

	"github.com/viriyahendarta/goreo/business"
	e "github.com/viriyahendarta/goreo/infra/error"
	userrepo "github.com/viriyahendarta/goreo/repository/user"
)

//Business holds contract of user business
type Business interface {
	business.Business
	GetUserProfile(ctx context.Context, userID int64) (interface{}, error)
}

type userBusiness struct {
	UserDB userrepo.Database
}

var bUser Business
var once sync.Once

//GetUserBusiness returns user business implementation
func GetUserBusiness(userDB userrepo.Database) Business {
	once.Do(func() {
		bUser = &userBusiness{
			UserDB: userDB,
		}
	})
	return bUser
}

//GetUserProfile find user profile by userID
func (b *userBusiness) GetUserProfile(ctx context.Context, userID int64) (interface{}, error) {
	return nil, e.New(ctx, e.ErrorInvalidAuthentication, "Error to query get user profile", errors.New("wkwkkw"))
	// mUser := b.UserDB.Find(userID)
	// return userdomain.BuildUserProfile(mUser), nil
}
