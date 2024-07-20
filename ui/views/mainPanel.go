package views

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type TPnlMain struct {
	*vcl.TPanel
}

func NewPnlMain(owner vcl.IComponent) *TPnlMain {
	c := new(TPnlMain)
	c.TPanel = vcl.NewPanel(owner)
	c.SetWidth(750)
	c.SetAlign(types.AlClient)
	c.SetBevelOuter(types.BvNone)
	c.SetClientWidth(750)
	c.SetParentBackground(false)
	c.SetParentColor(false)
	c.SetTabOrder(0)
	c.SetColor(0x342C28)
	return c

}

func (u *TPnlMain) OnActive(sender vcl.IObject) {

}
func (u *TPnlMain) SetCounter(sender vcl.IObject) {
}
