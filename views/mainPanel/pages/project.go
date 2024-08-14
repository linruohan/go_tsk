package pages

import (
	"github.com/ying32/govcl/vcl"
)

type TProjectPage struct {
	*vcl.TTabSheet
}

func NewProjectPage(owner vcl.IComponent) *TProjectPage {
	c := new(TProjectPage)

	c.TTabSheet = vcl.NewTabSheet(owner)
	c.SetVisible(false)
	c.SetCaption("Project")
	c.SetPageIndex(4)
	c.SetTabVisible(false)

	lbl := vcl.NewLabel(c)
	lbl.SetParent(c)
	lbl.SetCaption("Project")
	lbl.SetBounds(0, 0, 200, 25)
	return c
}
