public func reverseInteger(_ x: Int) -> Int {
    var x = x
    var result = 0
    
    while x != 0 {
        if result > Int.max / 10 || (result == Int.max / 10 && x > 7) {
            return 0
        }
        if result < Int.min / 10 || (result == Int.min / 10 && x < -8) {
            return 0
        }
        result = result * 10 + x % 10
        x /= 10
    }
    
    return result
}
