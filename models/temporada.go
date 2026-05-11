package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Temporada struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	NombreTemp   string             `bson:"nombre_temp"`
	Periodo      string             `bson:"periodo"`
	FechaInicio  time.Time          `bson:"fecha_inicio"`
	FechaFin     time.Time          `bson:"fecha_fin"`
	AuditoriaID  int                `bson:"auditoria_id"`
	Activo       bool               `bson:"activo"`

	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}
