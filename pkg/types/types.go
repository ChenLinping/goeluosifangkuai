// Package types 定义了俄罗斯方块游戏中使用的基础数据类型
package types

// Position 表示二维坐标位置
type Position struct {
	X, Y int
}

// Color 表示方块的颜色类型
type Color int

const (
	ColorEmpty Color = iota
	ColorI           // 青色 - I形方块
	ColorO           // 黄色 - O形方块
	ColorT           // 紫色 - T形方块
	ColorS           // 绿色 - S形方块
	ColorZ           // 红色 - Z形方块
	ColorJ           // 蓝色 - J形方块
	ColorL           // 橙色 - L形方块
)

// TetrominoType 表示俄罗斯方块的七种基本形状
type TetrominoType int

const (
	TetrominoI TetrominoType = iota
	TetrominoO
	TetrominoT
	TetrominoS
	TetrominoZ
	TetrominoJ
	TetrominoL
)

// Direction 表示旋转方向
type Direction int

const (
	DirectionNone Direction = iota
	DirectionLeft
	DirectionRight
)

// GameState 表示游戏状态
type GameState int

const (
	GameStateMenu GameState = iota
	GameStatePlaying
	GameStatePaused
	GameStateGameOver
)

// 游戏配置常量
const (
	BoardWidth  = 10 // 游戏棋盘宽度
	BoardHeight = 20 // 游戏棋盘高度

	// 游戏速度配置（毫秒）
	InitialDropInterval = 1000 // 初始下落间隔
	MinDropInterval     = 100  // 最小下落间隔
	FastDropInterval    = 50   // 快速下落间隔

	// 得分配置
	ScorePerLine         = 100 // 每消除一行的基础分数
	ScoreLevelMultiplier = 10  // 等级分数倍数
)
