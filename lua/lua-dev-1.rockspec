package = "lua"
version = "dev-1"
source = {
   url = "git+https://github.com/kellegous/leetcode.git"
}
description = {
   homepage = "https://github.com/kellegous/leetcode",
   license = "*** please specify a license ***"
}
dependencies = {
   "lua ~> 5.4"
}
build = {
   type = "builtin",
   modules = {
      main = "main.lua"
   }
}
