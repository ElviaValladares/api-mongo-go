package repository

import (
	"context"
	"time"

	"api-mongo-go/config"
	"api-mongo-go/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RolRepository struct{}

func (r RolRepository) Insert(rol models.Rol) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.DB.Collection("roles")

	_, err := collection.InsertOne(ctx, rol)
	return err
}

func (r RolRepository) FindAll() ([]models.Rol, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.DB.Collection("roles")

	cursor, err := collection.Find(ctx, bson.M{
		"deleted_at": bson.M{"$exists": false},
	})

	if err != nil {
		return nil, err
	}

	var roles []models.Rol
	err = cursor.All(ctx, &roles)

	return roles, err
}

func (r RolRepository) SoftDelete(id primitive.ObjectID) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.DB.Collection("roles")

	now := time.Now()

	_, err := collection.UpdateOne(ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"deleted_at": now}},
	)

	return err
}


func (r RolRepository) FindByID(id primitive.ObjectID) (*models.Rol, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.DB.Collection("roles")

	var rol models.Rol

	err := collection.FindOne(ctx, bson.M{
		"_id": id,
		"deleted_at": bson.M{"$exists": false},
	}).Decode(&rol)

	if err != nil {
		return nil, err
	}

	return &rol, nil
}

func (r RolRepository) Update(id primitive.ObjectID, update bson.M) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.DB.Collection("roles")

	_, err := collection.UpdateOne(ctx,
		bson.M{
			"_id": id,
			"deleted_at": bson.M{"$exists": false},
		},
		bson.M{"$set": update},
	)

	return err
}
