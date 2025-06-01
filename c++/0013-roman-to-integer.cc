#include <gtest/gtest.h>

#include <cassert>
#include <string_view>
#include <unordered_map>

namespace romanToInteger {

static const std::unordered_map<char, int> VALUES = {{'I', 1},   {'V', 5},
                                                     {'X', 10},  {'L', 50},
                                                     {'C', 100}, {'D', 500},
                                                     {'M', 1000}};
int solve(const std::string_view& s) {
  assert(s.length() > 0);
  int n = s.length();

  int result = 0;
  for (int i = 1; i < n; i++) {
    auto a = VALUES.at(s[i - 1]);
    auto b = VALUES.at(s[i]);
    if (a < b) {
      result -= a;
    } else {
      result += a;
    }
  }

  result += VALUES.at(s[n - 1]);

  return result;
}

}  // namespace romanToInteger

TEST(RomanToInteger, TestCases) {
  EXPECT_EQ(romanToInteger::solve("III"), 3);
  EXPECT_EQ(romanToInteger::solve("IV"), 4);
  EXPECT_EQ(romanToInteger::solve("IX"), 9);
  EXPECT_EQ(romanToInteger::solve("LVIII"), 58);
  EXPECT_EQ(romanToInteger::solve("MCMXCIV"), 1994);
}
