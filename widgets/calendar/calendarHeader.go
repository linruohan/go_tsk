package calendar

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
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
	m.TPanel.SetColor(0x342C28)
	m.TPanel.Font().SetName("LXGW WenKai GB")
	m.TPanel.Font().SetSize(12)
	m.TPanel.SetBorderStyle(types.None)
	m.TPanel.SetBorderWidth(0)
	m.TPanel.SetBounds(0, 0, 350, 25)

	m.monthLabel = vcl.NewLabel(m)
	m.monthLabel.SetParent(m)
	m.monthLabel.SetBounds(20, 0, 30, 25)
	m.monthLabel.Font().SetColor(colors.ClWhite)

	m.yearLabel = vcl.NewLabel(m)
	m.yearLabel.SetParent(m)
	m.yearLabel.SetBounds(100, 0, 50, 25)
	m.yearLabel.Font().SetColor(colors.ClWhite)

	m.leftBtn = vcl.NewXButton(m)
	m.leftBtn.SetParent(m)
	m.leftBtn.SetCaption("←")
	m.leftBtn.SetHoverColor(colors.ClSkyblue)
	m.leftBtn.SetBackColor(0x342C28)
	m.leftBtn.SetBounds(200, 0, 50, 25)
	m.leftBtn.SetNormalFontColor(colors.ClWhite)

	m.rightBtn = vcl.NewXButton(m)
	m.rightBtn.SetParent(m)
	m.rightBtn.SetCaption("→")
	m.rightBtn.SetNormalFontColor(colors.ClWhite)
	m.rightBtn.SetBackColor(0x342C28)
	m.rightBtn.SetHoverColor(colors.ClSkyblue)
	m.rightBtn.SetBounds(250, 0, 50, 25)

	m.centerBtn = vcl.NewXButton(m)
	m.centerBtn.SetParent(m)
	m.centerBtn.SetCaption("Today")
	m.centerBtn.SetHoverColor(colors.ClPurple)
	m.centerBtn.SetNormalFontColor(colors.ClWhite)
	m.centerBtn.SetBackColor(0x342C28)
	m.centerBtn.SetBounds(300, 0, 50, 25)

	return m
}
