// estacionamiento.go
package models

import (
	"sync"
	"time"
)

const (
	CapacidadMaxima = 20
)

// Estacionamiento representa el estacionamiento con cajones y una puerta compartida.
type Estacionamiento struct {
	cajones []bool
	mux     sync.Mutex
	puerta  sync.Mutex
}

// NuevoEstacionamiento crea un nuevo estacionamiento con la capacidad dada.
func NuevoEstacionamiento(capacidad int) *Estacionamiento {
	return &Estacionamiento{
		cajones: make([]bool, 0, capacidad),
	}
}

// EstacionamientoLleno verifica si el estacionamiento está lleno.
func (e *Estacionamiento) EstacionamientoLleno() bool {
	return len(e.cajones) >= CapacidadMaxima
}

// IntentarEstacionar intenta encontrar un cajón disponible en el estacionamiento.
func (e *Estacionamiento) IntentarEstacionar() int {
	e.mux.Lock()
	defer e.mux.Unlock()

	for i, ocupado := range e.cajones {
		if !ocupado {
			e.cajones[i] = true
			return i
		}
	}
	return -1
}

// LiberarCajon libera el cajón indicado en el estacionamiento.
func (e *Estacionamiento) LiberarCajon(indice int) {
	e.mux.Lock()
	defer e.mux.Unlock()

	if indice >= 0 && indice < len(e.cajones) {
		e.cajones[indice] = false
		println("Auto liberó el cajón:", indice)
	} else {
		println("Intento de liberar un cajón inválido:", indice)
	}
}

// AgregarAuto intenta agregar un auto al estacionamiento.
func (e *Estacionamiento) AgregarAuto(a Auto) {
	e.puerta.Lock()
	defer e.puerta.Unlock()

	e.mux.Lock()
	defer e.mux.Unlock()

	// Esperar hasta que haya cajones disponibles
	for e.EstacionamientoLleno() {
		println("El estacionamiento está lleno. Esperando...")
		e.mux.Unlock()
		time.Sleep(time.Second)
		e.mux.Lock()
	}

	// Intentar estacionar el auto
	indice := e.IntentarEstacionar()
	if indice != -1 {
		println("Auto estacionado en el cajón:", indice)
	} else {
		println("Error al intentar estacionar el auto. No se encontró un cajón disponible.")
	}
}
