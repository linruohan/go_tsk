package ui

import (
	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"go_tsk/ui/dialogs"
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
	m.ScreenCenter() //居中显示
	p := vcl.NewPanel(m)
	p.SetParent(m)
	p.SetBorderStyle(types.AlClient)
	p.SetBounds(0, 0, 1250, 800)
	datePicker := dialogs.NewDatePicker(m)
	datePicker.SetParent(m)

}
