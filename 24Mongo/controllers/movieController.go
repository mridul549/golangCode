package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	model "mongo/models"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://mridul0917:5aCPGaRdc4h2sEJz@cluster0.i9bhs8y.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

// most imp
var collection *mongo.Collection

// connecting to mongo
func init() {
	// client options
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongo
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")
	collection = client.Database(dbName).Collection(colName)

	// collection instance
	fmt.Println("Collection is ready")

	fmt.Println("Connected to MongoDB")
}

func insertOneMovie (movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in db with Id: ", inserted.InsertedID)
}

func updateOneMovie (movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Modified count: ", result.ModifiedCount)
}

func deleteOneMovie (movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully deleted ", result.DeletedCount)
}

func deleteAllMovies () {
	result, err := collection.DeleteMany(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully deleted ", result.DeletedCount)
}


func getAllMovies () []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M 
		err := cur.Decode(&movie)
		
		if err != nil {
			log.Fatal(err)
		}
		
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
	return movies
}

func GetAllMovies (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	movies :=getAllMovies()
	json.NewEncoder(w).Encode(movies)
}

func CreateMovie (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}
 
func MarkAsWatched (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PATCH")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovie (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PATCH")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PATCH")

	deleteAllMovies()
	json.NewEncoder(w).Encode("Deleted All movies")
}