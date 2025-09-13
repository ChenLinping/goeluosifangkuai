#!/bin/bash

# 同步代码到远程GitHub仓库脚本

set -e

echo "🚀 开始同步代码到远程仓库..."
echo "=================================="

# 检查是否在Git仓库中
if [ ! -d ".git" ]; then
    echo "❌ 错误：当前目录不是Git仓库"
    echo "请在项目根目录运行此脚本"
    exit 1
fi

# 显示当前Git状态
echo "📋 当前Git状态："
git status --short

# 检查是否有未提交的更改
if [ -n "$(git status --porcelain)" ]; then
    echo ""
    echo "📝 检测到未提交的更改，准备提交..."
    
    # 获取提交信息
    if [ -z "$1" ]; then
        echo "💬 请输入提交信息（或按Ctrl+C取消）："
        read -r commit_message
        if [ -z "$commit_message" ]; then
            echo "❌ 提交信息不能为空"
            exit 1
        fi
    else
        commit_message="$1"
    fi
    
    # 添加所有更改
    echo "📦 添加文件到暂存区..."
    git add .
    
    # 提交更改
    echo "💾 提交更改..."
    git commit -m "$commit_message"
    
    echo "✅ 代码已提交到本地仓库"
else
    echo "ℹ️  没有检测到未提交的更改"
fi

# 获取当前分支名
current_branch=$(git branch --show-current)
echo ""
echo "🌿 当前分支：$current_branch"

# 推送到远程仓库
echo "⬆️  推送代码到远程仓库..."
if git push origin "$current_branch"; then
    echo "✅ 代码已成功推送到远程仓库"
    
    # 显示远程仓库信息
    remote_url=$(git remote get-url origin)
    echo ""
    echo "🔗 远程仓库地址：$remote_url"
    
    # 显示最新提交信息
    echo ""
    echo "📊 最新提交信息："
    git log --oneline -n 5
    
else
    echo "❌ 推送失败，请检查网络连接和仓库权限"
    exit 1
fi

echo ""
echo "🎉 同步完成！"
echo "=================================="