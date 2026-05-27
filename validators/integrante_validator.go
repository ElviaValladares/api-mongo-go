package validators

import (
	"api-mongo-go/repository"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Valida que todos los IDs de roles existan en la base de datos
func ValidateRolesExist(roleIDs []string, rolRepo repository.RolRepository) ([]primitive.ObjectID, error) {
	if len(roleIDs) == 0 {
		return nil, errors.New("El integrante debe tener al menos un rol")
	}
	var objectIDs []primitive.ObjectID
	for _, id := range roleIDs {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, errors.New("ID de rol inválido: " + id)
		}
		rol, err := rolRepo.FindByID(objID)
		if err != nil || rol == nil {
			return nil, errors.New("El rol con ID " + id + " no existe")
		}
		objectIDs = append(objectIDs, objID)
	}
	return objectIDs, nil
}
