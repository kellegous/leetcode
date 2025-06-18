public func longestPalindromeSubstring(_ s: String) -> String {
    guard !s.isEmpty else { return "" }
    
    var start = 0
    var maxLength = 1

    // avoid looking for the start of chars w/ index
    let chars = Array(s)
    
    for i in 0..<s.count {
        // check for odd length palindromes
        let (l1, r1) = expandAroundCenter(chars, left: i, right: i)
        let len1 = r1 - l1 + 1
        if len1 > maxLength {
            maxLength = len1
            start = l1
        }
       
        // check for even length palindromes
        let (l2, r2) = expandAroundCenter(chars, left: i, right: i+1)
        let len2 = r2 - l2 + 1
        if len2 > maxLength {
            maxLength = len2
            start = l2
        }
    }

    let startIdx = s.index(s.startIndex, offsetBy: start)
    return String(s[startIdx..<s.index(startIdx, offsetBy: maxLength)])
}

func expandAroundCenter(_ s: [Character], left: Int, right: Int) -> (Int, Int) {
    var left = left;
    var right = right;
    while left >= 0 && right < s.count && s[left] == s[right] {
        left -= 1
        right += 1
    }
    return (left + 1, right - 1)
}
