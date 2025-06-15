.PHONY: all test clean build setup reset

# Default target
all: test

# Run all tests
test:
	swift test

# Build the project without running tests
build:
	swift build

# Clean build artifacts
clean:
	swift package clean

# Setup the project (resolve dependencies)
setup:
	swift package resolve

# Reset everything
reset: clean
	rm -rf .build

# Run tests with verbose output
test-verbose:
	swift test --verbose

# Run a specific test class (usage: make test-single CLASS=TwoSumTests)
test-single:
	swift test --filter $(CLASS)

# Generate Xcode project
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
	@echo "  test-single  - Run a specific test class (usage: make test-single CLASS=TwoSumTests)"
	@echo "  xcode       - Generate Xcode project"
	@echo "  help        - Show this help message" 