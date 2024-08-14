package pages

import (
	"github.com/ying32/govcl/vcl"
)

type TTodayPage struct {
	*vcl.TTabSheet
}

func NewTodayPage(owner vcl.IComponent) *TTodayPage {
	c := new(TTodayPage)

	c.TTabSheet = vcl.NewTabSheet(owner)
	c.SetVisible(false)
	c.SetCaption("Today")
	c.SetPageIndex(1)
	c.SetTabVisible(false)

	lbl := vcl.NewLabel(c)
	lbl.SetParent(c)
	lbl.SetCaption("Today")
	lbl.SetBounds(0, 0, 200, 25)
	return c
}
