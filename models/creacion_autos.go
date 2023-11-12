package models

import (
	"fmt"
	"math/rand"
	"time"
	"image/color"
)

// CrearAutos crea continuamente autos y los agrega al estacionamiento.
func CrearAutos(estacionamiento *Estacionamiento, updateViewFunc func(id int, clr color.Color)) {
	for i := 1; ; i++ {
		auto := Auto{ID: i}
		if estacionamiento.EstacionamientoLleno() {
			fmt.Println("El estacionamiento está lleno. Esperando...")
			time.Sleep(5 * time.Second) // Esperar antes de intentar nuevamente
			continue
		}
		estacionamiento.AgregarAuto(auto)

		// Simular tiempo de llegada con distribución de Poisson
		lambda := 0.2
		sleepTime := rand.ExpFloat64() / lambda
		time.Sleep(time.Duration(sleepTime) * time.Second)

		// Llamada a la función de actualización de la vista (por implementar)
		updateViewFunc(i, color.RGBA{R: 255, G: 255, B: 255, A: 255})
	}
}