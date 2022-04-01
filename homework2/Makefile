help: Makefile
	@echo "Выберите команду для выполенния:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## run: запустить приложение
run:
	go run ./main/main.go


## test: запустить unit-тесты
test:
	go test -cover ./sort/*

## build: запустить сборку приложения для текущей ОС
build:
	go build -o ./main ./main/main.go

## build-windows: запустить сборку приложения для Windows
build-windows:
	GOOS=windows go build -o ./main ./main/main.go
	file ./main/main.exe

## build-linux: запустить сборку приложения для Linux
build-linux:
	GOOS=linux go build -o ./main ./main/main.go
	file ./main/main