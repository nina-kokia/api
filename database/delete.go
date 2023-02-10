package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func Delete(coupon string) error {
	client, ctx := StartConnect()
	defer client.Disconnect(ctx)

	collection := client.Database("nica").Collection("coupon")
	filter := bson.M{"coupon": coupon}

	_, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		panic(err)
	}

	return nil
}
