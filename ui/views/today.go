package views

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types/colors"
)

type TTodayView struct {
	TUtilView
}

func NewTodayView(owner vcl.IComponent) *TTodayView {
	c := new(TTodayView)
	c.icon = "🐮"
	c.name = "Today"
	c.count = "10"
	c.fontColor = colors.ClGreen
	c.activeColor = colors.ClGreen
	c.NewView(owner, c.icon, c.name, c.count, c.fontColor, c.activeColor)
	return c
}
