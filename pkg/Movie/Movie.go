package movie

import (
	"context"
	"fmt"
	"log"

	cfg "Kanopy/pkg/Config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

func GetTest() {
	clientOptions := options.Client().ApplyURI(cfg.GetConfigApp().LinkToConnectionDB)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Succeeded")

	coll := client.Database("KanopyDB").Collection("Films")

	filter := bson.M{"film_id": 3}
	var result bson.M
	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	

	fmt.Println((result["Название"]))

}