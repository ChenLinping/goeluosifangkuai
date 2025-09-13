// Package game 提供俄罗斯方块的工厂函数
package game

import (
	"goeluosifangkuai/pkg/types"
	"math/rand"
	"time"
)

// TetrominoFactory 俄罗斯方块工厂
type TetrominoFactory struct {
	random *rand.Rand
}

// NewTetrominoFactory 创建新的方块工厂
func NewTetrominoFactory() *TetrominoFactory {
	return &TetrominoFactory{
		random: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// CreateRandomTetromino 创建随机的俄罗斯方块
func (f *TetrominoFactory) CreateRandomTetromino() Tetromino {
	tetrominoTypes := []types.TetrominoType{
		types.TetrominoI,
		types.TetrominoO,
		types.TetrominoT,
		types.TetrominoS,
		types.TetrominoZ,
		types.TetrominoJ,
		types.TetrominoL,
	}

	randomType := tetrominoTypes[f.random.Intn(len(tetrominoTypes))]
	return NewTetromino(randomType)
}

// CreateSpecificTetromino 创建指定类型的俄罗斯方块
func (f *TetrominoFactory) CreateSpecificTetromino(tetrominoType types.TetrominoType) Tetromino {
	return NewTetromino(tetrominoType)
}
