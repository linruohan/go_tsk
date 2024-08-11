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
	return c
}
