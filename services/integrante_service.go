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

type IntegranteService struct {
	repo repository.IntegranteRepository
}

func (s IntegranteService) Crear(dto dto.IntegranteDTO) error {

	if dto.ID == "" {
		return errors.New("id_integrante_liga es obligatorio")
	}

	integrante := models.IntegranteLiga{
		ID:              dto.ID,
		SecretPass:      dto.SecretPass,
		NombreCompleto:  dto.NombreCompleto,
		Fotografia:      dto.Fotografia,
		AuditoriaID:     dto.AuditoriaID,
		Activo:          dto.Activo,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	return s.repo.Insert(integrante)
}

func (s IntegranteService) Listar() ([]models.IntegranteLiga, error) {
	return s.repo.FindAll()
}


func (s IntegranteService) Eliminar(id string) error {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("id inválido")
	}

	return s.repo.SoftDelete(objectID)
}



func (s IntegranteService) ObtenerPorID(id string) (*models.IntegranteLiga, error) {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("id inválido")
	}

	return s.repo.FindByID(objectID)
}

func (s IntegranteService) Actualizar(id string, dto dto.IntegranteDTO) error {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("id inválido")
	}

	update := bson.M{
		"secret_pass": dto.SecretPass,
		"nombre_completo": dto.NombreCompleto,
		"fotografia": dto.Fotografia,
		"fecha_nacimiento": dto.FechaNacimiento,
		"activo": dto.Activo,
		"updated_at": time.Now(),
	}

	return s.repo.Update(objectID, update)
}

func (s IntegranteService) Login(id string, secretPass string) (*models.IntegranteLiga, error) {
    integrantes, err := s.repo.FindAll()
    if err != nil {
        return nil, errors.New("error al buscar integrantes")
    }
    for _, integrante := range integrantes {
        if integrante.ID == id && integrante.SecretPass == secretPass && integrante.DeletedAt == nil {
            return &integrante, nil
        }
    }
    return nil, errors.New("credenciales inválidas")
}