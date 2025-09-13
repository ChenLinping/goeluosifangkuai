// Package game_test 提供游戏逻辑的单元测试
package game

import (
	"goeluosifangkuai/pkg/types"
	"testing"
)

func TestNewBoard(t *testing.T) {
	board := NewBoard(10, 20)

	if board.GetWidth() != 10 {
		t.Errorf("期望棋盘宽度为 10，实际为 %d", board.GetWidth())
	}

	if board.GetHeight() != 20 {
		t.Errorf("期望棋盘高度为 20，实际为 %d", board.GetHeight())
	}

	// 检查初始状态所有单元格都是空的
	for y := 0; y < board.GetHeight(); y++ {
		for x := 0; x < board.GetWidth(); x++ {
			if board.GetCell(x, y) != types.ColorEmpty {
				t.Errorf("位置 (%d, %d) 应该是空的", x, y)
			}
		}
	}
}

func TestBoardSetAndGetCell(t *testing.T) {
	board := NewBoard(10, 20)

	// 设置一个单元格
	board.SetCell(5, 10, types.ColorI)

	// 验证单元格值
	if board.GetCell(5, 10) != types.ColorI {
		t.Errorf("期望位置 (5, 10) 的颜色为 %v，实际为 %v", types.ColorI, board.GetCell(5, 10))
	}

	// 测试边界外访问
	if board.GetCell(-1, 0) != types.ColorEmpty {
		t.Errorf("边界外访问应该返回空颜色")
	}

	if board.GetCell(15, 0) != types.ColorEmpty {
		t.Errorf("边界外访问应该返回空颜色")
	}
}

func TestNewTetromino(t *testing.T) {
	// 测试创建 I 形方块
	tetromino := NewTetromino(types.TetrominoI)

	if tetromino.GetType() != types.TetrominoI {
		t.Errorf("期望方块类型为 %v，实际为 %v", types.TetrominoI, tetromino.GetType())
	}

	if tetromino.GetColor() != types.ColorI {
		t.Errorf("期望方块颜色为 %v，实际为 %v", types.ColorI, tetromino.GetColor())
	}

	// 验证初始位置
	pos := tetromino.GetPosition()
	expectedX := types.BoardWidth / 2
	if pos.X != expectedX {
		t.Errorf("期望初始X位置为 %d，实际为 %d", expectedX, pos.X)
	}

	if pos.Y != 0 {
		t.Errorf("期望初始Y位置为 0，实际为 %d", pos.Y)
	}
}

func TestTetrominoRotation(t *testing.T) {
	tetromino := NewTetromino(types.TetrominoT)
	originalBlocks := tetromino.GetBlocks()

	// 向右旋转
	rotatedTetromino := tetromino.Rotate(types.DirectionRight)
	rotatedBlocks := rotatedTetromino.GetBlocks()

	// 验证方块确实发生了旋转（块的数量应该相同但位置不同）
	if len(rotatedBlocks) != len(originalBlocks) {
		t.Errorf("旋转后方块数量不应该改变")
	}

	// 验证至少有一个块的位置发生了变化
	different := false
	for i, block := range rotatedBlocks {
		if i < len(originalBlocks) && (block.X != originalBlocks[i].X || block.Y != originalBlocks[i].Y) {
			different = true
			break
		}
	}

	if !different {
		t.Errorf("旋转后方块位置应该发生变化")
	}
}

func TestTetrominoFactory(t *testing.T) {
	factory := NewTetrominoFactory()

	// 创建随机方块
	tetromino := factory.CreateRandomTetromino()
	if tetromino == nil {
		t.Errorf("工厂应该创建有效的方块")
	}

	// 创建特定类型的方块
	specificTetromino := factory.CreateSpecificTetromino(types.TetrominoO)
	if specificTetromino.GetType() != types.TetrominoO {
		t.Errorf("期望创建 O 形方块，实际为 %v", specificTetromino.GetType())
	}
}

func TestGameCreation(t *testing.T) {
	config := DefaultGameConfig()
	game := NewGame(config)

	if game.GetState() != types.GameStateMenu {
		t.Errorf("期望初始游戏状态为菜单，实际为 %v", game.GetState())
	}

	if game.GetScore() != 0 {
		t.Errorf("期望初始分数为 0，实际为 %d", game.GetScore())
	}

	if game.GetLevel() != 1 {
		t.Errorf("期望初始等级为 1，实际为 %d", game.GetLevel())
	}

	if game.GetBoard() == nil {
		t.Errorf("游戏应该有有效的棋盘")
	}
}
