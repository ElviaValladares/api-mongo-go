package dto

type TemporadaDTO struct {
	NombreTemp  string `json:"nombre_temp"`
	Periodo     string `json:"periodo"`
	FechaInicio string `json:"fecha_inicio"`
	FechaFin    string `json:"fecha_fin"`
	AuditoriaID int    `json:"auditoria_id"`
	Activo      bool   `json:"activo"`
}
