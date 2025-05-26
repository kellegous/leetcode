#include <gtest/gtest.h>

#include <string_view>
#include <unordered_map>

namespace longestSubstringWithoutRepeatingCharacters {

int solve(const std::string_view& s) {
  if (s.empty()) {
    return 0;
  }

  int max_len = 0;

  // start if the offset at which we last saw a repeated character
  int start = 0;
  std::unordered_map<char, int> char_positions;

  for (int i = 0; i < s.length(); ++i) {
    auto c = s[i];
    auto it = char_positions.find(c);
    // have we seen c since the last repeated character?
    if (it != char_positions.end() && it->second >= start) {
      start = it->second + 1;
    }

    char_positions[c] = i;

    max_len = std::max(max_len, i - start + 1);
  }

  return max_len;
}

}  // namespace longestSubstringWithoutRepeatingCharacters

TEST(LongestSubstringWithoutRepeatingCharacters, TestCases) {
  EXPECT_EQ(longestSubstringWithoutRepeatingCharacters::solve("abcabcbb"), 3);
  EXPECT_EQ(longestSubstringWithoutRepeatingCharacters::solve("bbbbb"), 1);
  EXPECT_EQ(longestSubstringWithoutRepeatingCharacters::solve("pwwkew"), 3);
}
