package navPanel

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
	"go_tsk/model"
	"go_tsk/views/navPanel/projectItem"
)

type TPnlProject struct {
	*vcl.TPanel
	computerPnl *vcl.TPanel

	projectsPnl     *vcl.TPanel
	projectItemList []*projectItem.TProject
}

func NewPnlProject(owner vcl.IComponent) *TPnlProject {
	c := new(TPnlProject)
	c.TPanel = vcl.NewPanel(owner)
	c.SetAlign(types.AlClient)
	c.SetBevelOuter(types.BvNone)
	c.SetParentBackground(false)
	c.SetParentColor(false)
	c.SetTabOrder(0)
	c.SetColor(0x342929)
	c.SetClientWidth(model.LeftWidth)
	c.SetClientHeight(380)
	c.SetBorderStyle(types.BvNone)
	c.SetBorderWidth(0)

	c.computerPnl = vcl.NewPanel(owner)
	c.computerPnl.SetParent(c)
	c.computerPnl.SetHeight(25)
	c.computerPnl.SetAlign(types.AlTop)
	l := vcl.NewLabel(c.computerPnl)
	l.SetParent(c.computerPnl)
	l.SetCaption("On this Computer")
	l.Font().SetSize(14)
	l.Font().SetColor(colors.ClWhite)
	l.SetHeight(25)
	l.SetWidth(50)
	l.SetBounds(0, 0, model.LeftWidth*0.9, 25)
	l.SetAlignment(types.TaCenter)
	l.SetLayout(types.TlCenter)

	lad := vcl.NewLabel(c.computerPnl)
	lad.SetHeight(25)
	lad.SetWidth(50)
	lad.SetParent(c.computerPnl)
	lad.SetCaption("➕")
	lad.Font().SetColor(colors.ClSkyblue)
	lad.SetAlign(types.AlRight)
	lad.SetBounds(model.LeftWidth*0.9, 0, model.LeftWidth*0.1, 25)

	c.projectsPnl = vcl.NewPanel(owner)
	c.projectsPnl.SetParent(c)
	c.projectsPnl.SetAlign(types.AlClient)

	c.AddProjectItem(c.projectsPnl, " ", "project案例", "today")
	c.AddProjectItem(c.projectsPnl, "📦", "Inbox", "tomorrow")
	c.AddProjectItem(c.projectsPnl, "🐮", "Today", "tomorrow")
	c.AddProjectItem(c.projectsPnl, "📅", "test project", "tomorrow")
	c.showProjectList()
	return c

}
func (c *TPnlProject) AddProjectItem(owner *vcl.TPanel, icon, name, due string) {
	pro := projectItem.NewProject(owner, icon, name, due)
	pro.SetParent(owner)
	c.projectItemList = append(c.projectItemList, pro)
}
func (c *TPnlProject) showProjectList() {
	for i, pro := range c.projectItemList {
		pro.SetHeight(25)
		pro.SetBounds(0, int32(int(25*i)), model.LeftWidth, 25)
		pro.SetAlign(types.AlTop)
	}
}
