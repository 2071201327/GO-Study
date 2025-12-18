# 练习 14: JSON 处理

## 目标
学习 Struct 和 JSON 之间的转换 (Marshal/Unmarshal)。

## 指南

1.  创建 `14-json` 目录。
2.  创建 `json.go`。
3.  定义结构体 `Response`:
    - `Message` (string, json tag: "message")
    - `Code` (int, json tag: "code")
4.  实现函数：
    - `Encode(r Response) ([]byte, error)`: 将结构体转换为 JSON。
    - `Decode(data []byte) (Response, error)`: 将 JSON 转换为结构体。

## 预期输出
测试通过。
