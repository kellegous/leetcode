#include <gtest/gtest.h>

#include <limits>

namespace reverseInteger {

int solve(int x) {
  int result = 0;

  while (x != 0) {
    // Check for overflow before multiplying
    if (result > std::numeric_limits<int>::max() / 10 ||
        (result == std::numeric_limits<int>::max() / 10 && x % 10 > 7)) {
      return 0;
    }
    if (result < std::numeric_limits<int>::min() / 10 ||
        (result == std::numeric_limits<int>::min() / 10 && x % 10 < -8)) {
      return 0;
    }

    result = result * 10 + x % 10;
    x /= 10;
  }

  return result;
}

}  // namespace reverseInteger

TEST(ReverseInteger, TestCases) {
  EXPECT_EQ(reverseInteger::solve(123), 321);
  EXPECT_EQ(reverseInteger::solve(-123), -321);
  EXPECT_EQ(reverseInteger::solve(120), 21);
}
