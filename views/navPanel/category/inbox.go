package category

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types/colors"
)

type TInboxCate struct {
	TUtilBtnCategory
}

func NewInboxCate(owner vcl.IComponent) *TInboxCate {
	c := new(TInboxCate)
	c.icon = "📦"
	c.name = "Inbox"
	c.count = "20"
	c.fontColor = colors.ClLegacySkyBlue
	c.activeColor = colors.ClLegacySkyBlue
	c.CategoryInit(owner, c.icon, c.name, c.count, c.fontColor, c.activeColor)
	return c
}
