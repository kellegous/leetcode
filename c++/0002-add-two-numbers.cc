#include <cassert>
#include <initializer_list>
#include <memory>
#include <string>
#include <utility>

#include <gtest/gtest.h>

namespace addTwoNumbers {

struct ListNode {
  int val;
  std::unique_ptr<ListNode> next;

  ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, std::unique_ptr<ListNode> next)
      : val(x), next(std::move(next)) {}

  [[nodiscard]] static std::unique_ptr<ListNode>
  create(const std::initializer_list<int> xs) {
    return create(xs.begin(), xs.end());
  }

  [[nodiscard]] static std::unique_ptr<ListNode>
  add(const std::unique_ptr<ListNode> &l1, const std::unique_ptr<ListNode> &l2,
      int carry) {
    std::unique_ptr<ListNode> dummy = std::make_unique<ListNode>(0);
    ListNode *current = dummy.get();

    const ListNode *p1 = l1.get();
    const ListNode *p2 = l2.get();

    while (p1 || p2 || carry) {
      int sum = carry;
      if (p1) {
        sum += p1->val;
        p1 = p1->next.get();
      }
      if (p2) {
        sum += p2->val;
        p2 = p2->next.get();
      }

      carry = sum / 10;
      current->next = std::make_unique<ListNode>(sum % 10);
      current = current->next.get();
    }

    return std::move(dummy->next);
  }

  [[nodiscard]] std::vector<int> values() const {
    std::vector<int> xs;
    const ListNode *curr = this;
    while (curr != nullptr) {
      xs.push_back(curr->val);
      curr = curr->next.get();
    }
    return xs;
  }

  [[nodiscard]] std::string ToString() const {
    std::string str;
    const ListNode *curr = this;
    while (curr != nullptr) {
      if (!str.empty()) {
        str.append(" -> ");
      }
      str.append(std::to_string(curr->val));
      curr = curr->next.get();
    }
    return str;
  }

private:
  [[nodiscard]] static std::unique_ptr<ListNode>
  create(std::initializer_list<int>::const_iterator head,
         std::initializer_list<int>::const_iterator tail) {
    if (head == tail) {
      return nullptr;
    }
    return std::make_unique<ListNode>(*head, create(std::next(head), tail));
  }
};

[[nodiscard]] std::unique_ptr<ListNode>
solve(const std::unique_ptr<ListNode> &l1,
      const std::unique_ptr<ListNode> &l2) {
  return ListNode::add(l1, l2, 0);
}

} // namespace addTwoNumbers

using addTwoNumbers::ListNode;

TEST(addTwoNumbers, TestCases) {
  EXPECT_EQ(addTwoNumbers::solve(ListNode::create({2, 4, 3}),
                                 ListNode::create({5, 6, 4}))
                ->values(),
            (std::vector<int>{7, 0, 8}));
  EXPECT_EQ(addTwoNumbers::solve(ListNode::create({0}), ListNode::create({0}))
                ->values(),
            (std::vector<int>{0}));
  EXPECT_EQ(addTwoNumbers::solve(ListNode::create({9, 9, 9, 9, 9, 9, 9}),
                                 ListNode::create({9, 9, 9, 9}))
                ->values(),
            (std::vector<int>{8, 9, 9, 9, 0, 0, 0, 1}));
}