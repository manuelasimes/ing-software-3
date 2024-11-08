package dto

type ViajeDto struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Destino     string `json:"destino"`
	Fecha       string `json:"fecha"`
	Descripcion string `json:"descripcion"`
	Puntaje     string `json:"puntaje"`
	Comentario  string `json:"comentario"`
}

type ViajesDto []ViajeDto
