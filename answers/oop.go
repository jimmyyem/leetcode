package answers

import "math"

//---------- 接 口 --------//
type Shape interface {
	Area() float64 //计算面积
	Perimeter() float64 //计算周长
}

//--------- 长方形 ----------//
type Rect struct {
	Width, Height float64
}

func (r *Rect) Area() float64 { //面积
	return r.Width * r.Height
}

func (r *Rect) Perimeter() float64 { //周长
	return 2*(r.Width + r.Height)
}

//----------- 圆  形 ----------//
type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 { //面积
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 { //周长
	return 2 * math.Pi * c.Radius
}

// ----------- 接口的使用 -----------//
//func interface_test() {
//	r := rect {width:2.9, height:4.8}
//	c := circle {radius:4.3}
//
//	s := []shape{&r, &c} //通过指针实现
//
//	for _, sh := range s {
//		fmt.Println(sh)
//		fmt.Println(sh.area())
//		fmt.Println(sh.perimeter())
//	}
//}