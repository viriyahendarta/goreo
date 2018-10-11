package user

import (
	"context"
	"sync"

	"github.com/viriyahendarta/goreo/business"
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
