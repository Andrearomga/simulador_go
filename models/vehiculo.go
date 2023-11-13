// models/vehiculo.go
package models

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Vehiculo representa un vehículo.
type Vehiculo struct {
	id int
}

func (v *Vehiculo) Llegar() {
	fmt.Printf("Vehículo %d ha llegado al estacionamiento.\n", v.id)
}

func (v *Vehiculo) EntrarAlEstacionamiento(sem *sync.Mutex) {
	sem.Lock()
	fmt.Printf("Vehículo %d intentando entrar al estacionamiento.\n", v.id)
	time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
	fmt.Printf("Vehículo %d ha entrado al estacionamiento.\n", v.id)
	sem.Unlock()
}

func (v *Vehiculo) Estacionar() {
	fmt.Printf("Vehículo %d buscando lugar de estacionamiento.\n", v.id)
	time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
	fmt.Printf("Vehículo %d estacionado.\n", v.id)
	time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
	fmt.Printf("Vehículo %d ha dejado el estacionamiento.\n", v.id)
}