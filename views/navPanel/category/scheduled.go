package category

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types/colors"
)

type TScheduledCate struct {
	TUtilCategory
}

func NewScheduledCate(owner vcl.IComponent) *TScheduledCate {
	c := new(TScheduledCate)
	c.icon = "📅"
	c.name = "Scheduled"
	c.count = "10"
	c.fontColor = colors.ClPurple + 80
	c.activeColor = colors.ClPurple + 80
	c.CategoryInit(owner, c.icon, c.name, c.count, c.fontColor, c.activeColor)
	return c
}
