// Package fyneui 提供基于Fyne的GUI界面组件
package fyneui

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"goeluosifangkuai/internal/game"
	"goeluosifangkuai/pkg/types"
)

// GameUI 游戏界面
type GameUI struct {
	app        fyne.App
	window     fyne.Window
	game       game.Game
	gameCanvas *fyne.Container
	infoPanel  *fyne.Container

	// 游戏画布
	boardCells [][]*canvas.Rectangle

	// 下一个方块预览
	nextPieceCells  [][]*canvas.Rectangle
	nextPieceCanvas *fyne.Container

	// 界面元素
	scoreLabel  *widget.Label
	levelLabel  *widget.Label
	linesLabel  *widget.Label
	nextPanel   *fyne.Container
	statusLabel *widget.Label

	// 控制按钮
	startButton   *widget.Button
	pauseButton   *widget.Button
	restartButton *widget.Button

	// 游戏状态
	isRunning bool
	isPaused  bool
}

// NewGameUI 创建新的游戏界面
func NewGameUI(app fyne.App) *GameUI {
	window := app.NewWindow("俄罗斯方块 - Tetris")
	window.Resize(fyne.NewSize(950, 750)) // 进一步增大窗口尺寸
	window.CenterOnScreen()

	// 创建游戏实例
	config := game.DefaultGameConfig()
	gameInstance := game.NewGame(config)

	ui := &GameUI{
		app:    app,
		window: window,
		game:   gameInstance,
	}

	ui.setupUI()
	ui.setupKeyboardEvents()
	return ui
}

// setupUI 设置用户界面
func (ui *GameUI) setupUI() {
	// 创建游戏棋盘
	ui.createGameBoard()

	// 创建信息面板
	ui.createInfoPanel()

	// 创建控制面板
	ui.createControlPanel()

	// 布局界面
	ui.layoutUI()
}

// createGameBoard 创建游戏棋盘
func (ui *GameUI) createGameBoard() {
	// 初始化棋盘单元格
	ui.boardCells = make([][]*canvas.Rectangle, types.BoardHeight)

	boardContainer := container.NewWithoutLayout()

	cellSize := float32(25)    // 调整回更合适的单元格大小
	boardMargin := float32(10) // 边距

	for y := 0; y < types.BoardHeight; y++ {
		ui.boardCells[y] = make([]*canvas.Rectangle, types.BoardWidth)
		for x := 0; x < types.BoardWidth; x++ {
			cell := canvas.NewRectangle(color.RGBA{40, 40, 40, 255})
			cell.StrokeColor = color.RGBA{100, 100, 100, 255}
			cell.StrokeWidth = 1

			// 设置位置和大小
			cell.Resize(fyne.NewSize(cellSize, cellSize))
			cell.Move(fyne.NewPos(
				boardMargin+float32(x)*cellSize,
				boardMargin+float32(y)*cellSize,
			))

			ui.boardCells[y][x] = cell
			boardContainer.Add(cell)
		}
	}

	// 设置棋盘容器大小
	boardSize := fyne.NewSize(
		float32(types.BoardWidth)*cellSize+boardMargin*2,
		float32(types.BoardHeight)*cellSize+boardMargin*2,
	)
	boardContainer.Resize(boardSize)

	// 直接使用棋盘容器，不再用Border包装
	ui.gameCanvas = boardContainer
}

// createNextPiecePreview 创建下一个方块预览区域
func (ui *GameUI) createNextPiecePreview() {
	// 初始化下一个方块预览单元格 (4x4 网格足够显示所有方块)
	previewSize := 4
	ui.nextPieceCells = make([][]*canvas.Rectangle, previewSize)

	nextContainer := container.NewWithoutLayout()

	cellSize := float32(15) // 较小的预览单元格
	margin := float32(5)

	for y := 0; y < previewSize; y++ {
		ui.nextPieceCells[y] = make([]*canvas.Rectangle, previewSize)
		for x := 0; x < previewSize; x++ {
			cell := canvas.NewRectangle(color.RGBA{30, 30, 30, 255}) // 深灰色背景
			cell.StrokeColor = color.RGBA{60, 60, 60, 255}
			cell.StrokeWidth = 1

			// 设置位置和大小
			cell.Resize(fyne.NewSize(cellSize, cellSize))
			cell.Move(fyne.NewPos(
				margin+float32(x)*cellSize,
				margin+float32(y)*cellSize,
			))

			ui.nextPieceCells[y][x] = cell
			nextContainer.Add(cell)
		}
	}

	// 设置预览区域大小
	previewCanvasSize := fyne.NewSize(
		float32(previewSize)*cellSize+margin*2,
		float32(previewSize)*cellSize+margin*2,
	)
	nextContainer.Resize(previewCanvasSize)

	ui.nextPieceCanvas = nextContainer
}

