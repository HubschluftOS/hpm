
BIN_NAME=hpm

.PHONY: all build clean

all: build

check_uname:
	@if [ "$$(uname)" != "Linux" ]; then exit 1; fi

build: check_uname
	@pwd
	@go build -o $(BIN_NAME)
	@echo -e "\n${BIN_NAME} is built!"
	@mv $(BIN_NAME) bin/
	@grep -rIL .

clean:
	@go clean

help:
	@echo -e "Usage: make <target>\n"
	@echo -e "Targets:"
	@echo -e "  build    Build this project"
	@echo -e "  clean    Clean this project"
