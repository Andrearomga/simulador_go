// vista.go
package views

import (
	"fmt"
	"sync"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"simulador/models"
    "fyne.io/fyne/v2"
)

func SimulacionEstacionamiento() {
	// Crear una aplicación y una ventana
	a := app.New()
	w := a.NewWindow("Simulación de Estacionamiento")

	// Crear un estacionamiento
	estacionamiento := models.NuevoEstacionamiento()

	// Crear un semáforo para controlar la entrada al estacionamiento
	sem := &sync.Mutex{}

	// Crear un canal para señalar cuando todos los vehículos han llegado
	done := make(chan bool)

	// Crear un widget de imagen para cada vehículo en la simulación
	carros := make([]*canvas.Image, models.NumVehiculos)
	for i := range carros {
		carros[i] = canvas.NewImageFromFile("assets/carro2.png") // Aquí es donde agregas la ruta a tu imagen
		carros[i].FillMode = canvas.ImageFillOriginal
		carros[i].Hide() // Ocultar la imagen inicialmente
	}

	// Crear y lanzar los vehículos
	go models.CrearVehiculos(sem, done, carros)

	// Crear un widget para cada cajón en el estacionamiento
	cajones := make([]*canvas.Rectangle, models.CapacidadMaxima)
	for i := range cajones {
		cajones[i] = canvas.NewRectangle(color.RGBA{0, 255, 0, 255}) // Verde para cajón libre
		cajones[i].SetMinSize(fyne.NewSize(10, 10)) // Establecer el tamaño del rectángulo
	}

	// Crear un botón que inicia la simulación
	boton := widget.NewButton("Iniciar Simulación", func() {
		// Actualizar el estado de los cajones
		for i, cajon := range estacionamiento.Cajones { // Asegúrate de que Cajones es un campo exportado en tu estructura Estacionamiento
			if cajon {
				cajones[i].FillColor = color.RGBA{255, 0, 0, 255} // Rojo para cajón ocupado
				cajones[i].Refresh()
			} else {
				cajones[i].FillColor = color.RGBA{0, 255, 0, 255} // Verde para cajón libre
				cajones[i].Refresh()
			}
		}

		// Esperar a que todos los vehículos hayan llegado
		<-done

		// Imprimir un mensaje final
		fmt.Println("Todos los vehículos han llegado.")
	})

	// Añadir el botón y los widgets de los cajones a la ventana y mostrarla
	contenedorCajones := container.NewHBox() // Contenedor para los cajones
	for _, cajon := range cajones {
		contenedorCajones.Add(cajon)
	}

	contenedorCarros := container.NewVBox() // Contenedor para los carros
	for _, carro := range carros {
		contenedorCarros.Add(carro)
	}

	contenedor := container.NewVBox(boton, contenedorCajones, contenedorCarros) // Añadir el contenedor de cajones y los carros después del botón
	w.SetContent(contenedor)
	w.ShowAndRun()
}