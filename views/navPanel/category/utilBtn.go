package category

import (
	"fmt"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"go_tsk/model"
)

type TUtilBtnCategory struct {
	*vcl.TSpeedButton
	fontColor   types.TColor
	activeColor types.TColor
	icon        string
	name        string
	count       string
	isActive    bool
}

func (u *TUtilBtnCategory) CategoryInit(owner vcl.IComponent, icon, name, count string, fontColor, activeColor types.TColor) {
	//c := new(TUtilBtnCategory)
	u.isActive = false
	u.icon = icon
	u.name = name
	u.count = count
	u.fontColor = fontColor
	u.activeColor = activeColor

	u.TSpeedButton = vcl.NewSpeedButton(owner)
	//u.SetBackColor(u.fontColor)
	//u.SetHoverColor(u.activeColor)
	u.Font().SetName(model.FontName)
	u.Font().SetSize(model.FontSize)
	u.Font().SetColor(u.fontColor)
	u.SetTransparent(true)
	u.SetFlat(true)

	//u.Font().SetColor(colors.ClBlack)
	u.SetBounds(0, 0, model.LeftWidth/2, 50)
	u.SetCap()
	u.SetOnMouseEnter(func(sender vcl.IObject) {
		//u.Parent().SetColor(colors.ClRed)
	})
	u.SetOnMouseLeave(func(sender vcl.IObject) {
		//u.Parent().SetColor(colors.ClGreen)
	})
}
func (u *TUtilBtnCategory) SetCap() {
	c := fmt.Sprintf("%-10s%s\n%s", u.icon, u.count, u.name)
	u.SetCaption(c)
}
func (u *TUtilBtnCategory) SetIcon(icon string) {
	u.icon = icon
}
func (u *TUtilBtnCategory) SetName(name string) {
	u.name = name
}
func (u *TUtilBtnCategory) SetCount(count string) {
	u.count = count
}
