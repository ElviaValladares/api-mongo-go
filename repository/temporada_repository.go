package repository

import (
	"context"
	"time"

	"api-mongo-go/config"
	"api-mongo-go/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TemporadaRepository struct{}

func (r TemporadaRepository) Insert(t models.Temporada) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.DB.Collection("temporadas")

	_, err := collection.InsertOne(ctx, t)
	return err
}

func (r TemporadaRepository) FindAll() ([]models.Temporada, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.DB.Collection("temporadas")

	cursor, err := collection.Find(ctx, bson.M{
		"deleted_at": bson.M{"$exists": false},
	})
	if err != nil {
		return nil, err
	}

	var temporadas []models.Temporada
	err = cursor.All(ctx, &temporadas)

	return temporadas, err
}

func (r TemporadaRepository) FindByID(id primitive.ObjectID) (*models.Temporada, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.DB.Collection("temporadas")

	var temporada models.Temporada

	err := collection.FindOne(ctx, bson.M{
		"_id": id,
		"deleted_at": bson.M{"$exists": false},
	}).Decode(&temporada)

	if err != nil {
		return nil, err
	}

	return &temporada, nil
}

func (r TemporadaRepository) Update(id primitive.ObjectID, update bson.M) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.DB.Collection("temporadas")

	_, err := collection.UpdateOne(ctx,
		bson.M{
			"_id": id,
			"deleted_at": bson.M{"$exists": false},
		},
		bson.M{"$set": update},
	)

	return err
}

func (r TemporadaRepository) SoftDelete(id primitive.ObjectID) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.DB.Collection("temporadas")

	now := time.Now()

	_, err := collection.UpdateOne(ctx,
		bson.M{
			"_id": id,
			"deleted_at": bson.M{"$exists": false},
		},
		bson.M{
			"$set": bson.M{
				"deleted_at": now,
				"updated_at": now,
			},
		},
	)

	return err
}