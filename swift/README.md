# LeetCode Swift Solutions

This directory contains LeetCode problem solutions implemented in Swift, organized similarly to the C++ solutions.

## Structure

- Each LeetCode problem is implemented in a separate Swift file in `Tests/LeetCodeSwiftTests/`
- Each file contains both the solution and test cases in a single file
- Files are named using the pattern: `{problem-number}{ProblemName}Tests.swift`
- All tests can be run with a single command

## File Structure

```
swift/
├── Package.swift                           # Swift Package Manager configuration
├── Makefile                               # Build and test commands
├── Sources/LeetCodeSwift/                # Main target (placeholder)
│   └── LeetCodeSwift.swift
└── Tests/LeetCodeSwiftTests/             # Problem solutions and tests
    ├── 0001TwoSumTests.swift
    ├── 0009PalindromeNumberTests.swift
    └── ... (more problems)
```

## Usage

### Running All Tests

```bash
# Using Make (recommended)
make test

# Or directly using Swift
swift test
```

### Running a Specific Problem

```bash
# Run only Two Sum tests
make test-single CLASS=TwoSumTests

# Or using Swift directly
swift test --filter TwoSumTests
```

### Other Commands

```bash
make build          # Build the project
make clean          # Clean build artifacts
make setup          # Resolve dependencies
make test-verbose   # Run tests with verbose output
make xcode          # Generate Xcode project
make help           # Show available commands
```

## Adding New Problems

1. Create a new file in `Tests/LeetCodeSwiftTests/` following the naming pattern
2. Use the following template:

```swift
import XCTest

final class {ProblemName}Tests: XCTestCase {

    // MARK: - Solution

    func solutionFunction() {
        // Your solution here
    }

    // MARK: - Test Cases

    func testCase1() {
        // Your test cases here
        XCTAssertEqual(expected, actual)
    }
}
```

## Example

See `0001TwoSumTests.swift` for a complete example that mirrors the C++ structure.

## Requirements

- Swift 5.9 or later
- macOS 10.15 or later
