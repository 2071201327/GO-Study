# 练习 11: Goroutines (协程)

## 目标
学习如何启动协程，并使用 `sync.WaitGroup` 等待协程结束。

## 指南

1.  创建 `11-goroutines` 目录。
2.  创建 `goroutines.go`。
3.  实现函数 `RunConcurrently(n int) int`:
    - 启动 `n` 个协程。
    - 每个协程完成一些工作（这里我们简单地忽略，只计数）。
    - 必须等待所有协程完成后函数才返回。
    - 返回实际执行的协程数量（应该是 `n`）。
4.  提示：使用 `sync.WaitGroup` 和 `atomic.AddInt32` (或者互斥锁) 来安全地计数，模拟并发聚合。

## 预期输出
测试通过。
