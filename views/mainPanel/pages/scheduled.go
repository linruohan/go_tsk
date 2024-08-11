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
	return c
}
