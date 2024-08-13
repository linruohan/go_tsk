package navPanel

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"go_tsk/model"
)

type TPnlNav struct {
	*vcl.TPanel
	PnlCategory *TPnlCategory //类别

	PnlProjects *TPnlProject //项目
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
	c.PnlCategory = NewPnlCategory(c)
	c.PnlCategory.SetParent(c)

	c.PnlProjects = NewPnlProject(c)
	c.PnlProjects.SetParent(c)

	return c

}

func (u *TPnlNav) OnActive(sender vcl.IObject) {

}
func (u *TPnlNav) SetCounter(sender vcl.IObject) {
}

func (u *TPnlNav) OnCateGoryClick(sender vcl.IObject) {

}
