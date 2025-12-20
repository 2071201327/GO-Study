# 练习 03: 控制结构

## 目标
掌握 `if/else`, `switch` 和 `for` 循环的使用。

## 指南

1.  创建 `03-control-structures` 目录并复制测试文件。
2.  创建 `control.go`，包名为 `control`。
3.  实现以下函数：
    - `CheckNumber(n int) string`: 如果 `n` 是正数返回 "positive"，负数返回 "negative"，零返回 "zero"。
    - `GetDayName(day int) string`: 使用 `switch`，1 返回 "Monday"，7 返回 "Sunday"，其他返回 "Unknown"。
    - `SumUpTo(n int) int`: 使用 `for` 循环计算 1 到 `n` 的和。

## 提示
Go 只有 `for` 循环，没有 `while`。
