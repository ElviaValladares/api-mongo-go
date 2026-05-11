package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Rol struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	NombreRol   string             `bson:"nombre_rol"`
	Descripcion string             `bson:"descripcion"`
	AuditoriaID int                `bson:"auditoria_id"`

	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}
