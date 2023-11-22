package models

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas" 
)

// Estacionamiento representa un estacionamiento con vehículos
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
	// Itera sobre cada elemento (booleano) en el slice e.Espacios
	for _, ocupado := range e.Espacios {
	// Verifica si el espacio no está ocupado (es decir, si está libre)
		if !ocupado {
			// Si encuentra al menos un espacio libre, devuelve falso
			return false
		}
	}
	// Si todos los espacios están ocupados, devuelve verdadero
	return true
}

func (e *Estacionamiento) EstacionarVehiculo(id int, vehiculo *Vehiculo) (estacionado bool) {
	// Bloquea el mutex para garantizar acceso exclusivo a la capacidad del estacionamiento
	e.MutexCapacidad.Lock()
	// Se asegura de que se desbloquee el mutex incluso si hay un error o una excepción
	defer e.MutexCapacidad.Unlock()
	// Verifica si el estacionamiento está lleno llamando a la función EstacionamientoLleno
	if e.EstacionamientoLleno() {
		fmt.Printf("Vehículo %d bloqueado, por el momento el estacionamiento se encuentra lleno.\n", id)
		// Retorna false porque el estacionamiento está lleno y el vehículo no puede estacionar
		return false
	}
	// Itera sobre los espacios del estacionamiento para encontrar uno libre
	for i, ocupado := range e.Espacios {
		// Verifica si el espacio actual no está ocupado (es decir, está libre)
		if !ocupado {
			// Marca el espacio como ocupado
			e.Espacios[i] = true
			// Mueve el vehículo a la posición en el estacionamiento y actualiza la posición de entrada
			vehiculo.Mover(fyne.NewPos(float32(160+20*e.Entrada), 200))
			// Actualiza el contador de entrada
			e.Entrada--
			// Imprime un mensaje indicando que el vehículo ha sido estacionado en el espacio i+1
			fmt.Printf("Vehículo %d estacionado en espacio %d.\n", id, i+1) // Suma 1 al índice
			// Mueve el vehículo en la interfaz gráfica
			MoverVehiculo(vehiculo, i)
			// Retorna true porque el vehículo ha sido estacionado exitosamente
			return true
		}
	}
	return false
}

// TiempoAleatorio genera una duración aleatoria entre 1 y 5 segundos
func TiempoAleatorio() time.Duration {
	min := 1
	max := 5
	return time.Duration(min+rand.Intn(max-min+1)) * time.Second
}


// VehiculoEntra simula la entrada de un vehículo al estacionamiento
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

// SalirVehiculo simula la salida de un vehículo del estacionamiento
func (e *Estacionamiento) SalirVehiculo(id int, imagen *canvas.Image, contenedor *fyne.Container) {
    e.MutexCapacidad.Lock()
    defer e.MutexCapacidad.Unlock()

    for i, ocupado := range e.Espacios {
        if ocupado {
            e.Espacios[i] = false
            break
        }
    }

    fmt.Printf("Vehículo %d se dirige a la salida\n", id)
    e.Salida++
    imagen.Move(fyne.NewPos(float32(20*e.Salida), 200))

    time.Sleep(time.Second * time.Duration(rand.Intn(2)))
    fmt.Printf("Vehículo %d sale del estacionamiento\n", id)
    contenedor.Remove(imagen)
    contenedor.Refresh()
    e.Salida--
}
