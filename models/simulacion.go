package models

import (
	"time"
	"math/rand"
)

func Simulacion(actualizarVista func(int)) {
	rand.Seed(time.Now().UnixNano())

	e := NuevoEstacionamiento(20, actualizarVista)

	for i := 0; i < 100; i++ {
		CrearAuto(i+1, e)
		
        // Los autos llegan en intervalos aleatorios.
        time.Sleep(time.Duration(rand.ExpFloat64()/float64(i+1)) * time.Second)
    }
}