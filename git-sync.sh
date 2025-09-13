#!/bin/bash

# 高级Git同步脚本
# 支持多种同步模式和选项

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# 脚本配置
SCRIPT_VERSION="1.0.0"
DEFAULT_BRANCH="master"
REMOTE_NAME="origin"

# 显示帮助信息
show_help() {
    echo -e "${BLUE}Git同步脚本 v${SCRIPT_VERSION}${NC}"
    echo "用法: $0 [选项] [提交信息]"
    echo ""
    echo "选项:"
    echo "  -h, --help          显示此帮助信息"
    echo "  -f, --force         强制推送（使用 git push --force-with-lease）"
    echo "  -p, --pull          推送前先拉取远程更新"
    echo "  -b, --branch <分支>  指定要推送的分支（默认：当前分支）"
    echo "  -r, --remote <远程>  指定远程仓库名称（默认：origin）"
    echo "  -d, --dry-run       只显示将要执行的操作，不实际执行"
    echo "  -q, --quiet         静默模式，减少输出"
    echo "  -s, --status        只显示Git状态，不执行同步"
    echo ""
    echo "示例:"
    echo "  $0                           # 基本同步"
    echo "  $0 \"修复了下一个方块预览功能\"   # 带提交信息的同步"
    echo "  $0 -p \"更新功能\"              # 先拉取再推送"
    echo "  $0 -f                        # 强制推送"
    echo "  $0 -b develop                # 推送到develop分支"
}

# 日志函数
log_info() { echo -e "${BLUE}ℹ️  $1${NC}"; }
log_success() { echo -e "${GREEN}✅ $1${NC}"; }
log_warning() { echo -e "${YELLOW}⚠️  $1${NC}"; }
log_error() { echo -e "${RED}❌ $1${NC}"; }

# 检查Git仓库
check_git_repo() {
    if [ ! -d ".git" ]; then
        log_error "当前目录不是Git仓库"
        log_info "请在项目根目录运行此脚本"
        exit 1
    fi
}

# 检查远程仓库连接
check_remote() {
    if ! git remote get-url "$REMOTE_NAME" >/dev/null 2>&1; then
        log_error "远程仓库 '$REMOTE_NAME' 不存在"
        log_info "请先添加远程仓库：git remote add $REMOTE_NAME <仓库URL>"
        exit 1
    fi
}

# 显示Git状态
show_status() {
    log_info "当前Git状态："
    git status --short
    
    local current_branch=$(git branch --show-current)
    log_info "当前分支：$current_branch"
    
    local remote_url=$(git remote get-url "$REMOTE_NAME")
    log_info "远程仓库：$remote_url"
    
    # 显示本地与远程的差异
    if git rev-parse --verify "$REMOTE_NAME/$current_branch" >/dev/null 2>&1; then
        local ahead=$(git rev-list --count "$REMOTE_NAME/$current_branch..HEAD")
        local behind=$(git rev-list --count "HEAD..$REMOTE_NAME/$current_branch")
        
        if [ "$ahead" -gt 0 ]; then
            log_warning "本地领先远程 $ahead 个提交"
        fi
        if [ "$behind" -gt 0 ]; then
            log_warning "本地落后远程 $behind 个提交"
        fi
        if [ "$ahead" -eq 0 ] && [ "$behind" -eq 0 ]; then
            log_success "本地与远程同步"
        fi
    else
        log_warning "远程分支不存在，这将是首次推送"
    fi
}

# 拉取远程更新
pull_remote() {
    log_info "拉取远程更新..."
    if git pull "$REMOTE_NAME" "$TARGET_BRANCH"; then
        log_success "远程更新已拉取"
    else
        log_error "拉取远程更新失败"
        exit 1
    fi
}

# 提交更改
commit_changes() {
    local commit_message="$1"
    
    if [ -n "$(git status --porcelain)" ]; then
        log_info "检测到未提交的更改"
        
        if [ -z "$commit_message" ]; then
            echo -n "💬 请输入提交信息：" >&2
            read -r commit_message
            if [ -z "$commit_message" ]; then
                log_error "提交信息不能为空"
                exit 1
            fi
        fi
        
        if [ "$DRY_RUN" = "true" ]; then
            log_info "[模拟] 将添加所有文件到暂存区"
            log_info "[模拟] 将提交：$commit_message"
        else
            log_info "添加文件到暂存区..."
            git add .
            
            log_info "提交更改..."
            git commit -m "$commit_message"
            log_success "代码已提交到本地仓库"
        fi
    else
        log_info "没有检测到未提交的更改"
    fi
}

# 推送到远程
push_to_remote() {
    local push_args="$REMOTE_NAME $TARGET_BRANCH"
    if [ "$FORCE_PUSH" = "true" ]; then
        push_args="$push_args --force-with-lease"
    fi
    
    if [ "$DRY_RUN" = "true" ]; then
        log_info "[模拟] 将推送到远程：git push $push_args"
    else
        log_info "推送代码到远程仓库..."
        if git push $push_args; then
            log_success "代码已成功推送到远程仓库"
            
            if [ "$QUIET" != "true" ]; then
                local remote_url=$(git remote get-url "$REMOTE_NAME")
                echo ""
                log_info "远程仓库地址：$remote_url"
                
                echo ""
                log_info "最新提交信息："
                git log --oneline -n 5 --color=always
            fi
        else
            log_error "推送失败，请检查网络连接和仓库权限"
            exit 1
        fi
    fi
}

# 解析命令行参数
FORCE_PUSH=false
PULL_FIRST=false
DRY_RUN=false
QUIET=false
STATUS_ONLY=false
COMMIT_MESSAGE=""
TARGET_BRANCH=""

while [[ $# -gt 0 ]]; do
    case $1 in
        -h|--help)
            show_help
            exit 0
            ;;
        -f|--force)
            FORCE_PUSH=true
            shift
            ;;
        -p|--pull)
            PULL_FIRST=true
            shift
            ;;
        -b|--branch)
            TARGET_BRANCH="$2"
            shift 2
            ;;
        -r|--remote)
            REMOTE_NAME="$2"
            shift 2
            ;;
        -d|--dry-run)
            DRY_RUN=true
            shift
            ;;
        -q|--quiet)
            QUIET=true
            shift
            ;;
        -s|--status)
            STATUS_ONLY=true
            shift
            ;;
        -*)
            log_error "未知选项：$1"
            show_help
            exit 1
            ;;
        *)
            COMMIT_MESSAGE="$1"
            shift
            ;;
    esac
done

# 主程序
main() {
    if [ "$QUIET" != "true" ]; then
        echo -e "${PURPLE}🚀 Git同步脚本 v${SCRIPT_VERSION}${NC}"
        echo "=================================="
    fi
    
    # 基本检查
    check_git_repo
    check_remote
    
    # 设置目标分支
    if [ -z "$TARGET_BRANCH" ]; then
        TARGET_BRANCH=$(git branch --show-current)
    fi
    
    # 显示状态
    if [ "$QUIET" != "true" ] || [ "$STATUS_ONLY" = "true" ]; then
        show_status
    fi
    
    # 如果只是查看状态，则退出
    if [ "$STATUS_ONLY" = "true" ]; then
        exit 0
    fi
    
    # 拉取远程更新（如果启用）
    if [ "$PULL_FIRST" = "true" ]; then
        pull_remote
    fi
    
    # 提交更改
    commit_changes "$COMMIT_MESSAGE"
    
    # 推送到远程
    push_to_remote
    
    if [ "$QUIET" != "true" ]; then
        echo ""
        log_success "同步完成！"
        echo "=================================="
    fi
}

# 运行主程序
main "$@"