// createInfoPanel 创建信息面板
func (ui *GameUI) createInfoPanel() {
	// 分数标签
	ui.scoreLabel = widget.NewLabel("分数: 0")
	ui.scoreLabel.TextStyle = fyne.TextStyle{Bold: true}

	// 等级标签
	ui.levelLabel = widget.NewLabel("等级: 1")
	ui.levelLabel.TextStyle = fyne.TextStyle{Bold: true}

	// 行数标签
	ui.linesLabel = widget.NewLabel("行数: 0")
	ui.linesLabel.TextStyle = fyne.TextStyle{Bold: true}

	// 状态标签
	ui.statusLabel = widget.NewLabel("准备开始")
	ui.statusLabel.TextStyle = fyne.TextStyle{Italic: true}

	// 创建下一个方块预览区域
	ui.createNextPiecePreview()

	// 下一个方块预览面板
	ui.nextPanel = container.NewVBox(
		widget.NewLabel("下一个方块:"),
		ui.nextPieceCanvas, // 使用创建的预览画布
	)

	// 组织信息面板
	ui.infoPanel = container.NewVBox(
		widget.NewCard("游戏信息", "", container.NewVBox(
			ui.scoreLabel,
			ui.levelLabel,
			ui.linesLabel,
			ui.statusLabel,
		)),
		widget.NewSeparator(),
		ui.nextPanel,
	)
}

// createControlPanel 创建控制面板
func (ui *GameUI) createControlPanel() {
	ui.startButton = widget.NewButton("开始游戏", ui.startGame)
	ui.pauseButton = widget.NewButton("暂停", ui.togglePause)
	ui.pauseButton.Disable()
	ui.restartButton = widget.NewButton("重新开始", ui.restartGame)
	ui.restartButton.Disable()
}

// layoutUI 布局界面
func (ui *GameUI) layoutUI() {
	// 顶部标题
	titleLabel := widget.NewLabel("🎮 俄罗斯方块")
	titleLabel.Alignment = fyne.TextAlignCenter
	titleLabel.TextStyle = fyne.TextStyle{Bold: true}

	// 为游戏棋盘添加边距和背景，防止重叠
	gameCanvasWrapper := container.NewVBox(
		widget.NewCard("", "", ui.gameCanvas),
	)

	// 信息面板也用卡片包装，确保清晰分离
	infoWrapper := container.NewVBox(
		ui.infoPanel,
	)

	// 使用HSplit但设置明确的分割比例
	gameMainArea := container.NewHSplit(
		gameCanvasWrapper, // 左侧：包装后的游戏棋盘
		infoWrapper,       // 右侧：包装后的信息面板
	)
	// 设置分割比例：70%给游戏区域，30%给信息面板
	gameMainArea.SetOffset(0.7)

	// 控制按钮容器 - 水平排列
	buttonContainer := container.NewHBox(
		ui.startButton,
		ui.pauseButton,
		ui.restartButton,
	)

	// 底部说明文字
	helpLabel := widget.NewLabel("使用 A/D 左右移动，W 旋转，S 下降，空格快速下降")
	helpLabel.Alignment = fyne.TextAlignCenter

	// 使用Border布局，确保游戏区域在中心，按钮在底部
	mainContainer := container.NewBorder(
		container.NewVBox(titleLabel, widget.NewSeparator()),                 // 顶部：标题
		container.NewVBox(widget.NewSeparator(), buttonContainer, helpLabel), // 底部：按钮+帮助
		nil,          // 左侧：无
		nil,          // 右侧：无
		gameMainArea, // 中心：游戏区域
	)

	ui.window.SetContent(mainContainer)
}

