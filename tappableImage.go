package fynitude

import (
	"image"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

/**
 * An image widget that can be clicked to trigger a custom function
 */
type TappableImageWidget struct {
	widget.BaseWidget

	Canvas   *canvas.Image
	Tapp func(pointEvent *fyne.PointEvent)
}

/**
 * Create the tappable image widget
 * @param image is the initial image
 * @param tapped is the custom event triggered
 * @return a configured TappableImageWidget
 */
func NewTappableImageWidget(image image.Image, tapped func(pointEvent *fyne.PointEvent)) *TappableImageWidget {	
	imageCanvas := canvas.NewImageFromImage(image)

	widget := &TappableImageWidget{
		Canvas: imageCanvas,
		Tapp: tapped,
	}

	widget.ExtendBaseWidget(widget)

	return widget
}

func (w *TappableImageWidget) Tapped(pointEvent *fyne.PointEvent) {
	if w.Tapp != nil {
		w.Tapp(pointEvent)
	}
}

func (w *TappableImageWidget) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(w.Canvas)
}
