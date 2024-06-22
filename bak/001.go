package bak

import (
	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/rtl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/win"
)

type TMainForm struct {
	*vcl.TForm
	Btn1     *vcl.TButton
	BtnMax   *vcl.TImageButton
	BtnMin   *vcl.TImageButton
	BtnClose *vcl.TImageButton
	BtnSkin  *vcl.TImageButton
}

type TAboutForm struct {
	*vcl.TForm
	Btn1 *vcl.TButton
}

var (
	MainForm *TMainForm
)

func (m *TMainForm) OnFormCreate(sender vcl.IObject) {
	path := rtl.ExtractFilePath(vcl.Application.ExeName())
	m.SetCaption("go_tsk")
	font := m.Font()
	font.SetStyle(types.FsBold)
	font.SetName("LXGW WenKai")
	m.ScreenCenter()
	m.SetWidth(1250)
	m.SetHeight(800)
	m.SetBorderStyle(types.BsSizeable)
	m.SetBorderWidth(100)
	m.ScreenCenter() //居中显示
	m.EnabledMaximize(true)
	m.SetDragMode(types.DmAutomatic)

	m.BtnClose = vcl.NewImageButton(m)
	m.BtnClose.SetParent(m)
	m.BtnClose.SetImageCount(4)
	m.BtnClose.SetAutoSize(true)
	m.BtnClose.Picture().LoadFromFile(path + "data/btn_close.png")
	m.BtnClose.SetHint("关闭")
	m.BtnClose.SetOnClick(func(vcl.IObject) {
		m.Close()
	})

	m.BtnMin = vcl.NewImageButton(m)
	m.BtnMin.SetParent(m)
	m.BtnMin.SetImageCount(4)
	m.BtnMin.SetAutoSize(true)
	m.BtnMin.Picture().LoadFromFile(path + "data/btn_min.png")
	m.BtnMin.SetLeft(m.BtnClose.Width())
	m.BtnMin.SetHint("最小化")
	m.BtnMin.SetOnClick(func(vcl.IObject) {
		m.SetWindowState(types.WsMinimized)
	})

	m.BtnMax = vcl.NewImageButton(m)
	m.BtnMax.SetParent(m)
	m.BtnMax.SetLeft(m.BtnClose.Width() + m.BtnMin.Width())
	m.BtnMax.SetImageCount(4)
	m.BtnMax.SetAutoSize(true)
	m.BtnMax.Picture().LoadFromFile(path + "data/btn_max.png")
	m.BtnMax.SetHint("最大化")
	m.BtnMax.SetOnClick(func(vcl.IObject) {
		if m.WindowState() == types.WsMaximized {
			// 最大化还原
			m.SetWindowState(types.WsNormal)
			m.SetBorderStyle(types.BsNone)
		} else {
			// 最大化
			m.SetBorderStyle(types.BsNone)
			m.SetWindowState(types.WsMaximized)
			//barHeight := win.GetSystemMetrics(w32.SM_CYCAPTION) + win.GetSystemMetrics(w32.SM_CYFRAME) + win.GetSystemMetrics(92)
			barHeight := win.GetSystemMetrics(4) + win.GetSystemMetrics(33) + win.GetSystemMetrics(92) + 10
			m.SetHeight(m.ClientHeight() - barHeight)
			print(m.ClientHeight())
			m.SetBounds(m.Left(), barHeight, m.ClientWidth(), m.ClientHeight())
		}
	})

	m.BtnSkin = vcl.NewImageButton(m)
	m.BtnSkin.SetParent(m)
	m.BtnSkin.SetImageCount(3)
	m.BtnSkin.SetAutoSize(true)
	m.BtnSkin.Picture().LoadFromFile(path + "data/btn_skin.png")
	m.BtnSkin.SetLeft(m.BtnMax.Left() + m.BtnMax.Width())
	m.BtnSkin.SetHint("皮肤")
	m.BtnSkin.SetOnClick(func(vcl.IObject) {
		vcl.ShowMessage("修改皮肤")
	})
}
