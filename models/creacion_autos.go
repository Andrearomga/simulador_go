package models

import (
	"fmt"
	"sync"
	"math/rand"
	"time"
	"fyne.io/fyne/v2/canvas"
)

const (
	NumVehiculos = 100
)

func CrearVehiculos(sem *sync.Mutex, done chan<- bool, carros []*canvas.Image, updates chan<- string) {
	for i := 1; i <= NumVehiculos; i++ {
		v := &Vehiculo{id: i}
		v.Llegar()

		// Mostrar la imagen del carro en la GUI
		carros[i-1].Show()

		go func(v *Vehiculo) {
			v.EntrarAlEstacionamiento(sem)
			v.Estacionar()
			updates <- fmt.Sprintf("Vehículo %d estacionado", v.id)
		}(v)

		interval := time.Duration(rand.ExpFloat64()*float64(time.Second))
		time.Sleep(interval)
	}

	done <- true
}