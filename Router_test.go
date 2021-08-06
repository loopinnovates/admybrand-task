package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestGetAllUsers(t *testing.T) {
	router := Router()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGetUserByID(t *testing.T) {
	router := Router()
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/user/610c73b2537eda325d239d45", nil)
	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func TestCreateUser(t *testing.T) {
	router := Router()
	rr := httptest.NewRecorder()
	data := User{
		Name:        "dummy",
		Dob:         "2000-01-01",
		Address:     "dummy's Address",
		Description: "dummy's Description",
	}
	json_data, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", "/adduser", bytes.NewBuffer(json_data))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(rr, req)

	assert.Equal(t, 202, rr.Code)
}

func TestDeleteAnUser(t *testing.T) {
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
			Key:   "name",
			Value: "dummy",
		}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	router := Router()
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/deleteuser/"+result.ID, nil)
	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func TestUpdateUser(t *testing.T) {

	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI("mongodb://localhost:27017"),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	db := client.Database("admybrand")
	coll := db.Collection("User")

	var result User
	coll.FindOne(
		context.Background(),
		bson.D{{
			Key:   "name",
			Value: "dummy",
		}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	router := Router()
	rr := httptest.NewRecorder()
	data := User{
		ID:          result.ID,
		Name:        "dummyUpdated",
		Dob:         "2000-01-01",
		Address:     "dummyUpdated's Address",
		Description: "dummyUpdated's Description",
	}
	json_data, _ := json.Marshal(data)
	req, _ := http.NewRequest("PUT", "/updateuser", bytes.NewBuffer(json_data))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}
