# 练习 10: 包 (Packages)

## 目标
理解包的可见性（导出 vs 未导出）。

## 指南

1.  创建 `10-packages` 目录。
2.  由于 Go 的包管理限制，在同一个目录下测试包可见性比较特殊。本练习我们模拟创建两个文件。
3.  创建 `mypkg.go`，包名为 `mypkg` (注意：测试文件通常属于 `mypkg` 或 `mypkg_test`)。为了简化，我们这里都使用 package `packages`。
4.  实现：
    - 公有函数 `PublicFunc() string`: 返回 "public"。
    - 私有函数 `privateFunc() string`: 返回 "private"。
5.  注意：在同一个包内，私有函数是可见的。这个练习主要验证你能否正确命名（大写开头为导出）。

## 预期输出
测试通过。
