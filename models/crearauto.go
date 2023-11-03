package models

import (
	"time"
)

func CrearAutos(estacionamiento *Estacionamiento) {
	for i := 1; i <= 100; i++ {
		auto := Auto{ID: i}
		estacionamiento.AgregarAuto(auto)
		time.Sleep(time.Second) // Simular tiempo de llegada
	}
}