import Testing
import LeetCode


@Test("Test AddTwoNumbers", arguments: [
    (ListNode.create(2, 4, 3), ListNode.create(5, 6, 4), ListNode.create(7, 0, 8)),
    (ListNode.create(0), ListNode.create(0), ListNode.create(0)),
    (ListNode.create(9,9,9,9,9,9,9), ListNode.create(9,9,9,9), ListNode.create(8, 9, 9, 9, 0, 0, 0, 1))
])
func testAddTwoNumbers(a: ListNode?, b: ListNode?, expected: ListNode?) async throws {
    #expect(addTwoNumbers(a, b) == expected)
}
