TARGET=fake-poloniex-api

default: build run

build:
	go build -o $(TARGET)

run:
	./$(TARGET)

