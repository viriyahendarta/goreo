package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/viriyahendarta/goreo/config"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	userrepo "github.com/viriyahendarta/goreo/repository/user"
	businessresource "github.com/viriyahendarta/goreo/resource/business"
	repositoryresource "github.com/viriyahendarta/goreo/resource/repository"
	serviceresource "github.com/viriyahendarta/goreo/resource/service"
	"github.com/viriyahendarta/goreo/server"
)

//init initialize the application
func init() {
	config.Init()
}

func main() {
	// connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	dbConfig := config.Get().CoreDatabase
	dbConnection, err := sql.Open(dbConfig.Driver, dbConfig.URL)
	if err != nil {
		log.Fatal(err)
	}
	if err := dbConnection.Ping(); err != nil {
		log.Fatalln("Failed to ping core database")
	}

	repoResource := initRepositoryResource(dbConnection)
	businessResource := initBusinessResource(repoResource)
	serviceResource := initServiceResource(businessResource)

	arg := "http"
	if len(os.Args) >= 2 {
		arg = os.Args[1]
	}

	switch arg {
	default:
		startHTTPServer(serviceResource)
	}
}

func startHTTPServer(serviceResource *serviceresource.Resource) error {
	httpServer := server.InitHTTPServer(serviceResource, config.Get().HTTPServer.Port)
	return httpServer.Run()
}

func initRepositoryResource(dbConnection *sql.DB) *repositoryresource.Resource {
	return &repositoryresource.Resource{
		UserDB: dbConnection,
	}
}

func initBusinessResource(repoResource *repositoryresource.Resource) *businessresource.Resource {
	return &businessresource.Resource{
		UserDatabase: userrepo.GetDatabase(repoResource.UserDB),
	}
}

func initServiceResource(businessResource *businessresource.Resource) *serviceresource.Resource {
	return &serviceresource.Resource{
		BusinessResource: businessResource,
		Router:           mux.NewRouter(),
	}
}
