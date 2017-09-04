TARGET=fake-poloniex-api
MAIN=main.go

default: build format run

build: $(MAIN)
	go build -o $(TARGET)

format: $(MAIN)
	gofmt -w $(MAIN)

run:
	./$(TARGET)

