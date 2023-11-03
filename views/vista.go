package views

import (
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/widget"
    "image/color"
    "simulador/models" // Asegúrate de usar el paquete correcto para tus modelos
)

func IniciarVentana() {
    a := app.New()
    w := a.NewWindow("Estacionamiento")

    espacios := make([]fyne.CanvasObject, 20)
for i := range espacios {
    rect := canvas.NewRectangle(color.RGBA{R: 100, G: 100, B: 100, A: 255})
    rect.SetMinSize(fyne.NewSize(50, 20)) // Cambiar el ancho a 50 y el alto a 20
    espacios[i] = rect
}

	autosImagenes := []string{
		"assets/carro220.png",
		//"assets/carro2.png",
		// Agrega más rutas de imágenes según sea necesario
	}

    // Declarar parkingContainer antes de actualizarVista
    var parkingContainer *fyne.Container

	// Función para actualizar la vista
	actualizarVista := func(numAutos int) {
		for i := 0; i < numAutos; i++ {
			auto := canvas.NewImageFromFile(autosImagenes[i % len(autosImagenes)])
			auto.FillMode = canvas.ImageFillOriginal
			espacios[i] = auto
		}
		for i := numAutos; i < len(espacios); i++ {
			rect := canvas.NewRectangle(color.RGBA{R: 100, G: 100, B: 100, A: 255})
			rect.SetMinSize(fyne.NewSize(50, 20)) // Cambiar el ancho a 50 y el alto a 20
			espacios[i] = rect
		}
		canvas.Refresh(parkingContainer)
	}

    // Inicializar parkingContainer después de actualizarVista
    parkingContainer = container.NewGridWrap(fyne.NewSize(60, 100), espacios...)

	// Botón para iniciar la simulación
	iniciarButton := widget.NewButton("Iniciar", func() {
		go models.Simulacion(actualizarVista)
	})

    // Crear un contenedor vertical (VBox) para organizar el contenedor de estacionamiento y la imagen del auto
    vboxContainer := container.NewVBox(parkingContainer, iniciarButton)

    // Agregar el contenedor al contenido de la ventana
    w.SetContent(vboxContainer)

	w.ShowAndRun()
}