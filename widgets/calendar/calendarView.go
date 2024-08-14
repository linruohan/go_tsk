package calendar

import (
	"fmt"
	"github.com/6tail/lunar-go/HolidayUtil"
	"github.com/6tail/lunar-go/calendar"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
	"go_tsk/model"
	"go_tsk/utils"
	"strconv"
	"time"
)

type TView struct {
	*vcl.TPanel
	days        []*TDay
	currentDate time.Time
	day         int
}

func NewView(owner vcl.IComponent) *TView {
	m := new(TView)
	m.TPanel = vcl.NewPanel(owner)
	m.TPanel.SetColor(model.BackColor)
	m.TPanel.Font().SetName(model.FontName)
	m.TPanel.Font().SetSize(model.FontSize)
	m.TPanel.SetBorderStyle(types.None)
	m.TPanel.SetBorderWidth(2)

	m.TPanel.SetBounds(0, 0, 50, 50)
	tmpDay := time.Now()
	m.currentDate = utils.GenerateDate(tmpDay, tmpDay.Day())
	var row int32 = 0
	var col int32 = 0
	for i := 0; i < 42; i++ {
		day := NewDay(owner)
		day.SetParent(m)
		day.SetBounds(col*50, row*50, 50, 50)
		day.solarLabel.SetOnClick(m.OnViewDayClick)

		col = col + 1
		if col != 0 && col%7 == 0 {
			row = row + 1
			col = 0
		}
		m.days = append(m.days, day)
	}
	//每个月的第一天
	firstOfMonth := utils.GenerateDate(m.currentDate, 1)
	// 每个月的最后一天
	endDay := firstOfMonth.AddDate(0, 1, -1).Day()
	// 每个月第一天的星期几
	startPos := int(firstOfMonth.Weekday()) - 1
	m.fillGridDays(startPos, endDay, m.currentDate)
	return m
}
func (m *TView) fillGridDays(startPos, endDay int, day time.Time) {
	var dayNum = 1
	for i := 0; i < 42; i++ {
		var item = m.days[i]
		item.solarLabel.SetCaption("")
		item.lunarLabel.SetCaption("")
		item.flagLabel.SetCaption("")
		if i < startPos || i >= endDay+startPos {
			continue
		}
		currentDate := utils.GenerateDate(day, dayNum)
		lunarDay := calendar.NewLunarFromDate(currentDate)
		holiday := HolidayUtil.GetHolidayByYmd(currentDate.Year(), int(currentDate.Month()), dayNum)

		item.solarLabel.SetCaption(fmt.Sprintf("%d", dayNum))
		item.lunarLabel.SetCaption(lunarDay.GetDayInChinese())
		//today
		if currentDate.Equal(m.currentDate) {
			item.SetColor(model.BackColor + 20)
		}
		//周末
		if currentDate.Weekday() == 0 || currentDate.Weekday() == 6 {
			item.flagLabel.SetCaption("休")
			item.flagLabel.SetColor(colors.ClGreen)
		}
		//假日
		if holiday != nil && holiday.GetName() != "" {
			if holiday.IsWork() {
				fmt.Printf("%v\n", holiday.GetName())
				item.flagLabel.SetCaption("班")
				item.flagLabel.SetColor(colors.ClRed)
			} else {
				item.flagLabel.SetCaption("休")
				item.flagLabel.SetColor(colors.ClGreen)
			}
			item.lunarLabel.SetCaption(holiday.GetName())
			item.lunarLabel.Font().SetColor(colors.ClYellow)
		}

		dayNum += 1
	}
}
func (m *TView) ClearSelect() {
	for idx, day := range m.days {
		if day.Color() == colors.ClGreen || day.solarLabel.Color() == colors.ClGreen {
			m.days[idx].SetColor(model.BackColor)
		}
	}
}
func (m *TView) OnViewDayClick(sender vcl.IObject) {
	m.ClearSelect()
	clickBtn := vcl.AsButton(sender)
	m.day, _ = strconv.Atoi(clickBtn.Caption())
	// p:solarLabel->Panel    p.p view.panel  p.p.p calendar.panel  p.p.p
	clickBtn.Parent().Parent().Parent().Parent().Hide()

}
