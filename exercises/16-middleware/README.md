# 练习 16: 中间件 (Middleware)

## 目标
学习如何编写 HTTP 中间件，拦截请求和响应。

## 指南

1.  创建 `16-middleware` 目录。
2.  创建 `middleware.go`。
3.  实现函数 `LoggingMiddleware(next http.Handler) http.Handler`:
    - 在调用 `next` 之前，记录一条日志（这里我们改为给响应头添加 `X-Logged: true` 以便测试验证）。
    - 调用 `next.ServeHTTP`。

## 预期输出
测试通过。
