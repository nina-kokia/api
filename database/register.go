package database

import (
	"api/schemas"
	"context"
)

func Register(id string, ranking string) error {
	client, ctx := StartConnect()
	defer client.Disconnect(ctx)

	collection := client.Database("nica").Collection("nina")
	filter := schemas.Register{Id: id, Ranking: ranking}

	_, err := collection.InsertOne(context.Background(), filter)
	if err != nil {
		panic(err)
	}

	return err
}

func RegisterCoupon(coupon string) error {
	client, ctx := StartConnect()
	defer client.Disconnect(ctx)

	collection := client.Database("nica").Collection("coupon")
	filter := schemas.RegisterCoupon{Coupon: coupon}

	_, err := collection.InsertOne(context.Background(), filter)
	if err != nil {
		panic(err)
	}

	return err
}
