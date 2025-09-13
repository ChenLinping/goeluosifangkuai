// Package game 实现游戏主控制器逻辑
package game

import (
	"goeluosifangkuai/pkg/types"
)

// gameImpl 是 Game 接口的具体实现
type gameImpl struct {
	state            types.GameState
	board            Board
	currentTetromino Tetromino
	nextTetromino    Tetromino
	factory          *TetrominoFactory

	// 游戏统计
	score        int
	level        int
	linesCleared int

	// 游戏时间控制
	dropTimer    int
	dropInterval int

	// 游戏配置
	config GameConfig
}

// GameConfig 游戏配置
type GameConfig struct {
	BoardWidth           int
	BoardHeight          int
	InitialDropInterval  int
	MinDropInterval      int
	FastDropInterval     int
	ScorePerLine         int
	ScoreLevelMultiplier int
	LinesPerLevel        int
}

// DefaultGameConfig 返回默认游戏配置
func DefaultGameConfig() GameConfig {
	return GameConfig{
		BoardWidth:           types.BoardWidth,
		BoardHeight:          types.BoardHeight,
		InitialDropInterval:  types.InitialDropInterval,
		MinDropInterval:      types.MinDropInterval,
		FastDropInterval:     types.FastDropInterval,
		ScorePerLine:         types.ScorePerLine,
		ScoreLevelMultiplier: types.ScoreLevelMultiplier,
		LinesPerLevel:        10,
	}
}

// NewGame 创建新的游戏实例
func NewGame(config GameConfig) Game {
	factory := NewTetrominoFactory()
	board := NewBoard(config.BoardWidth, config.BoardHeight)

	game := &gameImpl{
		state:        types.GameStateMenu,
		board:        board,
		factory:      factory,
		config:       config,
		score:        0,
		level:        1,
		linesCleared: 0,
		dropTimer:    0,
		dropInterval: config.InitialDropInterval,
	}

	game.generateNextTetromino()
	game.spawnNewTetromino()

	return game
}

// GetState 返回当前游戏状态
func (g *gameImpl) GetState() types.GameState {
	return g.state
}

// SetState 设置游戏状态
func (g *gameImpl) SetState(state types.GameState) {
	g.state = state
}

// GetBoard 返回游戏棋盘
func (g *gameImpl) GetBoard() Board {
	return g.board
}

// GetCurrentTetromino 返回当前正在下落的方块
func (g *gameImpl) GetCurrentTetromino() Tetromino {
	return g.currentTetromino
}

// GetNextTetromino 返回下一个方块
func (g *gameImpl) GetNextTetromino() Tetromino {
	return g.nextTetromino
}

// GetScore 返回当前分数
func (g *gameImpl) GetScore() int {
	return g.score
}

// GetLevel 返回当前等级
func (g *gameImpl) GetLevel() int {
	return g.level
}

// GetLinesCleared 返回已消除的行数
func (g *gameImpl) GetLinesCleared() int {
	return g.linesCleared
}

// MoveTetromino 移动当前方块
func (g *gameImpl) MoveTetromino(dx, dy int) bool {
	if g.state != types.GameStatePlaying || g.currentTetromino == nil {
		return false
	}

	// 创建新位置的方块副本
	newTetromino := g.currentTetromino.Clone()
	currentPos := newTetromino.GetPosition()
	newPos := types.Position{
		X: currentPos.X + dx,
		Y: currentPos.Y + dy,
	}
	newTetromino.SetPosition(newPos)

	// 检查新位置是否有效
	if g.board.IsValidPosition(newTetromino) {
		g.currentTetromino = newTetromino
		return true
	}

	return false
}

// RotateTetromino 旋转当前方块
func (g *gameImpl) RotateTetromino(direction types.Direction) bool {
	if g.state != types.GameStatePlaying || g.currentTetromino == nil {
		return false
	}

	// 尝试旋转
	rotatedTetromino := g.currentTetromino.Rotate(direction)

	// 检查旋转后的位置是否有效
	if g.board.IsValidPosition(rotatedTetromino) {
		g.currentTetromino = rotatedTetromino
		return true
	}

	// 如果直接旋转失败，尝试踢墙算法（简单版本）
	return g.tryWallKick(rotatedTetromino)
}

