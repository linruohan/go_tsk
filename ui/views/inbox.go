package views

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types/colors"
)

type TInboxView struct {
	TUtilView
}

func NewInboxView(owner vcl.IComponent) *TInboxView {
	c := new(TInboxView)
	c.icon = "📦"
	c.name = "Inbox"
	c.count = "20"
	c.fontColor = colors.ClLegacySkyBlue
	c.activeColor = colors.ClLegacySkyBlue
	c.NewView(owner, c.icon, c.name, c.count, c.fontColor, c.activeColor)
	return c
}
