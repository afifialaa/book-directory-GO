package dbConnection

import (
	  "go.mongodb.org/mongo-driver/bson"
	  "go.mongodb.org/mongo-driver/mongo"
	  "go.mongodb.org/mongo-driver/mongo/options"

	  "github.com/afifialaa/validation"

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

var userCollection *mongo.Collection

func Connect() {
	fmt.Println("#Connect")
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
	userCollection = client.Database("private").Collection("users")
	fmt.Println(userCollection)
}

// Insert new user
func SaveUser(user *validation.User_type) bool {
	insertResult, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println("mongodb error " ,err.Error())
		return false
	}

	fmt.Println("#saveuser -> user was created: ", insertResult.InsertedID)
	return true;
}

func FindUser(user *validation.User_login_type)bool{
	var result validation.User_type

	filter := bson.D{{"email", user.Email}}
	err := userCollection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		fmt.Println("user was not found")
		return false
	}

	return true
}
