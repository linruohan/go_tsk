package navPanel

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"go_tsk/model"
	"go_tsk/views/navPanel/category"
)

type TPnlCategory struct {
	*vcl.TPanel
	Inbox     *category.TInboxCate
	Today     *category.TTodayCate
	Scheduled *category.TScheduledCate
	Labels    *category.TLabelsCate
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
	c.categoryInit()

	return c

}

func (cate *TPnlCategory) categoryInit() {
	cate.Inbox = category.NewInboxCate(cate)
	cate.Inbox.SetParent(cate)
	cate.Inbox.SetBounds(0, 0, model.LeftWidth/2, 50)
	cate.Inbox.SetTag(0)

	cate.Today = category.NewTodayCate(cate)
	cate.Today.SetParent(cate)
	cate.Today.SetBounds(model.LeftWidth/2, 0, model.LeftWidth/2, 50)
	cate.Today.SetTag(1)

	cate.Scheduled = category.NewScheduledCate(cate)
	cate.Scheduled.SetParent(cate)
	cate.Scheduled.SetBounds(0, 50, model.LeftWidth/2, 50)
	cate.Scheduled.SetTag(2)

	cate.Labels = category.NewLabelsCate(cate)
	cate.Labels.SetParent(cate)
	cate.Labels.SetBounds(model.LeftWidth/2, 50, model.LeftWidth/2, 50)
	cate.Labels.SetTag(3)

}
