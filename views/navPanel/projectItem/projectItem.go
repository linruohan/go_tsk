package projectItem

import (
	"github.com/ying32/govcl/vcl"
)

type TProject struct {
	TUtilBtnProject
}

func NewProject(owner vcl.IComponent, icon, name, due string) *TProject {
	c := new(TProject)
	c.icon = icon
	c.name = name
	c.due = due
	c.ProjectInit(owner, c.icon, c.name, c.due)
	return c
}
