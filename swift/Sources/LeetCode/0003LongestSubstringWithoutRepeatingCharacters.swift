
public func longestSubstringWithoutRepeatingCharacters(_ s: String) -> Int {
    guard !s.isEmpty else {
        return 0
    }
    
    var maxLength = 0
    var start = 0
    var charPositions = Dictionary<Character, Int>()
    
    for (index, char) in s.enumerated() {
        let pos = charPositions[char, default: -1]
        if (pos != -1 && pos >= start) {
            start = pos + 1
        }
        
        charPositions[char] = index
        
        maxLength = max(maxLength, index - start + 1)
    }
    
    return maxLength
}
