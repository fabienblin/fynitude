package fynitude

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

/**
 * A Select that triggers the options function before opening
 * This allows you to implement the options function
 */
type DynamicSelectWidget struct {
	widget.Select

	Options func() []string
}

/**
 * Create the dynamic select
 * @param options is the function that refreshes the list of options
 * @param changed is the triggered function when an option is selected
 * @return a configured DynamicSelectWidget
 */
func NewDynamicSelectWidget(options func() []string, changed func(string)) *DynamicSelectWidget {
	w := &DynamicSelectWidget{Options: options}
	w.Select.OnChanged = changed

	w.ExtendBaseWidget(w)

	return w
}

func (w *DynamicSelectWidget) Tapped(point *fyne.PointEvent) {
	if w.Options != nil {
		w.Select.Options = w.Options()
	}
	w.Select.Tapped(point)
}
