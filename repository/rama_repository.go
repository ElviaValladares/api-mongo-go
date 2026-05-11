package repository

import (
	"context"
	"time"

	"api-mongo-go/config"
	"api-mongo-go/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RamaRepository struct{}

func (r RamaRepository) Insert(rama models.Rama) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.DB.Collection("ramas")

	_, err := collection.InsertOne(ctx, rama)
	return err
}

func (r RamaRepository) FindAll() ([]models.Rama, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.DB.Collection("ramas")

	cursor, err := collection.Find(ctx, bson.M{
		"deleted_at": bson.M{"$exists": false},
	})
	if err != nil {
		return nil, err
	}

	var ramas []models.Rama
	err = cursor.All(ctx, &ramas)

	return ramas, err
}

func (r RamaRepository) SoftDelete(id primitive.ObjectID) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.DB.Collection("ramas")

	now := time.Now()

	_, err := collection.UpdateOne(ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"deleted_at": now}},
	)

	return err
}


func (r RamaRepository) FindByID(id primitive.ObjectID) (*models.Rama, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.DB.Collection("ramas")

	var rama models.Rama

	err := collection.FindOne(ctx, bson.M{
		"_id": id,
		"deleted_at": bson.M{"$exists": false},
	}).Decode(&rama)

	if err != nil {
		return nil, err
	}

	return &rama, nil
}

func (r RamaRepository) Update(id primitive.ObjectID, update bson.M) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.DB.Collection("ramas")

	_, err := collection.UpdateOne(ctx,
		bson.M{
			"_id": id,
			"deleted_at": bson.M{"$exists": false},
		},
		bson.M{"$set": update},
	)

	return err
}