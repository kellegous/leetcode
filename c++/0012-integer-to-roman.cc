#include <gtest/gtest.h>

#include <string>
#include <utility>
#include <vector>

namespace integerToRoman {

std::string solve(int num) {
  std::string result;

  // Define the Roman numeral symbols and their values
  const std::vector<std::pair<int, std::string>> values = {
      {1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"}, {100, "C"},
      {90, "XC"},  {50, "L"},   {40, "XL"}, {10, "X"},   {9, "IX"},
      {5, "V"},    {4, "IV"},   {1, "I"}
  };

  // Iterate through the values from largest to smallest
  for (const auto& [value, symbol] : values) {
    // subtract the value from num until it's less than the value
    while (num >= value) {
      result += symbol;
      num -= value;
    }
  }

  return result;
}

}  // namespace integerToRoman

TEST(IntegerToRoman, TestCases) {
  EXPECT_EQ(integerToRoman::solve(3), "III");
  EXPECT_EQ(integerToRoman::solve(4), "IV");
  EXPECT_EQ(integerToRoman::solve(9), "IX");
  EXPECT_EQ(integerToRoman::solve(3749), "MMMDCCXLIX");
  EXPECT_EQ(integerToRoman::solve(58), "LVIII");
  EXPECT_EQ(integerToRoman::solve(1994), "MCMXCIV");
}