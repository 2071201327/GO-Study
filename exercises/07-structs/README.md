# 练习 07: 结构体 (Structs)

## 目标
学习结构体的定义、初始化和嵌套。

## 指南

1.  创建 `07-structs` 目录。
2.  创建 `structs.go`。
3.  定义结构体 `User`:
    - Name (string)
    - Age (int)
    - Address (Address struct)
4.  定义结构体 `Address`:
    - City (string)
5.  实现函数：
    - `NewUser(name string, age int, city string) User`: 返回初始化的 User。
    - `UpdateAge(u *User, newAge int)`: 修改 User 的年龄（注意指针）。

## 预期输出
测试通过。
