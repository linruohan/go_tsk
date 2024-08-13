package ui

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
	"go_tsk/model"
	"go_tsk/views/mainPanel"
	"go_tsk/views/mainPanel/pages"
	"go_tsk/views/navPanel"
	"go_tsk/views/navPanel/projectItem"
)

type TMainForm struct {
	*vcl.TForm
	pnlNav       *navPanel.TPnlNav
	pnlMain      *mainPanel.TPnlMain
	projectTag   int
	projectPages []*pages.TProjectPage
}

var (
	MainForm *TMainForm
)

func (m *TMainForm) OnFormCreate(sender vcl.IObject) {
	m.SetLeft(0)
	m.SetTop(0)
	m.HorzScrollBar().SetVisible(false)
	m.VertScrollBar().SetVisible(false)
	m.SetAlphaBlend(true)
	m.SetAlphaBlendValue(250)
	m.SetBorderStyle(types.BsToolWindow)
	m.SetCaption("go_tsk")
	m.SetClientWidth(1000)
	m.SetClientHeight(580)
	m.SetDefaultMonitor(types.DmMainForm)
	m.Font().SetColor(colors.ClWhite)
	//m.Font().SetHeight(11)
	m.Font().SetName("LXGW WenKai GB")
	m.SetPosition(types.PoScreenCenter)

	m.pnlNav = navPanel.NewPnlNav(m)
	m.pnlNav.SetParent(m)
	//nav绑定函数
	m.pnlNav.PnlCategory.Inbox.SetOnClick(m.OnNavClick)
	m.pnlNav.PnlCategory.Today.SetOnClick(m.OnNavClick)
	m.pnlNav.PnlCategory.Scheduled.SetOnClick(m.OnNavClick)
	m.pnlNav.PnlCategory.Labels.SetOnClick(m.OnNavClick)
	for _, project := range m.pnlNav.PnlProjects.ProjectItemList {
		project.SetOnClick(m.OnNavClick)
	}
	m.projectTag = 3
	m.pnlMain = mainPanel.NewPnlMain(m)
	m.pnlMain.SetParent(m)

	m.AddProject(m.pnlNav.PnlProjects.ProjectsPnl, "🎈", "Project Name", "today")
	m.pnlNav.PnlProjects.AddLbl.SetOnClick(func(sender vcl.IObject) {
		m.AddProject(m.pnlNav.PnlProjects.ProjectsPnl, "🎈", "Project Name", "today")
	})

}
func (m *TMainForm) OnNavClick(sender vcl.IObject) {
	tmp := vcl.AsSpeedButton(sender)
	m.pnlMain.PageCtl.SetActivePageIndex(int32(tmp.Tag()))
	//fmt.Printf("index:%d;tag:%d\n", m.pnlMain.PageCtl.ActivePageIndex(), tmp.Tag())
}
func (m *TMainForm) AddProject(owner *vcl.TPanel, icon, name, due string) {
	m.projectTag += 1
	pro := projectItem.NewProject(owner, icon, name, due)
	pro.SetParent(owner)
	pro.SetAlign(types.AlTop)
	pro.SetTag(m.projectTag)
	pro.SetOnClick(m.OnNavClick)
	m.pnlNav.PnlProjects.ProjectItemList = append(m.pnlNav.PnlProjects.ProjectItemList, pro)
	m.showProjectList()
	//	add pages
	projectPage := pages.NewProjectPage(m.pnlMain.PageCtl.TPageControl)
	projectPage.SetParent(m.pnlMain.PageCtl.TPageControl)
	projectPage.SetPageIndex(int32(m.projectTag))
	projectPage.SetOnMouseMove(m.pnlMain.OnPnlMainMouseMove)
	projectPage.SetOnMouseUp(m.pnlMain.OnPnlMainMouseUp)
	projectPage.SetOnMouseDown(m.pnlMain.OnPnlMainMouseDown)
	m.projectPages = append(m.projectPages, projectPage)
}
func (m *TMainForm) showProjectList() {
	for i, pro := range m.pnlNav.PnlProjects.ProjectItemList {
		pro.SetHeight(25)
		pro.SetBounds(0, int32(int(25*i)), model.LeftWidth, 25)
		pro.SetAlign(types.AlTop)
	}
}
