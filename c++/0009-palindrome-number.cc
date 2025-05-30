#include <gtest/gtest.h>

namespace palindromeNumber {

bool solve(int x) {
  // Negative numbers are not palindromes
  if (x < 0) {
    return false;
  }

  // Single digit numbers are palindromes
  if (x < 10) {
    return true;
  }

  // If number ends with 0, it can't be a palindrome
  // (except for 0 itself, which we handled above)
  if (x % 10 == 0) {
    return false;
  }

  int reversed = 0;
  // We only need to reverse half the number
  while (x > reversed) {
    reversed = reversed * 10 + x % 10;
    x /= 10;
  }

  // When the original number is less than or equal to the reversed number,
  // we've processed half or more of the digits
  // For even length numbers: x == reversed
  // For odd length numbers: x == reversed/10
  return x == reversed || x == reversed / 10;
}

}  // namespace palindromeNumber

TEST(PalindromeNumber, TestCases) {
  EXPECT_TRUE(palindromeNumber::solve(121));
  EXPECT_FALSE(palindromeNumber::solve(-121));
  EXPECT_FALSE(palindromeNumber::solve(10));
  EXPECT_TRUE(palindromeNumber::solve(0));
}