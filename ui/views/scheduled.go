package views

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types/colors"
)

type TScheduledView struct {
	TUtilView
}

func NewScheduledView(owner vcl.IComponent) *TScheduledView {
	c := new(TScheduledView)
	c.icon = "📅"
	c.name = "Scheduled"
	c.count = "10"
	c.fontColor = colors.ClPurple + 80
	c.activeColor = colors.ClPurple + 80
	c.NewView(owner, c.icon, c.name, c.count, c.fontColor, c.activeColor)
	return c
}