// setupKeyboardEvents 设置键盘事件
func (ui *GameUI) setupKeyboardEvents() {
	ui.window.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		if !ui.isRunning || ui.isPaused {
			return
		}

		switch event.Name {
		case fyne.KeyA:
			ui.game.MoveTetromino(-1, 0)
		case fyne.KeyD:
			ui.game.MoveTetromino(1, 0)
		case fyne.KeyS:
			ui.game.MoveTetromino(0, 1)
		case fyne.KeyW:
			ui.game.RotateTetromino(types.DirectionRight)
		case fyne.KeySpace:
			ui.game.DropTetromino()
		case fyne.KeyP:
			ui.togglePause()
		}

		ui.updateDisplay()
	})
}

// startGame 开始游戏
func (ui *GameUI) startGame() {
	// 如果是游戏结束后重新开始，需要重置游戏
	if ui.game.GetState() == types.GameStateGameOver {
		ui.game.Reset()
	}

	ui.game.SetState(types.GameStatePlaying)
	ui.isRunning = true
	ui.isPaused = false

	ui.startButton.Disable()
	ui.startButton.SetText("开始游戏") // 重置按钮文字
	ui.pauseButton.Enable()
	ui.restartButton.Enable()

	ui.statusLabel.SetText("游戏进行中")

	// 更新显示后再启动定时器
	ui.updateDisplay()

	// 使用定时器而不是单独的goroutine
	ui.startGameTimer()
}

// startGameTimer 启动游戏定时器
func (ui *GameUI) startGameTimer() {
	ticker := time.NewTicker(time.Millisecond * 500)
	lastUpdate := time.Now()

	go func() {
		defer ticker.Stop()
		for ui.isRunning {
			select {
			case <-ticker.C:
				if !ui.isPaused {
					now := time.Now()
					deltaTime := int(now.Sub(lastUpdate) / time.Millisecond)
					lastUpdate = now

					// 更新游戏状态
					ui.game.Update(deltaTime)

					// 检查游戏结束
					if ui.game.GetState() == types.GameStateGameOver {
						ui.handleGameOver()
						return
					}

					// 更新显示
					ui.updateDisplay()
				}
			}
		}
	}()
}

// togglePause 切换暂停状态
func (ui *GameUI) togglePause() {
	if !ui.isRunning {
		return
	}

	ui.isPaused = !ui.isPaused

	if ui.isPaused {
		ui.game.SetState(types.GameStatePaused)
		ui.pauseButton.SetText("继续")
		ui.statusLabel.SetText("游戏已暂停")
	} else {
		ui.game.SetState(types.GameStatePlaying)
		ui.pauseButton.SetText("暂停")
		ui.statusLabel.SetText("游戏进行中")
	}
}

// restartGame 重新开始游戏
func (ui *GameUI) restartGame() {
	ui.game.Reset()
	ui.game.SetState(types.GameStatePlaying)
	ui.isRunning = true
	ui.isPaused = false

	ui.startButton.Disable()
	ui.pauseButton.Enable()
	ui.pauseButton.SetText("暂停")
	ui.restartButton.Enable()

	ui.statusLabel.SetText("游戏进行中")
	ui.updateDisplay()

	// 重新启动游戏循环
	ui.startGameTimer()
}

// handleGameOver 处理游戏结束
func (ui *GameUI) handleGameOver() {
	ui.isRunning = false
	ui.isPaused = false

	ui.startButton.Enable()
	ui.startButton.SetText("重新开始")
	ui.pauseButton.Disable()
	ui.restartButton.Enable()

	ui.statusLabel.SetText(fmt.Sprintf("游戏结束！最终分数: %d", ui.game.GetScore()))
}

// updateDisplay 更新显示
func (ui *GameUI) updateDisplay() {
	// 在主UI线程中更新游戏信息
	fyne.DoAndWait(func() {
		ui.scoreLabel.SetText(fmt.Sprintf("分数: %d", ui.game.GetScore()))
		ui.levelLabel.SetText(fmt.Sprintf("等级: %d", ui.game.GetLevel()))
		ui.linesLabel.SetText(fmt.Sprintf("行数: %d", ui.game.GetLinesCleared()))
	})

	// 更新棋盘显示
	ui.updateBoard()

	// 更新下一个方块预览
	ui.updateNextPiece()
}

