#include <gtest/gtest.h>

#include <string_view>

namespace stringToIntegerAtoi {

int solve(const std::string_view& s) {
  int i = 0;
  int sign = 1;
  long result = 0;

  // discard any leading whitespace
  while (i < s.length() && std::isspace(s[i])) {
    i++;
  }

  // is this a sign character?
  if (i < s.length() && (s[i] == '-' || s[i] == '+')) {
    sign = (s[i] == '-') ? -1 : 1;
    i++;
  }

  // from now on, we're only interested in digits
  while (i < s.length() && std::isdigit(s[i])) {
    result = result * 10 + (s[i] - '0');

    // how to handle overflow/underflow cases? clamp, I guess.
    if (sign * result > std::numeric_limits<int>::max()) {
      return std::numeric_limits<int>::max();
    }
    if (sign * result < std::numeric_limits<int>::min()) {
      return std::numeric_limits<int>::min();
    }

    i++;
  }

  return sign * result;
}

}  // namespace stringToIntegerAtoi

TEST(StringToIntegerAtoi, TestCases) {
  EXPECT_EQ(stringToIntegerAtoi::solve("42"), 42);
  EXPECT_EQ(stringToIntegerAtoi::solve("   -042"), -42);
  EXPECT_EQ(stringToIntegerAtoi::solve("1337c0d3"), 1337);
  EXPECT_EQ(stringToIntegerAtoi::solve("0-1"), 0);
}