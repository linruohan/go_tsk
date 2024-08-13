package pages

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types/colors"
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

	lbl := vcl.NewLabel(c)
	lbl.SetParent(c)
	lbl.SetCaption("Today")
	lbl.Font().SetColor(colors.ClRed)
	lbl.Font().SetSize(18)
	return c
}
