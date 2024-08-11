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
	ProjectTabSheet   *TProjectPage
}

func NewPageCtl(owner vcl.IComponent) *TPageCtl {
	c := new(TPageCtl)

	c.TPageControl = vcl.NewPageControl(owner)
	c.TPageControl.SetLeft(10)
	c.TPageControl.SetTop(10)
	c.TPageControl.SetWidth(744)
	c.TPageControl.SetHeight(550)

	c.TPageControl.SetControlStyle(types.TsFlatButtons)
	c.TPageControl.SetAlign(types.AlClient)
	c.TPageControl.SetTabOrder(0)
	c.TPageControl.SetColor(0x342C29)

	c.TodayTabSheet = NewTodayPage(c)
	c.TodayTabSheet.SetParent(c)

	c.InboxTabSheet = NewInboxPage(c)
	c.InboxTabSheet.SetParent(c)

	c.ScheduledTabSheet = NewScheduledPage(c)
	c.ScheduledTabSheet.SetParent(c)

	c.LabelsTabSheet = NewLabelsPage(c)
	c.LabelsTabSheet.SetParent(c)

	c.ProjectTabSheet = NewProjectPage(c)
	c.ProjectTabSheet.SetParent(c)

	return c
}
