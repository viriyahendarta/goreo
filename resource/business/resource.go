package business

import userrepo "github.com/viriyahendarta/goreo/repository/user"

//Resource holds resources needed for business
type Resource struct {
	UserDatabase userrepo.Database
}
