#include <gtest/gtest.h>

#include <string>
#include <string_view>

namespace longestCommonPrefix {

std::string solve(const std::vector<std::string_view>& strs) {
  if (strs.empty()) {
    return "";
  }

  const std::string_view& first = strs[0];
  for (size_t i = 0, n = first.length(); i < n; i++) {
    char current_char = first[i];
    for (size_t j = 1, m = strs.size(); j < m; j++) {
      if (i >= strs[j].length() || strs[j][i] != current_char) {
        return std::string(first.substr(0, i));
      }
    }
  }

  return std::string(first);
}

}  // namespace longestCommonPrefix

TEST(LongestCommonPrefix, TestCases) {
  EXPECT_EQ(longestCommonPrefix::solve({"flower", "flow", "flight"}), "fl");
  EXPECT_EQ(longestCommonPrefix::solve({"dog", "racecar", "car"}), "");
}