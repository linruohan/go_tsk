package category

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types/colors"
)

type TTodayCate struct {
	TUtilBtnCategory
}

func NewTodayCate(owner vcl.IComponent) *TTodayCate {
	c := new(TTodayCate)
	c.icon = "🐮"
	c.name = "Today"
	c.count = "10"
	c.fontColor = colors.ClGreen
	c.activeColor = colors.ClGreen
	c.CategoryInit(owner, c.icon, c.name, c.count, c.fontColor, c.activeColor)
	return c
}
