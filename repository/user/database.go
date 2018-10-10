package user

import (
	"database/sql"
	"sync"

	"github.com/viriyahendarta/goreo/model"
	"github.com/viriyahendarta/goreo/repository"
)

//Database holds contracts from repository database and user database it self
type Database interface {
	repository.Database
	Find(id int64) *model.User
}

type userDatabase struct {
	Conn *sql.DB
}

var db Database
var once sync.Once

//GetDatabase receives database connection for user database
//and returns user database
func GetDatabase(conn *sql.DB) Database {
	once.Do(func() {
		db = &userDatabase{
			Conn: conn,
		}
	})

	return db
}

func (d *userDatabase) Begin() {
	// d.Conn.B
}

//Find finds user by user_id
func (d *userDatabase) Find(id int64) *model.User {
	return new(model.User)
}
