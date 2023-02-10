package database

import (
	"api/schemas"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func Find(id string) (*schemas.Register, error) {
	client, ctx := StartConnect()
	defer client.Disconnect(ctx)

	collection := client.Database("nica").Collection("nina")
	filter := bson.M{"id": id}

	var request schemas.Register
	err := collection.FindOne(context.Background(), filter).Decode(&request)
	if err != nil {
		fmt.Println("Esse GameID não existe.")
	}

	return &request, nil
}
