import Testing
import LeetCode

@Test("Test MedianOfTwoSortedArrays", arguments: [
    ([1, 3], [2], 2.0),
    ([1, 2], [3, 4], 2.5),
    ([], [1], 1.0),
    ([1], [], 1.0),
    ([1, 2, 3], [4, 5, 6], 3.5),
    ([1, 2, 3, 4], [5, 6, 7, 8], 4.5),
    ([1, 3, 5, 7], [2, 4, 6, 8], 4.5),
    ([1, 2, 3, 4, 5], [6, 7, 8, 9, 10], 5.5),
])
func testMedianOfTwoSortedArrays(a: [Int], b: [Int], expected: Double) throws {
    #expect(abs(medianOfTwoSortedArrays(a, b) - expected) < 0.000001)
}
