#!/bin/bash

# 俄罗斯方块游戏启动脚本

set -e

echo "🎮 俄罗斯方块游戏 (Native Version)"
echo "=================================="

# 检查Go是否安装
if ! command -v go &> /dev/null; then
    echo "❌ 错误：未检测到Go语言环境"
    echo "请先安装Go 1.19或更高版本"
    exit 1
fi

# 检查Go版本
GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
echo "✅ 检测到Go版本: $GO_VERSION"

echo ""
echo "🚀 正在启动Fyne原生桌面应用..."

# 检查是否已构建
if [ -f "bin/tetris-native" ]; then
    echo "🎮 正在运行游戏..."
    ./bin/tetris-native
else
    echo "📦 正在构建游戏..."
    make build
    echo "🎮 正在运行游戏..."
    ./bin/tetris-native
fi

echo ""
echo "🎯 游戏结束！感谢游玩俄罗斯方块！"