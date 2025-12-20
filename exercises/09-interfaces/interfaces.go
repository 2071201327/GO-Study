package interfaces

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	// TODO: 实现该方法
	panic("TODO")
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	// TODO: 实现该方法
	panic("TODO")
}

// TotalArea 计算所有形状的总面积
func TotalArea(shapes []Shape) float64 {
	// TODO: 实现该方法
	panic("TODO")
}
