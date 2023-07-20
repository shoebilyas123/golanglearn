package dbConnect

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

const connectionString = "mongodb+srv://shoebilyas432:WERV7JQvW8Z6Gede@snaps-cluster.fbtao.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"


var collection *mongo.Connection

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	HandleError(err)
	fmt.Println("Connection to MongoDB instance successful")
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection Instance is Ready")
}