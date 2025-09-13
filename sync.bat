@echo off
REM Windows版本的同步脚本

echo 🚀 开始同步代码到远程仓库...
echo ==================================

REM 检查是否在Git仓库中
if not exist ".git" (
    echo ❌ 错误：当前目录不是Git仓库
    echo 请在项目根目录运行此脚本
    pause
    exit /b 1
)

REM 显示当前Git状态
echo 📋 当前Git状态：
git status --short

REM 检查是否有未提交的更改
git status --porcelain > temp_status.txt
for /f %%i in ("temp_status.txt") do set size=%%~zi
del temp_status.txt

if %size% gtr 0 (
    echo.
    echo 📝 检测到未提交的更改，准备提交...
    
    REM 获取提交信息
    if "%~1"=="" (
        set /p commit_message="💬 请输入提交信息: "
        if "!commit_message!"=="" (
            echo ❌ 提交信息不能为空
            pause
            exit /b 1
        )
    ) else (
        set commit_message=%~1
    )
    
    REM 添加所有更改
    echo 📦 添加文件到暂存区...
    git add .
    
    REM 提交更改
    echo 💾 提交更改...
    git commit -m "!commit_message!"
    
    echo ✅ 代码已提交到本地仓库
) else (
    echo ℹ️  没有检测到未提交的更改
)

REM 获取当前分支名
for /f "tokens=*" %%i in ('git branch --show-current') do set current_branch=%%i
echo.
echo 🌿 当前分支：%current_branch%

REM 推送到远程仓库
echo ⬆️  推送代码到远程仓库...
git push origin %current_branch%

if %errorlevel%==0 (
    echo ✅ 代码已成功推送到远程仓库
    
    REM 显示远程仓库信息
    for /f "tokens=*" %%i in ('git remote get-url origin') do set remote_url=%%i
    echo.
    echo 🔗 远程仓库地址：%remote_url%
    
    REM 显示最新提交信息
    echo.
    echo 📊 最新提交信息：
    git log --oneline -n 5
    
) else (
    echo ❌ 推送失败，请检查网络连接和仓库权限
    pause
    exit /b 1
)

echo.
echo 🎉 同步完成！
echo ==================================
pause