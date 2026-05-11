package dto

type RolDTO struct {
	NombreRol   string `json:"nombre_rol"`
	Descripcion string `json:"descripcion"`
	AuditoriaID int    `json:"auditoria_id"`
}
