package services

import (
	"time"

	"api-mongo-go/dto"
	"api-mongo-go/models"
	"api-mongo-go/repository"
	"api-mongo-go/validators"

	"errors"


	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"


)

type TemporadaService struct {
	repo repository.TemporadaRepository
}

func (s TemporadaService) Crear(dto dto.TemporadaDTO) error {

	if err := validators.ValidarTemporada(dto); err != nil {
		return err
	}

	t := models.Temporada{
		NombreTemp:  dto.NombreTemp,
		Periodo:     dto.Periodo,
		Activo:      dto.Activo,
		AuditoriaID: dto.AuditoriaID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return s.repo.Insert(t)
}

func (s TemporadaService) Listar() ([]models.Temporada, error) {
	return s.repo.FindAll()
}

func (s TemporadaService) ObtenerPorID(id string) (*models.Temporada, error) {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("id inválido")
	}

	return s.repo.FindByID(objectID)
}

func (s TemporadaService) Actualizar(id string, dto dto.TemporadaDTO) error {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("id inválido")
	}

	update := bson.M{
		"nombre_temp": dto.NombreTemp,
		"periodo": dto.Periodo,
		"auditoria_id": dto.AuditoriaID,
		"activo": dto.Activo,
		"updated_at": time.Now(),
	}

	return s.repo.Update(objectID, update)
}

func (s TemporadaService) Eliminar(id string) error {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("id inválido")
	}

	return s.repo.SoftDelete(objectID)
}