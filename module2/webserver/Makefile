run:
	go run ./

build_image:
	docker build -t kuokou1/go-web:v0.1  ./ --no-cache

clean:
	go clean

fmt:
	go fmt ./
	go vet ./

help:
	@echo "make run: 运行程序"
	@echo "make build_image: 打包Docker镜像"
	@echo "make clean: 清理"
	@echo "make fmt: 格式化"
