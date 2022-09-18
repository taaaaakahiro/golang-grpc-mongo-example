package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/taaaaakahiro/grpc-example/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	//if crash the code, get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("hw")

	// connect to database
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Blog Service Started")
	collection := client.Database(os.Getenv("DATABASE")).Collection("user")
	var user entity.User
	err = collection.FindOne(context.TODO(), bson.D{{"id", 1}}).Decode(&user)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println(user)

}
