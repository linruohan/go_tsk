package mainPanel

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"go_tsk/views/mainPanel/pages"
	"runtime"
)

type TPnlMain struct {
	*vcl.TPanel
	PageCtl      *pages.TPageCtl
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

	c.PageCtl = pages.NewPageCtl(c)
	c.PageCtl.SetParent(c)
	c.bindDragFunc()
	return c

}

func (u *TPnlMain) OnFormConstrainedResize(sender vcl.IObject, minWidth, minHeight, maxWidth, maxHeight *int32) {
	if runtime.GOOS == "windows" {
		//if u.Parent().Parent().WindowState() == types.WsMaximized {
		//	*maxHeight = vcl.Screen.WorkAreaHeight()
		//}
	}
	*minWidth = 400
	*minHeight = 400
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
func (u *TPnlMain) bindDragFunc() {
	var i int32
	for i = 0; i < u.PageCtl.PageCount(); i++ {
		sheet := u.PageCtl.Pages(i)
		sheet.SetOnMouseMove(u.OnPnlMainMouseMove)
		sheet.SetOnMouseDown(u.OnPnlMainMouseDown)
		sheet.SetOnMouseUp(u.OnPnlMainMouseUp)
	}
}
