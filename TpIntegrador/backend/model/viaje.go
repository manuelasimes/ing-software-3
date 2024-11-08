package model

type Viaje struct {
	ID          int    `gorm:"primaryKey"`
	Nombre      string `gorm:"type:varchar(350);not null"`
	Destino     string `gorm:"type:varchar(350);not null"`
	Fecha       string `gorm:"type:text"`
	descripcion string `gorm:"type:varchar(350);not null"`
	Puntaje     int    `gorm:"type:int"`
	Comentario  string `gorm:"type:varchar(350);not null"`
}

type Viajes []Viaje
