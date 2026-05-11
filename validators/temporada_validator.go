package validators

import (
	"errors"
	"api-mongo-go/dto"
)

func ValidarTemporada(t dto.TemporadaDTO) error {

	if t.NombreTemp == "" {
		return errors.New("nombre_temp es obligatorio")
	}

	if t.Periodo == "" {
		return errors.New("periodo es obligatorio")
	}

	return nil
}
