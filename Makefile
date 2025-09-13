# Makefile for Tetris Game (Native Version Only)

# 变量定义
BINARY_NAME=tetris-native
MAIN_PACKAGE=./cmd/tetris-native
BUILD_DIR=./bin
GO_VERSION=1.19

# 默认目标
.PHONY: help
help: ## 显示帮助信息
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## 构建游戏
	@echo "Building Tetris Native Game..."
	@mkdir -p $(BUILD_DIR)
	@go build -ldflags="-s -w" -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

.PHONY: run
run: ## 运行游戏
	@echo "Running Tetris Native Game..."
	@go run $(MAIN_PACKAGE)

.PHONY: clean
clean: ## 清理构建文件
	@echo "Cleaning build files..."
	@rm -rf $(BUILD_DIR)
	@go clean

.PHONY: test
test: ## 运行所有测试
	@echo "Running tests..."
	@go test -v ./...

.PHONY: test-coverage
test-coverage: ## 运行测试并生成覆盖率报告
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

.PHONY: lint
lint: ## 运行代码检查
	@echo "Running linter..."
	@gofmt -l .
	@go vet ./...

.PHONY: fmt
fmt: ## 格式化代码
	@echo "Formatting code..."
	@go fmt ./...

.PHONY: mod-tidy
mod-tidy: ## 整理模块依赖
	@echo "Tidying modules..."
	@go mod tidy

.PHONY: install
install: build ## 安装到 GOPATH/bin
	@echo "Installing $(BINARY_NAME)..."
	@cp $(BUILD_DIR)/$(BINARY_NAME) $(GOPATH)/bin/

.PHONY: dist
dist: ## 构建多平台发布版本
	@echo "Building distribution packages..."
	@mkdir -p $(BUILD_DIR)/dist
	
	# Linux AMD64
	@GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/dist/$(BINARY_NAME)-linux-amd64 $(MAIN_PACKAGE)
	
	# Windows AMD64
	@GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/dist/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PACKAGE)
	
	# macOS AMD64
	@GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/dist/$(BINARY_NAME)-darwin-amd64 $(MAIN_PACKAGE)
	
	# macOS ARM64 (Apple Silicon)
	@GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o $(BUILD_DIR)/dist/$(BINARY_NAME)-darwin-arm64 $(MAIN_PACKAGE)
	
	@echo "Distribution packages built in $(BUILD_DIR)/dist/"

.PHONY: check
check: fmt lint test ## 运行所有检查

.PHONY: dev
dev: ## 开发模式运行（带实时重载）
	@echo "Starting development mode..."
	@which air > /dev/null || (echo "Installing air..." && go install github.com/cosmtrek/air@latest)
	@air

.PHONY: deps
deps: ## 安装开发依赖
	@echo "Installing development dependencies..."
	@go install github.com/cosmtrek/air@latest

# 显示项目信息
.PHONY: info
info: ## 显示项目信息
	@echo "Project: Tetris Game (Native Version)"
	@echo "Go version: $(shell go version)"
	@echo "Module: $(shell go list -m)"
	@echo "Build dir: $(BUILD_DIR)"
	@echo "Binary name: $(BINARY_NAME)"