// updateBoard 更新棋盘显示
func (ui *GameUI) updateBoard() {
	board := ui.game.GetBoard()
	currentTetromino := ui.game.GetCurrentTetromino()

	// 创建渲染缓冲区
	buffer := make([][]types.Color, types.BoardHeight)
	for i := range buffer {
		buffer[i] = make([]types.Color, types.BoardWidth)
		for j := range buffer[i] {
			buffer[i][j] = board.GetCell(j, i)
		}
	}

	// 渲染当前方块
	if currentTetromino != nil {
		position := currentTetromino.GetPosition()
		blocks := currentTetromino.GetBlocks()
		tetrominoColor := currentTetromino.GetColor()

		for _, block := range blocks {
			x := position.X + block.X
			y := position.Y + block.Y

			if x >= 0 && x < types.BoardWidth && y >= 0 && y < types.BoardHeight {
				buffer[y][x] = tetrominoColor
			}
		}
	}

	// 更新单元格颜色 - 在主UI线程中执行
	fyne.DoAndWait(func() {
		for y := 0; y < types.BoardHeight; y++ {
			for x := 0; x < types.BoardWidth; x++ {
				cellColor := ui.getColorForType(buffer[y][x])
				ui.boardCells[y][x].FillColor = cellColor
				ui.boardCells[y][x].Refresh()
			}
		}
	})
}

// updateNextPiece 更新下一个方块预览
func (ui *GameUI) updateNextPiece() {
	nextTetromino := ui.game.GetNextTetromino()
	if nextTetromino == nil {
		return
	}

	// 获取下一个方块的块位置
	blocks := nextTetromino.GetBlocks()
	nextColor := nextTetromino.GetColor()
	uiColor := ui.getColorForType(nextColor)

	// 计算方块在预览区域的中心位置
	// 找到方块的边界
	minX, maxX := blocks[0].X, blocks[0].X
	minY, maxY := blocks[0].Y, blocks[0].Y
	for _, block := range blocks {
		if block.X < minX {
			minX = block.X
		}
		if block.X > maxX {
			maxX = block.X
		}
		if block.Y < minY {
			minY = block.Y
		}
		if block.Y > maxY {
			maxY = block.Y
		}
	}

	// 计算偏移量以将方块居中显示
	offsetX := (4 - (maxX - minX + 1)) / 2
	offsetY := (4 - (maxY - minY + 1)) / 2

	// 在主UI线程中更新预览区域
	fyne.DoAndWait(func() {
		// 清空预览区域
		for y := 0; y < 4; y++ {
			for x := 0; x < 4; x++ {
				ui.nextPieceCells[y][x].FillColor = color.RGBA{30, 30, 30, 255} // 深灰色背景
			}
		}

		// 渲染下一个方块
		for _, block := range blocks {
			x := block.X - minX + offsetX
			y := block.Y - minY + offsetY

			if x >= 0 && x < 4 && y >= 0 && y < 4 {
				ui.nextPieceCells[y][x].FillColor = uiColor
			}
		}

		// 刷新所有预览单元格
		for y := 0; y < 4; y++ {
			for x := 0; x < 4; x++ {
				ui.nextPieceCells[y][x].Refresh()
			}
		}
	})
}

// getColorForType 根据方块类型获取颜色
func (ui *GameUI) getColorForType(colorType types.Color) color.Color {
	switch colorType {
	case types.ColorEmpty:
		return color.RGBA{40, 40, 40, 255} // 深灰色背景
	case types.ColorI:
		return color.RGBA{0, 255, 255, 255} // 青色
	case types.ColorO:
		return color.RGBA{255, 255, 0, 255} // 黄色
	case types.ColorT:
		return color.RGBA{128, 0, 128, 255} // 紫色
	case types.ColorS:
		return color.RGBA{0, 255, 0, 255} // 绿色
	case types.ColorZ:
		return color.RGBA{255, 0, 0, 255} // 红色
	case types.ColorJ:
		return color.RGBA{0, 0, 255, 255} // 蓝色
	case types.ColorL:
		return color.RGBA{255, 165, 0, 255} // 橙色
	default:
		return color.RGBA{40, 40, 40, 255}
	}
}

// Show 显示窗口
func (ui *GameUI) Show() {
	ui.window.ShowAndRun()
}

// Close 关闭窗口
func (ui *GameUI) Close() {
	ui.isRunning = false
	ui.window.Close()
}
