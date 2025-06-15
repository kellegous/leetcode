public func isPalindrome(_ x: Int) -> Bool {
    if x < 0 {
        return false
    }

    if x < 10 {
        return true
    }

    var original = x
    var reversed = 0

    while original > 0 {
        reversed = reversed * 10 + original % 10
        original /= 10
    }

    return x == reversed
}
