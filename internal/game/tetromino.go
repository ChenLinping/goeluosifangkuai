// Package game 实现俄罗斯方块的具体结构和旋转逻辑
package game

import (
	"goeluosifangkuai/pkg/types"
)

// tetromino 是 Tetromino 接口的具体实现
type tetromino struct {
	tetrominoType types.TetrominoType
	color         types.Color
	position      types.Position
	rotation      int                // 当前旋转状态 (0-3)
	blocks        [][]types.Position // 四个旋转状态的方块位置
}

// GetType 返回方块类型
func (t *tetromino) GetType() types.TetrominoType {
	return t.tetrominoType
}

// GetColor 返回方块颜色
func (t *tetromino) GetColor() types.Color {
	return t.color
}

// GetPosition 返回方块当前位置
func (t *tetromino) GetPosition() types.Position {
	return t.position
}

// SetPosition 设置方块位置
func (t *tetromino) SetPosition(pos types.Position) {
	t.position = pos
}

// GetBlocks 返回方块的所有组成块的相对位置
func (t *tetromino) GetBlocks() []types.Position {
	if t.rotation >= 0 && t.rotation < len(t.blocks) {
		return t.blocks[t.rotation]
	}
	return t.blocks[0]
}

// Rotate 旋转方块
func (t *tetromino) Rotate(direction types.Direction) Tetromino {
	newTetromino := t.Clone().(*tetromino)

	switch direction {
	case types.DirectionLeft:
		newTetromino.rotation = (newTetromino.rotation + 3) % 4
	case types.DirectionRight:
		newTetromino.rotation = (newTetromino.rotation + 1) % 4
	}

	return newTetromino
}

// Clone 克隆方块
func (t *tetromino) Clone() Tetromino {
	newBlocks := make([][]types.Position, len(t.blocks))
	for i, rotationBlocks := range t.blocks {
		newBlocks[i] = make([]types.Position, len(rotationBlocks))
		copy(newBlocks[i], rotationBlocks)
	}

	return &tetromino{
		tetrominoType: t.tetrominoType,
		color:         t.color,
		position:      t.position,
		rotation:      t.rotation,
		blocks:        newBlocks,
	}
}

// 方块形状定义：每个方块有4个旋转状态，每个状态包含4个位置
var tetrominoShapes = map[types.TetrominoType][][]types.Position{
	// I 形方块
	types.TetrominoI: {
		{{X: -1, Y: 0}, {X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}}, // 水平
		{{X: 0, Y: -1}, {X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}}, // 垂直
		{{X: -1, Y: 0}, {X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}}, // 水平
		{{X: 0, Y: -1}, {X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}}, // 垂直
	},

	// O 形方块
	types.TetrominoO: {
		{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}},
		{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}},
		{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}},
		{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}},
	},

	// T 形方块
	types.TetrominoT: {
		{{X: -1, Y: 0}, {X: 0, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 1}},  // ┴
		{{X: 0, Y: -1}, {X: 0, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}}, // ├
		{{X: -1, Y: 0}, {X: 0, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: -1}}, // ┬
		{{X: 0, Y: -1}, {X: 0, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 0}},  // ┤
	},

	// S 形方块
	types.TetrominoS: {
		{{X: -1, Y: 1}, {X: 0, Y: 1}, {X: 0, Y: 0}, {X: 1, Y: 0}},
		{{X: 0, Y: -1}, {X: 0, Y: 0}, {X: 1, Y: 0}, {X: 1, Y: 1}},
		{{X: -1, Y: 1}, {X: 0, Y: 1}, {X: 0, Y: 0}, {X: 1, Y: 0}},
		{{X: 0, Y: -1}, {X: 0, Y: 0}, {X: 1, Y: 0}, {X: 1, Y: 1}},
	},

	// Z 形方块
	types.TetrominoZ: {
		{{X: -1, Y: 0}, {X: 0, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}},
		{{X: 1, Y: -1}, {X: 1, Y: 0}, {X: 0, Y: 0}, {X: 0, Y: 1}},
		{{X: -1, Y: 0}, {X: 0, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}},
		{{X: 1, Y: -1}, {X: 1, Y: 0}, {X: 0, Y: 0}, {X: 0, Y: 1}},
	},

	// J 形方块
	types.TetrominoJ: {
		{{X: -1, Y: 0}, {X: 0, Y: 0}, {X: 1, Y: 0}, {X: -1, Y: 1}},
		{{X: 0, Y: -1}, {X: 0, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: -1}},
		{{X: -1, Y: 0}, {X: 0, Y: 0}, {X: 1, Y: 0}, {X: 1, Y: -1}},
		{{X: 0, Y: -1}, {X: 0, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}},
	},

	// L 形方块
	types.TetrominoL: {
		{{X: -1, Y: 0}, {X: 0, Y: 0}, {X: 1, Y: 0}, {X: 1, Y: 1}},
		{{X: 0, Y: -1}, {X: 0, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 1}},
		{{X: -1, Y: 0}, {X: 0, Y: 0}, {X: 1, Y: 0}, {X: -1, Y: -1}},
		{{X: 0, Y: -1}, {X: 0, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: -1}},
	},
}

// 方块类型到颜色的映射
var tetrominoColors = map[types.TetrominoType]types.Color{
	types.TetrominoI: types.ColorI,
	types.TetrominoO: types.ColorO,
	types.TetrominoT: types.ColorT,
	types.TetrominoS: types.ColorS,
	types.TetrominoZ: types.ColorZ,
	types.TetrominoJ: types.ColorJ,
	types.TetrominoL: types.ColorL,
}

// NewTetromino 创建新的俄罗斯方块
func NewTetromino(tetrominoType types.TetrominoType) Tetromino {
	shapes, exists := tetrominoShapes[tetrominoType]
	if !exists {
		// 默认创建 I 形方块
		shapes = tetrominoShapes[types.TetrominoI]
		tetrominoType = types.TetrominoI
	}

	color, exists := tetrominoColors[tetrominoType]
	if !exists {
		color = types.ColorI
	}

	// 深拷贝形状数据
	blocks := make([][]types.Position, len(shapes))
	for i, shape := range shapes {
		blocks[i] = make([]types.Position, len(shape))
		copy(blocks[i], shape)
	}

	return &tetromino{
		tetrominoType: tetrominoType,
		color:         color,
		position:      types.Position{X: types.BoardWidth / 2, Y: 0},
		rotation:      0,
		blocks:        blocks,
	}
}
