import XCTest

final class TwoSumTests: XCTestCase {
    
    // MARK: - Solution
    
    func twoSum(_ nums: [Int], _ target: Int) -> [Int] {
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
    
    // MARK: - Test Cases
    
    func testCase1() {
        let result = twoSum([2, 7, 11, 15], 9)
        XCTAssertEqual(result, [0, 1])
    }
    
    func testCase2() {
        let result = twoSum([3, 2, 4], 6)
        XCTAssertEqual(result, [1, 2])
    }
    
    func testCase3() {
        let result = twoSum([3, 3], 6)
        XCTAssertEqual(result, [0, 1])
    }
} 