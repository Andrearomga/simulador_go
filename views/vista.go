package views

import (
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/layout"
    "fyne.io/fyne/v2"
    "simulador/models"
    "time"
)

func IniciarVentana() {
    myApp := app.New()
    myWindow := myApp.NewWindow("Estacionamiento")
    myWindow.SetFixedSize(true)

    // Cargar la imagen del estacionamiento
    estacionamientoImagen := canvas.NewImageFromFile("assets/estacion.png")
    estacionamientoImagen.FillMode = canvas.ImageFillOriginal

    // Cargar la imagen del carrito
    autoImagen := canvas.NewImageFromFile("assets/carro220.png")
    autoImagen.FillMode = canvas.ImageFillOriginal

    vista := container.NewHBox(layout.NewSpacer(), estacionamientoImagen, autoImagen, layout.NewSpacer())
    iniciarBoton := widget.NewButton("Iniciar", func() {
        // Iniciar la simulación de llegada continua de vehículos desde el paquete models
        go models.SimularEstacionamiento(100, 20)

        // Definir la posición inicial del carro
        posX := float32(0)
        posY := float32(0)

        // Crear una función para actualizar la posición del carro
        actualizarPosicion := func() {
            for {
                // Actualizar la posición del carro (simulación de movimiento)
                posX += 2
                if posX > float32(estacionamientoImagen.Size().Width) {
                    posX = 0
                }

                // Actualizar la posición de la imagen del carro en la vista
                autoImagen.Move(fyne.NewPos(posX, posY))
                myWindow.Canvas().Refresh(autoImagen)
                time.Sleep(100 * time.Millisecond) // Controla la velocidad de movimiento
            }
        }

        // Iniciar la función de actualización de posición en segundo plano
        go actualizarPosicion()
    })

    vistaConBoton := container.NewVBox(iniciarBoton, vista)

    myWindow.SetContent(vistaConBoton)
    myWindow.Resize(fyne.NewSize(500, estacionamientoImagen.Size().Height))
    myWindow.ShowAndRun()
}
