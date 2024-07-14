package dateTimePicker

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"time"
)

type TTimePicker struct {
	*vcl.TPanel
	time time.Time
}

func NewTimePicker(owner vcl.IComponent) *TTimePicker {
	c := new(TTimePicker)
	c.TPanel = vcl.NewPanel(owner)
	c.TPanel.SetBorderStyle(types.None)
	c.TPanel.SetBorderWidth(0)

	return c
}
