// auto.go
package models

import (
	"math/rand"
	
	"time"
)

// Auto representa un vehículo en el estacionamiento.
type Auto struct {
	ID int
}

// Estacionar intenta estacionar el vehículo en el estacionamiento.
func (a Auto) Estacionar(e *Estacionamiento) {
	// Intentar estacionar el auto
	indice := e.IntentarEstacionar()
	if indice != -1 {
		// Simular el tiempo que el auto ocupa el cajón
		time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)

		// Liberar el cajón cuando el tiempo de estacionamiento ha pasado
		e.LiberarCajon(indice)
	} else {
		// Manejar el caso en el que no se pudo estacionar (por ejemplo, estacionamiento lleno)
		// Puedes agregar lógica adicional aquí según tus necesidades
		println("No hay cajones disponibles para estacionar el auto", a.ID)
	}
}