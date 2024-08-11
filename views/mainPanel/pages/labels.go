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
	return c
}
