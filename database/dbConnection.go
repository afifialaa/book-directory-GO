package dbConnection

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/afifialaa/REST-GO/models"

	"context"
	"fmt"
	"log"
)

var booksCollection *mongo.Collection

func Connect() {
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

	// Set database and collection
	booksCollection = client.Database("library").Collection("books")
	return
}

type BookType struct {
	BookID             string
	Title              string
	Authors            string
	Average_rating     string
	Isbn               string
	Isb13              string
	Language_code      string
	Ratings_count      string
	Text_reviews_count string
	Publication_dates  string
	Publisher          string
}

func SearchByAuthor(author string)[]models.BookType{
	filter := bson.D{{"authors", author}}

	// Finding multiple documents returns a cursor
	cursor, err := booksCollection.Find(context.TODO(), filter)

	if err != nil {
		fmt.Println(err)
	}

	var result []models.BookType

	// Iterate over the cursor and decode each document
	for cursor.Next(context.TODO()) {
		var book models.BookType

		err := cursor.Decode(&book)

		if err != nil {
			fmt.Println(err)
		}

		result = append(result, book)
	}

	cursor.Close(context.TODO())
	return result
}


func SearchByID(bookId string) models.BookType {
	filter := bson.D{{"bookID", bookId}}

	var result models.BookType
	err := booksCollection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			fmt.Println("There are no document")
		}
	}
	fmt.Println(result.BookID)
	return result
}

func SearchByTitle(bookTitle string) BookType {
	filter := bson.D{{"title", bookTitle}}

	var result BookType
	err := booksCollection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			fmt.Println("There are no document")
		}
	}
	fmt.Println("result is " , result)
	return result
}
