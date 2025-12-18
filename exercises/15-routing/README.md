# 练习 15: 路由 (Routing)

## 目标
实现简单的根据 URL 路径分配不同的处理逻辑。

## 指南

1.  创建 `15-routing` 目录。
2.  创建 `routing.go`。
3.  实现 `ServeHTTP` 方法使结构体 `Router` 实现 `http.Handler` 接口。
4.  逻辑：
    - 如果路径是 `/`，返回 "Home"
    - 如果路径是 `/about`，返回 "About"
    - 其他返回 404 "Not Found"

## 预期输出
测试通过。
