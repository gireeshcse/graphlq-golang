package main

import (
	"log"
	"net/http"
	// "os"
	"strconv"

	c "github.com/gireeshcse/graphlq-golang/config"
	"github.com/spf13/viper"

	database "github.com/gireeshcse/graphlq-golang/internal/pkg/db/mysql"
	"github.com/gireeshcse/graphlq-golang/internal/auth"
	"github.com/gireeshcse/graphlq-golang/internal/pkg/jwt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gireeshcse/graphlq-golang/graph"
	"github.com/gireeshcse/graphlq-golang/graph/generated"

	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {

	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration c.Configurations

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}
	
	// Set undefined variables
	viper.SetDefault("database.dbname", "hackernews")
	viper.SetDefault("server.port", defaultPort)

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Printf("Unable to decode into struct, %v", err)
	}

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = defaultPort
	// }

	database.InitDB(configuration.Database.DBHost,configuration.Database.DBPort,configuration.Database.DBUser,configuration.Database.DBPassword,configuration.Database.DBName)
	database.Migrate()
	jwt.InitJWT(configuration.JWT.JWTSecret)

	router := chi.NewRouter()
	router.Use(auth.Middleware())
	

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", strconv.Itoa(configuration.Server.Port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(configuration.Server.Port), router))
}
