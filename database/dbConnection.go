package dbConnection

import (
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
	"fmt"
	"context"
	"log"
)

type UserType struct{
	firstName string
	lastName string
	email string
	password string
}

func Connect(){
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// collection
	collection := client.Database("private").Collection("users")
	fmt.Println(collection)

	// Mock 
	filter := bson.D{{"email", "afifi@gmail.com"}}

	var result UserType
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)
}