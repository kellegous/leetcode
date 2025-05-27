#include <gtest/gtest.h>

#include <cassert>
#include <string>
#include <string_view>

namespace zigzagConversion {

std::string solve(const std::string_view& s, int num_rows) {
  assert(num_rows > 0);
  if (num_rows == 1 || s.length() <= num_rows) {
    return std::string(s);
  }

  std::string result;
  result.reserve(s.length());

  // Calculate the step size between characters in the same row
  int step = 2 * (num_rows - 1);

  // For each row
  for (int row = 0; row < num_rows; ++row) {
    // For each character in this row
    for (int i = row; i < s.length(); i += step) {
      result.push_back(s[i]);

      // For middle rows, add the diagonal character
      if (row > 0 && row < num_rows - 1) {
        int diagonal = i + step - 2 * row;
        if (diagonal < s.length()) {
          result.push_back(s[diagonal]);
        }
      }
    }
  }

  return result;
}

}  // namespace zigzagConversion

TEST(ZigzagConversion, TestCases) {
  // Basic test cases
  EXPECT_EQ(zigzagConversion::solve("PAYPALISHIRING", 3), "PAHNAPLSIIGYIR");
  EXPECT_EQ(zigzagConversion::solve("PAYPALISHIRING", 4), "PINALSIGYAHRPI");

  // Edge cases
  EXPECT_EQ(zigzagConversion::solve("A", 1), "A");
  EXPECT_EQ(zigzagConversion::solve("", 3), "");
  EXPECT_EQ(zigzagConversion::solve("ABC", 2), "ACB");
  EXPECT_EQ(zigzagConversion::solve("ABCD", 2), "ACBD");

  // Single row case
  EXPECT_EQ(zigzagConversion::solve("ABCDEF", 1), "ABCDEF");
}