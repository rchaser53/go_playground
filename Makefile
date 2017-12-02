.PHONY: build clean

default: build

OUTPUT_DIR=build/goPlayground

clean:
	rm -rf build && mkdir build && touch build/.gitkeep

build: clean
	go build -o $(OUTPUT_DIR)