// Package game 实现游戏棋盘的逻辑
package game

import (
	"goeluosifangkuai/pkg/types"
)

// board 是 Board 接口的具体实现
type board struct {
	width  int
	height int
	cells  [][]types.Color
}

// NewBoard 创建新的游戏棋盘
func NewBoard(width, height int) Board {
	cells := make([][]types.Color, height)
	for i := range cells {
		cells[i] = make([]types.Color, width)
		for j := range cells[i] {
			cells[i][j] = types.ColorEmpty
		}
	}

	return &board{
		width:  width,
		height: height,
		cells:  cells,
	}
}

// GetWidth 返回棋盘宽度
func (b *board) GetWidth() int {
	return b.width
}

// GetHeight 返回棋盘高度
func (b *board) GetHeight() int {
	return b.height
}

// GetCell 获取指定位置的单元格颜色
func (b *board) GetCell(x, y int) types.Color {
	if x < 0 || x >= b.width || y < 0 || y >= b.height {
		return types.ColorEmpty
	}
	return b.cells[y][x]
}

// SetCell 设置指定位置的单元格颜色
func (b *board) SetCell(x, y int, color types.Color) {
	if x >= 0 && x < b.width && y >= 0 && y < b.height {
		b.cells[y][x] = color
	}
}

// IsValidPosition 检查方块是否可以放置在指定位置
func (b *board) IsValidPosition(tetromino Tetromino) bool {
	position := tetromino.GetPosition()
	blocks := tetromino.GetBlocks()

	for _, block := range blocks {
		x := position.X + block.X
		y := position.Y + block.Y

		// 检查边界
		if x < 0 || x >= b.width || y >= b.height {
			return false
		}

		// 检查是否与已有方块冲突（忽略棋盘上方的区域）
		if y >= 0 && b.cells[y][x] != types.ColorEmpty {
			return false
		}
	}

	return true
}

// PlaceTetromino 将方块放置到棋盘上
func (b *board) PlaceTetromino(tetromino Tetromino) {
	position := tetromino.GetPosition()
	blocks := tetromino.GetBlocks()
	color := tetromino.GetColor()

	for _, block := range blocks {
		x := position.X + block.X
		y := position.Y + block.Y

		// 只在有效范围内放置方块
		if x >= 0 && x < b.width && y >= 0 && y < b.height {
			b.cells[y][x] = color
		}
	}
}

// ClearLines 清除已满的行，返回清除的行数
func (b *board) ClearLines() int {
	clearedLines := 0

	// 从下往上检查每一行
	for y := b.height - 1; y >= 0; y-- {
		if b.isLineFull(y) {
			b.clearLine(y)
			clearedLines++
			y++ // 重新检查这一行（因为上面的行下移了）
		}
	}

	return clearedLines
}

// isLineFull 检查指定行是否已满
func (b *board) isLineFull(y int) bool {
	if y < 0 || y >= b.height {
		return false
	}

	for x := 0; x < b.width; x++ {
		if b.cells[y][x] == types.ColorEmpty {
			return false
		}
	}

	return true
}

// clearLine 清除指定行并将上面的行下移
func (b *board) clearLine(lineY int) {
	// 将指定行上面的所有行向下移动一行
	for y := lineY; y > 0; y-- {
		for x := 0; x < b.width; x++ {
			b.cells[y][x] = b.cells[y-1][x]
		}
	}

	// 清空顶部行
	for x := 0; x < b.width; x++ {
		b.cells[0][x] = types.ColorEmpty
	}
}

// IsGameOver 检查游戏是否结束
func (b *board) IsGameOver() bool {
	// 检查顶部几行是否有方块
	for y := 0; y < 4; y++ {
		for x := 0; x < b.width; x++ {
			if b.cells[y][x] != types.ColorEmpty {
				return true
			}
		}
	}

	return false
}

// Clear 清空棋盘
func (b *board) Clear() {
	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			b.cells[y][x] = types.ColorEmpty
		}
	}
}

// GetAllCells 获取所有单元格（用于渲染）
func (b *board) GetAllCells() [][]types.Color {
	// 返回副本以避免外部修改
	result := make([][]types.Color, b.height)
	for i := range result {
		result[i] = make([]types.Color, b.width)
		copy(result[i], b.cells[i])
	}
	return result
}
