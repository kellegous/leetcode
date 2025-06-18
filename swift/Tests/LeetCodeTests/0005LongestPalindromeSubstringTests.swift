import Testing
import LeetCode

@Test("Test LongestPalindromicSubstring", arguments: [
    ("babad", "bab"),
    ("cbbd", "bb"),
    ("a", "a"),
    ("", ""),
])
func testLongestPalindromeSubstring(_ input: String, _ expected: String) async throws {
    #expect(LeetCode.longestPalindromeSubstring(input) == expected)
}
