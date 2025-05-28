#include <gtest/gtest.h>

#include <string>
#include <string_view>

namespace longestPalindromicSubstring {

std::pair<int, int> expand_around_center(
    const std::string_view& s, int left, int right
) {
  while (left >= 0 && right < s.length() && s[left] == s[right]) {
    left--;
    right++;
  }
  return {left + 1, right - 1};
}

std::string solve(const std::string_view& s) {
  if (s.empty()) {
    return "";
  }

  int start = 0;
  int max_len = 1;

  for (int i = 0; i < s.length(); ++i) {
    // Check for odd length palindromes
    auto [l1, r1] = expand_around_center(s, i, i);
    int len1 = r1 - l1 + 1;
    if (len1 > max_len) {
      max_len = len1;
      start = l1;
    }

    // Check for even length palindromes
    auto [l2, r2] = expand_around_center(s, i, i + 1);
    int len2 = r2 - l2 + 1;
    if (len2 > max_len) {
      max_len = len2;
      start = l2;
    }
  }

  return std::string(s.substr(start, max_len));
}

}  // namespace longestPalindromicSubstring

TEST(LongestPalindromicSubstring, TestCases) {
  EXPECT_EQ(longestPalindromicSubstring::solve("babad"), "bab");
  EXPECT_EQ(longestPalindromicSubstring::solve("cbbd"), "bb");

  EXPECT_EQ(longestPalindromicSubstring::solve("a"), "a");
  EXPECT_EQ(longestPalindromicSubstring::solve(""), "");
}