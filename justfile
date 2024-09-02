# Print this help message.
help:
    @just --list

# Compile the project into a single binary.
build:
    go build \
    -o ./bin/tg \
    -ldflags="-X 'main.version=$(git describe --tags --always)'" \
    ./cmd/app/main.go

# Install the binary for the current user.
[unix]
install: build
    mkdir -p ~/.local/bin
    cp --remove-destination ./bin/tg ~/.local/bin/tg
