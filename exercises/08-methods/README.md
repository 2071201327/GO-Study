# 练习 08: 方法 (Methods)

## 目标
理解 Go 方法及其接收者（Receiver），包括值接收者和指针接收者。

## 指南

1.  创建 `08-methods` 目录。
2.  创建 `methods.go`。
3.  定义结构体 `Counter`，包含字段 `Value` (int)。
4.  实现方法：
    - `(c Counter) GetValue() int`: 返回当前值。
    - `(c *Counter) Increment()`: 将 Value 增加 1。
5.  思考：如果 `Increment` 使用值接收者会发生什么？

## 预期输出
测试通过。
