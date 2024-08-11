package navPanel

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
	"go_tsk/model"
)

type TPnlNav struct {
	*vcl.TPanel
	pnlCategory *TPnlCategory //类别

	pnlProjects *TPnlProject //项目
}

func NewPnlNav(owner vcl.IComponent) *TPnlNav {
	c := new(TPnlNav)
	c.TPanel = vcl.NewPanel(owner)

	c.SetWidth(model.LeftWidth)
	c.SetAlign(types.AlLeft)
	c.SetBevelOuter(types.BvNone)
	c.SetClientWidth(model.LeftWidth)
	c.SetParentBackground(false)
	c.SetParentColor(false)
	c.SetTabOrder(1)
	c.SetColor(0x342929)
	c.pnlCategory = NewPnlCategory(c)
	c.pnlCategory.SetParent(c)

	c.pnlProjects = NewPnlProject(c)
	c.pnlProjects.SetParent(c)

	return c

}

func (u *TPnlNav) OnActive(sender vcl.IObject) {

}
func (u *TPnlNav) SetCounter(sender vcl.IObject) {
}
func (u *TPnlNav) on1(sender vcl.IObject) {
	u.SetColor(colors.ClGreen)
}
func (u *TPnlNav) on2(sender vcl.IObject) {
	u.SetColor(colors.ClRed)
}
