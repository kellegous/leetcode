.PHONY: all test clean build setup reset

all: test

test:
	swift test

build:
	swift build

clean:
	swift package clean

setup:
	swift package resolve

reset: clean
	rm -rf .build

test-verbose:
	swift test --verbose

xcode:
	swift package generate-xcodeproj

# Help
help:
	@echo "Available targets:"
	@echo "  all         - Run all tests (default)"
	@echo "  test        - Run all tests"
	@echo "  build       - Build the project"
	@echo "  clean       - Clean build artifacts"
	@echo "  setup       - Resolve package dependencies"
	@echo "  reset       - Clean everything"
	@echo "  test-verbose - Run tests with verbose output"
	@echo "  xcode       - Generate Xcode project"
	@echo "  help        - Show this help message" 
