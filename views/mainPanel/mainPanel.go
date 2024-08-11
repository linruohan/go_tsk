package mainPanel

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"go_tsk/views/mainPanel/pages"
)

type TPnlMain struct {
	*vcl.TPanel
	pageCtl      *pages.TPageCtl
	isMouseDown  bool
	mouseDownPos types.TPoint
}

func NewPnlMain(owner vcl.IComponent) *TPnlMain {
	c := new(TPnlMain)
	c.TPanel = vcl.NewPanel(owner)
	c.SetWidth(750)
	c.SetAlign(types.AlClient)
	c.SetBevelOuter(types.BvNone)
	c.SetClientWidth(750)
	c.SetClientHeight(500)
	c.SetParentBackground(false)
	c.SetParentColor(false)
	c.SetTabOrder(0)
	c.SetColor(0x342C29)

	c.SetOnMouseMove(c.OnPnlMainMouseMove)
	c.SetOnMouseDown(c.OnPnlMainMouseDown)
	c.SetOnMouseUp(c.OnPnlMainMouseUp)

	c.pageCtl = pages.NewPageCtl(c)
	c.pageCtl.SetParent(c)
	c.pageCtl.SetActivePage(c.pageCtl.TodayTabSheet)
	c.pageCtl.TodayTabSheet.Show()

	return c

}

func (u *TPnlMain) OnPnlMainMouseDown(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	if button == types.MbLeft {
		u.isMouseDown = true
		u.mouseDownPos.X = x
		u.mouseDownPos.Y = y
	}
}
func (u *TPnlMain) OnPnlMainMouseMove(sender vcl.IObject, shift types.TShiftState, x, y int32) {
	if u.isMouseDown {
		parent := u.Parent()
		parent.SetLeft(parent.Left() + (x - u.mouseDownPos.X))
		parent.SetTop(parent.Top() + (y - u.mouseDownPos.Y))
	}
}
func (u *TPnlMain) OnPnlMainMouseUp(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	u.isMouseDown = false
}
