# 练习 02: 变量与基本类型

## 目标
理解 Go 语言中的变量声明、基本类型（字符串、整数、布尔值）以及零值的概念。

## 指南

1.  在你的用户目录下创建 `02-variables-types` 目录。
2.  复制本目录下的 `variables_test.go`。
3.  创建 `variables.go` 并实现 `variables` 包。
4.  实现以下函数/变量以通过测试：
    - `GetName()`: 返回字符串 "Gopher"。
    - `GetAge()`: 返回整数 `10`。
    - `IsCool()`: 返回布尔值 `true`。
    - `GetZeroValues()`: 返回一个整数和一个字符串的零值（分别为 `0` 和 `""`）。

## 预期输出
`go test -v` 应该全部通过。
