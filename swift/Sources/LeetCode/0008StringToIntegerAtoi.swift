public func stringToIntegerAtoi(_ s: String) -> Int {
    let chars = Array(s)
    
    var i = 0
    var sign = 1
    var result  = 0

    // discard leading whitespace
    while i < chars.count, chars[i].isWhitespace {
        i += 1
    }
    
    // look for sign character
    if i < chars.count, chars[i] == "-" || chars[i] == "+" {
        sign = chars[i] == "-" ? -1 : 1
        i += 1
    }
    
    while i < chars.count, chars[i].isNumber {
        result = result * 10 + chars[i].wholeNumberValue!
        i += 1
    }
    
    return sign * result
}
