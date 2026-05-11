package services

import (
	"errors"
	"time"

	"api-mongo-go/dto"
	"api-mongo-go/models"
	"api-mongo-go/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RamaService struct {
	repo repository.RamaRepository
}

func (s RamaService) Crear(dto dto.RamaDTO) error {

	if dto.NombreRama == "" {
		return errors.New("nombre_rama es obligatorio")
	}

	rama := models.Rama{
		NombreRama:  dto.NombreRama,
		Descripcion: dto.Descripcion,
		AuditoriaID: dto.AuditoriaID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return s.repo.Insert(rama)
}

func (s RamaService) Listar() ([]models.Rama, error) {
	return s.repo.FindAll()
}

func (s RamaService) Eliminar(id string) error {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("id inválido")
	}

	return s.repo.SoftDelete(objectID)
}



func (s RamaService) ObtenerPorID(id string) (*models.Rama, error) {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("id inválido")
	}

	return s.repo.FindByID(objectID)
}

func (s RamaService) Actualizar(id string, dto dto.RamaDTO) error {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("id inválido")
	}

	update := bson.M{
		"nombre_rama": dto.NombreRama,
		"descripcion": dto.Descripcion,
		"updated_at": time.Now(),
	}

	return s.repo.Update(objectID, update)
}