package dto

type IntegranteDTO struct {
	ID              string `json:"id_integrante_liga"`
	SecretPass      string `json:"secret_pass"`
	NombreCompleto  string `json:"nombre_completo"`
	Fotografia      string `json:"fotografia"`
	FechaNacimiento string `json:"fecha_nacimiento"`
	AuditoriaID     int    `json:"auditoria_id"`
	Activo          bool   `json:"activo"`
}
