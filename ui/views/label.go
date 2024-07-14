package views

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types/colors"
)

type TLabelView struct {
	TUtilView
}

func NewLabelView(owner vcl.IComponent) *TLabelView {
	c := new(TLabelView)
	c.icon = "🏷"
	c.name = "Label"
	c.count = "10"
	c.fontColor = colors.ClYellow
	c.activeColor = colors.ClYellow

	c.NewView(owner, c.icon, c.name, c.count, c.fontColor, c.activeColor)
	return c
}
