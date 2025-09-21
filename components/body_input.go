package components

import (
	"fyne.io/fyne/v2/widget"
)

// type BodyEntry struct {
// 	Container     *fyne.Container
// 	rowsContainer *fyne.Container
// 	rows          []struct {
// 		Key   *widget.Entry
// 		Value *widget.Entry
// 	}
// }

func NewRequestBody() *widget.Entry {

	bodyEntry := widget.NewMultiLineEntry()
	bodyEntry.SetPlaceHolder("Enter json input here...")
	// return bodyEntry

	// b := &BodyEntry{
	// 	rowsContainer: container.NewVBox(),
	// }

	// content := container.NewStack()
	// radio := widget.NewRadioGroup([]string{"raw", "form-data"}, func(selected string) {
	// 	switch selected {
	// 	case "raw":
	// 		bodyEntry := widget.NewMultiLineEntry()
	// 		bodyEntry.SetPlaceHolder("Enter json input here...")
	// 		content.Objects = []fyne.CanvasObject{bodyEntry}
	// 	case "form-data":
	// 		content.Objects = []fyne.CanvasObject{widget.NewLabel("Form Data Tab")}
	// 	}
	// 	content.Refresh()
	// })
	// radio.Horizontal = true
	// radio.SetSelected("raw")

	// b.Container = container.NewVBox(radio, content)

	// return b
	return bodyEntry
}