// tryWallKick 尝试踢墙算法
func (g *gameImpl) tryWallKick(rotatedTetromino Tetromino) bool {
	// 简单的踢墙尝试：左右各一格
	kickTests := []types.Position{
		{X: -1, Y: 0}, // 向左踢
		{X: 1, Y: 0},  // 向右踢
		{X: 0, Y: -1}, // 向上踢
	}

	originalPos := rotatedTetromino.GetPosition()

	for _, kick := range kickTests {
		testPos := types.Position{
			X: originalPos.X + kick.X,
			Y: originalPos.Y + kick.Y,
		}
		rotatedTetromino.SetPosition(testPos)

		if g.board.IsValidPosition(rotatedTetromino) {
			g.currentTetromino = rotatedTetromino
			return true
		}
	}

	return false
}

// DropTetromino 快速下落当前方块
func (g *gameImpl) DropTetromino() {
	if g.state != types.GameStatePlaying || g.currentTetromino == nil {
		return
	}

	// 持续向下移动直到无法移动
	for g.MoveTetromino(0, 1) {
		// 增加快速下落的分数奖励
		g.score += 2
	}

	// 立即固定方块
	g.lockCurrentTetromino()
}

// Update 更新游戏状态（用于游戏循环）
func (g *gameImpl) Update(deltaTime int) bool {
	if g.state != types.GameStatePlaying {
		return false
	}

	// 更新下落计时器
	g.dropTimer += deltaTime

	// 检查是否需要自动下落
	if g.dropTimer >= g.dropInterval {
		g.dropTimer = 0

		// 尝试向下移动
		if !g.MoveTetromino(0, 1) {
			// 无法向下移动，固定当前方块
			g.lockCurrentTetromino()
		}
	}

	return true
}

// lockCurrentTetromino 固定当前方块到棋盘
func (g *gameImpl) lockCurrentTetromino() {
	if g.currentTetromino == nil {
		return
	}

	// 将方块放置到棋盘上
	g.board.PlaceTetromino(g.currentTetromino)

	// 清除完整的行
	clearedLines := g.board.ClearLines()
	if clearedLines > 0 {
		g.updateScore(clearedLines)
		g.updateLevel()
	}

	// 检查游戏是否结束
	if g.board.IsGameOver() {
		g.state = types.GameStateGameOver
		return
	}

	// 生成新的方块
	g.spawnNewTetromino()
}

// updateScore 更新分数
func (g *gameImpl) updateScore(clearedLines int) {
	baseScore := clearedLines * g.config.ScorePerLine

	// 连击奖励
	multiplier := 1
	switch clearedLines {
	case 2:
		multiplier = 3 // 双消
	case 3:
		multiplier = 5 // 三消
	case 4:
		multiplier = 8 // 四消（Tetris）
	}

	g.score += baseScore * multiplier * g.level
	g.linesCleared += clearedLines
}

// updateLevel 更新等级和下落速度
func (g *gameImpl) updateLevel() {
	newLevel := (g.linesCleared / g.config.LinesPerLevel) + 1
	if newLevel > g.level {
		g.level = newLevel
		// 增加下落速度
		g.dropInterval = g.config.InitialDropInterval - (g.level-1)*50
		if g.dropInterval < g.config.MinDropInterval {
			g.dropInterval = g.config.MinDropInterval
		}
	}
}

// spawnNewTetromino 生成新的当前方块
func (g *gameImpl) spawnNewTetromino() {
	g.currentTetromino = g.nextTetromino
	g.generateNextTetromino()

	// 检查新方块是否可以放置
	if g.currentTetromino != nil && !g.board.IsValidPosition(g.currentTetromino) {
		g.state = types.GameStateGameOver
	}
}

// generateNextTetromino 生成下一个方块
func (g *gameImpl) generateNextTetromino() {
	g.nextTetromino = g.factory.CreateRandomTetromino()
}

// Reset 重置游戏
func (g *gameImpl) Reset() {
	g.state = types.GameStateMenu
	g.board.Clear()
	g.score = 0
	g.level = 1
	g.linesCleared = 0
	g.dropTimer = 0
	g.dropInterval = g.config.InitialDropInterval

	g.generateNextTetromino()
	g.spawnNewTetromino()
}
