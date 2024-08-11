package category

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type TUtilCategory struct {
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

func (u *TUtilCategory) CategoryInit(owner vcl.IComponent, icon, name, count string, fontColor, activeColor types.TColor) {
	//c := new(TUtilCategory)
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
	u.TPanel.SetBounds(0, 0, 100, 50)
	//u.TPanel.SetOnClick(u.OnActive)

	u.iconLabel = vcl.NewLabel(u)
	u.iconLabel.SetParent(u)
	u.iconLabel.SetCaption(u.icon)
	//u.iconLabel.SetAlign(types.AlClient)
	u.iconLabel.SetAlignment(types.TaLeftJustify)
	u.iconLabel.SetLayout(types.TlTop)
	u.iconLabel.SetTransparent(true)
	//u.iconLabel.SetOnClick(u.OnActive)
	u.iconLabel.SetBounds(20, 0, 50, 25)

	u.nameLabel = vcl.NewLabel(u)
	u.nameLabel.SetParent(u)
	u.nameLabel.SetCaption(u.name)
	u.nameLabel.SetAlignment(types.TaLeftJustify)
	u.nameLabel.SetLayout(types.TlBottom)
	u.nameLabel.SetTransparent(true)
	u.nameLabel.Font().SetSize(15)
	//u.nameLabel.SetOnClick(u.OnActive)
	u.nameLabel.SetBounds(10, 20, 50, 25)

	u.countLabel = vcl.NewLabel(u)
	u.countLabel.SetParent(u)
	u.countLabel.SetCaption(u.count)
	u.countLabel.SetAlign(types.AlClient)
	u.countLabel.SetAlignment(types.TaRightJustify)
	u.countLabel.SetLayout(types.TlTop)
	u.countLabel.SetTransparent(true)
	//u.countLabel.SetOnClick(u.OnActive)
	u.countLabel.SetBounds(60, 0, 50, 25)
}

func (u *TUtilCategory) OnActive(sender vcl.IObject) {
	if u.isActive {
		u.TPanel.SetColor(0x342C28)
		u.isActive = false
		return
	}
	u.TPanel.SetColor(0x342C28 + 80)
	u.isActive = true
}
func (u *TUtilCategory) SetCounter(sender vcl.IObject, cnt string) {
	u.countLabel.SetCaption(cnt)
}
