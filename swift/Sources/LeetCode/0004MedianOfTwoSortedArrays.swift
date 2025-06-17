public func medianOfTwoSortedArrays(_ a: [Int], _ b: [Int]) -> Double {
    let m = a.count
    let n = b.count
    guard m <= n else { return medianOfTwoSortedArrays(b, a)}
    
    var l = 0
    var r = m
    
    while l <= r {
        let partitionX = (l + r) / 2
        let partitionY = (m + n + 1) / 2 - partitionX
  
        let maxLeftX = (partitionX == 0) ? Int.min : a[partitionX - 1]
        let minRightX = (partitionX == m) ? Int.max : a[partitionX]
        let maxLeftY = (partitionY == 0) ? Int.min : b[partitionY - 1]
        let minRightY = (partitionY == n) ? Int.max : b[partitionY]
        
        if maxLeftX <= minRightY && maxLeftY <= minRightX {
            if (m + n) % 2 == 1 {
                return Double(max(maxLeftX, maxLeftY))
            }
            
            return Double(max(maxLeftX, maxLeftY) + min(minRightX, minRightY)) / 2.0
        } else if maxLeftX > minRightY {
            r = partitionX - 1
        } else {
            l = partitionX + 1
        }
    }
    fatalError("unreachable")
}
