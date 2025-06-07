#include <gtest/gtest.h>

#include <algorithm>
#include <vector>

namespace threeSum {

struct Triple {
  int i, j, k;

  Triple() = default;
  Triple(int i, int j, int k) : i(i), j(j), k(k) {}

  bool operator==(const Triple& other) const = default;
  bool operator<(const Triple& other) const {
    return i < other.i || (i == other.i && j < other.j) ||
           (i == other.i && j == other.j && k < other.k);
  }
};

std::vector<Triple> solve(std::vector<int>& nums) {
  std::vector<Triple> result;
  std::sort(nums.begin(), nums.end());

  for (int i = 0; i < nums.size() - 2; i++) {
    // Skip duplicates for i
    if (i > 0 && nums[i] == nums[i - 1]) {
      continue;
    }

    int left = i + 1;
    int right = nums.size() - 1;

    while (left < right) {
      int sum = nums[i] + nums[left] + nums[right];

      if (sum < 0) {
        left++;
      } else if (sum > 0) {
        right--;
      } else {
        result.push_back(Triple(nums[i], nums[left], nums[right]));

        // Skip duplicates for left
        while (left < right && nums[left] == nums[left + 1]) {
          left++;
        }
        // Skip duplicates for right
        while (left < right && nums[right] == nums[right - 1]) {
          right--;
        }

        left++;
        right--;
      }
    }
  }

  return result;
}

}  // namespace threeSum

TEST(ThreeSum, TestCases) {
  {
    std::vector<int> nums{-1, 0, 1, 2, -1, -4};
    auto result = threeSum::solve(nums);
    std::sort(result.begin(), result.end());
    EXPECT_EQ(
        threeSum::solve(nums),
        std::vector<threeSum::Triple>(
            {threeSum::Triple{-1, -1, 2}, threeSum::Triple{-1, 0, 1}}
        )
    );
  }

  {
    std::vector<int> nums{0, 0, 0};
    auto result = threeSum::solve(nums);
    std::sort(result.begin(), result.end());
    EXPECT_EQ(
        result, std::vector<threeSum::Triple>({threeSum::Triple{0, 0, 0}})
    );
  }

  {
    std::vector<int> nums{0, 1, 1};
    auto result = threeSum::solve(nums);
    std::sort(result.begin(), result.end());
    EXPECT_EQ(result, std::vector<threeSum::Triple>({}));
  }
}