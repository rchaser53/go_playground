.PHONY: build clean

default: build

OUTPUT_PATH=build/goPlayground

clean:
	rm -rf build && mkdir build && touch build/.gitkeep

build: clean
	go build -o $(OUTPUT_PATH)

run:
	$(OUTPUT_PATH)