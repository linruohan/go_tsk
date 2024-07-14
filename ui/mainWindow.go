package ui

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"go_tsk/ui/views"
)

type TMainForm struct {
	*vcl.TForm
}

var (
	MainForm *TMainForm
)

func (m *TMainForm) OnFormCreate(sender vcl.IObject) {
	m.SetCaption("go_tsk")
	m.SetWidth(1250)
	m.SetHeight(800)
	m.SetColor(0x342C28)
	m.ScreenCenter() //居中显示
	p := vcl.NewPanel(m)
	p.SetParent(m)
	p.SetBorderStyle(types.AlClient)
	p.SetBorderWidth(2)

	//datePicker := dialogs.NewDatePicker(m)
	//datePicker.SetParent(m)
	//datePicker.SetControlStyle(types.FsSystemStayOnTop)
	//datePicker.SetBounds(0, 0, 350, 25)

	inbox := views.NewInboxView(m)
	inbox.SetParent(m)
	inbox.SetBounds(0, 25, 120, 50)
	today := views.NewTodayView(m)
	today.SetParent(m)
	today.SetBounds(120, 25, 120, 50)
	scheduled := views.NewScheduledView(m)
	scheduled.SetParent(m)
	scheduled.SetBounds(0, 75, 120, 50)
	label := views.NewLabelView(m)
	label.SetParent(m)
	label.SetBounds(120, 75, 120, 50)

}
