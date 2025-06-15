import Testing
import LeetCode

@Test("Test Palindrome Number", arguments: [
    (121, true),
    (-121, false),
    (10, false),
    (0, true),
    (1, true)
])
func testPalindromeNumber(_ input: Int, _ expected: Bool) throws {
    #expect(LeetCode.isPalindrome(input) == expected)
}
