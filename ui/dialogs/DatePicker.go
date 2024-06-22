package dialogs

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"go_tsk/widgets/calendar"
)

type TDatePicker struct {
	*vcl.TPanel
	c *calendar.TCalendar
	e *vcl.TEdit
}

func NewDatePicker(owner vcl.IComponent) *TDatePicker {
	datePicker := new(TDatePicker)
	datePicker.TPanel = vcl.NewPanel(owner)
	datePicker.TPanel.SetBorderStyle(types.None)
	datePicker.TPanel.SetBorderWidth(0)
	datePicker.TPanel.SetBorderStyle(types.AlClient)
	datePicker.TPanel.SetBounds(0, 0, 1250, 800)

	datePicker.e = vcl.NewEdit(datePicker)
	datePicker.e.SetParent(datePicker)
	datePicker.e.SetOnClick(datePicker.onEditClick)
	datePicker.e.SetClientWidth(350)

	datePicker.c = calendar.NewCalendar(datePicker)
	datePicker.c.SetParent(datePicker)
	datePicker.c.SetBounds(0, datePicker.e.ClientHeight()+10, 350, 350)
	datePicker.c.Hide()
	return datePicker
}

func (datePicker *TDatePicker) onEditClick(sender vcl.IObject) {
	datePicker.c.Show()
	go func() {
		for {
			if !datePicker.c.Visible() {
				datePicker.e.SetText(datePicker.c.GetDay())
				break
			}
		}
	}()
}
