package models

import (
    "time"
    "math/rand"
)

type Auto struct {
	ID int
}

func CrearAuto(id int, estacionamiento *Estacionamiento) {
    auto := Auto{ID: id} 

    go func() {
        for {
            select {
            case <-time.After(time.Duration(rand.Intn(5)+1) * time.Second):
                // Intentar entrar al estacionamiento
                if estacionamiento.Entrar(auto) {
                    // Esperar un tiempo aleatorio entre 1 y 5 segundos
                    time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)

                    // Intentar salir del estacionamiento
                    estacionamiento.Salir(auto)
                } else {
                    // Bloquear el vehículo hasta que haya un espacio disponible
                    <-estacionamiento.EspacioDisponible
                }
            default:
                continue
            }
        }
    }()
}