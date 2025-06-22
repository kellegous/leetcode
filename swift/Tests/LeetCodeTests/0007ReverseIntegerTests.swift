import Testing
import LeetCode

@Test("Test ReverseInteger", arguments: [
    (123, 321),
    (-123, -321),
    (120, 21),
])
func testReverseInteger(_ x: Int, expected: Int) {
    #expect(LeetCode.reverseInteger(x) == expected)
}
