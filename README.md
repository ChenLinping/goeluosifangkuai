# 🎮 俄罗斯方块游戏 (Tetris Game)

一个用Go语言开发的现代化俄罗斯方块游戏，采用Fyne框架构建原生桌面应用。

## ✨ 特性

- 🚀 **原生桌面应用** - 基于Fyne框架的跨平台GUI
- 🎯 **经典游戏玩法** - 完整的俄罗斯方块游戏逻辑
- 🌈 **精美界面** - 现代化的用户界面设计
- ⌨️ **键盘控制** - 流畅的键盘操作体验
- 📊 **游戏统计** - 实时分数、等级和行数统计
- 🎵 **游戏状态** - 开始、暂停、重新开始功能

## 🚀 快速开始

### 环境要求

- Go 1.19 或更高版本
- 支持的操作系统：Windows、macOS、Linux

### 安装依赖

```bash
# 克隆项目
git clone <repository-url>
cd goeluosifangkuai

# 下载依赖
go mod tidy
```

### 运行游戏

```bash
# 方式一：使用启动脚本（推荐）
./play.sh

# 方式二：使用 Makefile
make run

# 方式三：直接运行
go run cmd/tetris-native/main.go
```

### 构建游戏

```bash
# 构建可执行文件
make build

# 运行已构建的游戏
./bin/tetris-native
```

## 🎮 游戏控制

| 按键 | 功能 |
|------|------|
| **A** | 向左移动 |
| **D** | 向右移动 |
| **S** | 向下移动 |
| **W** | 顺时针旋转 |
| **空格** | 快速下降 |
| **P** | 暂停/继续 |

## 🏗️ 项目结构

```
.
├── cmd/
│   └── tetris-native/          # 原生桌面版本入口
├── internal/
│   ├── fyneui/                 # Fyne GUI界面组件
│   └── game/                   # 核心游戏逻辑
├── pkg/
│   └── types/                  # 类型定义
├── Makefile                    # 构建脚本
├── play.sh                     # 启动脚本
└── README.md                   # 项目说明
```

## 🛠️ 开发

### 运行测试

```bash
# 运行所有测试
make test

# 生成测试覆盖率报告
make test-coverage
```

### 代码格式化

```bash
# 格式化代码
make fmt

# 运行代码检查
make lint

# 运行所有检查
make check
```

### 构建多平台版本

```bash
# 构建多平台发布版本
make dist
```

## 📝 License

MIT License

## 🤝 贡献

欢迎提交Issue和Pull Request！

---

**享受游戏！🎉**