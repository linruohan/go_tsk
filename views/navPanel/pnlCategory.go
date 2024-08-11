package navPanel

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
	"go_tsk/model"
	"go_tsk/views/navPanel/category"
)

type TPnlCategory struct {
	*vcl.TPanel
	//默认
	inboxBtn1     *vcl.TSpeedButton
	todayBtn1     *vcl.TSpeedButton
	scheduledBtn1 *vcl.TSpeedButton
	labelsBtn1    *vcl.TSpeedButton
	//new
	inboxBtn     *category.TInboxCate
	todayBtn     *category.TTodayCate
	scheduledBtn *category.TScheduledCate
	labelsBtn    *category.TLabelsCate
}

func NewPnlCategory(owner vcl.IComponent) *TPnlCategory {
	c := new(TPnlCategory)
	c.TPanel = vcl.NewPanel(owner)
	c.SetAlign(types.AlTop)
	c.SetBevelOuter(types.BvNone)
	c.SetParentBackground(false)
	c.SetParentColor(false)
	c.SetTabOrder(0)
	c.SetColor(0x342929)
	c.SetClientWidth(model.LeftWidth)
	c.SetClientHeight(100)
	c.SetBorderWidth(0)

	//init category
	c.categoryInit1()

	return c

}

func (c *TPnlCategory) categoryInit1() {
	c.inboxBtn = category.NewInboxCate(c)
	c.inboxBtn.SetParent(c)
	c.inboxBtn.SetBounds(0, 0, model.LeftWidth/2, 50)

	c.todayBtn = category.NewTodayCate(c)
	c.todayBtn.SetParent(c)
	c.todayBtn.SetBounds(model.LeftWidth/2, 0, model.LeftWidth/2, 50)

	c.scheduledBtn = category.NewScheduledCate(c)
	c.scheduledBtn.SetParent(c)
	c.scheduledBtn.SetBounds(0, 50, model.LeftWidth/2, 50)

	c.labelsBtn = category.NewLabelsCate(c)
	c.labelsBtn.SetParent(c)
	c.labelsBtn.SetBounds(model.LeftWidth/2, 50, model.LeftWidth/2, 50)

}
func (c *TPnlCategory) categoryInit() {
	c.inboxBtn1 = vcl.NewSpeedButton(c)
	c.inboxBtn1.SetParent(c)
	c.inboxBtn1.SetCaption("Inbox")
	c.inboxBtn1.Font().SetColor(colors.ClBlue)
	c.inboxBtn1.SetBounds(0, 0, 100, 100)

	c.todayBtn1 = vcl.NewSpeedButton(c)
	c.todayBtn1.SetParent(c)
	c.todayBtn1.SetCaption("Today")
	c.todayBtn1.Font().SetColor(colors.ClGreen)
	c.todayBtn1.SetBounds(100, 0, 100, 100)

	c.scheduledBtn1 = vcl.NewSpeedButton(c)
	c.scheduledBtn1.SetParent(c)
	c.scheduledBtn1.SetCaption("Scheduled")
	c.scheduledBtn1.Font().SetColor(colors.ClPink)
	c.scheduledBtn1.SetBounds(0, 100, 100, 100)

	c.labelsBtn1 = vcl.NewSpeedButton(c)
	c.labelsBtn1.SetParent(c)
	c.labelsBtn1.SetCaption("Labels")
	c.labelsBtn1.Font().SetColor(colors.ClGray)
	c.labelsBtn1.SetBounds(100, 100, 100, 100)
}
