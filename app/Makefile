# Simple Makefile for a Go project

# Build the application
all: build

build: build-internal
	@echo "Building..."
	@go build -o main cmd/api/main.go

# Run the application
run: build-internal
	@go run cmd/api/main.go

# Test the application
test:
	@echo "Testing..."
	@go test . -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main
	@rm -f ./cmd/web/static/htmx.min.js
	@rm -f ./cmd/web/static/styles.css

# required build steps
build-internal:
	@templ generate
	@cp ./node_modules/htmx.org/dist/htmx.min.js ./cmd/web/static/htmx.min.js
	@npm run build

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: all build run test clean
