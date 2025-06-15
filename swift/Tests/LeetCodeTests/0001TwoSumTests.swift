import LeetCode
import Testing

@Test("Test TwoSum", arguments: [
    ([2, 7, 11, 15], 9, [0, 1]),
    ([3, 2, 4], 6, [1, 2]),
    ([3, 3], 6, [0, 1])
])
func testTwoSum(_ nums: [Int], _ target: Int, _ expected: [Int]) throws {
    #expect(LeetCode.twoSum(nums, target) == expected)
}
