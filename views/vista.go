package views



import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2"
	"simulador/models"
	"fyne.io/fyne/v2/layout"
)

func IniciarVentana() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Estacionamiento")
	myWindow.SetFixedSize(true)

	// Crear imágenes para el estacionamiento y el auto
	estacionamientoImagen := canvas.NewImageFromFile("assets/estacion.png")
	estacionamientoImagen.FillMode = canvas.ImageFillOriginal

	autoImagen := canvas.NewImageFromFile("assets/carro220.png")
	autoImagen.FillMode = canvas.ImageFillOriginal

	// Crear una vista que contiene la imagen del estacionamiento y el auto
	vista := container.NewHBox(layout.NewSpacer(), autoImagen, estacionamientoImagen, layout.NewSpacer())
	iniciarBoton := widget.NewButton("Iniciar", func() {
		// Iniciar la simulación de llegada continua de vehículos desde el paquete models
		go models.SimularEstacionamiento(100, 20) // Aquí está la corrección
	})

	// Crear una vista vertical que contiene el botón en la parte superior y la vista en la parte inferior
	vistaConBoton := container.NewVBox(iniciarBoton, vista)

	myWindow.SetContent(vistaConBoton)
	myWindow.Resize(fyne.NewSize(500, estacionamientoImagen.Size().Height))
	myWindow.ShowAndRun()
}