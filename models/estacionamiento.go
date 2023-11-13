package models

import (
	"math/rand"
	"time"
	"fmt"
	"sync"
)

type Estacionamiento struct {
	Cajones []bool
	muxCajones sync.Mutex
	puertaEntrada sync.Mutex
	puertaSalida sync.Mutex
	semEspaciosDisponibles chan bool
}

const (
	CapacidadMaxima = 20
)

func NuevoEstacionamiento() *Estacionamiento {
	return &Estacionamiento{
		Cajones: make([]bool, CapacidadMaxima),
		semEspaciosDisponibles: make(chan bool, CapacidadMaxima),
	}
}

func (e *Estacionamiento) LlegadaVehiculo(id int) {
	fmt.Printf("Vehículo %d llegando al estacionamiento\n", id)

	// Esperar a que haya espacio en el estacionamiento
	<-e.semEspaciosDisponibles

	e.puertaEntrada.Lock()
	fmt.Printf("Vehículo %d entrando al estacionamiento\n", id)

	indiceCajon := e.intentarEstacionar()
	if indiceCajon != -1 {
		fmt.Printf("Vehículo %d estacionado en el cajón %d\n", id, indiceCajon)
		e.ocuparCajon(indiceCajon)
		tiempoAleatorio := time.Duration(rand.Intn(5) + 1)
		time.Sleep(tiempoAleatorio * time.Second)
		e.desocuparCajon(indiceCajon)
		fmt.Printf("Vehículo %d ha dejado el cajón %d\n", id, indiceCajon)
	} else {
		fmt.Printf("Vehículo %d no pudo encontrar un cajón disponible y se va\n", id)
	}

	e.puertaEntrada.Unlock()

	e.puertaSalida.Lock()
	fmt.Printf("Vehículo %d saliendo del estacionamiento\n", id)
	e.puertaSalida.Unlock()

	// Indicar que hay un espacio disponible
	e.semEspaciosDisponibles <- true
}

func (e *Estacionamiento) estacionamientoLleno() bool {
	for _, ocupado := range e.Cajones {
		if !ocupado {
			return false
		}
	}
	return true
}

func (e *Estacionamiento) intentarEstacionar() int {
	for i, ocupado := range e.Cajones {
		if !ocupado {
			return i
		}
	}
	return -1
}

func (e *Estacionamiento) ocuparCajon(indice int) {
	e.muxCajones.Lock()
	defer e.muxCajones.Unlock()
	e.Cajones[indice] = true
}

func (e *Estacionamiento) desocuparCajon(indice int) {
	e.muxCajones.Lock()
	defer e.muxCajones.Unlock()
	e.Cajones[indice] = false
}