# 练习 01: Hello World

## 目标
实现一个返回问候字符串的函数，以验证你的 Go 环境设置正确。

## 指南

1.  在你的用户目录下创建一个名为 `01-hello-world` 的目录 (例如 `lingfenglong/01-hello-world`)。
2.  将本目录下的测试文件 `hello_test.go` 复制到你的新目录中。
3.  在你的目录下创建一个名为 `hello.go` 的文件。
4.  实现一个名为 `hello` 的包 (`package hello`)。
5.  实现一个函数 `Greeting() string`，使其返回 "Hello, Go Study!"。
6.  运行测试: `go test -v`

## 预期输出
```
=== RUN   TestGreeting
    hello_test.go:9: Greeting() = "Hello, Go Study!"; want "Hello, Go Study!"
--- PASS: TestGreeting (0.00s)
PASS
```
