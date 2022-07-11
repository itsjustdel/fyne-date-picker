package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	// c := newCalendar(time.Now(), onSelectedDate)
	// c.showAtPos(w.Canvas(), fyne.NewPos(50, 0))
	NewCalendarAtPos(w.Canvas(), fyne.NewPos(50, 50), time.Now(), onDateSelected)

	w.SetContent(widget.NewLabel("Calendar"))
	w.Resize(fyne.NewSize(400, 400))
	w.ShowAndRun()
}

func onDateSelected(t time.Time) {
	fmt.Println(t)
}
