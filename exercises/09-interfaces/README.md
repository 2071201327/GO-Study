# 练习 09: 接口 (Interfaces)

## 目标
学习接口的定义和实现，以及依赖注入的基本概念。

## 指南

1.  创建 `09-interfaces` 目录。
2.  创建 `interfaces.go`。
3.  定义接口 `Shape`:
    - `Area() float64`
4.  定义结构体 `Rectangle` (Width, Height float64) 和 `Circle` (Radius float64)。
5.  让 `Rectangle` 和 `Circle` 都实现 `Shape` 接口。
6.  实现函数 `TotalArea(shapes []Shape) float64`，计算所有形状的面积之和。

## 提示
使用 `math.Pi`。

## 预期输出
测试通过。
