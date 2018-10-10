package repository

import "database/sql"

//Resource holds objects needed for repositories
type Resource struct {
	UserDB *sql.DB
}
