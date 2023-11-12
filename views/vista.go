package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func MostrarEstacionamiento() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Rectángulos con Imagen en Fyne")

	// Crear rectángulos
	rectangles := make([]fyne.CanvasObject, 20)

	for i := 0; i < 20; i++ {
		rect := canvas.NewRectangle(theme.BackgroundColor())
		rect.FillColor = theme.PrimaryColor()
		x := float32(i * 70) // Aumenta la separación entre los rectángulos
		rect.Move(fyne.NewPos(x, 0)) // Cambia la posición en y a 0 para que estén en una sola fila
		rect.Resize(fyne.NewSize(60, 80))
		rectangles[i] = rect
	}

	// Crear imagen del carro
	carroImage := canvas.NewImageFromFile("assets/carro2.png")
	carroImage.FillMode = canvas.ImageFillOriginal
	carroImage.Resize(fyne.NewSize(60, 80)) // Ajusta el tamaño de la imagen del carro según tus necesidades

	// Centrar la imagen del carro en el eje Y
	windowHeight := myWindow.Canvas().Size().Height
	carroY := (windowHeight - carroImage.Size().Height) / 2
	carroImage.Move(fyne.NewPos(0, carroY))

	// Crear contenedor para los rectángulos y la imagen del carro
	content := container.NewVBox(container.NewWithoutLayout(rectangles...), carroImage)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
