package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IntegranteLiga struct {
	ID              string               `bson:"_id"` // usuario app
	SecretPass      string               `bson:"secret_pass,omitempty"`
	NombreCompleto  string               `bson:"nombre_completo"`
	Fotografia      string               `bson:"fotografia,omitempty"`
	FechaNacimiento time.Time            `bson:"fecha_nacimiento"`
	AuditoriaID     int                  `bson:"auditoria_id"`
	Activo          bool                 `bson:"activo"`
	Roles           []primitive.ObjectID `bson:"roles" json:"roles"`

	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}
