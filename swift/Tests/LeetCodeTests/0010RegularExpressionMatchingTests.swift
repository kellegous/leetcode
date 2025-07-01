import Testing
import LeetCode

@Test("Test RegularExpressionMatching", arguments: [
    ("aa", "a*", true),
    ("aa", "a", false),
    ("ab", ".*", true),
    ("aab", "c*a*b", true),
    ("mississippi", "mis*is*p*.", false),
])
func testRegularExpressionMatching(_ s: String, _ p: String, _ expected: Bool) async throws {
    #expect(regularExpressionMatching(s, p) == expected)
}
