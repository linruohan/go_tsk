package calendar

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"go_tsk/model"
)

var weekendZH = []string{"一", "二", "三", "四", "五", "六", "日"}
var weekendEN = []string{"Mo", "Tu", "We", "Th", "Fr", "Sa", "Su"}

type TWeek struct {
	*vcl.TPanel
	labels []*vcl.TLabel
}

func NewWeek(owner vcl.IComponent) *TWeek {
	m := new(TWeek)
	m.TPanel = vcl.NewPanel(owner)
	m.TPanel.SetColor(model.BackColor)
	m.TPanel.Font().SetName(model.FontName)
	m.TPanel.Font().SetSize(model.FontSize)
	m.TPanel.SetBorderStyle(types.None)
	m.TPanel.SetBorderWidth(0)
	m.TPanel.SetBounds(0, 0, 350, 25)

	for i := int32(0); i < 7; i++ {
		label := vcl.NewLabel(owner)
		label.SetParent(m)
		label.SetCaption(weekendZH[i])
		label.SetBounds(i*50+15, 0, 50, 25)
		label.SetAlignment(types.TaCenter)
		label.SetLayout(types.TlCenter)
		label.Font().SetColor(model.FontColor)
		m.labels = append(m.labels, label)
	}
	return m
}
