import Testing
import LeetCode

@Test("Test ZigzagConversion", arguments: [
    ("PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"),
    ("PAYPALISHIRING", 4, "PINALSIGYAHRPI"),
    ("A", 1, "A"),
    ("", 3, ""),
    ("ABC", 2, "ACB"),
    ("ABCD", 2, "ACBD"),
    ("ABCDEF", 1, "ABCDEF"),
])
func testZigzagConversion(_ s: String, _ rows: Int, _ expected: String) async throws {
    #expect(LeetCode.zigzagConversion(s, rows) == expected)
}
