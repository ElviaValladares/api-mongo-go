package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Rama struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	NombreRama  string             `bson:"nombre_rama"`
	Descripcion string             `bson:"descripcion"`
	AuditoriaID int                `bson:"auditoria_id"`

	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}
