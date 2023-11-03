package models

import (
	"sync"
)

type Estacionamiento struct {
	capacidad          int
	autos              chan Auto
	EspacioDisponible  chan bool
	mutex              sync.Mutex // Mutex para sincronizar el acceso a la entrada/salida del estacionamiento.
	actualizarVista    func(int)  // Función para actualizar la vista
}

func NuevoEstacionamiento(capacidad int, actualizarVista func(int)) *Estacionamiento {
	return &Estacionamiento{
		capacidad:         capacidad,
		autos:             make(chan Auto, capacidad),
		EspacioDisponible: make(chan bool, capacidad),
		actualizarVista:   actualizarVista,
	}
}

func (e *Estacionamiento) Entrar(a Auto) bool {
	e.mutex.Lock() // Bloquear la entrada/salida del estacionamiento.
	defer e.mutex.Unlock()

	if len(e.autos) < e.capacidad {
		e.autos <- a
		e.actualizarVista(len(e.autos)) // Actualizar la vista
		return true
	}
	return false
}

func (e *Estacionamiento) Salir(a Auto) {
	e.mutex.Lock() // Bloquear la entrada/salida del estacionamiento.
	defer e.mutex.Unlock()

	for i := 0; i < len(e.autos); i++ {
		if auto := <-e.autos; auto.ID == a.ID {
			e.EspacioDisponible <- true
			e.actualizarVista(len(e.autos)) // Actualizar la vista
			break
		} else {
			e.autos <- auto // Si el auto que salió del canal no es el que queremos, lo devolvemos al canal.
		}
	}
}