package main

import (
	"log"
	"os"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/vinidel/weight-api/db"
	"github.com/vinidel/weight-api/app"
	"github.com/gorilla/handlers"
)

func main() {
	database, err := db.SetupDB()
	port := os.Getenv("PORT")
	router :=  mux.NewRouter().StrictSlash(true)
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
	)

	router.Use(cors)

	log.Println("This is the port", port)
	if err != nil {
		log.Fatal("Database connection failed: %s", err.Error())
	}

	app := &app.App{
		Router:  router,
		DBClient: database,
	}

	app.SetupRouter()

	// log.Fatal(http.ListenAndServe(":8080", app.Router))
	log.Fatal(http.ListenAndServe(":" + port, router))
}
