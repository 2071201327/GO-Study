# 练习 13: HTTP Server

## 目标
学习使用 `net/http` 包创建一个简单的 HTTP 服务器。

## 指南

1.  创建 `13-http-server` 目录。
2.  创建 `server.go`。
3.  实现函数 `HelloHandler(w http.ResponseWriter, r *http.Request)`:
    - 写入响应 "Hello, Web!"。
4.  在测试中，我们将使用 `httptest` 来验证 Handler，而不需要真正启动服务器监听端口。

## 预期输出
测试通过。
