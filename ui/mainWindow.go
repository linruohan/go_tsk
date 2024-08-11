package ui

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
	"go_tsk/views/mainPanel"
	"go_tsk/views/navPanel"
)

type TMainForm struct {
	*vcl.TForm
	pnlNav  *navPanel.TPnlNav
	pnlMain *mainPanel.TPnlMain
}

var (
	MainForm *TMainForm
)

func (m *TMainForm) OnFormCreate(sender vcl.IObject) {
	m.SetLeft(0)
	m.SetTop(0)
	m.HorzScrollBar().SetVisible(false)
	m.VertScrollBar().SetVisible(false)
	m.SetAlphaBlend(true)
	m.SetAlphaBlendValue(250)
	m.SetBorderStyle(types.BsToolWindow)
	m.SetCaption("go_tsk")
	m.SetClientWidth(1000)
	m.SetClientHeight(580)
	m.SetDefaultMonitor(types.DmMainForm)
	m.Font().SetColor(colors.ClWhite)
	//m.Font().SetHeight(11)
	m.Font().SetName("LXGW WenKai GB")
	m.SetPosition(types.PoScreenCenter)

	m.pnlNav = navPanel.NewPnlNav(m)
	m.pnlNav.SetParent(m)

	m.pnlMain = mainPanel.NewPnlMain(m)
	m.pnlMain.SetParent(m)
}
