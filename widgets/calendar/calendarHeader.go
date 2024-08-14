package calendar

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
	"go_tsk/model"
)

type THeader struct {
	*vcl.TPanel
	monthLabel *vcl.TLabel
	yearLabel  *vcl.TLabel
	leftBtn    *vcl.TXButton
	rightBtn   *vcl.TXButton
	centerBtn  *vcl.TXButton
}

func NewHeader(owner vcl.IComponent) *THeader {
	m := new(THeader)
	m.TPanel = vcl.NewPanel(owner)
	m.TPanel.SetColor(model.BackColor)
	m.TPanel.Font().SetName(model.FontName)
	m.TPanel.Font().SetSize(model.FontSize)
	m.TPanel.SetBorderStyle(types.None)
	m.TPanel.SetBorderWidth(0)
	m.TPanel.SetBounds(0, 0, 350, 25)

	m.monthLabel = vcl.NewLabel(m)
	m.monthLabel.SetParent(m)
	m.monthLabel.SetBounds(20, 0, 30, 25)
	m.monthLabel.Font().SetColor(model.FontColor)

	m.yearLabel = vcl.NewLabel(m)
	m.yearLabel.SetParent(m)
	m.yearLabel.SetBounds(100, 0, 50, 25)
	m.yearLabel.Font().SetColor(model.FontColor)

	m.leftBtn = vcl.NewXButton(m)
	m.leftBtn.SetParent(m)
	m.leftBtn.SetCaption("←")
	m.leftBtn.SetHoverColor(colors.ClSkyblue)
	m.leftBtn.SetBackColor(model.BackColor)
	m.leftBtn.SetBounds(200, 0, 50, 25)
	m.leftBtn.SetNormalFontColor(model.FontColor)

	m.rightBtn = vcl.NewXButton(m)
	m.rightBtn.SetParent(m)
	m.rightBtn.SetCaption("→")
	m.rightBtn.SetNormalFontColor(model.FontColor)
	m.rightBtn.SetBackColor(model.BackColor)
	m.rightBtn.SetHoverColor(colors.ClSkyblue)
	m.rightBtn.SetBounds(250, 0, 50, 25)

	m.centerBtn = vcl.NewXButton(m)
	m.centerBtn.SetParent(m)
	m.centerBtn.SetCaption("Today")
	m.centerBtn.SetHoverColor(colors.ClPurple)
	m.centerBtn.SetNormalFontColor(model.FontColor)
	m.centerBtn.SetBackColor(model.BackColor)
	m.centerBtn.SetBounds(300, 0, 50, 25)

	return m
}
