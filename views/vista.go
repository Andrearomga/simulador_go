package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"math/rand"
	"sync"
	"time"
	"simulador/models"
)

const (
	anchoEspacio     = 60
	altoEspacio      = 40
	separadorEspacio = 50
)

func IniciarLlegada(numVehiculos int) {
	appPrincipal := app.New()
	rand.Seed(time.Now().UnixNano())

	estacionamiento := models.NuevoEstacionamiento(20)
	ventana, contenedor := configurarVentana(appPrincipal, numVehiculos, estacionamiento)
	ventana.ShowAndRun()
	ejecutarEstacionamiento(ventana, estacionamiento, contenedor, numVehiculos)
}

func ejecutarEstacionamiento(ventana fyne.Window, estacionamiento *models.Estacionamiento, contenedor *fyne.Container, numVehiculos int) {
	var coord sync.WaitGroup

	// Muestra el mensaje "Esperando inicio de la simulación"
	mensajeEspera := widget.NewLabel("Esperando inicio de la simulación")
	popUp := container.NewVBox(mensajeEspera)
	ventana.SetContent(popUp)
	ventana.Resize(fyne.NewSize(800, 500))

	// Cierra el mensaje después de 2 segundos
	go func() {
		time.Sleep(2 * time.Second)
		ventana.SetContent(container.NewVBox(contenedor, layout.NewSpacer()))
		ventana.Resize(fyne.NewSize(800, 500))
	}()

	// Inicia la simulación después de presionar el botón "Iniciar"
	go func() {
		time.Sleep(2 * time.Second)
		for i := 1; i <= numVehiculos; i++ {
			coord.Add(1)
			go estacionamiento.VehiculoEntra(i, &coord, contenedor, "assets/carro2200.png")
			time.Sleep(time.Duration(rand.ExpFloat64() * float64(time.Second)))
		}
		coord.Wait()
	}()
}

func configurarVentana(appPrincipal fyne.App, numVehiculos int, estacionamiento *models.Estacionamiento) (fyne.Window, *fyne.Container) {
	ventana := appPrincipal.NewWindow("Simulador-Estacionamiento")
	ventana.Resize(fyne.NewSize(800, 500))

	cont := container.NewWithoutLayout()

	for i := 0; i < 20; i++ {
		rectangulo := canvas.NewRectangle(color.RGBA{R: 0, G: 255, B: 0, A: 150}) // Color verde semi-transparente
		rectangulo.Resize(fyne.NewSize(float32(anchoEspacio), float32(altoEspacio)))
		rectangulo.FillColor = color.RGBA{R: 200, G: 200, B: 200, A: 150}        // Color de fondo más claro
		rectangulo.StrokeColor = color.Black                                       // Color del borde
		rectangulo.StrokeWidth = 1                                                // Ancho del borde

		cont.Add(rectangulo)
		models.MoverVehiculo(&models.Vehiculo{ObjetoVehiculo: rectangulo}, i)
	}

	botonIniciar := widget.NewButton("Iniciar", func() {
		go ejecutarEstacionamiento(ventana, estacionamiento, cont, numVehiculos)
	})

	ventana.SetContent(container.NewVBox(
		cont,
		layout.NewSpacer(),
		botonIniciar,
	))

	return ventana, cont
}
