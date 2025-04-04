local lfs = require("lfs")

local function get_script_directory()
    local str = debug.getinfo(1, "S").source:sub(2)
    local dir = str:match("(.*/)") or "./"
	return dir:gsub("/+$", "")
end

local function get_problem_dirs(root)
	local iter, obj = lfs.dir(root)
	return function()
		local file = iter(obj)
		while file and not string.find(file, "^p%d%d%d%d$") do
			file = iter(obj)
		end
		return file
	end
end

function abspath(path)
	local here = lfs.currentdir()
	lfs.chdir(path)
	local abs = lfs.currentdir()
	lfs.chdir(here)
	return abs
end

function path_join(...)
    local separator = package.config:sub(1,1)
    return table.concat({...}, separator)
end

function run_test(root, dir)
	local lua = path_join(root, "lua")
	lfs.chdir(dir)
	local status = os.execute(lua .. " tests.lua")
	lfs.chdir(root)
	return status
end

local root = get_script_directory()
for problem_dir in get_problem_dirs(root) do
	local test_file = path_join(root, problem_dir, "tests.lua")
	if lfs.attributes(test_file) then
		local test = require(problem_dir .. "/tests")
		test()
	end
end