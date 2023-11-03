package models

import (
	"sync"
)

const (
	CapacidadMaxima = 20
)

type Estacionamiento struct {
	sem     *sync.WaitGroup
	cajones []sync.Mutex
	contar  int
	mux     sync.Mutex
}

func NuevoEstacionamiento(capacidad int) *Estacionamiento {
	return &Estacionamiento{
		sem:     &sync.WaitGroup{},
		cajones: make([]sync.Mutex, capacidad),
	}
}

func (e *Estacionamiento) AgregarAuto(a Auto) {
	e.mux.Lock()
	if e.contar >= CapacidadMaxima {
		println("El estacionamiento está lleno. Esperando...")
		e.mux.Unlock()
		return
	}
	e.contar++
	e.mux.Unlock()

	e.sem.Add(1)
	e.sem.Wait()
	for i := range e.cajones {
		if e.IntentarEstacionar(i) {
			go a.Estacionar(e, i)
			break
		}
	}
}

func (e *Estacionamiento) IntentarEstacionar(indice int) bool {
	if e.cajones[indice].TryLock() {
		return true
	}
	return false
}

func (e *Estacionamiento) LiberarCajon(indice int) {
	e.cajones[indice].Unlock()
	e.mux.Lock()
	e.contar--
	e.mux.Unlock()
	e.sem.Done()
}