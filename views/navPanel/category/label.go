package category

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types/colors"
)

type TLabelsCate struct {
	TUtilBtnCategory
}

func NewLabelsCate(owner vcl.IComponent) *TLabelsCate {
	c := new(TLabelsCate)
	c.icon = "🏷"
	c.name = "Label"
	c.count = "10"
	c.fontColor = colors.ClGray
	c.activeColor = colors.ClGray

	c.CategoryInit(owner, c.icon, c.name, c.count, c.fontColor, c.activeColor)
	return c
}
