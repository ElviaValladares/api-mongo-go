package repository

import (
	"context"
	"time"

	"api-mongo-go/config"
	"api-mongo-go/models"

	"go.mongodb.org/mongo-driver/bson"
)

type IntegranteRepository struct{}

func (r IntegranteRepository) Insert(i models.IntegranteLiga) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.DB.Collection("integrantes")

	_, err := collection.InsertOne(ctx, i)
	return err
}

func (r IntegranteRepository) FindAll() ([]models.IntegranteLiga, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.DB.Collection("integrantes")

	cursor, err := collection.Find(ctx, bson.M{
		"deleted_at": bson.M{"$exists": false},
	})

	if err != nil {
		return nil, err
	}

	var integrantes []models.IntegranteLiga
	err = cursor.All(ctx, &integrantes)

	return integrantes, err
}

func (r IntegranteRepository) SoftDelete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := config.DB.Collection("integrantes")
	now := time.Now()
	_, err := collection.UpdateOne(ctx,
		bson.M{"_id": id, "deleted_at": bson.M{"$exists": false}},
		bson.M{"$set": bson.M{"deleted_at": now}},
	)
	return err
}

//

func (r IntegranteRepository) FindByID(id string) (*models.IntegranteLiga, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := config.DB.Collection("integrantes")
	var integrante models.IntegranteLiga
	err := collection.FindOne(ctx, bson.M{
		"_id":        id,
		"deleted_at": bson.M{"$exists": false},
	}).Decode(&integrante)
	if err != nil {
		return nil, err
	}
	return &integrante, nil
}

func (r IntegranteRepository) Update(id string, update bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := config.DB.Collection("integrantes")
	_, err := collection.UpdateOne(ctx,
		bson.M{
			"_id":        id,
			"deleted_at": bson.M{"$exists": false},
		},
		bson.M{"$set": update},
	)
	return err
}
