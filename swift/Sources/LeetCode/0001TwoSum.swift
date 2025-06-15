public func twoSum(_ nums: [Int], _ target: Int) -> [Int] {
    var indexMap: [Int: Int] = [:]
    
    for (i, num) in nums.enumerated() {
        let complement = target - num
        if let j = indexMap[complement] {
            return [j, i]
        }
        indexMap[num] = i
    }
    
    fatalError("No valid solution found")
}