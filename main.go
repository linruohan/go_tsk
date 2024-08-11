package main

import (
	"fmt"
	// 如果使用自定义的syso文件，不要引用这个包 github.com/ying32/govcl/pkgs/winappres
	// 不带uac的
	//import _ "github.com/ying32/govcl/pkgs/winappres"
	// 带uac的
	//import _ "github.com/ying32/govcl/pkgs/winappres/uac"
	"github.com/ying32/govcl/vcl"
	"go_tsk/ui"
)

func main() {
	// 异常捕获
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Exception: ", err)
			vcl.ShowMessage(err.(error).Error())
		}
	}()

	vcl.Application.Initialize()
	vcl.Application.SetOnException(func(sender vcl.IObject, e *vcl.Exception) {
		fmt.Println("exception:", e.ClassName())
	})
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.SetScaled(true)
	vcl.Application.CreateForm(&ui.MainForm)
	//vcl.Application.CreateForm(&ui.AboutForm)
	vcl.Application.Run()
}
