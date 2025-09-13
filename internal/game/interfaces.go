// Package game 定义了俄罗斯方块游戏的核心接口和数据结构
package game

import "goeluosifangkuai/pkg/types"

// Tetromino 表示一个俄罗斯方块
type Tetromino interface {
	// GetType 返回方块类型
	GetType() types.TetrominoType

	// GetColor 返回方块颜色
	GetColor() types.Color

	// GetPosition 返回方块当前位置
	GetPosition() types.Position

	// SetPosition 设置方块位置
	SetPosition(pos types.Position)

	// GetBlocks 返回方块的所有组成块的相对位置
	GetBlocks() []types.Position

	// Rotate 旋转方块
	Rotate(direction types.Direction) Tetromino

	// Clone 克隆方块
	Clone() Tetromino
}

// Board 表示游戏棋盘
type Board interface {
	// GetWidth 返回棋盘宽度
	GetWidth() int

	// GetHeight 返回棋盘高度
	GetHeight() int

	// GetCell 获取指定位置的单元格颜色
	GetCell(x, y int) types.Color

	// SetCell 设置指定位置的单元格颜色
	SetCell(x, y int, color types.Color)

	// IsValidPosition 检查方块是否可以放置在指定位置
	IsValidPosition(tetromino Tetromino) bool

	// PlaceTetromino 将方块放置到棋盘上
	PlaceTetromino(tetromino Tetromino)

	// ClearLines 清除已满的行，返回清除的行数
	ClearLines() int

	// IsGameOver 检查游戏是否结束
	IsGameOver() bool

	// Clear 清空棋盘
	Clear()
}

// Game 表示游戏主控制器
type Game interface {
	// GetState 返回当前游戏状态
	GetState() types.GameState

	// SetState 设置游戏状态
	SetState(state types.GameState)

	// GetBoard 返回游戏棋盘
	GetBoard() Board

	// GetCurrentTetromino 返回当前正在下落的方块
	GetCurrentTetromino() Tetromino

	// GetNextTetromino 返回下一个方块
	GetNextTetromino() Tetromino

	// GetScore 返回当前分数
	GetScore() int

	// GetLevel 返回当前等级
	GetLevel() int

	// GetLinesCleared 返回已消除的行数
	GetLinesCleared() int

	// MoveTetromino 移动当前方块
	MoveTetromino(dx, dy int) bool

	// RotateTetromino 旋转当前方块
	RotateTetromino(direction types.Direction) bool

	// DropTetromino 快速下落当前方块
	DropTetromino()

	// Update 更新游戏状态（用于游戏循环）
	Update(deltaTime int) bool

	// Reset 重置游戏
	Reset()
}

// GameStats 游戏统计信息
type GameStats struct {
	Score        int
	Level        int
	LinesCleared int
	ElapsedTime  int // 游戏时间（秒）
}
