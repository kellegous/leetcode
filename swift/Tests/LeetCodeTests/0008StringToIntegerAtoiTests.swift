import Testing
import LeetCode

@Test("Test StringToIntegerAtoi", arguments: [
    ("42", 42),
    ("  -042", -42),
    ("1337c0d3", 1337),
    ("0-1", 0),
])
func testStringToIntegerAtoi(_ input: String, _ expected: Int) async throws {
    #expect(stringToIntegerAtoi(input) == expected)
}
