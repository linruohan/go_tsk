package calendar

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
	"go_tsk/model"
)

type TDay struct {
	*vcl.TPanel
	lunarLabel *vcl.TLabel
	solarLabel *vcl.TLabel
	flagLabel  *vcl.TLabel
}

func NewDay(owner vcl.IComponent) *TDay {
	m := new(TDay)
	m.TPanel = vcl.NewPanel(owner)
	m.TPanel.SetColor(model.BackColor)
	m.TPanel.Font().SetName(model.FontName)
	m.TPanel.Font().SetSize(model.FontSize)
	m.TPanel.SetBorderStyle(types.None)
	m.TPanel.SetBorderWidth(1)
	m.TPanel.SetOnClick(m.OnDayClick)
	m.TPanel.SetBounds(0, 0, 50, 50)

	//阳历
	m.solarLabel = vcl.NewLabel(m)
	m.solarLabel.SetAlign(types.AlClient)
	m.solarLabel.SetParent(m)
	m.solarLabel.SetBounds(5, 5, 35, 35)
	m.solarLabel.Font().SetColor(colors.ClWheat)
	m.solarLabel.Font().SetSize(15)
	m.solarLabel.SetAlignment(types.TaCenter)
	m.solarLabel.SetLayout(types.TlCenter)

	m.solarLabel.SetOnClick(m.OnDayClick)
	//农历
	m.lunarLabel = vcl.NewLabel(m)
	m.lunarLabel.SetParent(m)
	m.lunarLabel.SetBounds(10, 35, 25, 10)
	m.lunarLabel.Font().SetColor(colors.ClSkyblue)
	m.lunarLabel.Font().SetSize(10)
	m.lunarLabel.SetAlignment(types.TaCenter)
	m.lunarLabel.SetLayout(types.TlCenter)

	m.flagLabel = vcl.NewLabel(m)
	m.flagLabel.SetParent(m)
	m.flagLabel.SetBounds(35, 0, 15, 15)
	m.flagLabel.Font().SetSize(10)
	m.flagLabel.Font().SetColor(model.FontColor)

	return m
}
func (d *TDay) OnDayClick(sender vcl.IObject) {
	d.SetColor(colors.ClGreen)
}
