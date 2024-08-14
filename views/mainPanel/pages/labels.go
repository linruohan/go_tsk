package pages

import (
	"github.com/ying32/govcl/vcl"
)

type TLabelsPage struct {
	*vcl.TTabSheet
}

func NewLabelsPage(owner vcl.IComponent) *TLabelsPage {
	c := new(TLabelsPage)
	c.TTabSheet = vcl.NewTabSheet(owner)
	c.SetVisible(false)
	c.SetCaption("Labels")
	c.SetPageIndex(3)
	c.SetTabVisible(false)

	lbl := vcl.NewLabel(c)
	lbl.SetParent(c)
	lbl.SetCaption("Labels")
	lbl.SetBounds(0, 0, 200, 25)
	return c
}
