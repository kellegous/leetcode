local luaunit = require("luaunit")
local solution = require("p0001/solution")

TestTwoSum = {}

function TestTwoSum:testExample1()
    local result = solution.two_sum({2, 7, 11, 15}, 9)
    luaunit.assertEquals(result, {1, 2})
end

function TestTwoSum:testExample2()
    local result = solution.two_sum({3, 2, 4}, 6)
    luaunit.assertEquals(result, {2, 3})
end

function TestTwoSum:testExample3()
    local result = solution.two_sum({3, 3}, 6)
    luaunit.assertEquals(result, {1, 2})
end

function TestTwoSum:testNoSolution()
    local result = solution.two_sum({1, 2, 3}, 7)
    luaunit.assertNil(result)
end

-- Run the test cases
os.exit(luaunit.LuaUnit.run())
