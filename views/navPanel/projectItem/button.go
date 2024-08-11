package projectItem

import (
	"fmt"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
)

type TUtilBtnProject struct {
	*vcl.TSpeedButton
	icon string
	name string
	due  string
}

func (u *TUtilBtnProject) ProjectInit(owner vcl.IComponent, icon, name, due string) {
	u.TSpeedButton = vcl.NewSpeedButton(owner)
	u.TSpeedButton.Font().SetName("LXGW WenKai GB")
	u.TSpeedButton.SetColor(0x342C28)
	u.TSpeedButton.Font().SetSize(12)
	u.TSpeedButton.Font().SetColor(colors.ClWhite)
	u.TSpeedButton.SetBounds(0, 0, 200, 25)
	u.TSpeedButton.SetAlign(types.AlClient)
	u.SetAlign(types.AlTop)
	u.SetFlat(true)
	u.SetSpacing(-1)
	u.SetTransparent(true)

	u.icon = icon
	u.name = name
	u.due = due
	captionName := fmt.Sprintf("%-20s %s %s", u.icon, u.name, u.due)
	u.TSpeedButton.SetCaption(captionName)

}

func (u *TUtilBtnProject) setCaption() {
	captionName := fmt.Sprintf("%-20s %s %s", u.icon, u.name, u.due)
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
