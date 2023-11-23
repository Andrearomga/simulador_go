package models

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas" 
)


type Estacionamiento struct {
	Espacios       []bool
	MutexCapacidad *sync.Mutex
	Entrada        int 
	Salida         int
}

// NuevoEstacionamiento crea un nuevo estacionamiento con la cantidad de espacios especificada
func NuevoEstacionamiento(numEspacios int) *Estacionamiento {
	return &Estacionamiento{
		Espacios:       make([]bool, numEspacios),
		MutexCapacidad: &sync.Mutex{},
	}
}

// EstacionamientoLleno verifica si todos los espacios del estacionamiento están ocupados
func (e *Estacionamiento) EstacionamientoLleno() bool {
	
	for _, ocupado := range e.Espacios {
	
		if !ocupado {
			
			return false
		}
	}
	
	return true
}

func (e *Estacionamiento) EstacionarVehiculo(id int, vehiculo *Vehiculo) (estacionado bool) {
	e.MutexCapacidad.Lock()
	
	defer e.MutexCapacidad.Unlock()

	if e.EstacionamientoLleno() {
		fmt.Printf("Vehículo %d bloqueado, por el momento el estacionamiento se encuentra lleno.\n", id)
		return false
	}
	
	for i, ocupado := range e.Espacios {
		if !ocupado {
			
			e.Espacios[i] = true
			vehiculo.Mover(fyne.NewPos(float32(160+20*e.Entrada), 200))
			e.Entrada--
			fmt.Printf("Vehículo %d estacionado en espacio %d.\n", id, i+1) 
			MoverVehiculo(vehiculo, i)
			return true
		}
	}
	return false
}


func TiempoAleatorio() time.Duration {
	min := 1
	max := 5
	return time.Duration(min+rand.Intn(max-min+1)) * time.Second
}



func (e *Estacionamiento) VehiculoEntra(id int, wg *sync.WaitGroup, contenedor *fyne.Container, imagenPath string) {
    defer wg.Done() 
    fmt.Printf("Vehículo %d llega.\n", id) 
    e.Entrada++ 

    vehiculo := NuevoVehiculo(contenedor, imagenPath) 

    for !e.EstacionarVehiculo(id, vehiculo) {
        time.Sleep(100 * time.Millisecond) 
    }

    time.Sleep(TiempoAleatorio()) 

    e.SalirVehiculo(id, vehiculo.Imagen, contenedor) 
}

func (e *Estacionamiento) SalirVehiculo(id int, imagen *canvas.Image, contenedor *fyne.Container) {
    e.MutexCapacidad.Lock()
    defer e.MutexCapacidad.Unlock() 
  
    for i, ocupado := range e.Espacios {
        if ocupado {
            e.Espacios[i] = false
            break
        }
    }

    fmt.Printf("El Vehículo %d se dirige a la salida\n", id) 
    e.Salida++ 
    imagen.Move(fyne.NewPos(float32(20*e.Salida), 200))

    time.Sleep(time.Second * time.Duration(rand.Intn(2))) 

    fmt.Printf("Vehículo %d saliendo del estacionamiento\n", id) 
    contenedor.Remove(imagen) 
    contenedor.Refresh() 
    e.Salida--
}
