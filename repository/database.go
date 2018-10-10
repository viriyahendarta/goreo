package repository

//Database is a contract to database implementation
type Database interface {
	Begin()
}
