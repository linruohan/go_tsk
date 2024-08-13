package pages

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type TPageCtl struct {
	*vcl.TPageControl

	TodayTabSheet     *TTodayPage
	InboxTabSheet     *TInboxPage
	ScheduledTabSheet *TScheduledPage
	LabelsTabSheet    *TLabelsPage
}

func NewPageCtl(owner vcl.IComponent) *TPageCtl {
	c := new(TPageCtl)

	c.TPageControl = vcl.NewPageControl(owner)
	c.TPageControl.SetLeft(10)
	c.TPageControl.SetTop(10)
	c.TPageControl.SetWidth(744)
	c.TPageControl.SetHeight(550)

	c.TPageControl.SetControlStyle(types.TsNone)
	c.TPageControl.SetAlign(types.AlClient)
	c.TPageControl.SetTabOrder(0)

	c.InboxTabSheet = NewInboxPage(c)
	c.InboxTabSheet.SetParent(c)

	c.TodayTabSheet = NewTodayPage(c)
	c.TodayTabSheet.SetParent(c)

	c.ScheduledTabSheet = NewScheduledPage(c)
	c.ScheduledTabSheet.SetParent(c)

	c.LabelsTabSheet = NewLabelsPage(c)
	c.LabelsTabSheet.SetParent(c)

	return c
}

func (f *TPageCtl) hideAllTab() {
	var i int32
	for i = 0; i < f.PageCount(); i++ {
		sheet := f.Pages(i)
		sheet.SetTabVisible(false)
		sheet.SetVisible(false)
	}
}

/*func (f *TPageCtl) setPage(idx int32) {
	if idx != 0 && idx != -1 && idx != 1 {
		return
	}
	if idx == 0 {
		f.SetActivePageIndex(0)
		sheet := f.Pages(0)
		sheet.SetVisible(true)
		return
	}
	curIdx := f.ActivePageIndex()
	sheet := f.Pages(curIdx)
	sheet.SetVisible(false)
	f.SetActivePageIndex(curIdx + idx)
	sheet = f.Pages(curIdx + idx)
	sheet.SetVisible(true)
}*/

func (f *TPageCtl) OnActPagePrevExecute(sender vcl.IObject) {
	f.SetActivePageIndex(f.ActivePageIndex() - 1)
	//f.setPage(-1)
}

func (f *TPageCtl) OnActPagePrevUpdate(sender vcl.IObject) {
	vcl.AsAction(sender).SetEnabled(f.ActivePageIndex() > 0)
}

func (f *TPageCtl) OnActPageNextExecute(sender vcl.IObject) {
	f.SetActivePageIndex(f.ActivePageIndex() + 1)
	//f.setPage(1)
}

func (f *TPageCtl) OnActPageNextUpdate(sender vcl.IObject) {
	vcl.AsAction(sender).SetEnabled(f.ActivePageIndex() < f.PageCount()-1)
}
