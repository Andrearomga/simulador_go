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

func IniciarLlegada(numVehiculos int) {
	appMain := app.New()
	rand.Seed(time.Now().UnixNano())

	park := models.NuevoEstacionamiento(20)
	ventana, contenedor := configurarVentana(appMain)
	ejecutarEstacionamiento(park, contenedor, numVehiculos)

	ventana.ShowAndRun()
}

func configurarVentana(appMain fyne.App) (fyne.Window, *fyne.Container) {
	vent := appMain.NewWindow("Estacionamiento Up xD")
	vent.Resize(fyne.NewSize(800, 500))
	vent.SetFixedSize(true)

	imgRecurso, _ := fyne.LoadResourceFromPath("assets/bg.png")
	img := canvas.NewImageFromResource(imgRecurso)
	img.Resize(fyne.NewSize(800, 500))

	cont := container.NewWithoutLayout(img)

	// Agrega círculos y rectángulos para representar los espacios del estacionamiento
	numEspacios := 20
	espacioWidth := 40
	for i := 0; i < numEspacios; i++ {
		circulo := canvas.NewCircle(color.Black)
		circulo.Resize(fyne.NewSize(20, 20))
		cont.Add(circulo)
		circulo.Move(fyne.NewPos(float32(160+20*i), 200))

		rectangulo := canvas.NewRectangle(color.RGBA{R: 0, G: 255, B: 0, A: 150}) // Color verde semi-transparente
		rectangulo.Resize(fyne.NewSize(float32(espacioWidth), 20))
		cont.Add(rectangulo)
		rectangulo.Move(fyne.NewPos(float32(160+20*i), 200))
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
