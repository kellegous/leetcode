// swift-tools-version: 5.9

import PackageDescription

let package = Package(
    name: "LeetCodeSwift",
    platforms: [
        .macOS(.v10_15)
    ],
    products: [
        .library(
            name: "LeetCodeSwift",
            targets: ["LeetCodeSwift"]),
    ],
    targets: [
        .target(
            name: "LeetCodeSwift",
            dependencies: []),
        .testTarget(
            name: "LeetCodeSwiftTests",
            dependencies: ["LeetCodeSwift"]),
    ]
) 