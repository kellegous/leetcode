#include <cassert>
#include <unordered_map>
#include <vector>

#include <gtest/gtest.h>

namespace twoSum {

std::vector<int> solve(const std::vector<int> &nums, int target) {
  std::unordered_map<int, int> idx;
  for (int i = 0, n = nums.size(); i < n; i++) {
    auto j = idx.find(target - nums[i]);
    if (j != idx.end()) {
      return {j->second, i};
    }
    idx[nums[i]] = i;
  }

  assert(false);
}

} // namespace twoSum

TEST(TwoSum, TestCases) {
  EXPECT_EQ(twoSum::solve({2, 7, 11, 15}, 9), (std::vector<int>{0, 1}));
  EXPECT_EQ(twoSum::solve({3, 2, 4}, 6), (std::vector<int>{1, 2}));
  EXPECT_EQ(twoSum::solve({3, 3}, 6), (std::vector<int>{0, 1}));
}
