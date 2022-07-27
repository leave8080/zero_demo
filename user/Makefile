.PHONY: all build clean run check cover lint docker help
APP          := zero_demo
VERSION      := v1.0.0-beta
#COMMIT       := $(shell git rev-parse --short HEAD)
#CLI          := ./cmd/fogcloud
#DIST         := ./dist
BUILD_DATE   := `date +%FT%T%z`

#all: check build

api:
	goctl api go -api *.api -dir ./ --style=goZero
clean:
	@go clean
	#rm --fore "xx.out"

test:
	@go test

run:
	@go run user.go -f etc/user-api.yaml

help:
	@echo "make 格式化go代码 并编译生成二进制文件"
	@echo "make build 编译go代码生成二进制文件"
	@echo "make clean 清理中间目标文件"
	@echo "make test 执行测试case"
	@echo "make check 格式化go代码"
	@echo "make cover 检查测试覆盖率"
	@echo "make run 直接运行程序"
	@echo "make lint 执行代码检查"
	@echo "make docker 构建docker镜像"
