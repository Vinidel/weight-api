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
	
	log.Println("This is the port", port)
	if err != nil {
		log.Fatal("Database connection failed: %s", err.Error())
	}

	app := &app.App{
		Router:   mux.NewRouter().StrictSlash(true),
		DBClient: database,
	}

	app.SetupRouter()

	// log.Fatal(http.ListenAndServe(":8080", app.Router))
	log.Fatal(http.ListenAndServe(":" + port, handlers.CORS()(app.Router)))
}
