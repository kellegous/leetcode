public func regularExpressionMatching(_ s: String, _ p: String) -> Bool {
    let sChars = Array(s)
    let pChars = Array(p)
    let sLen = sChars.count
    let pLen = pChars.count
    
    // dp[i][j] represents whether s[0...i-1] matches p[0...j-1]
    var dp = Array(repeating: Array(repeating: false, count: pLen + 1), count: sLen + 1)
    
    // Base case: empty string matches empty pattern
    dp[0][0] = true
    
    // Handle patterns like "a*b*c*" that can match empty string
    if pLen >= 2 {
        for j in 2...pLen {
            if pChars[j-1] == "*" {
                dp[0][j] = dp[0][j-2]
            }
        }
    }
    
    // Fill the DP table
    for i in 1...sLen {
        for j in 1...pLen {
            let currentChar = sChars[i-1]
            let patternChar = pChars[j-1]
            
            if patternChar == "*" {
                // Current pattern character is '*'
                let prevPatternChar = pChars[j-2]
                
                // Option 1: Match zero occurrences of the character before '*'
                dp[i][j] = dp[i][j-2]
                
                // Option 2: Match one or more occurrences if characters match
                if !dp[i][j] && (prevPatternChar == currentChar || prevPatternChar == ".") {
                    dp[i][j] = dp[i-1][j]
                }
            } else {
                // Current pattern character is not '*'
                // Characters must match (either exact match or '.')
                if patternChar == currentChar || patternChar == "." {
                    dp[i][j] = dp[i-1][j-1]
                }
            }
        }
    }
    
    return dp[sLen][pLen]
}

