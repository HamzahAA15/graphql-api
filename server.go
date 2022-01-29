package main

import (
	"log"
	"net/http"
	"os"
	"sirclo/gql/config"
	"sirclo/gql/controller/graph"
	auth "sirclo/gql/repository/authmiddleware"
	"sirclo/gql/repository/book"
	"sirclo/gql/repository/user"
	"sirclo/gql/util"
	"sirclo/gql/util/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

const defaultPort = "8080"

// type Auth struct {
// 	auth auth.Middleware
// }

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := chi.NewRouter()

	//load config if available or set to defaults
	config := config.GetConfig()

	//initialize database connection based on given config
	db := util.MysqlDriver(config)

	// var midwr Auth
	router.Use(auth.Middleware())
	userRepo := user.NewRepositoryUser(db)
	bookRepo := book.NewRepositoryBook(db)
	client := graph.NewResolver(userRepo, bookRepo)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: client}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
