package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Get All User's Data from DB and send it in respose
func GetAllUsers(c *gin.Context) {
	client, err := mongo.Connect( // Creating DB Connection
		context.TODO(),
		options.Client().ApplyURI("mongodb://localhost:27017"),
	)
	if err != nil { // Checking Whehter there is any error while connecting
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil) // Pinging DB to Confirm DB Connection
	if err != nil {                        // Checking Whether there is any error while pinging the DB
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	db := client.Database("admybrand") // Creating or Accessing DB
	coll := db.Collection("User")      // Creating or Accessing Collection
	cursor, err := coll.Find(          // Getting all Data from Collection
		context.Background(),
		bson.D{},
	)
	if err != nil { // Checking wether any error occured while getting data from collection
		log.Fatal(err)
	}
	var result []User

	err = cursor.All(context.TODO(), &result)
	if err != nil { // Getting all data into User Object and also checking for error
		log.Fatal(err)
	}

	err = client.Disconnect(context.TODO()) // Closing Connection of DB
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, result) // Sending Response to User
}

// Getting User from front end and Inserting it in DB
func CreateUser(c *gin.Context) {
	var newUser User                             // Creating User Struct Obj
	if err := c.BindJSON(&newUser); err != nil { // Checking for error and Getting Params
		log.Fatal(err)
		return
	}
	client, err := mongo.Connect( // Creating DB Connection
		context.TODO(),
		options.Client().ApplyURI("mongodb://localhost:27017"),
	)
	if err != nil { // Checking Whehter there is any error while connecting
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil) // Pinging DB to Confirm DB Connection
	if err != nil {                        // Checking Whether there is any error while pinging the DB
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	db := client.Database("admybrand") // Creating or Accessing DB
	coll := db.Collection("User")      // Creating or Accessing collection

	// Inserting a Document
	result, err := coll.InsertOne(context.Background(), bson.D{
		{
			Key:   "name",
			Value: newUser.Name,
		},
		{
			Key:   "Dob",
			Value: newUser.Dob,
		},
		{
			Key:   "Address",
			Value: newUser.Address,
		},
		{
			Key:   "Description",
			Value: newUser.Description,
		},
		{
			Key:   "createdAt",
			Value: time.Now(),
		},
	})

	if err != nil { // Checking Whether any error occured during insertion of data
		log.Fatal(err)
	}

	err = client.Disconnect(context.TODO()) // Closing Connection of DB
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusAccepted, result) // Sending Response to User
}

// Get a specific user using it's ID
func GetUserByID(c *gin.Context) {
	ID, _ := primitive.ObjectIDFromHex(c.Param("ID")) // Getting Param and converting it into MongoDB's ObjectID
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI("mongodb://localhost:27017"),
	) // Creating DB Connection
	if err != nil { // Checking Whehter there is any error while connecting
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil) // Pinging DB to Confirm DB Connection
	if err != nil {                        // Checking Whether there is any error while pinging the DB
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	db := client.Database("admybrand") // Creating or Accessing DB
	coll := db.Collection("User")      // Creating or Accessing Collection

	var result User // User Struct obj
	coll.FindOne(   // Finding record based on ID
		context.Background(),
		bson.D{{
			Key:   "_id",
			Value: ID,
		}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, result)
}

// Update existing user's data
func UpdateUser(c *gin.Context) {
	var newUser User                             // Creating User Struct Obj
	if err := c.BindJSON(&newUser); err != nil { // Checking for error and Getting Params
		log.Fatal(err)
		return
	}
	client, err := mongo.Connect( // Creating DB Connection
		context.TODO(),
		options.Client().ApplyURI("mongodb://localhost:27017"),
	)
	if err != nil { // Checking Whehter there is any error while connecting
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil) // Pinging DB to Confirm DB Connection
	if err != nil {                        // Checking Whether there is any error while pinging the DB
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	db := client.Database("admybrand") // Creating or Accessing DB
	coll := db.Collection("User")      // Creating or Accessing collection

	ID, _ := primitive.ObjectIDFromHex(newUser.ID)
	result, err := coll.UpdateOne(
		context.Background(),
		bson.D{
			{Key: "_id", Value: ID},
		},
		bson.D{
			{"$set", bson.D{
				{Key: "name", Value: newUser.Name},
				{Key: "Dob", Value: newUser.Dob},
				{Key: "Address", Value: newUser.Address},
				{Key: "Description", Value: newUser.Description},
				{Key: "createdAt", Value: time.Now()},
			}},
		})
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, result)
}

// Delete Existing User
func DeleteAnUser(c *gin.Context) {
	ID, _ := primitive.ObjectIDFromHex(c.Param("ID")) // Getting Param and converting it into MongoDB's ObjectID
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI("mongodb://localhost:27017"),
	) // Creating DB Connection
	if err != nil { // Checking Whehter there is any error while connecting
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil) // Pinging DB to Confirm DB Connection
	if err != nil {                        // Checking Whether there is any error while pinging the DB
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	db := client.Database("admybrand") // Creating or Accessing DB
	coll := db.Collection("User")      // Creating or Accessing Collection

	result, err := coll.DeleteOne(context.Background(), bson.D{{Key: "_id", Value: ID}})
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, result)
}
