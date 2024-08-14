package pages

import (
	"github.com/ying32/govcl/vcl"
)

type TScheduledPage struct {
	*vcl.TTabSheet
}

func NewScheduledPage(owner vcl.IComponent) *TScheduledPage {
	c := new(TScheduledPage)

	c.TTabSheet = vcl.NewTabSheet(owner)
	c.SetVisible(false)
	c.SetCaption("Scheduled")
	c.SetPageIndex(2)
	c.SetTabVisible(false)

	lbl := vcl.NewLabel(c)
	lbl.SetParent(c)
	lbl.SetCaption("Scheduled")
	lbl.SetBounds(0, 0, 200, 25)
	return c
}
