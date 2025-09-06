package main

import (
	"context"
	"fmt"
	"go_tuts/routes"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Student struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

var Database []Student = []Student{}

func main() {
	db := "tuts"
	client := LoadMongoDB("mongodb://localhost:27017")
	fmt.Printf("Succesfully connect to mongodb")

	route := gin.Default()
	routes.UserRoute(route, client, db)
	routes.ProductRoute(route)

	route.Run(":8080")

}

// context, Future (async await)
func LoadMongoDB(mongoUrl string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		panic(err)
	}
	//mysql client
	//postgres client
	//mongo client  (create , read, update ,delete )
	return client

}
