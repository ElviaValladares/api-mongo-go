package services

import (
	"time"

	"api-mongo-go/dto"
	"api-mongo-go/models"
	"api-mongo-go/repository"

	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

type RolService struct {
	repo repository.RolRepository
}

func (s RolService) Crear(dto dto.RolDTO) error {

	rol := models.Rol{
		NombreRol:   dto.NombreRol,
		Descripcion: dto.Descripcion,
		AuditoriaID: dto.AuditoriaID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return s.repo.Insert(rol)
}

func (s RolService) Listar() ([]models.Rol, error) {
	return s.repo.FindAll()
}


func (s RolService) Eliminar(id string) error {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("id inválido")
	}

	return s.repo.SoftDelete(objectID)
}



func (s RolService) ObtenerPorID(id string) (*models.Rol, error) {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("id inválido")
	}

	return s.repo.FindByID(objectID)
}

func (s RolService) Actualizar(id string, dto dto.RolDTO) error {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("id inválido")
	}

	update := bson.M{
		"nombre_rol": dto.NombreRol,
		"descripcion": dto.Descripcion,
		"updated_at": time.Now(),
	}

	return s.repo.Update(objectID, update)
}