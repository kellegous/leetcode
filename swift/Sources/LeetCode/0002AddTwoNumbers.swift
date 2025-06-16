public class ListNode: Equatable, CustomStringConvertible {
    var val: Int
    var next: ListNode?
    
    init(_ val: Int, _ next: ListNode? = nil) {
        self.val = val
        self.next = next
    }
 
	static func add(l1: ListNode?, l2: ListNode?, carry: Int) -> ListNode? {
		if l1 == nil && l2 == nil {
			return nil
		}

		var value = carry
		var p1 = l1
		var p2 = l2

        if let l1 = l1 {
            value += l1.val
            p1 = l1.next
        }

        if let l2 = l2 {
            value += l2.val
            p2 = l2.next
        }

		let root = ListNode(value % 10)
        var carry = value / 10
		carry = value / 10

		var current = root

		while p1 != nil || p2 != nil || carry != 0 {
			var sum = carry
            if let p = p1 {
                sum += p.val
                p1 = p.next
            }
            if let p = p2 {
                sum += p.val
                p2 = p.next
            }

			carry = sum / 10
            let next = ListNode(sum % 10)
			current.next = next
			current = next
		}

		return root
	}
   
    static func create(_ slice: ArraySlice<Int>) -> ListNode? {
        guard let first = slice.first else { return nil }
        return ListNode(first, create(slice.dropFirst()))
    }
    
    public static func create(_ arr: Int...) -> ListNode? {
        .create(arr[...])
    }
    
    public static func == (lhs: ListNode, rhs: ListNode) -> Bool {
        guard lhs.val == rhs.val else { return false }
        
        switch (lhs.next, rhs.next) {
            case (.none, .none):
            return true
        case (.some(let l), .some(let r)):
            return l == r
        default:
            return false
        }
    }
    
    public var description: String {
        var current: ListNode? = self
        var result = ""
        while let node = current {
            result += "\(node.val)"
            if node.next != nil {
                result += " -> "
            }
            current = node.next
        }
        return "[\(result)]"
    }
}

public func addTwoNumbers(_ l1: ListNode?, _ l2: ListNode?) -> ListNode? {
    return ListNode.add(l1: l1, l2: l2, carry: 0)
}
