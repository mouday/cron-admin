
# 编译到 Linux
.PHONY: build-linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go dist -o ./dist/cron-admin-linux ./src/main.go 

# 编译到 macOS
# make build-darwin
.PHONY: build-darwin
build-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./dist/cron-admin-darwin ./src/main.go

# 编译到 windows
.PHONY: build-windows
build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./dist/cron-admin-windows.exe ./src/main.go 

# 编译到 全部平台
.PHONY: build-all
build-all:
	make clean
	mkdir -p ./dist
	make build-linux
	make build-darwin
	make build-windows

.PHONY: clean
clean:
	rm -rf ./dist

.PHONY: dev
dev:
	# go run ./src/main.go
	air