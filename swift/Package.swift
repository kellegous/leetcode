// swift-tools-version: 5.9

import PackageDescription

let package = Package(
    name: "LeetCode",
    platforms: [
        .macOS(.v10_15)
    ],
    products: [
        .library(
            name: "LeetCode",
            targets: ["LeetCode"]),
    ],
    targets: [
        .target(
            name: "LeetCode",
            dependencies: []),
        .testTarget(
            name: "LeetCodeTests",
            dependencies: ["LeetCode"]),
    ]
) 
