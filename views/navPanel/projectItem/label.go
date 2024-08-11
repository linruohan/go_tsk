package projectItem

import (
	"fmt"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
)

type TUtilLblProject struct {
	*vcl.TPanel
	iconLabel *vcl.TLabel // icon
	nameLabel *vcl.TLabel // projectItem name
	dueLabel  *vcl.TLabel // due string
	icon      string
	name      string
	due       string
}

func (u *TUtilLblProject) ProjectInit(owner vcl.IComponent, icon, name, due string) {

	u.TPanel = vcl.NewPanel(owner)
	u.TPanel.SetBorderStyle(types.None)
	u.TPanel.SetBorderWidth(0)
	u.TPanel.Font().SetName("LXGW WenKai GB")
	u.TPanel.SetColor(0x342C28)
	u.TPanel.Font().SetSize(12)
	u.TPanel.Font().SetColor(colors.ClWhite)
	u.TPanel.SetBounds(0, 0, 200, 50)
	u.TPanel.SetAlign(types.AlClient)

	u.icon = icon
	u.name = name
	u.due = due

	u.iconLabel = vcl.NewLabel(u)
	u.iconLabel.SetParent(u)
	u.iconLabel.SetCaption(u.icon)
	u.iconLabel.SetBounds(0, 0, 20, 25)
	u.iconLabel.SetAlign(types.AlLeft)
	u.iconLabel.SetAlignment(types.TaCenter)
	u.iconLabel.SetLayout(types.TlCenter)

	u.nameLabel = vcl.NewLabel(u)
	u.nameLabel.SetParent(u)
	u.nameLabel.SetAlign(types.AlClient)
	u.nameLabel.SetCaption(u.name)
	u.nameLabel.SetBounds(20, 0, 130, 25)
	u.nameLabel.SetAlignment(types.TaCenter)
	u.nameLabel.SetLayout(types.TlCenter)

	u.dueLabel = vcl.NewLabel(u)
	u.dueLabel.SetParent(u)
	u.dueLabel.SetCaption(u.due)
	u.dueLabel.SetAlign(types.AlRight)
	u.dueLabel.SetBounds(150, 0, 50, 25)
}
func (u *TUtilLblProject) setCaption() {
	captionName := fmt.Sprintf("%-20s %s %s", u.icon, u.name, u.due)
	u.SetCaption(captionName)
}
func (u *TUtilLblProject) SetDue(due string) {
	u.dueLabel.SetCaption(due)
}
func (u *TUtilLblProject) SetProjectName(name string) {
	u.nameLabel.SetCaption(name)
}
func (u *TUtilLblProject) SetIcon(icon string) {
	u.iconLabel.SetCaption(icon)
}
