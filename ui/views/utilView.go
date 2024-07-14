package views

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type TUtilView struct {
	*vcl.TPanel
	iconLabel   *vcl.TLabel
	nameLabel   *vcl.TLabel
	countLabel  *vcl.TLabel
	fontColor   types.TColor
	activeColor types.TColor
	icon        string
	name        string
	count       string
	isActive    bool
}

func (u *TUtilView) NewView(owner vcl.IComponent, icon, name, count string, fontColor, activeColor types.TColor) {
	//c := new(TUtilView)
	u.isActive = false
	u.icon = icon
	u.name = name
	u.count = count
	u.fontColor = fontColor
	u.activeColor = activeColor

	u.TPanel = vcl.NewPanel(owner)
	//u.TPanel.SetBorderStyle(types.BsSingle)
	//u.TPanel.SetBorderWidth(1)
	u.TPanel.SetBorderStyle(types.None)
	u.TPanel.SetBorderWidth(0)
	u.TPanel.Font().SetName("LXGW WenKai GB")
	u.TPanel.SetColor(0x342C28)
	u.TPanel.Font().SetSize(12)
	u.TPanel.Font().SetColor(u.fontColor)
	u.TPanel.SetBounds(0, 0, 120, 50)
	u.TPanel.SetOnClick(u.OnActive)

	u.iconLabel = vcl.NewLabel(u)
	u.iconLabel.SetParent(u)
	u.iconLabel.SetCaption(u.icon)
	//u.iconLabel.SetAlign(types.AlClient)
	u.iconLabel.SetAlignment(types.TaLeftJustify)
	u.iconLabel.SetLayout(types.TlTop)
	u.iconLabel.SetTransparent(true)
	u.iconLabel.SetOnClick(u.OnActive)
	u.iconLabel.SetBounds(20, 0, 60, 25)

	u.nameLabel = vcl.NewLabel(u)
	u.nameLabel.SetParent(u)
	u.nameLabel.SetCaption(u.name)
	u.nameLabel.SetAlignment(types.TaLeftJustify)
	u.nameLabel.SetLayout(types.TlBottom)
	u.nameLabel.SetTransparent(true)
	u.nameLabel.Font().SetSize(15)
	u.nameLabel.SetOnClick(u.OnActive)
	u.nameLabel.SetBounds(10, 20, 60, 25)

	u.countLabel = vcl.NewLabel(u)
	u.countLabel.SetParent(u)
	u.countLabel.SetCaption(u.count)
	u.countLabel.SetAlign(types.AlClient)
	u.countLabel.SetAlignment(types.TaRightJustify)
	u.countLabel.SetLayout(types.TlTop)
	u.countLabel.SetTransparent(true)
	u.countLabel.SetOnClick(u.OnActive)
	u.countLabel.SetBounds(60, 0, 60, 25)
}

func (u *TUtilView) OnActive(sender vcl.IObject) {
	if u.isActive {
		u.TPanel.SetColor(0x342C28)
		u.isActive = false
		return
	}
	u.TPanel.SetColor(0x342C28 + 80)
	u.isActive = true
	u.SetCounter(u)
}
func (u *TUtilView) SetCounter(sender vcl.IObject) {
	u.countLabel.SetCaption("112")
}
