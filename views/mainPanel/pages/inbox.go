package pages

import (
	"github.com/ying32/govcl/vcl"
)

type TInboxPage struct {
	*vcl.TTabSheet
}

func NewInboxPage(owner vcl.IComponent) *TInboxPage {
	c := new(TInboxPage)

	c.TTabSheet = vcl.NewTabSheet(owner)
	c.SetVisible(false)
	c.SetCaption("Inbox")
	c.SetPageIndex(0)
	c.SetTabVisible(false)

	lbl := vcl.NewLabel(c)
	lbl.SetParent(c)
	lbl.SetCaption("Inbox")
	lbl.SetBounds(0, 0, 200, 25)
	return c
}
