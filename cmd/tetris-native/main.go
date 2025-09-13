package main

import (
	"fyne.io/fyne/v2/app"
	"goeluosifangkuai/internal/fyneui"
)

func main() {
	// 创建Fyne应用
	myApp := app.New()
	myApp.SetIcon(nil) // 可以设置应用图标

	// 创建游戏UI
	gameUI := fyneui.NewGameUI(myApp)

	// 显示并运行应用
	gameUI.Show()
}