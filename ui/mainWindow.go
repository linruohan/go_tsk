package ui

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
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
	m.SetWidth(950)
	m.SetHeight(580)
	m.ScreenCenter()
	m.HorzScrollBar().SetVisible(false)
	m.VertScrollBar().SetVisible(false)
	m.SetAlphaBlend(true)
	m.SetBorderStyle(types.BsToolWindow)
	m.SetClientWidth(950)
	m.SetClientHeight(580)

	m.SetColor(0x342C28)
	m.SetDefaultMonitor(types.DmMainForm)
	m.Font().SetColor(colors.ClWhite)

	//m.ScreenCenter() //居中显示
	pnlNav := views.NewPnlNav(m)
	pnlNav.SetParent(m)

	pnlMain := views.NewPnlMain(m)
	pnlMain.SetParent(m)

	//p.SetBorderStyle(types.AlClient)
	//p.SetBorderWidth(2)

	//datePicker := dialogs.NewDatePicker(m)
	//datePicker.SetParent(m)
	//datePicker.SetControlStyle(types.FsSystemStayOnTop)
	//datePicker.SetBounds(0, 0, 350, 25)

	//inbox := views.NewInboxView(m)
	//inbox.SetParent(m)
	//inbox.SetBounds(0, 25, 120, 50)
	//today := views.NewTodayView(m)
	//today.SetParent(m)
	//today.SetBounds(120, 25, 120, 50)
	//scheduled := views.NewScheduledView(m)
	//scheduled.SetParent(m)
	//scheduled.SetBounds(0, 75, 120, 50)
	//label := views.NewLabelView(m)
	//label.SetParent(m)
	//label.SetBounds(120, 75, 120, 50)
	//g := vcl.NewComboBox(m)
	//g.SetParent(m)
	//
	//t := vcl.NewToggleBox(m)
	//t.SetParent(m)
	//t.SetCaption("inbox")
	//t.SetOnClick(m.on1)
	//t.Hide()
	//g.AddItem("111", t)
	//
	//t1 := vcl.NewToggleBox(m)
	//t1.SetParent(m)
	//t1.SetCaption("today")
	//t1.SetOnClick(m.on2)
	//t1.Hide()
	//g.AddItem("222", t1)

}
func (m *TMainForm) on1(sender vcl.IObject) {
	m.SetColor(colors.ClGreen)
}
func (m *TMainForm) on2(sender vcl.IObject) {
	m.SetColor(colors.ClRed)
}
