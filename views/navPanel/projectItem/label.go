package projectItem

import (
	"fmt"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"go_tsk/model"
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

func (lbl *TUtilLblProject) ProjectInit(owner vcl.IComponent, icon, name, due string) {

	lbl.TPanel = vcl.NewPanel(owner)
	lbl.TPanel.SetBorderStyle(types.None)
	lbl.TPanel.SetBorderWidth(0)
	lbl.TPanel.Font().SetName(model.FontName)
	lbl.TPanel.SetColor(model.BackColor)
	lbl.TPanel.Font().SetSize(model.FontSize)
	lbl.TPanel.Font().SetColor(model.FontColor)
	lbl.TPanel.SetBounds(0, 0, model.LeftWidth, 50)
	lbl.TPanel.SetAlign(types.AlClient)

	lbl.icon = icon
	lbl.name = name
	lbl.due = due

	lbl.iconLabel = vcl.NewLabel(lbl)
	lbl.iconLabel.SetParent(lbl)
	lbl.iconLabel.SetCaption(lbl.icon)
	lbl.iconLabel.SetBounds(0, 0, 20, 25)
	lbl.iconLabel.SetAlign(types.AlLeft)
	lbl.iconLabel.SetAlignment(types.TaCenter)
	lbl.iconLabel.SetLayout(types.TlCenter)

	lbl.nameLabel = vcl.NewLabel(lbl)
	lbl.nameLabel.SetParent(lbl)
	lbl.nameLabel.SetAlign(types.AlClient)
	lbl.nameLabel.SetCaption(lbl.name)
	lbl.nameLabel.SetBounds(20, 0, model.LeftWidth-70, 25)
	lbl.nameLabel.SetAlignment(types.TaCenter)
	lbl.nameLabel.SetLayout(types.TlCenter)

	lbl.dueLabel = vcl.NewLabel(lbl)
	lbl.dueLabel.SetParent(lbl)
	lbl.dueLabel.SetCaption(lbl.due)
	lbl.dueLabel.SetAlign(types.AlRight)
	lbl.dueLabel.SetBounds(model.LeftWidth-50, 0, 50, 25)
}
func (lbl *TUtilLblProject) setCaption() {
	captionName := fmt.Sprintf("%-10s %s %s", lbl.icon, lbl.name, lbl.due)
	lbl.SetCaption(captionName)
}
func (lbl *TUtilLblProject) SetDue(due string) {
	lbl.dueLabel.SetCaption(due)
}
func (lbl *TUtilLblProject) SetProjectName(name string) {
	lbl.nameLabel.SetCaption(name)
}
func (lbl *TUtilLblProject) SetIcon(icon string) {
	lbl.iconLabel.SetCaption(icon)
}
