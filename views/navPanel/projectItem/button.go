package projectItem

import (
	"fmt"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
)

type TUtilBtnProject struct {
	*vcl.TXButton
	icon string
	name string
	due  string
}

func (u *TUtilBtnProject) ProjectInit(owner vcl.IComponent, icon, name, due string) {
	u.TXButton = vcl.NewXButton(owner)
	u.TXButton.Font().SetName("LXGW WenKai GB")
	u.TXButton.Font().SetSize(12)
	u.TXButton.Font().SetColor(colors.ClWhite)
	u.TXButton.SetBounds(-10, 0, 200, 25)
	u.TXButton.SetAlign(types.AlClient)
	u.SetAlign(types.AlTop)
	u.SetBackColor(colors.ClLegacySkyBlue)
	u.SetHoverColor(colors.ClMoneyGreen)

	u.icon = icon
	u.name = name
	u.due = due
	u.setCaption()

}

func (u *TUtilBtnProject) setCaption() {
	//d := model.LeftWidth - 15
	captionName := fmt.Sprintf(`%-5s %s %10s`, u.icon, u.name, u.due)
	u.SetCaption(captionName)
}
func (u *TUtilBtnProject) SetDue(due string) {
	u.due = due
	u.setCaption()
}
func (u *TUtilBtnProject) SetProjectName(name string) {
	u.name = name
	u.setCaption()
}
func (u *TUtilBtnProject) SetIcon(icon string) {
	u.icon = icon
	u.setCaption()
}
