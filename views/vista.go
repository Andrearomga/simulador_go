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
	myWindow := myApp.NewWindow("Rectángulos en Fyne")

	rectangles := make([]fyne.CanvasObject, 20)

	for i := 0; i < 20; i++ {
		rect := canvas.NewRectangle(theme.BackgroundColor())
		rect.FillColor = theme.PrimaryColor()
		x := float32(i * 70) // Aumenta la separación entre los rectángulos
		rect.Move(fyne.NewPos(x, 0)) // Cambia la posición en y a 0 para que estén en una sola fila
		rect.Resize(fyne.NewSize(60, 80))
		rectangles[i] = rect
	}

	content := container.NewWithoutLayout(rectangles...)
	myWindow.SetContent(content)

	myWindow.ShowAndRun()
}