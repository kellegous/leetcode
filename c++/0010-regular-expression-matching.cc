#include <gtest/gtest.h>

#include <string_view>

namespace regularExpressionMatching {

bool solve(const std::string_view& s, const std::string_view& p) {
  // Base case: if pattern is empty, string should also be empty
  if (p.empty()) {
    return s.empty();
  }

  // Check if first character matches
  bool firstMatch = !s.empty() && (p[0] == s[0] || p[0] == '.');

  // If pattern has '*' as second character
  if (p.length() >= 2 && p[1] == '*') {
    // Try matching zero occurrences of the current pattern
    if (solve(s, p.substr(2))) {
      return true;
    }
    // If that doesn't work, try matching one or more occurrences
    return firstMatch && solve(s.substr(1), p);
  }

  // If no '*', just check first character match and continue with rest
  return firstMatch && solve(s.substr(1), p.substr(1));
}

}  // namespace regularExpressionMatching

TEST(RegularExpressionMatching, TestCases) {
  EXPECT_TRUE(regularExpressionMatching::solve("aa", "a*"));
  EXPECT_FALSE(regularExpressionMatching::solve("aa", "a"));
  EXPECT_TRUE(regularExpressionMatching::solve("ab", ".*"));
  EXPECT_TRUE(regularExpressionMatching::solve("aab", "c*a*b"));
  EXPECT_FALSE(regularExpressionMatching::solve("mississippi", "mis*is*p*."));
}