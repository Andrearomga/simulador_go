package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	//"image/color"
)

// Vehiculo representa un vehículo en el estacionamiento
type Vehiculo struct {
	ObjetoVehiculo fyne.CanvasObject
	Imagen         *canvas.Image
}


// NuevoVehiculo crea un nuevo vehículo y lo agrega al contenedor proporcionado
func NuevoVehiculo(contenedor *fyne.Container, imagenPath string) *Vehiculo {
	imagen := canvas.NewImageFromFile(imagenPath)
	imagen.Resize(fyne.NewSize(40, 40)) 

	contenedor.Add(imagen)
	imagen.Move(fyne.NewPos(30, 260))
	contenedor.Refresh()

	return &Vehiculo{
		ObjetoVehiculo: imagen,
		Imagen:         imagen,
	}
}



// MoverVehiculo mueve el vehículo a la posición calculada en base al espacio proporcionado
func MoverVehiculo(vehiculo *Vehiculo, espacio int) {
	x := 200
	y := 200
	
	x = 40*espacio + x
	y = 400

	vehiculo.ObjetoVehiculo.Move(fyne.NewPos(float32(x), float32(y)))
}

// Mover mueve el vehículo a la posición especificada
func (v *Vehiculo) Mover(pos fyne.Position) {
	v.ObjetoVehiculo.Move(pos)
}