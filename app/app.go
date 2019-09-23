package app

import (
	"encoding/json"
	"net/http"
	"github.com/vinidel/weight-api/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"	
	"context"
	"log"
)

type App struct {
	Router *mux.Router
	DBClient *mongo.Client
}

type Pong struct {
	Message string `json:"message"`
}

var weights []models.Weight
var history models.History

func (app *App) SetupRouter() {
	app.Router.Methods("GET").Path("/api/ping").HandlerFunc(app.pong)
	app.Router.Methods("GET").Path("/api/weights").HandlerFunc(app.getWeightHistory)
	app.Router.Methods("POST").Path("/api/weights").HandlerFunc(app.postWeight)
}

// Testing connection
func (app *App) pong(w http.ResponseWriter, r *http.Request) {
	message := &Pong{Message: "pong"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func (app *App) getWeightHistory(w http.ResponseWriter, r *http.Request) {
	collection := app.DBClient.Database("weight-api").Collection("weights")
	
	// Pass these options to the Find method
	findOptions := options.Find()
	var results []*models.Weight

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
			log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
			
			// create a value into which the single document can be decoded
			var elem models.Weight
			err := cur.Decode(&elem)
			if err != nil {
					log.Fatal(err)
			}

			results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
			log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	log.Printf("Found multiple documents (array of pointers): %+v\n", results)

	history := results
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}

func (app *App) postWeight(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var weight models.Weight
	_ = json.NewDecoder(r.Body).Decode(&weight)
	log.Println("This is the createdAt ", weight.CreatedAt)
	collection := app.DBClient.Database("weight-api").Collection("weights")

	result, err := collection.InsertOne(context.TODO(), weight)

	if err != nil {
		log.Println("There was an error inserting weight {}", err.Error())
	} 

	log.Println("Inserted multiple documents: ", result.InsertedID)

	json.NewEncoder(w).Encode(weight)
}