package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SyedAanif/rss-feed-aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // just need unused library DB connector driver
)

func main(){
	fmt.Println("*** Welcome to RSS(Rich Site Survey) Feed Aggregator! ***")
	
	// go get github.com/joho/godotenv --> get env variables
	// go mod vendor --> local copy
	// else OS works on exported ENV

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == ""{
		log.Fatal("PORT is not found in the environment")
	}

	fmt.Println("Port:",portString)

	dbURL := os.Getenv("DB_URL")
	if dbURL == ""{
		log.Fatal("DB_URL is not found in the environment")
	}

	fmt.Println("DB_URL:",dbURL)

	// Connect to DB
	conn, e := sql.Open("postgres",dbURL)
	if e != nil {
		log.Fatal("Could not connect to DB:",e)
	}

	// Generate access to DB
	apiCfg := apiConfig{
		DB: database.New(conn), // Convert sql_DB queries to DB_queries
	}

	// CHI router is light-weight standard GO router/web-server
	router := chi.NewRouter()

	// Allow CORS for access via browser
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Hook-up a path pattern to a request handler
	v1Router := chi.NewRouter()
	// v1Router.HandleFunc("/healthz",handlerReadiness) // Handles all HTTP verbs
	v1Router.Get("/healthz", handlerReadiness) // Only HTTP GET verb
	v1Router.Get("/err", handleError)
	v1Router.Post("/users", apiCfg.handlerCreateUser) // using pointer to gain access to HTTP handler

	// Mount V1 router under sub-path of V1 on main chi-router
	router.Mount("/v1",v1Router)

	// Create a server over the router and port using pointer
	server := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	log.Printf("Sever starting on port %v", portString)

	// Handles HTTP requests, thus blocking
	err := server.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}
}

// Hold connection to a DB
type apiConfig struct {
	DB *database.Queries
}