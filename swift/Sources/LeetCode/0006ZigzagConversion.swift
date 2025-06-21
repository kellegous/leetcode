public func zigzagConversion(_ s: String, _ rows: Int) -> String {
    precondition(rows > 0)
    guard rows != 1 && s.count > rows else { return s }

    var result = ""
    result.reserveCapacity(s.count)

    let step = 2 * (rows - 1)

    let chars = Array(s)
    
    for row in 0..<rows {
        for i in stride(from: row, to: chars.count, by: step) {
            result.append(chars[i])
           
            // for middle rows, add the diagonal character
            if row > 0 && row < rows - 1 {
                let diagonal = i + step - 2 * row
                if diagonal < chars.count {
                    result.append(chars[diagonal])
                }
            }
        }
    }
    
    return result
}
