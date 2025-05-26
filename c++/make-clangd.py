#!/usr/bin/env python3

import os
import subprocess

import yaml


def get_cflags_for(pkg):
    result = subprocess.run(
        ["pkg-config", "--cflags", pkg], capture_output=True, text=True
    )
    return result.stdout.strip().split(" ")


contents = {
    "CompileFlags": {
        "Add": [
            "-std=c++23",
        ]
        + get_cflags_for("gtest_main")
    }
}
with open(".clangd", "w") as f:
    yaml.dump(contents, f)
