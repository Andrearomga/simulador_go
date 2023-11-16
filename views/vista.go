package views

import (
	"image/color"
	"math/rand"
	"simulador/models"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/app"
)

const (
	espacioWidth     = 40
	espacioHeight    = 20
	espacioSeparador = 10
)

func IniciarLlegada(numVehiculos int) {
	appMain := app.New()
	rand.Seed(time.Now().UnixNano())

	park := models.NuevoEstacionamiento(20)
	ventana, contenedor := configurarVentana(appMain, numVehiculos, park)
	ejecutarEstacionamiento(park, contenedor, numVehiculos)

	ventana.ShowAndRun()
}

func configurarVentana(appMain fyne.App, numVehiculos int, park *models.Estacionamiento) (fyne.Window, *fyne.Container) {
	vent := appMain.NewWindow("Estacionamiento Up xD")
	vent.Resize(fyne.NewSize(800, 500))

	// Habilita la opción de maximizar la ventana y muestra el ícono de la ventana
	vent.SetMaster()

	imgRecurso, _ := fyne.LoadResourceFromPath("assets/bg.png")
	img := canvas.NewImageFromResource(imgRecurso)
	img.Resize(fyne.NewSize(800, 500))

	cont := container.NewWithoutLayout(img)

	// Agrega círculos y rectángulos para representar los espacios del estacionamiento
	for i := 0; i < 20; i++ { // Mantenido en 20 para tener una sola fila
		circulo := canvas.NewCircle(color.Black)
		circulo.Resize(fyne.NewSize(20, 20))
		cont.Add(circulo)
		moverVehiculo(circulo, i)
	}

	for i := 0; i < 20; i++ { // Mantenido en 20 para tener una sola fila
		rectangulo := canvas.NewRectangle(color.RGBA{R: 0, G: 255, B: 0, A: 150}) // Color verde semi-transparente
		rectangulo.Resize(fyne.NewSize(espacioWidth, espacioHeight))
		cont.Add(rectangulo)
		moverVehiculo(rectangulo, i)
	}

	vent.SetContent(cont)
	return vent, cont
}

func ejecutarEstacionamiento(park *models.Estacionamiento, contenedor *fyne.Container, numVehiculos int) {
	var coord sync.WaitGroup
	go func() {
		time.Sleep(2 * time.Second)
		for i := 1; i <= numVehiculos; i++ {
			coord.Add(1)
			go park.VehiculoEntra(i, &coord, contenedor)
			time.Sleep(time.Duration(rand.ExpFloat64() * float64(time.Second)))
		}
		coord.Wait()
	}()
}

func moverVehiculo(carro fyne.CanvasObject, espacio int) {
	x := 200
	y := 200

	// Calcula las coordenadas en línea para una sola fila
	x = espacioWidth*espacio + x
	y = 400

	carro.Move(fyne.NewPos(float32(x), float32(y)))
}
