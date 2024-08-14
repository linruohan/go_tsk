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

	ProjectsPnl     *vcl.TPanel
	ProjectItemList []*projectItem.TProject
	AddPro          *vcl.TSpeedButton
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
	l.Font().SetName("Yuppy SC Regular")
	l.Font().SetSize(model.FontSize + 3)
	l.Font().SetColor(model.FontColor)
	l.SetHeight(25)
	l.SetWidth(model.LeftWidth - 50)
	l.SetBounds(20, 0, model.LeftWidth*0.9, 25)
	l.SetAlignment(types.TaCenter)
	l.SetLayout(types.TlCenter)

	c.AddPro = vcl.NewSpeedButton(c.computerPnl)
	c.AddPro.SetParent(c.computerPnl)
	c.AddPro.SetHeight(25)
	c.AddPro.SetWidth(50)
	c.AddPro.SetParent(c.computerPnl)
	c.AddPro.SetCaption("➕")
	c.AddPro.SetFlat(true)
	c.AddPro.SetTransparent(true)
	c.AddPro.Font().SetColor(colors.ClSkyblue)
	c.AddPro.SetAlign(types.AlRight)
	c.AddPro.SetBounds(model.LeftWidth*0.9, 0, model.LeftWidth*0.1, 25)

	c.ProjectsPnl = vcl.NewPanel(owner)
	c.ProjectsPnl.SetParent(c)
	c.ProjectsPnl.SetAlign(types.AlClient)

	//c.AddProjectItem(c.ProjectsPnl, " ", "project案例", "today")
	//c.AddProjectItem(c.ProjectsPnl, "📦", "Inbox", "tomorrow")
	//c.AddProjectItem(c.ProjectsPnl, "🐮", "Today", "tomorrow")
	//c.AddProjectItem(c.ProjectsPnl, "📅", "test project", "tomorrow")
	//c.showProjectList()
	return c

}
