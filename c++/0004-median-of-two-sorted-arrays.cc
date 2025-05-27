#include <gtest/gtest.h>

#include <algorithm>
#include <limits>
#include <utility>
#include <vector>

namespace medianOfTwoSortedArrays {

double solve(const std::vector<int>& vals1, const std::vector<int>& vals2) {
  int m = vals1.size();
  int n = vals2.size();

  // ensure that nums1 is the smaller array
  if (m > n) {
    return solve(vals2, vals1);
  }

  int l = 0;
  int r = m;

  while (l <= r) {
    int partitionX = (l + r) / 2;
    int partitionY = (m + n + 1) / 2 - partitionX;

    // Find the elements around the partition
    int maxLeftX = (partitionX == 0) ? std::numeric_limits<int>::min()
                                     : vals1[partitionX - 1];
    int minRightX =
        (partitionX == m) ? std::numeric_limits<int>::max() : vals1[partitionX];

    int maxLeftY = (partitionY == 0) ? std::numeric_limits<int>::min()
                                     : vals2[partitionY - 1];
    int minRightY =
        (partitionY == n) ? std::numeric_limits<int>::max() : vals2[partitionY];

    // Check if we found the correct partition
    if (maxLeftX <= minRightY && maxLeftY <= minRightX) {
      // If total length is odd
      if ((m + n) % 2 == 1) {
        return std::max(maxLeftX, maxLeftY);
      }
      // If total length is even
      return (std::max(maxLeftX, maxLeftY) + std::min(minRightX, minRightY)) /
             2.0;
    } else if (maxLeftX > minRightY) {
      r = partitionX - 1;
    } else {
      l = partitionX + 1;
    }
  }

  std::unreachable();
}

}  // namespace medianOfTwoSortedArrays

TEST(MedianOfTwoSortedArrays, TestCases) {
  EXPECT_EQ(medianOfTwoSortedArrays::solve({1, 3}, {2}), 2.0);
  EXPECT_EQ(medianOfTwoSortedArrays::solve({1, 2}, {3, 4}), 2.5);

  EXPECT_EQ(medianOfTwoSortedArrays::solve({}, {1}), 1.0);
  EXPECT_EQ(medianOfTwoSortedArrays::solve({1}, {}), 1.0);
  EXPECT_EQ(medianOfTwoSortedArrays::solve({1, 2, 3}, {4, 5, 6}), 3.5);
  EXPECT_EQ(medianOfTwoSortedArrays::solve({1, 2, 3, 4}, {5, 6, 7, 8}), 4.5);
  EXPECT_EQ(medianOfTwoSortedArrays::solve({1, 3, 5, 7}, {2, 4, 6, 8}), 4.5);
  EXPECT_EQ(
      medianOfTwoSortedArrays::solve({1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}), 5.5
  );
}
