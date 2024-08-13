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
	return c
}
