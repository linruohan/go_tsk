package views

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type TPnlNav struct {
	*vcl.TPanel
	pnlUser       *vcl.TPanel
	pnlNavButtons *vcl.TPanel
}

func NewPnlNav(owner vcl.IComponent) *TPnlNav {
	c := new(TPnlNav)
	c.TPanel = vcl.NewPanel(owner)
	c.SetWidth(200)
	c.SetAlign(types.AlLeft)
	c.SetBevelOuter(types.BvNone)
	c.SetClientWidth(200)
	c.SetParentBackground(false)
	c.SetParentColor(false)
	c.SetTabOrder(1)
	c.SetColor(0x342C29)

	return c

}

func (u *TPnlNav) OnActive(sender vcl.IObject) {

}
func (u *TPnlNav) SetCounter(sender vcl.IObject) {
}
