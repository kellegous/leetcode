import XCTest

final class PalindromeNumberTests: XCTestCase {
    
    // MARK: - Solution
    
    func isPalindrome(_ x: Int) -> Bool {
        // Negative numbers are not palindromes
        if x < 0 {
            return false
        }
        
        // Single digit numbers are palindromes
        if x < 10 {
            return true
        }
        
        // Reverse the number and compare
        var original = x
        var reversed = 0
        
        while original > 0 {
            reversed = reversed * 10 + original % 10
            original /= 10
        }
        
        return x == reversed
    }
    
    // MARK: - Test Cases
    
    func testCase1() {
        XCTAssertTrue(isPalindrome(121))
    }
    
    func testCase2() {
        XCTAssertFalse(isPalindrome(-121))
    }
    
    func testCase3() {
        XCTAssertFalse(isPalindrome(10))
    }
    
    func testCase4() {
        XCTAssertTrue(isPalindrome(0))
    }
    
    func testCase5() {
        XCTAssertTrue(isPalindrome(1))
    }
} 