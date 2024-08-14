package calendar

import (
	"fmt"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"go_tsk/model"
	"go_tsk/utils"
	"time"
)

type TCalendar struct {
	*vcl.TPanel
	year        int
	month       int
	day         int
	header      *THeader
	view        *TView
	week        *TWeek
	currentDate time.Time
}

func NewCalendar(owner vcl.IComponent) *TCalendar {
	c := new(TCalendar)
	c.TPanel = vcl.NewPanel(owner)
	c.TPanel.SetBorderStyle(types.None)
	c.TPanel.SetBorderWidth(0)

	p := vcl.NewPanel(c)
	p.SetParent(c)
	p.SetBorderStyle(types.AlClient)
	p.SetColor(model.BackColor)
	p.SetBounds(0, 0, 350, 350)
	p.Font().SetName(model.FontName)
	p.Font().SetSize(model.FontSize)

	c.currentDate = time.Now()
	c.year = c.currentDate.Year()
	c.month = int(c.currentDate.Month())

	c.header = NewHeader(c)
	c.header.SetParent(p)
	c.header.SetBounds(0, 0, 350, 25)
	c.header.monthLabel.SetCaption(fmt.Sprintf("%d月", int32(c.currentDate.Month())))
	c.header.yearLabel.SetCaption(fmt.Sprintf("%d年", c.currentDate.Year()))
	c.header.leftBtn.SetOnClick(c.onLeftClick)
	c.header.centerBtn.SetOnClick(c.onCenterClick)
	c.header.rightBtn.SetOnClick(c.onRightClick)

	c.week = NewWeek(c)
	c.week.SetParent(p)
	c.week.SetBounds(0, 25, 350, 25)

	c.view = NewView(c)
	c.view.SetParent(p)
	c.view.SetBounds(0, 50, 350, 300)
	return c
}

func (c *TCalendar) fillDays(date time.Time) {
	firstOfMonth := utils.GenerateDate(date, 1)
	// 每个月的最后一天
	endDay := firstOfMonth.AddDate(0, 1, -1).Day()
	// 每个月第一天的星期几
	startPos := int(firstOfMonth.Weekday()) - 1
	if startPos < 0 {
		startPos += 7
	}
	startPos = (startPos + 7) % 7
	c.header.yearLabel.SetCaption(fmt.Sprintf("%d年", c.year))
	c.header.monthLabel.SetCaption(fmt.Sprintf("%d月", c.month))
	c.view.fillGridDays(startPos, endDay, date)
}

func (c *TCalendar) onLeftClick(sender vcl.IObject) {
	c.month -= 1
	if c.month < 1 {
		c.year -= 1
		c.month = 12
	}
	date := time.Date(c.year, time.Month(c.month), c.day, 0, 0, 0, 0, time.Local)
	c.fillDays(date)
}

func (c *TCalendar) onRightClick(sender vcl.IObject) {
	c.month = c.month + 1
	if c.month > 12 {
		c.year += 1
		c.month = 1
	}
	date := time.Date(c.year, time.Month(c.month), c.day, 0, 0, 0, 0, time.Local)
	c.fillDays(date)
}

func (c *TCalendar) onCenterClick(sender vcl.IObject) {
	c.year = c.currentDate.Year()
	c.month = int(c.currentDate.Month())
	c.day = c.currentDate.Day()
	c.fillDays(c.currentDate)
}
func (c *TCalendar) GetDay() string {
	if c.view.day != 0 {
		c.day = c.view.day
	}
	return fmt.Sprintf("%4d-%02d-%02d", c.year, c.month, c.day)
}
