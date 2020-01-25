package dao

import (
	"context"
	"fmt"
	"log"
	"time"

	. "top-ten-stonks/models"

	mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type PapersDAO struct {
	URI string
}

var BooksCollection *mongo.Collection
var ctx context.Context

const (
	COLLECTION = "papers"
	DATABASE   = "redventures"
)

func (p *PapersDAO) Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, error := mongo.NewClient(options.Client().ApplyURI(p.URI))
	error = client.Connect(ctx)

	//Checking the connection
	error = client.Ping(context.TODO(), nil)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("Database connected")

	//Specify your respective collection
	BooksCollection = client.Database(DATABASE).Collection(COLLECTION)
}

func (m *PapersDAO) Create(paper Paper) {
	_, err := BooksCollection.InsertOne(ctx, bson.M{"name": paper.Name, "dayPerCent": paper.DayPerCent, "value": paper.Value, "corp": paper.Corp})
	if err != nil {
		log.Fatal(err)
	}
}
