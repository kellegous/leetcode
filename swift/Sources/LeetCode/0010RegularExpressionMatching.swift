public func regularExpressionMatching(_ s: String, _ p: String) -> Bool {
    return regularExpressionMatching(s[...], p[...])
}

private func regularExpressionMatching(_ s: Substring, _ p: Substring) -> Bool {
    // Base case: if pattern is empty, string should also be empty
    if p.isEmpty {
        return s.isEmpty
    }
    
    // Check if first character matches
    let firstMatch = !s.isEmpty && (p.first! == s.first! || p.first! == ".")
    
    // If pattern has '*' as second character
    if p.count >= 2 && p[p.index(p.startIndex, offsetBy: 1)] == "*" {
        // Try matching zero occurrences of the current pattern
        if regularExpressionMatching(s, p.dropFirst(2)) {
            return true
        }
        // If that doesn't work, try matching one or more occurrences
        return firstMatch && regularExpressionMatching(s.dropFirst(), p)
    }
    
    // If no '*', just check first character match and continue with rest
    return firstMatch && regularExpressionMatching(s.dropFirst(), p.dropFirst())
}

