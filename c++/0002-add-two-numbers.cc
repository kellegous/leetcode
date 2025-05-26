#include <gtest/gtest.h>

#include <cassert>
#include <initializer_list>
#include <memory>
#include <string>
#include <utility>

namespace addTwoNumbers {

struct ListNode {
  int val;
  std::unique_ptr<ListNode> next;

  ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, std::unique_ptr<ListNode> next)
      : val(x), next(std::move(next)) {}

  [[nodiscard]] static std::unique_ptr<ListNode> create(
      const std::initializer_list<int> xs
  ) {
    return create(xs.begin(), xs.end());
  }

  [[nodiscard]] static std::unique_ptr<ListNode> add(
      const std::unique_ptr<ListNode>& l1, const std::unique_ptr<ListNode>& l2,
      int carry
  ) {
    if (!l1 && !l2) {
      return nullptr;
    }

    auto value = carry;
    const ListNode* p1 = nullptr;
    const ListNode* p2 = nullptr;
    if (l1) {
      value += l1->val;
      p1 = l1->next.get();
    }
    if (l2) {
      value += l2->val;
      p2 = l2->next.get();
    }
    std::unique_ptr<ListNode> root = std::make_unique<ListNode>(value % 10);
    carry = value / 10;

    ListNode* current = root.get();

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

    return std::move(root);
  }

  static bool Equal(const ListNode* l1, const ListNode* l2) {
    while (l1 || l2) {
      if (!l1 || !l2) {
        return false;
      }
      if (l1->val != l2->val) {
        return false;
      }
      l1 = l1->next.get();
      l2 = l2->next.get();
    }
    return true;
  }

  [[nodiscard]] std::string ToString() const {
    std::string str;
    const ListNode* curr = this;
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
  [[nodiscard]] static std::unique_ptr<ListNode> create(
      std::initializer_list<int>::const_iterator head,
      std::initializer_list<int>::const_iterator tail
  ) {
    if (head == tail) {
      return nullptr;
    }
    return std::make_unique<ListNode>(*head, create(std::next(head), tail));
  }
};

[[nodiscard]] std::unique_ptr<ListNode> solve(
    const std::unique_ptr<ListNode>& l1, const std::unique_ptr<ListNode>& l2
) {
  return ListNode::add(l1, l2, 0);
}

}  // namespace addTwoNumbers

using addTwoNumbers::ListNode;

TEST(addTwoNumbers, TestCases) {
  EXPECT_TRUE(ListNode::Equal(
      addTwoNumbers::solve(
          ListNode::create({2, 4, 3}), ListNode::create({5, 6, 4})
      )
          .get(),
      ListNode::create({7, 0, 8}).get()
  ));

  EXPECT_TRUE(ListNode::Equal(
      addTwoNumbers::solve(ListNode::create({0}), ListNode::create({0})).get(),
      ListNode::create({0}).get()
  ));

  EXPECT_TRUE(ListNode::Equal(
      addTwoNumbers::solve(
          ListNode::create({9, 9, 9, 9, 9, 9, 9}),
          ListNode::create({9, 9, 9, 9})
      )
          .get(),
      ListNode::create({8, 9, 9, 9, 0, 0, 0, 1}).get()
  ));
}