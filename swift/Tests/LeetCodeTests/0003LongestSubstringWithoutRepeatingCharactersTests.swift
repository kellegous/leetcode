import Testing
import LeetCode

@Test("Test LongestSubstringWithoutRepeatingCharacters", arguments: [
    ("abcabcbb", 3),
    ("bbbbb", 1),
    ("pwwkew", 3)
])
func testLongestSubstringWithoutRepeatingCharacters(s: String, expected: Int) async throws {
    #expect(LeetCode.longestSubstringWithoutRepeatingCharacters(s) == expected)
}
