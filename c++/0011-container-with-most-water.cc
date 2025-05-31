#include <gtest/gtest.h>

namespace containerWithMostWater {

int solve(const std::vector<int>& height) {
  int l = 0;
  int r = height.size() - 1;
  int maxArea = 0;

  while (l < r) {
    // current area between l and r
    int area = std::min(height[l], height[r]) * (r - l);

    // keep up with the max area
    maxArea = std::max(maxArea, area);

    // move the smaller height pointer
    if (height[l] < height[r]) {
      l++;
    } else {
      r--;
    }
  }

  return maxArea;
}

}  // namespace containerWithMostWater

TEST(ContainerWithMostWater, TestCases) {
  EXPECT_EQ(containerWithMostWater::solve({1, 8, 6, 2, 5, 4, 8, 3, 7}), 49);
  EXPECT_EQ(containerWithMostWater::solve({1, 1}), 1);
}